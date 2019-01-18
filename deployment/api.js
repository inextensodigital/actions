const { owner, repo, ref } = require("./tools");

module.exports = kit => {
  const api = {
    createDeploymentFromRef: async add => {
      const { data: deployment } = await kit.repos.createDeployment({
        owner,
        repo,
        ref,
        ...add
      });

      return deployment;
    },
    getReleaseByTag: async (tag) => {
      const { data: release } = await kit.repos.getReleaseByTag({
        owner,
        repo,
        tag
      })

      return release
    }
    appendToReleaseBody: async (tag, contents, mark = 'DEPLOY') => {
      const release = await api.getReleaseByTag(tag)

      const result = await kit.repos.updateRelease({
        owner,
        repo,
        release_id: release.id,
        body: `${release.body}<!-- ${mark}_BEGIN -->
${contents}
<!-- ${mark}_END -->`
      });

      return true
    }
  };

  return api;
};
