const octokit = require("@octokit/rest")();

octokit.authenticate({
  type: "token",
  token: process.env.GITHUB_TOKEN
});
const create = async () => {
  const tag = "asimplerelease2"; //process.env.GITHUB_REF.split("/")[2];
  const owner = process.env.GITHUB_REPOSITORY.split("/", 1)[0];
  const repo = process.env.GITHUB_REPOSITORY.substring(owner.length + 1);
  const ref = process.env.GITHUB_SHA;

  // create deployment
  const { data: deployment } = await octokit.repos.createDeployment({
    owner,
    repo,
    ref,
    auto_merge: false,
    required_contexts: []
  });

  console.log(deployment);

  // update matching release
  try {
    const { data: release } = await octokit.repos.getReleaseByTag({
      owner,
      repo,
      tag
    });

    const result = await octokit.repos.updateRelease({
      owner,
      repo,
      release_id: release.id,
      body: `${release.body}
<!-- DEPLOY_BEGIN -->
## Deploy to production

[![Deploy to prod](https://img.shields.io/badge/Deploy%20to-Production-blue.svg?style=for-the-badge)](https://deploy.emeabridge.eu/${owner}/${repo}/${
        deployment.id
      }/${release.id})
<!-- DEPLOY_END -->`
    });

    console.log(result);
  } catch (e) {
    console.error(`Release with tag ${tag} not found for ${owner}/${repo}`);
  }
};
console.log(process.env);
if (process.argv[2] === "create") {
  create().then(null, (...args) => {
    console.error(args);
  });
}
