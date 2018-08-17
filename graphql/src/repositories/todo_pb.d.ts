// package: mindflash.todo
// file: todo.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Todo extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getComplete(): boolean;
  setComplete(value: boolean): void;

  hasCompletedAt(): boolean;
  clearCompletedAt(): void;
  getCompletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCompletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getDescription(): string;
  setDescription(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
    id: string,
    complete: boolean,
    completedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    description: string,
    title: string,
  }
}

export class CompleteTodoInput extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CompleteTodoInput.AsObject;
  static toObject(includeInstance: boolean, msg: CompleteTodoInput): CompleteTodoInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CompleteTodoInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CompleteTodoInput;
  static deserializeBinaryFromReader(message: CompleteTodoInput, reader: jspb.BinaryReader): CompleteTodoInput;
}

export namespace CompleteTodoInput {
  export type AsObject = {
    id: string,
  }
}

export class CompleteTodoOutput extends jspb.Message {
  hasTodo(): boolean;
  clearTodo(): void;
  getTodo(): Todo | undefined;
  setTodo(value?: Todo): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CompleteTodoOutput.AsObject;
  static toObject(includeInstance: boolean, msg: CompleteTodoOutput): CompleteTodoOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CompleteTodoOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CompleteTodoOutput;
  static deserializeBinaryFromReader(message: CompleteTodoOutput, reader: jspb.BinaryReader): CompleteTodoOutput;
}

export namespace CompleteTodoOutput {
  export type AsObject = {
    todo?: Todo.AsObject,
  }
}

export class CreateTodoInput extends jspb.Message {
  getDescription(): string;
  setDescription(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoInput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoInput): CreateTodoInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateTodoInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoInput;
  static deserializeBinaryFromReader(message: CreateTodoInput, reader: jspb.BinaryReader): CreateTodoInput;
}

export namespace CreateTodoInput {
  export type AsObject = {
    description: string,
    title: string,
  }
}

export class CreateTodoOutput extends jspb.Message {
  hasTodo(): boolean;
  clearTodo(): void;
  getTodo(): Todo | undefined;
  setTodo(value?: Todo): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoOutput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoOutput): CreateTodoOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateTodoOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoOutput;
  static deserializeBinaryFromReader(message: CreateTodoOutput, reader: jspb.BinaryReader): CreateTodoOutput;
}

export namespace CreateTodoOutput {
  export type AsObject = {
    todo?: Todo.AsObject,
  }
}

export class ListTodosInput extends jspb.Message {
  getAfter(): string;
  setAfter(value: string): void;

  getFirst(): number;
  setFirst(value: number): void;

  getQuery(): string;
  setQuery(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosInput.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosInput): ListTodosInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListTodosInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosInput;
  static deserializeBinaryFromReader(message: ListTodosInput, reader: jspb.BinaryReader): ListTodosInput;
}

export namespace ListTodosInput {
  export type AsObject = {
    after: string,
    first: number,
    query: string,
  }
}

export class ListTodosOutput extends jspb.Message {
  clearTodosList(): void;
  getTodosList(): Array<Todo>;
  setTodosList(value: Array<Todo>): void;
  addTodos(value?: Todo, index?: number): Todo;

  getTotal(): number;
  setTotal(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosOutput.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosOutput): ListTodosOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListTodosOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosOutput;
  static deserializeBinaryFromReader(message: ListTodosOutput, reader: jspb.BinaryReader): ListTodosOutput;
}

export namespace ListTodosOutput {
  export type AsObject = {
    todosList: Array<Todo.AsObject>,
    total: number,
  }
}

