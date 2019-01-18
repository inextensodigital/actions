const { context } = require("./tools");

module.exports = async () => {
  console.log(await context.readEvent());
  console.log(process.env);
};
