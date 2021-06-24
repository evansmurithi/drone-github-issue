# drone-github-issue

[![Build Status](https://cloud.drone.io/api/badges/evansmurithi/drone-github-issue/status.svg?ref=refs/heads/main)](https://cloud.drone.io/evansmurithi/drone-github-issue)

Drone plugin to create GitHub Issues.

## Build

With `Go` installed:

```sh
go get -u -v github.com/evansmurithi/drone-github-issue
```

or build the binary with the following command:

```sh
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-github-issue ./cmd/drone-github-issue
```

## Docker

Build the Docker image with the following command:

```sh
docker build \
    --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
    --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
    --file docker/Dockerfile.linux.amd64 --tag evansmurithi/drone-github-issue .
```

## Usage

```sh
docker run --rm \
    -e DRONE_BUILD_EVENT=tag \
    -e DRONE_REPO_OWNER=octocat \
    -e DRONE_REPO_NAME=foo \
    -e DRONE_COMMIT_REF=refs/heads/main \
    -e GITHUB_ISSUE_API_KEY=xxxx \
    -e GITHUB_ISSUE_TITLE="issue title" \
    -e GITHUB_ISSUE_BODY="issue body" \
    -e GITHUB_ISSUE_ASSIGNEES="user1,user2" \
    -e GITHUB_ISSUE_LABELS="bug" \
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    evansmurithi/drone-github-issue
```
