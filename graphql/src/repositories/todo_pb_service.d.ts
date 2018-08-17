// package: mindflash.todo
// file: todo.proto

import * as todo_pb from "./todo_pb";
import {grpc} from "grpc-web-client";

type TodosCompleteTodo = {
  readonly methodName: string;
  readonly service: typeof Todos;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof todo_pb.CompleteTodoInput;
  readonly responseType: typeof todo_pb.CompleteTodoOutput;
};

type TodosCreateTodo = {
  readonly methodName: string;
  readonly service: typeof Todos;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof todo_pb.CreateTodoInput;
  readonly responseType: typeof todo_pb.CreateTodoOutput;
};

type TodosListTodos = {
  readonly methodName: string;
  readonly service: typeof Todos;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof todo_pb.ListTodosInput;
  readonly responseType: typeof todo_pb.ListTodosOutput;
};

export class Todos {
  static readonly serviceName: string;
  static readonly CompleteTodo: TodosCompleteTodo;
  static readonly CreateTodo: TodosCreateTodo;
  static readonly ListTodos: TodosListTodos;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class TodosClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  completeTodo(
    requestMessage: todo_pb.CompleteTodoInput,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: todo_pb.CompleteTodoOutput|null) => void
  ): void;
  completeTodo(
    requestMessage: todo_pb.CompleteTodoInput,
    callback: (error: ServiceError, responseMessage: todo_pb.CompleteTodoOutput|null) => void
  ): void;
  createTodo(
    requestMessage: todo_pb.CreateTodoInput,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: todo_pb.CreateTodoOutput|null) => void
  ): void;
  createTodo(
    requestMessage: todo_pb.CreateTodoInput,
    callback: (error: ServiceError, responseMessage: todo_pb.CreateTodoOutput|null) => void
  ): void;
  listTodos(
    requestMessage: todo_pb.ListTodosInput,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: todo_pb.ListTodosOutput|null) => void
  ): void;
  listTodos(
    requestMessage: todo_pb.ListTodosInput,
    callback: (error: ServiceError, responseMessage: todo_pb.ListTodosOutput|null) => void
  ): void;
}

