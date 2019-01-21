const { readFileSync } = require("fs");
const template = readFileSync("./functions/deploy/confirm.html", "utf8");
const encodeData = data => {
  return Object.keys(data)
    .map(function(key) {
      return [key, data[key]].map(encodeURIComponent).join("=");
    })
    .join("&");
};

module.exports = (
  { queryStringParameters, requestContext: { domainName, path } },
  context,
  callback
) => {
  console.log(
    JSON.stringify(queryStringParameters),
    { domainName, path },
    context,
    callback
  );
  const url = `https://${domainName}${path.replace(
    "deploy",
    "confirm"
  )}?${encodeData(queryStringParameters)}`;

  console.log(url);

  callback(null, {
    statusCode: 200,
    headers: {
      "Content-Type": "text/html; charset=utf-8"
    },
    body: template
      .replace("__LINK__", url)
      .replace(
        "__BACK_LINK__",
        `https://github.com/${queryStringParameters.owner}/${
          queryStringParameters.repo
        }/releases/tag/${queryStringParameters.tag}`
      )
      .replace(
        "__APP__",
        `${queryStringParameters.owner}/${queryStringParameters.repo}`
      )
  });
  return;
};
