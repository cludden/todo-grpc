import * as grpc from 'grpc'

import { Config } from '../types';
import { TodosClient } from './todo_pb_service' // tslint:disable-line

export const inject = {
  name: 'repositories/todo',
  require: ['config'],
};

export default async function (config: Config) {
  const client = new TodosClient(config.todos.endpoint, {
    transport: grpc.credentials.createInsecure(),
  })

  console.log(config);
}