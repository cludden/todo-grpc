import * as bunyan from 'bunyan'
import { BaseContext } from 'koa';
import { DataTypeAbstract, DefineAttributeColumnOptions } from "sequelize";

export interface Config {
  elasticsearch: {
    aws?: boolean,
    host: string | string[],
  },
  graphql: {
    enable_graphiql?: boolean,
    enable_mocking?: boolean,
    enable_voyager?: boolean,
  },
  http: {
    port: number,
  },
  log: {
    level: bunyan.LogLevelString
  },
  mysql: {
    database: string,
    host: string,
    password: string,
    port?: number,
    sync: boolean,
    user: string,
  }
}

export interface GraphqlContext {
  id: string;
  log: bunyan;
}

declare global {
  type SequelizeAttributes<T extends { [key: string]: any }> = {
    [P in keyof T]: string | DataTypeAbstract | DefineAttributeColumnOptions;
  };
}

declare module 'koa' {
  interface BaseContext {
    state: {
      context: GraphqlContext;
      id: string;
      log: bunyan;
    }
  }
}