import * as Koa from 'koa';

import container from './container';

export default async function main() {
  try {
    await container.load<Koa>('http');
    console.log('success!');
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
}

if (require.main === module) {
  main();
}
