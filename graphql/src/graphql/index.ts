import { fromCallback } from 'bluebird';
import * as fs from 'fs';
import * as glob from 'glob';
import { addMockFunctionsToSchema, makeExecutableSchema, IMockFn, IResolvers } from 'graphql-tools';
import * as path from 'path';

import { Config } from '../types';

export const inject = {
  name: 'graphql',
  require: ['container!', 'config', 'any!graphql/[^\/]+/resolver'],
};

export default async function (container: any, config: Config, resolvers: IResolvers[]) {
  // load modular schemas
  const typeDefs = await fromCallback<string[]>(done => glob(`**/*.graphql`, {
    cwd: __dirname,
    matchBase: true,
  }, done))
    .map((f: string) => {
      const file = path.join(__dirname, `./${f}`);
      return fromCallback(done => fs.readFile(file, 'utf8', done));
    });

  // build executable schema
  const schema = makeExecutableSchema({ typeDefs, resolvers });

  // enable mocks if applicable
  if (config.graphql.enable_mocking === true) {
    const mockDefs = await container.load('any!graphql/[^\/]+/mock') as { [k: string]: IMockFn }[];
    const mocks = mockDefs.reduce((acc, m) => Object.assign(acc, m), {});
    addMockFunctionsToSchema({ schema, mocks });
  }

  return schema;
}
