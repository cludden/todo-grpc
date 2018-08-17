import * as Router from 'koa-router';

export const inject = {
  name: 'http/routes/health',
  require: [],
};

export default function () {
  // create new router
  const router = new Router();

  // register healthcheck endpoint
  router.get('/health', async (ctx) => {
    ctx.status = 200;
    ctx.body = { meta: { status: 'healthy' } };
  });

  return router;
}
