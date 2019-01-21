const crypto = require("crypto");

const { owner, repo, ref, refName, context } = require("./tools");
const api = require("./api");

const encodeData = data => {
  return Object.keys(data)
    .map(function(key) {
      return [key, data[key]].map(encodeURIComponent).join("=");
    })
    .join("&");
};

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

  const sign = crypto.createSign("RSA-SHA256");
  sign.update(owner + repo + deploy.id + refName);

  const url = `https://auto-deploy.inextenso.io/deploy?${encodeData({
    owner,
    repo,
    deploy: deploy.id,
    tag: refName,
    sign: sign.sign(process.env.PRIVATE_KEY, "hex")
  })}`;
  await api.appendToReleaseBody(
    refName,
    `## Deploy to production :rocket:

[![Deploy to prod](https://img.shields.io/badge/Deploy%20to-Production-orange.svg?style=for-the-badge)](${url})`
  );
};
