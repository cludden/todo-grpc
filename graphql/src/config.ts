import * as ajv from 'ajv';
import * as config from 'config';

import { Config } from './types';

export const inject = {
  name: 'config',
  require: ['validation'],
};

export default function (v: ajv.Ajv): Config {
  const parsed = config.util.toObject();
  if (!v.validate('config', parsed)) {
    throw new Error(`invalid configuration: ${v.errorsText(v.errors)}`);
  }
  return parsed as Config;
}
