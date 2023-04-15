package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/egorgasay/dockerdb/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"grpc-storage/internal/grpc-storage/transaction-logger/service"
	"log"
	"strings"
	"sync"

	_ "github.com/jackc/pgx"
)

type TransactionLogger struct {
	db   *sql.DB
	mu   sync.Mutex
	path string

	events chan service.Event
	errors chan error
}

const insertQuery = `UPSERT INTO transactions (event_type, key, value) VALUES 
                                                      ($1, $2, $3), ($4, $5, $6), ($7, $8, $9), ($10, $11, $12), 
                                                      ($13, $14, $15), ($16, $17, $18), ($19, $20, $21), ($22, $23, $24),
                                                      ($25, $26, $27), ($28, $29, $30), ($31, $32, $33), ($34, $35, $36),
                                                      ($37, $38, $39), ($40, $41, $42), ($43, $44, $45), ($46, $47, $48), 
                                                      ($49, $50, $51), ($52, $53, $54), ($55, $56, $57), ($58, $59, $60) `

var insertEvent *sql.Stmt

func NewLogger(path string) (*TransactionLogger, error) {
	path = strings.TrimRight(path, "/") + "/transactionLoggerDB"
	cfg := dockerdb.CustomDB{
		DB: dockerdb.DB{
			Name:     "adm",
			User:     "adm",
			Password: "adm",
		},
		Port:   "334",
		Vendor: dockerdb.Postgres15,
	}

	vdb, err := dockerdb.New(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(vdb.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/sqlite3_tlogger",
		"sqlite", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	insertEvent, err = vdb.DB.Prepare(insertQuery)
	if err != nil {
		return nil, err
	}

	return &TransactionLogger{
		db:   vdb.DB,
		path: path,
	}, nil
}

func (t *TransactionLogger) WriteSet(key, value string) {
	t.events <- service.Event{EventType: service.Set, Key: key, Value: value}
}

func (t *TransactionLogger) WriteDelete(key string) {
	t.events <- service.Event{EventType: service.Delete, Key: key}
}

func (t *TransactionLogger) Run() {
	events := make(chan service.Event, 20)

	t.events = events
	var data = make([]service.Event, 0, 20)
	var dataBackup = make([]service.Event, 20)
	go func() {
		for e := range events {
			data = append(data, e)
			if len(data) == 20 {
				copy(dataBackup, data)
				go t.flash(dataBackup)
				data = data[:0]
			}
		}
	}()
}

// flash grabs 20 events and saves them to the db.
func (t *TransactionLogger) flash(data []service.Event) {
	//t.mu.Lock()
	//defer t.mu.Unlock()
	errorsch := make(chan error, 20)
	t.errors = errorsch

	var anys = make([]any, 0, len(data)*3)

	for _, e := range data {
		anys = append(anys, e.EventType, e.Key, e.Value)
	}

	_, err := insertEvent.Exec(anys...)
	if err != nil {
		log.Println("Run:", err)
		errorsch <- err
	}
}

func (t *TransactionLogger) Err() <-chan error {
	return t.errors
}

func (t *TransactionLogger) ReadEvents() (<-chan service.Event, <-chan error) {
	outEvent := make(chan service.Event)
	outError := make(chan error, 1)

	rows, err := t.db.Query("SELECT event_type, key, value FROM transactions")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		outError <- err
		return outEvent, outError
	}

	go func() {
		var event service.Event
		defer close(outEvent)
		defer close(outError)

		for rows.Next() {
			err = rows.Scan(&event.EventType, &event.Key, &event.Value)
			if err != nil {
				log.Println("ReadEvents:", err)
				continue
			}

			outEvent <- event
		}
	}()

	return outEvent, outError
}

func (t *TransactionLogger) Clear() error {
	_, err := t.db.Exec("DELETE FROM transactions")
	if err != nil {
		return err
	}

	return nil
}
