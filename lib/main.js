const fs = require('fs/promises')
const path = require('path')
const { promisify } = require('util')
require('./wasm_exec')

;(async () => {
  const go = new Go()
  const bytes = new Uint8Array(await fs.readFile(path.resolve(__dirname, './main.wasm')));
  const wasm = new WebAssembly.Instance(new WebAssembly.Module(bytes), go.importObject);

  go.run(wasm)

  const p = promisify(_p)
  // const transfer = promisify(_transfer)
  const isAnimated = promisify(_isAnimated)

  await p('aaaaaa')
  console.log('done aaaaa')

  const file = await fs.readFile(path.resolve(__dirname, './animated.gif'))
  const isA = await isAnimated(file)
  console.log('isA:', isA)

  console.log('done all')
})()
