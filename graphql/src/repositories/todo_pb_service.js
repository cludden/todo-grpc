// package: mindflash.todo
// file: todo.proto

var todo_pb = require("./todo_pb");
var grpc = require("grpc-web-client").grpc;

var Todos = (function () {
  function Todos() {}
  Todos.serviceName = "mindflash.todo.Todos";
  return Todos;
}());

Todos.CompleteTodo = {
  methodName: "CompleteTodo",
  service: Todos,
  requestStream: false,
  responseStream: false,
  requestType: todo_pb.CompleteTodoInput,
  responseType: todo_pb.CompleteTodoOutput
};

Todos.CreateTodo = {
  methodName: "CreateTodo",
  service: Todos,
  requestStream: false,
  responseStream: false,
  requestType: todo_pb.CreateTodoInput,
  responseType: todo_pb.CreateTodoOutput
};

Todos.ListTodos = {
  methodName: "ListTodos",
  service: Todos,
  requestStream: false,
  responseStream: false,
  requestType: todo_pb.ListTodosInput,
  responseType: todo_pb.ListTodosOutput
};

exports.Todos = Todos;

function TodosClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

TodosClient.prototype.completeTodo = function completeTodo(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Todos.CompleteTodo, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

TodosClient.prototype.createTodo = function createTodo(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Todos.CreateTodo, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

TodosClient.prototype.listTodos = function listTodos(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Todos.ListTodos, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.TodosClient = TodosClient;

