import * as Ajv from 'ajv';
import * as fs from 'fs';
import * as path from 'path';

export const inject = {
  name: 'validation',
};

const schemas = ['config'];

export default function () {
  const v = new Ajv({
    $data: true,
    coerceTypes: 'array',
    useDefaults: true,
    schemas: schemas.reduce((acc, s) => {
      const filename = path.join(__dirname, `../schemas/${s}.json`);
      Object.assign(acc, { [s]: JSON.parse(fs.readFileSync(filename, 'utf8')) });
      return acc;
    }, {}),
  });

  return v;
}
