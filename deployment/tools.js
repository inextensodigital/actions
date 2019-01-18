const octokit = require("@octokit/rest");
const fs = require("fs");
const { promisify } = require("util");

const writeFile = promisify(fs.writeFile);
const readFile = promisify(fs.readFile);

const owner = process.env.GITHUB_REPOSITORY.split("/", 1)[0];
const repo = process.env.GITHUB_REPOSITORY.substring(owner.length + 1);
const ref = process.env.GITHUB_REF;
const refName = process.env.GITHUB_REF.split("/")[2];
const token = process.env.GITHUB_TOKEN;
const eventPath = process.env.GITHUB_EVENT_PATH;
const home = process.env.HOME;

module.exports = {
  owner,
  repo,
  ref,
  refName,
  token,
  context: {
    repo: add => ({
      repo,
      owner,
      ...add
    }),
    octokit: config => {
      const client = octokit({ ...config });
      client.authenticate({
        type: "token",
        token
      });

      return client;
    },
    writeJSON: (name, obj) =>
      writeFile(`${home}/${name}.json`, JSON.stringify(obj)),
    readJSON: async name =>
      JSON.parse(await readFile(`${home}/${name}.json`, "utf8")),
    readEvent: async () => JSON.parse(await readFile(eventPath, "utf8")),
    slackMessage: obj => writeFile(`./slack.json`, JSON.stringify(obj))
  }
};
