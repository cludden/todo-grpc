import { BaseContext } from 'koa';

export const inject = {
  name: 'http/middleware/context',
};

export const GRAPHQL_CONTEXT = 'GRAPHQL_CONTEXT';

export default function () {
  return async function graphqlContext(ctx: BaseContext, next: () => Promise<any>) {
    const context = {
      id: ctx.state.id,
      log: ctx.state.log,
    };
    ctx.state.context = context;
    return next();
  };
}
