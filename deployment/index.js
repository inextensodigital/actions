const octokit = require("@octokit/rest")();

octokit.authenticate({
  type: "token",
  token: process.env.GITHUB_TOKEN
});
const create = async () => {
  const tag = process.env.GITHUB_REF.split("/")[2];
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

  const result = await octokit.repos.getReleaseByTag({ owner, repo, tag });
  console.log(result);
};
console.log(process.env);
if (process.argv[2] === "create") {
  create();
}
