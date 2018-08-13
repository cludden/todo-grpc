package local

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"todo-grpc/proto"
	"todo-grpc/todo"

	"github.com/boltdb/bolt"
)

// Repository provides access to a local todos store
type Repository struct {
	bucket []byte
	db     *bolt.DB
	n      int64
}

// NewRepository creates a new repository value
func NewRepository() (*Repository, error) {
	// open db
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", db)
	}

	// create and return repository
	r := Repository{
		bucket: []byte(`todos`),
		db:     db,
	}
	return &r, nil
}

// CreateTodo creates a new todo
func (r *Repository) CreateTodo(ctx context.Context, in *todo.Todo) error {
	// marshal json
	raw, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %v", err)
	}

	// insert json
	err = r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(r.bucket)
		return bucket.Put([]byte(in.ID), raw)
	})
	if err != nil {
		return err
	}

	// increment count
	atomic.AddInt64(&r.n, 1)
	return nil
}

// ListTodos retrieves a paginated list of todos
func (r *Repository) ListTodos(ctx context.Context, in *proto.ListTodosInput) (*todo.ListTodosOutput, error) {
	var out todo.ListTodosOutput

	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(r.bucket)
		cursor := bucket.Cursor()

		// retrieve first todo
		var k, v []byte
		if in.After != "" {
			after := []byte(in.After)
			if val := bucket.Get(after); val == nil {
				return fmt.Errorf("invalid cursor: %s", in.After)
			}
			cursor.Seek(after)
			k, v = cursor.Next()
		} else {
			k, v := cursor.First()
		}
		if k == nil {
			return nil
		}
		var t todo.Todo
		if err := json.Unmarshal(v, &t); err != nil {
			return err
		}
		out.Todos = append(out.Todos, &t)

		// retrieve next page of keys
		n := 1
		for k, v := cursor.Next(); k != nil && n < in.First; k, v = cursor.Next() {

		}
	})
	return &out, err
}
