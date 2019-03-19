# GITHUB WORKFLOW

Manage Github Action workflows and actions by cli. Allow you to script edition.

### Available commands:
- [x] github_workflow initialize
- [x] github_workflow lint

- [x] github_workflow workflow ls [name] --on="pull_request"
- [x] github_workflow workflow create <name> <on> <action>
- [x] github_workflow workflow add <name> <action>
- [x] github_workflow workflow rename <old> <new>
- [ ] github_workflow workflow rm <name> --actions=true|false
- [ ] github_workflow workflow merge <name> --filter=

- [x] github_workflow action ls <name> --filter=""
- [x] github_workflow action create <name> <uses> --env=foo=bar --env=bar=baz --secrets=FOO --secrets=BAR
- [x] github_workflow action rename <old> <new>
- [x] github_workflow action rm <name>


### Example for creating a new action on "pull_request"
```
action_name="Auto create master → dev PRs"
action_image="inextensodigital/actions/create-pull-request@master"

github_workflow initialize
github_workflow action ls "$action_image"
github_workflow action create "$action_name" "$action_image" "GITHUB_TOKEN"


if [ $workflow_name = $(github_worklow workflow ls --on="pull_request" ] ; then
    github_worklow workflow add "$workflow_name" "$action_name"
else
     github_workflow workflow create "On pull request" "pull_request" "$action_name"
fi
```

