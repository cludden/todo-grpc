
export const inject = {
  name: 'graphql/utils',
  type: 'constructor',
};

export default class Utils {
  decodeGlobalId(globalId: string) {
    const decoded = Buffer.from(globalId, 'base64').toString('utf8').split('/');
    if (decoded.length !== 2) {
      throw new Error('invalid-global-id');
    }
    return { type: decoded[0], id: decoded[1] };
  }

  encodeGlobalId(type: string, id: string) {
    return Buffer.from(`${type}/${id}`, 'utf8').toString('base64');
  }
}
