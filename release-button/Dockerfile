FROM node:10

LABEL "com.github.actions.name"="Github Deployment"
LABEL "com.github.actions.description"="Create deployment button"
LABEL "com.github.actions.icon"="upload-cloud"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/inextensodigital/actions"
LABEL "homepage"="http://github.com/inextensodigital"
LABEL "maintainer"="IED <contact@inextenso.digital>"

RUN mkdir -p /app
ADD package.json /app/package.json
RUN cd /app && yarn
ADD . /app

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
