import { graphiqlKoa, graphqlKoa } from 'apollo-server-koa';
import { GraphQLSchema } from 'graphql';
import voyager from 'graphql-voyager/middleware/koa';
import * as koa from 'koa';
import * as Router from 'koa-router';

import { Config } from '../../types';

export const inject = {
  name: 'http/routes/graphql',
  require: ['config', 'graphql', 'http/middleware/context'],
};

// tslint:disable-next-line
export default function (config: Config, schema: GraphQLSchema, context: koa.Middleware) {
  // create new router
  const router = new Router();

  // define graphql handler
  router.post('/graphql', context, graphqlKoa(async (ctx: koa.BaseContext) => ({
    schema,
    context: ctx.state.context,
  })));

  // define graphiql handler if configured
  if (config.graphql.enable_graphiql === true) {
    router.get('/graphiql', graphiqlKoa({
      endpointURL: '/graphql',
    }));
  }
  if (config.graphql.enable_voyager === true) {
    router.all('/voyager', voyager({
      endpointUrl: '/graphql',
    }));
  }

  return router;
}
