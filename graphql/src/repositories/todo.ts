import { Config } from '../types';
import * as blah from './todo_grpc_pb' // tslint:disable-line

export const inject = {
  name: 'repositories/todo',
  require: ['config'],
};

export default async function(config: Config) {
  blah.
  console.log(config);
}