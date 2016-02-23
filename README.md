# service-status-go

This library provides a small set of reusable features that most of our microservices need.

## BuildInfo
*TODO fix passing of flags, doesn't seem to work at the moment !*

This depends on the -ldflags feature which allows you to set vars. As an example build this project as follows:
```
go get github.com/Financial-Times/service-status-go
cd  github.com/Financial-Times/service-status-go
flags="-X github.com/Financial-Times/service-status-go/buildinfo/version=0.0.1 -X github.com/Financial-Times/service-status-go/buildinfo/repository=https://github.com/Financial-Times/service-status-go.git -X github.com/Financial-Times/service-status-go/buildinfo/commit=`git rev-parse HEAD` -X github.com/Financial-Times/service-status-go/buildinfo/builder=`go version` -X github.com/Financial-Times/service-status-go/buildinfo/dateTime=`date +%Y%m%d%H%M%S`"
go install -ldflags $flags
service-status-go
```
