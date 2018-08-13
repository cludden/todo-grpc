package todo

import (
	"time"
	"todo-grpc/proto"

	"github.com/golang/protobuf/ptypes"
	null "gopkg.in/guregu/null.v3"
)

// Todo domain type
type Todo struct {
	ID          string
	Complete    bool
	CompletedAt null.Time
	CreatedAt   time.Time
	Description null.String
	Title       string
}

// MarshalPB returns a valid protobuf message
func (t *Todo) MarshalPB() (*proto.Todo, error) {
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
