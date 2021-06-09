# node-midec
Multi-image detector (Wasm). cf. Animated GIF, APNG, Animated WebP, Animated HEIF/AVIF.

Pure Webassembly + JavaScript port of [midec](https://github.com/sapphi-red/midec).

## Installation
```shell
$ npm i node-midec
```

## Example
```js
import { isAnimated } from 'node-midec'
import fs from 'fs/promises'

const filePath = new URL('./animated.gif', import.meta.url)
const file = await fs.readFile(filePath)
const result = await isAnimated(file)
console.log(`animated.gif is ${result ? '' : 'not '}an animated image.`)
```

## Build
```shell
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./lib
$ GOOS=js GOARCH=wasm go build -o lib/main.wasm
```
