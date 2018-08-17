import { methodNotAllowed, notImplemented } from 'boom';
import * as bunyan from 'bunyan';
import * as http from 'http';
import * as Koa from 'koa';
import * as body from 'koa-bodyparser';
import * as Router from 'koa-router';

import { Config } from '../types';
import { MiddlewareMap } from './middleware';

export const inject = {
  name: 'http',
  type: 'constructor',
  init: 'listen',
  require: ['config', 'log', 'any!^http/routes/.+', 'http/middleware'],
};

export default class Http {
  public app: Koa;

  constructor(readonly config: Config, private log: bunyan, routes: Router[], mw: MiddlewareMap) {
    // create http app
    const app = new Koa();

    // register global middleware
    app.use(body());
    app.use(mw.requestId);
    app.use(mw.logger);

    // register routes
    routes.forEach((router) => {
      app.use(router.routes());
      app.use(router.allowedMethods({
        methodNotAllowed,
        notImplemented,
        throw: true,
      }));
    });

    this.app = app;
  }

  // listen binds the http server to the configured port and begins listening for
  // connections
  listen() {
    return new Promise<http.Server>((resolve) => {
      const server = this.app.listen(this.config.http.port, () => {
        this.log.info(`http server listening on port ${this.config.http.port}`);
        resolve(server);
      });
    });
  }
}
