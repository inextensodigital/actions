const { context } = require("./tools");

export default async () => {
  console.log(await context.readEvent());
  console.log(process.env);
};
