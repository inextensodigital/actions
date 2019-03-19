# GITHUB WORKFLOW

Manage Github Action workflows and actions by cli. Allows you to script edition.

### Available commands:

- [x] github-workflow initialize
- [x] github-workflow lint
- [x] github-workflow workflow ls [ID][--on="pull_request"]
- [x] github-workflow workflow create ID ON [--resolve=<action>]
- [x] github-workflow workflow add ID --resolve=<action>
- [x] github-workflow workflow rename SOURCE TARGET
- [ ] github-workflow workflow rm NAME
- [ ] github-workflow workflow merge ID [--on=]

- [x] github-workflow action ls [ID]
- [x] github-workflow action create ID USE [--env=<env_name>=<env_value> --secret=<secret_name>]
- [x] github-workflow action rename SOURCE TARGET
- [x] github-workflow action rm ID

### Example for creating a new action on "pull_request"

```
action_name="Auto create master → dev PRs"
action_image="inextensodigital/actions/create-pull-request@master"

github-workflow initialize
github-workflow action ls "$action_image"
github-workflow action create "$action_name" "$action_image" --secret=GITHUB_TOKEN


if [ $workflow_name = $(github-workflow workflow ls --on="pull_request" | head -n 1) ] ; then
    github-workflow workflow add "$workflow_name" "$action_name"
else
     github-workflow workflow create "On pull request" "pull_request" --resolve="$action_name"
fi
```
