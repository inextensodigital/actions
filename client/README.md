# GITHUB WORKFLOW


```

- [ ] github_workflow initialize
- [x] github_workflow workflow ls [name] --filter='on=="pull_request"'
- [x] github_workflow workflow create <name> <on> --action=
- [x] github_workflow workflow add <name> --action=
- [x] # github_workflow workflow rename <old> <new>
- [ ] # github_workflow workflow rm <name> --actions=true|false
- [ ] # github_workflow workflow merge <name> --filter=

- [x] github_workflow action ls [name] --filter=""
- [x] github_workflow action create <name> <uses> --env= # --secrets
- [x] # github_workflow action rename <old> <new>
- [x] # github_workflow action rm <name>
- [x] # github_workflow lint



action_name="Auto create master → dev PRs"
action_image="inextensodigital/actions/create-pull-request@master"

github_workflow initialize
github_workflow action ls --filter='uses=="$action_image"' | jq -e --slurp '.[0]' || \
github_workflow action create  "$action_name" "$action_image" --env="GITHUB_TOKEN"


if [ $workflow_name = $(github_worklow workflow ls --filter="on==pull_request" | jq --slurp '.[0].key') ] ; then
    github_worklow workflow add "$workflow_name" --action="Auto create master → dev PRs"
else
     github_workflow workflow create "On pull request" "pull_request" --action="$action_name"

fi
```
