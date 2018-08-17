import * as faker from 'faker';
import { GraphQLResolveInfo } from 'graphql';
import * as Koa from 'koa';

export const inject = {
  name: 'graphql/date-time/mock',
};

export default {
  DateTime(obj: any, _: { [key: string]: any }, __: Koa.Context, info: GraphQLResolveInfo) {
    if (Object.prototype.hasOwnProperty.call(info, 'fieldName')) {
      if (Object.prototype.hasOwnProperty.call(obj, info.fieldName)) {
        if (obj[info.fieldName] instanceof Date) {
          return obj[info.fieldName].toISOString();
        }
        return obj[info.fieldName];
      }
    }
    return faker.date.past();
  },
};
