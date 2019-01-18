const func = require(`./${process.argv[2]}`);
try {
  func();
} catch (e) {
  console.error(e);
  process.exit(1);
}
