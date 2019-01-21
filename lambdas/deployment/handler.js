/* eslint-disable global-require, import/no-dynamic-require */
const handler = {
  get: (target, name) => require(`./functions/${name}`)
};

module.exports = new Proxy({}, handler);
