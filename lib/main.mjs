import fs from 'fs/promises'
import { promisify } from 'util'
import './wasm_exec.js'

const WASM_PATH = './main.wasm'

const go = new Go()

const mod = await WebAssembly.compile(await fs.readFile(new URL(WASM_PATH, import.meta.url)))
const wasm = await WebAssembly.instantiate(mod, go.importObject)

go.run(wasm)

export const isAnimated = promisify(_isAnimated)
