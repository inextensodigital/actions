const { context } = require("./tools");
const api = require("./api");

module.exports = async () => {
  const event = await context.readEvent();
  if (event.deployment) {
    await api.createDeploymentStatus(event.deployment.id, process.argv[3]);
  }
};
