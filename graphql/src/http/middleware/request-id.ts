import { Context } from 'koa';
import { v4 } from 'uuid';

export const inject = {
  name: 'http/middleware/request-id',
  type: 'object',
};
// set request id in context state
export default async function (ctx: Context, next: () => Promise<any>) {
  let id: string = ctx.headers['x-request-id'];
  if (!id) {
    id = v4();
  }
  ctx.state.id = id;
  return next();
}
