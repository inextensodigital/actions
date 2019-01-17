const octokit = require("@octokit/rest")();

const create = async () => {
  console.log(process.env);
  const [owner, repo] = process.env.GITHUB_REPOSITORY.split("/", 2);
  const ref = process.env.GITHUB_SHA;

  const deployment = await octokit.repos.createDeployment({
    owner,
    repo,
    ref,
    auto_merge: false
  });

  const result = await octokit.repos.createDeploymentStatus({
    owner,
    repo,
    deployment_id: deployment.id,
    state: "pending"
  });
};

if (process.argv[2] === "create") {
  create();
}
