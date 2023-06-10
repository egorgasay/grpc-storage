package usecase

import (
	"context"
	"errors"
	"fmt"
	"itisadb/internal/memory-balancer/servers"
)

func (uc *UseCase) Set(ctx context.Context, key, val string, serverNumber int32, uniques bool) (int32, error) {
	if uc.servers.Len() == 0 {
		return 0, ErrNoServers
	}

	if serverNumber == setToAll {
		failedServers := uc.servers.SetToAll(ctx, key, val, uniques)
		if len(failedServers) != 0 {
			return setToAll, fmt.Errorf("some servers wouldn't get values: %v", failedServers)
		}
		return setToAll, nil
	}

	var cl *servers.Server
	var ok bool

	if serverNumber > 0 {
		cl, ok = uc.servers.GetServerByID(serverNumber)
		if !ok || cl == nil {
			return 0, ErrUnknownServer
		}
	} else {
		cl, ok = uc.servers.GetServer()
		if !ok || cl == nil {
			return 0, ErrNoServers
		}
	}

	err := cl.Set(context.Background(), key, val, uniques)
	if err != nil {
		return 0, err
	}

	return cl.GetNumber(), nil
}

var ErrNoServers = errors.New("no servers available")
var ErrNotFound = errors.New("key not found")

func (uc *UseCase) Get(ctx context.Context, key string, serverNumber int32) (string, error) {
	if uc.servers.Len() == 0 {
		return "", ErrNoServers
	}

	if serverNumber == searchEverywhere {
		value, err := uc.servers.DeepSearch(ctx, key)
		if errors.Is(err, servers.ErrNotFound) {
			return "", ErrNotFound
		}
		return value, err
	} else if !uc.servers.Exists(serverNumber) {
		return "", ErrNotFound
	}

	cl, ok := uc.servers.GetServerByID(serverNumber)
	if !ok || cl == nil {
		return "", ErrUnknownServer
	}

	res, err := cl.Get(context.Background(), key)
	if err == nil {
		cl.ResetTries()
		return res.Value, nil
	}

	uc.logger.Warn(err.Error())

	cl.IncTries()

	if cl.GetTries() > 2 {
		err = uc.Disconnect(ctx, cl.GetNumber())
		if err != nil {
			uc.logger.Warn(err.Error())
		}
	}

	return "", ErrNotFound
}

func (uc *UseCase) Connect(address string, available, total uint64, server int32) (int32, error) {
	uc.logger.Info("New request for connect from " + address)
	number, err := uc.servers.AddServer(address, available, total, server)
	if err != nil {
		uc.logger.Warn(err.Error())
		return 0, err
	}

	return number, nil
}

func (uc *UseCase) Disconnect(ctx context.Context, number int32) error {
	ch := make(chan struct{})
	uc.pool <- struct{}{}
	go func() { // TODO: add pool
		uc.servers.Disconnect(number)
		close(ch)
		<-uc.pool
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (uc *UseCase) Servers() []string {
	return uc.servers.GetServers()
}

func (uc *UseCase) Delete(ctx context.Context, key string, num int32) (err error) {
	ch := make(chan struct{}) // TODO: handle possible memory leak

	uc.pool <- struct{}{}
	go func() {
		err = uc.delete(ctx, key, num)
		close(ch)
		<-uc.pool
	}()

	select {
	case <-ch:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (uc *UseCase) delete(ctx context.Context, key string, num int32) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if num == 0 {
		num = setToAll
	}

	switch num {
	case setToAll:
		// TODO: delete from setToAll servers
	}

	cl, ok := uc.servers.GetServerByID(num)
	if !ok || cl == nil {
		return ErrUnknownServer
	}

	err := cl.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("error while deleting value: %w", err)
	}
	return nil
}
