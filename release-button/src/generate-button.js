const crypto = require('crypto')
const stringify = require('querystring')
const program = require('commander')
const { parseArgv } = require('./utils')

function generateButton(deployId, repository, ref) {
  const { PRIVATE_KEY } = process.env

  if (!PRIVATE_KEY) {
    console.log(
      'Environment variable PRIVATE_KEY is required to sign de release',
    )
    process.exit(1)
  }

  const [owner, repo] = repository.split('/')
  const [, , refName] = ref.split('/')

  const sign = crypto.createSign('RSA-SHA256')

  const query = stringify({
    owner,
    repo,
    deploy: deployId,
    tag: refName,
    sign: sign.sign(PRIVATE_KEY, 'hex'),
  })

  const url = `https://auto-deploy.inextenso.io/deploy?${query}`
  const img =
    'https://img.shields.io/badge/Deploy%20to-Production-orange.svg?style=for-the-badge'

  return `[![Deploy to prod](${img})](${url})`
}

program
  .option('-d, --deploy-id <id>', 'deploy id')
  .option(
    '-g, --repository <github repository>',
    'the github repository (owner/repository)',
  )
  .option('-r, --ref <github reference>', 'the github ref')

module.exports = function() {
  try {
    const { deployId, repository, ref } = parseArgv(program)
    generateButton(deployId, repository, ref)
  } catch (error) {
    console.log(`${error} \n`)
    program.outputHelp()
  }
}
