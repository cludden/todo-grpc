import * as bunyan from 'bunyan';
import * as fs from 'fs';
import * as path from 'path';

import { Config } from './types';

export const inject = {
  name: 'log',
  require: ['config'],
};

// load package.json via fs and create default logger
const pkg = JSON.parse(fs.readFileSync(path.join(__dirname, '../package.json'), 'utf8'));
export const log = bunyan.createLogger({ name: pkg.name, version: pkg.version });

export default function (config: Config) {
  return log.child(config.log);
}
