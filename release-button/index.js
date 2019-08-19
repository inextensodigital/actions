/* eslint-disable global-require, import/no-dynamic-require */
const func = require(`./src/${process.argv[2]}`)

try {
  func()
} catch (e) {
  console.error(e)
  process.exit(1)
}
