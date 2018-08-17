// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var todo_pb = require('./todo_pb.js');
var google_api_annotations_pb = require('./google/api/annotations_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_mindflash_todo_CompleteTodoInput(arg) {
  if (!(arg instanceof todo_pb.CompleteTodoInput)) {
    throw new Error('Expected argument of type mindflash.todo.CompleteTodoInput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_CompleteTodoInput(buffer_arg) {
  return todo_pb.CompleteTodoInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mindflash_todo_CompleteTodoOutput(arg) {
  if (!(arg instanceof todo_pb.CompleteTodoOutput)) {
    throw new Error('Expected argument of type mindflash.todo.CompleteTodoOutput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_CompleteTodoOutput(buffer_arg) {
  return todo_pb.CompleteTodoOutput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mindflash_todo_CreateTodoInput(arg) {
  if (!(arg instanceof todo_pb.CreateTodoInput)) {
    throw new Error('Expected argument of type mindflash.todo.CreateTodoInput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_CreateTodoInput(buffer_arg) {
  return todo_pb.CreateTodoInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mindflash_todo_CreateTodoOutput(arg) {
  if (!(arg instanceof todo_pb.CreateTodoOutput)) {
    throw new Error('Expected argument of type mindflash.todo.CreateTodoOutput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_CreateTodoOutput(buffer_arg) {
  return todo_pb.CreateTodoOutput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mindflash_todo_ListTodosInput(arg) {
  if (!(arg instanceof todo_pb.ListTodosInput)) {
    throw new Error('Expected argument of type mindflash.todo.ListTodosInput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_ListTodosInput(buffer_arg) {
  return todo_pb.ListTodosInput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_mindflash_todo_ListTodosOutput(arg) {
  if (!(arg instanceof todo_pb.ListTodosOutput)) {
    throw new Error('Expected argument of type mindflash.todo.ListTodosOutput');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_mindflash_todo_ListTodosOutput(buffer_arg) {
  return todo_pb.ListTodosOutput.deserializeBinary(new Uint8Array(buffer_arg));
}


var TodosService = exports.TodosService = {
  // Mark an existing Todo as completed
  completeTodo: {
    path: '/mindflash.todo.Todos/CompleteTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.CompleteTodoInput,
    responseType: todo_pb.CompleteTodoOutput,
    requestSerialize: serialize_mindflash_todo_CompleteTodoInput,
    requestDeserialize: deserialize_mindflash_todo_CompleteTodoInput,
    responseSerialize: serialize_mindflash_todo_CompleteTodoOutput,
    responseDeserialize: deserialize_mindflash_todo_CompleteTodoOutput,
  },
  // Creates a new Todo
  createTodo: {
    path: '/mindflash.todo.Todos/CreateTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.CreateTodoInput,
    responseType: todo_pb.CreateTodoOutput,
    requestSerialize: serialize_mindflash_todo_CreateTodoInput,
    requestDeserialize: deserialize_mindflash_todo_CreateTodoInput,
    responseSerialize: serialize_mindflash_todo_CreateTodoOutput,
    responseDeserialize: deserialize_mindflash_todo_CreateTodoOutput,
  },
  // Retrieves a paginated list of Todos
  listTodos: {
    path: '/mindflash.todo.Todos/ListTodos',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.ListTodosInput,
    responseType: todo_pb.ListTodosOutput,
    requestSerialize: serialize_mindflash_todo_ListTodosInput,
    requestDeserialize: deserialize_mindflash_todo_ListTodosInput,
    responseSerialize: serialize_mindflash_todo_ListTodosOutput,
    responseDeserialize: deserialize_mindflash_todo_ListTodosOutput,
  },
};

exports.TodosClient = grpc.makeGenericClientConstructor(TodosService);
