const { ref, refName, context } = require("./tools");
const _api = require("./api");

const octokit = context.octokit();
const api = _api(octokit);

export default async () => {
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
  const url = `https://deploy.emeabridge.eu/${owner}/${repo}/${deployment.id}`;
  await api.appendToReleaseBody(
    refName,
    `## Deploy to production

[![Deploy to prod](https://img.shields.io/badge/Deploy%20to-Production-orange.svg?style=for-the-badge)](${url})`
  );
};
