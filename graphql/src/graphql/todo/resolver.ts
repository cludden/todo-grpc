import { pipe, prop } from 'ramda';

import Utils from '../../graphql/utils';
import { SparkAttributes } from '../../repositories/todo_pb';

export const inject = {
  name: 'graphql/spark/resolver',
  require: ['graphql/utils', 'sparks'],
};

export default function (utils: Utils, sparks: SparksService) {
  return {
    Query: {
      async todos(_: any, args: { input?: ListSparksInput }) {
        return sparks.listSparks(args.input || {});
      },
    },

    Todo: {
      id: pipe<any, any, string>(
        prop('id'),
        utils.encodeGlobalId.bind(utils, 'todo'),
      ),
      cover: (spark: SparkAttributes) => spark.cover ? { id: spark.cover } : null,
      createdAt: prop('created_at'),
    },
  };
}


