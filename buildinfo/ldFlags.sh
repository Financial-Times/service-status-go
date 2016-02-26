#!/bin/sh
# The following is a sample of how to create the flags that are passed to `go install`
# see the top level README.md for details

package="github.com/Financial-Times/service-status-go/buildinfo."
version="version=$(git describe --tag 2> /dev/null)"
dateTime="dateTime=$(date -u +%Y%m%d%H%M%S)"
repository="repository=$(git config --get remote.origin.url)"
revision="revision=$(git rev-parse HEAD)"
builder="builder=$(go version)"

for flag in "$version" "$dateTime" "$repository" "$revision" "$builder"
do
  set -- $flag
  ldflag="-X '"${package}${flag}"'"
  flags="$flags $ldflag"
done

echo $flags
exit 0
