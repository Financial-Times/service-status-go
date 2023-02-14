[![CircleCI](https://dl.circleci.com/status-badge/img/gh/Financial-Times/service-status-go/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/Financial-Times/service-status-go/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/Financial-Times/service-status-go/badge.svg?branch=master)](https://coveralls.io/github/Financial-Times/service-status-go?branch=master)

# service-status-go

This library provides a small set of reusable features that most of our microservices need.

## BuildInfo
*TODO fix passing of flags, doesn't seem to work at the moment !*

This depends on the -ldflags feature which allows you to set vars. As an example build this project could be:
```shell
    go get github.com/Financial-Times/service-status-go
    cd ${GOPATH}/src/github.com/Financial-Times/service-status-go
    flags="$(${GOPATH}/src/github.com/Financial-Times/service-status-go/buildinfo/ldFlags.sh)"
    go install -a -ldflags="${flags}"
    service-status-go | jq '.'
```
