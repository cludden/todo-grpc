import * as bunyan from 'bunyan';
import { Context } from 'koa';

export const inject = {
  name: 'http/middleware/logger',
  require: ['log'],
};

export default function (log: bunyan) {
  // define a request scoped logger instance
  return async function (ctx: Context, next: () => Promise<any>) {
    const { id } = ctx.state;
    const child = log.child({ req_id: id });
    ctx.state.log = child;
    return next();
  };
}
