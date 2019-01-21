const crypto = require("crypto");
const octokit = require("@octokit/rest")();
const { readFileSync } = require("fs");
const template = readFileSync("./functions/confirm/ok.html", "utf8");

module.exports = async ({ queryStringParameters }, context, callback) => {
  if (queryStringParameters) {
    const { owner, repo, deploy, tag, sign } = queryStringParameters;
    if (owner && repo && deploy && sign && tag) {
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
          headers: {
            "Content-Type": "text/html; charset=utf-8"
          },
          body: template.replace(
            "__LINK__",
            `https://github.com/${owner}/${repo}/deployments`
          )
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
