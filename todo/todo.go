package todo

import (
	"errors"
	"time"
	"todo-grpc/proto"

	"github.com/golang/protobuf/ptypes"
	null "gopkg.in/guregu/null.v3"
)

// Todo domain type
type Todo struct {
	ID          string      `json:"id"`
	Complete    bool        `json:"complete"`
	CompletedAt null.Time   `json:"completed_at"`
	CreatedAt   time.Time   `json:"created_at"`
	Description null.String `json:"description"`
	Title       string      `json:"title"`
}

// MarshalPB returns a valid protobuf message
func (t *Todo) MarshalPB() (*proto.Todo, error) {
	if t == nil {
		return nil, errors.New("unable to marshal nil todo")
	}
	todo := proto.Todo{
		Id:          t.ID,
		Complete:    t.Complete,
		Description: t.Description.ValueOrZero(),
		Title:       t.Title,
	}

	if t.CompletedAt.Valid {
		completedAt, err := ptypes.TimestampProto(t.CompletedAt.ValueOrZero())
		if err != nil {
			return nil, err
		}
		todo.CompletedAt = completedAt
	}

	createdAt, err := ptypes.TimestampProto(t.CreatedAt)
	if err != nil {
		return nil, err
	}
	todo.CreatedAt = createdAt

	return &todo, nil
}

// UnmarshalTodoPB unmarshals the protobuf message into a valid domain type
func UnmarshalTodoPB(t *proto.Todo) (*Todo, error) {
	todo := Todo{
		ID:          t.GetId(),
		Complete:    t.GetComplete(),
		Description: null.StringFrom(t.GetDescription()),
		Title:       t.GetTitle(),
	}

	completedAt, err := ptypes.Timestamp(t.GetCompletedAt())
	if err != nil {
		return nil, err
	}
	todo.CompletedAt = null.TimeFrom(completedAt)

	createdAt, err := ptypes.Timestamp(t.GetCreatedAt())
	if err != nil {
		return nil, err
	}
	todo.CreatedAt = createdAt

	return &todo, nil
}
