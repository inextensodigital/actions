const { context } = require("./tools");

module.exports = async () => {
  const type = process.argv[3];
  const value = process.argv[4];
  let pass = false;

  switch (type) {
    case "deployment_status":
      const event = await context.readEvent();
      if (event.deployment_status) {
        pass = event.deployment_status.state === value;
      }
      break;
  }

  process.exit(pass ? 0 : 78);
};
