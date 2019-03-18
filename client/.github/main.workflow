workflow "on pull request merge, delete the branch" {
  on = "pull_request"
  resolves = ["branch cleanup"]
}

action "branch cleanup" {
  uses = "jessfraz/branch-cleanup-action@master"
  secrets = ["GITHUB_TOKEN"]
}

workflow "Prepare production deploy" {
  on = "release"
  resolves = [
    "Prepare deploy Notification",
  ]
}

action "Docker Login" {
  uses = "actions/docker/login@c08a5fc9e0286844156fefff2c141072048141f6"
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Build" {
  uses = "inextensodigital/actions/make@master"
  secrets = ["COMPOSER_AUTH"]
  args = "actions-build"
  needs = ["Docker Login"]
}

action "Push" {
  uses = "inextensodigital/actions/make@master"
  args = "actions-push"
  needs = ["Build"]
}

action "Create deployment & update release" {
  uses = "inextensodigital/actions/deployment@master"
  needs = ["Push"]
  secrets = [
    "GITHUB_TOKEN",
    "PRIVATE_KEY",
  ]
  args = "create"
}

action "Prepare deploy Notification" {
  needs = "Create deployment & update release"
  uses = "apex/actions/slack@master"
  secrets = ["SLACK_WEBHOOK_URL"]
  env = {
    SLACK_CHANNEL = "#publishers"
  }
}

# after check !
workflow "Production deploy" {
  on = "deployment_status"
  resolves = ["Update deployment"]
}

action "Filter in progress deployment" {
  uses = "inextensodigital/actions/deployment@master"
  args = "filter deployment_status in_progress"
}

action "Deploy to prod" {
  uses = "inextensodigital/actions/make@master"
  args = "actions-deploy"
  needs = ["Filter in progress deployment"]
  secrets = ["COMPOSER_AUTH", "GITHUB_TOKEN", "K8S_AUTH_CERT", "K8S_SERVER", "K8S_CLIENT_CERT", "K8S_CLIENT_KEY", "K8S_USERNAME"]
  env = {
    DESIRED_VERSION = "v2.12.3"
  }
}

action "Update deployment" {
  uses = "inextensodigital/actions/deployment@master"
  needs = ["Deploy to prod"]
  args = "update_deployment success"
  secrets = ["GITHUB_TOKEN"]
}
