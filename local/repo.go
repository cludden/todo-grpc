// Package local provides types and methods for interacting with a local persistence layer
package local

import (
	"context"
	"encoding/json"
	"fmt"
	"todo-grpc/proto"
	"todo-grpc/todo"
	"todo-grpc/validation"

	"github.com/boltdb/bolt"
)

// Config describes the input to a NewRepository operation
type Config struct {
	Path string `validate:"required"`
}

// Repository provides access to a local todos store
type Repository struct {
	bucket []byte
	db     *bolt.DB
}

// NewRepository creates a new repository value
func NewRepository(c *Config) (*Repository, error) {
	// validate config
	if err := validation.Validate.Struct(c); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// open db
	db, err := bolt.Open(c.Path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", db)
	}

	// create todos bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(`todos`))
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("error creating bucket: %v", err)
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
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(r.bucket)
		return bucket.Put([]byte(in.ID), raw)
	})
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
			k, v = cursor.First()
		}
		if k == nil {
			return nil
		}
		var t todo.Todo
		if err := json.Unmarshal(v, &t); err != nil {
			return err
		}
		out.Todos = append(out.Todos, &t)

		// iterate until page is full or no more todos exist
		n := int32(1)
		for k, v := cursor.Next(); k != nil && n < in.First; k, v = cursor.Next() {
			var t todo.Todo
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			out.Todos = append(out.Todos, &t)
		}
		out.Total = int64(bucket.Stats().KeyN)
		return nil
	})
	return &out, err
}
