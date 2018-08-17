/// <reference path="./container.d.ts" />
import * as Container from 'app-container';

const container = new Container({
  namespace: 'inject',
  defaults: {
    singleton: true,
  },
});

container.glob('**/*.js', {
  cwd: __dirname,
  ignore: ['container.js', 'index.js'],
});

export default container;
