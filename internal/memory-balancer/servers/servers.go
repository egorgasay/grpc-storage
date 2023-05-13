package servers

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"itisadb/pkg/api/storage"
	"os"
	"strconv"
	"sync"
)

type Servers struct {
	servers map[int32]*Server
	freeID  int32
	sync.RWMutex
}

var ErrNotFound = errors.New("the value was not found")

func New() (*Servers, error) {
	f, err := os.OpenFile("servers", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var number uint32 = 0
	s := make(map[int32]*Server, 10)
	servers := &Servers{
		servers: s,
		freeID:  int32(number + 1),
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		_, err = f.WriteString("1")
		return servers, err
	}

	line := scanner.Text()
	atoi, err := strconv.Atoi(line)
	if err != nil {
		return nil, fmt.Errorf("can't get the last used number: %w", err)
	}

	servers.freeID = int32(atoi) + 1

	return servers, nil
}

func (s *Servers) GetClient() (*Server, bool) {
	s.RLock()
	defer s.RUnlock()

	max := 0.0
	var serverNumber int32 = 0

	for num, cl := range s.servers {
		r := cl.GetRAM()
		if float64(r.available)/float64(r.total)*100 > max {
			serverNumber = num
		}
	}

	cl, ok := s.servers[serverNumber]
	return cl, ok
}

func (s *Servers) Len() int32 {
	s.RLock()
	defer s.RUnlock()
	return int32(len(s.servers))
}

func (s *Servers) AddClient(address string, available, total uint64, server int32) (int32, error) {
	s.Lock()
	defer s.Unlock()

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}

	cl := storage.NewStorageClient(conn)

	var stClient = &Server{
		storage: cl,
		ram:     RAM{available: available, total: total},
		mu:      &sync.RWMutex{},
	}

	if server != 0 {
		stClient.number = server
		if server > s.freeID {
			s.freeID = server + 1
		}
	} else {
		f, err := os.OpenFile("servers", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return 0, err
		}
		defer f.Close()

		stClient.number = s.freeID
		s.freeID++
		_, err = f.WriteString(fmt.Sprintf("%d", s.freeID))
		if err != nil {
			return 0, fmt.Errorf("can't save last id: %w", err)
		}
	}

	s.servers[stClient.number] = stClient

	return stClient.number, nil
}

func (s *Servers) Disconnect(number int32) {
	s.Lock()
	defer s.Unlock()
	delete(s.servers, number)
}

func (s *Servers) GetServers() []string {
	s.RLock()
	defer s.RUnlock()

	var servers = make([]string, 0, 5)
	for _, cl := range s.servers {
		r := cl.GetRAM()
		servers = append(servers, fmt.Sprintf("s#%d Avaliable: %d MB, Total: %d MB", cl.GetNumber(), r.total, r.available))
	}

	return servers
}

func (s *Servers) DeepSearch(ctx context.Context, key string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var out = make(chan string, 1)
	defer close(out)

	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	// TODO: Add pull of goroutines
	var once = sync.Once{}
	for _, cl := range s.servers {
		c := cl
		go func() {
			defer wg.Done()
			c.find(ctxCancel, key, out, &once)
		}()
	}

	allIsDone := make(chan struct{})

	go func() {
		defer close(allIsDone)
		wg.Wait()
		allIsDone <- struct{}{}
	}()

	select {
	case v := <-out:
		cancel()
		return v, nil
	case <-allIsDone:
		return "", ErrNotFound
	}
}

func (s *Server) find(ctx context.Context, key string, out chan<- string, once *sync.Once) {
	get, err := s.Get(ctx, key)
	if err != nil {
		return
	}

	once.Do(func() {
		out <- get.Value
	})
}

func (s *Servers) GetClientByID(number int32) (*Server, bool) {
	s.RLock()
	defer s.RUnlock()
	srv, ok := s.servers[number]
	return srv, ok
}

func (s *Servers) Exists(number int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.servers[number]
	return ok
}

func (s *Servers) SetToAll(ctx context.Context, key, val string, uniques bool) []int32 {
	var failedServers = make([]int32, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup

	s.RLock()
	defer s.RUnlock()

	wg.Add(len(s.servers))
	for n, serv := range s.servers {
		go func(server *Server, number int32) {
			defer wg.Done()
			err := server.Set(ctx, key, val, uniques)
			if err != nil {
				if server.GetTries() > 2 {
					delete(s.servers, number)
				}
				server.IncTries()
				mu.Lock()
				failedServers = append(failedServers, number)
				mu.Unlock()
				return
			}

			server.ResetTries()
		}(serv, n)
	}
	wg.Wait()

	return failedServers
}
