
declare module 'app-container' {
  import * as glob from 'glob'

  interface Options {
    namespace: string,
    defaults?: {
      singleton?: Boolean,
      type?: string,
    }
  }

  class Container {
    constructor(options: Options)

    glob(pattern: string, options: glob.IOptions): void

    load<T>(deps: string | string[] | { [key: string]: string }): Promise<T>
  }

  var appcontainer: {
    new(options: Options): Container
  }

  export = appcontainer
}

