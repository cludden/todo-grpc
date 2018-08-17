import * as chance from 'chance';
import * as faker from 'faker';
import { MockList } from 'graphql-tools';
import { get } from 'lodash';

export const inject = {
  name: 'graphql/mock-utils',
};

const DEFAULT_PAGE_SIZE = 20;

export const connection = (xmin = 1, max = 50) =>
  (_: any, args: { input: { first: number, after: number } }) => {
    const total = faker.random.number({ max, min: xmin });
    const first = get(args, 'first') || get(args, 'input.first') || DEFAULT_PAGE_SIZE;
    const min = total < first ? total : first;
    return {
      edges: () => new MockList(min),
      pageInfo: {
        total,
        cursor: chance.Chance().apple_token(),
        hasNextPage: min === total || min < first ? false : true,
      },
    };
  };

export const edge = () => () => ({ cursor: chance.Chance().apple_token() });

