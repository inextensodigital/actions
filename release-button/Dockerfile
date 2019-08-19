FROM node:10-alpine

WORKDIR /app
ADD package.json /app/package.json
RUN yarn
ADD . .

ENTRYPOINT ["./deploy_button"]
