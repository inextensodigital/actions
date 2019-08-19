const camelCase = require('camelcase')

function parseArgv(program) {
  program.parse(process.argv)
  program.options
    .filter(o => {
      return o.required
    })
    .forEach(option => {
      const key = camelCase(option.long, { pascalCase: false })
      if (!program[key]) {
        throw new Error(`Missing value for option ${option.flags}`)
      }
    })

  return program
}

module.exports = {
  parseArgv,
}
