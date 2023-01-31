# Contributing to the API

## Community Discussions:

If you have ideas that you would like to socialize/discuss with 
You can reach the maintainers of this repository at:
- Slack: [#topology-aware-scheduling](https://kubernetes.slack.com/archives/C012XSGFZQE)

## Submitting a PR to this repo

Feel free to create an issue/ submit a PR to this repo for community discussion.

## Verify autogenerate code in the repo is upto-date
Before submitting a PR, please run the following commands to verify that the auto-generated code is upto-date
```
    go mod vendor
    ./hack/verify-codegen.sh
```

## Updating auto-generated code

Run the following commands to regenerate auto-generated code:
```
    go mod tidy
    go mod vendor
    ./hack/update-codegen.sh
```

