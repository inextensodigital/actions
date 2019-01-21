const crypto = require("crypto");
const octokit = require("@octokit/rest")();

module.exports = async ({ queryStringParameters }, context, callback) => {
  if (queryStringParameters) {
    const { owner, repo, deploy, tag, sign } = queryStringParameters;
    if (owner && repo && deploy && sign) {
      const verify = crypto.createVerify("SHA256");
      verify.update(owner + repo + deploy + tag);

      if (verify.verify(process.env.PUBLIC_KEY, sign, "hex")) {
        octokit.authenticate({
          type: "token",
          token: process.env.GITHUB_TOKEN
        });

        const { data: stats } = await octokit.repos.createDeploymentStatus({
          owner,
          repo,
          deployment_id: deploy,
          state: "in_progress",
          headers: {
            accept: "application/vnd.github.flash-preview+json"
          }
        });

        console.log(stats);

        callback(null, {
          statusCode: 200,
          body: "🚀"
        });
        return;
      }
    }
  }

  callback(null, {
    statusCode: 400,
    body: `Missing params`
  });
};
