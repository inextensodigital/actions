const { owner, repo, ref, refName, context } = require("./tools");
const api = require("./api");

module.exports = async () => {
  const deploy = await api.createDeploymentFromRef({
    auto_merge: false,
    required_contexts: [],
    payload: JSON.stringify({
      ref,
      tag: refName
    }),
    description: `Production deploy for tag ${refName}`
  });

  await context.writeJSON("deployment", deploy);
  const url = `https://deploy.emeabridge.eu/${owner}/${repo}/${deploy.id}`;
  await api.appendToReleaseBody(
    refName,
    `## Deploy to production :rocket:

[![Deploy to prod](https://img.shields.io/badge/Deploy%20to-Production-orange.svg?style=for-the-badge)](${url})`
  );
};
