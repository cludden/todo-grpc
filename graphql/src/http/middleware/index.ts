import { Middleware } from 'koa';

export const inject = {
  name: 'http/middleware',
  require: {
    context: 'http/middleware/context',
    logger: 'http/middleware/logger',
    requestId: 'http/middleware/request-id',
  },
};

export interface MiddlewareMap {
  logger: Middleware;
  requestId: Middleware;
}

export default function (mw: MiddlewareMap) {
  return mw;
}
