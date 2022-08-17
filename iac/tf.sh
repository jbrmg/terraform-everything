#!/bin/bash

# try to find out provider version to use
# the provider version must be equal to the tag of the docker-image to use
# reads all lines until the script finds the provider declaration
# then, the script reads the next line which will contain the version number
while IFS= read -r line; do
  if [[ $found ]]; then
    # get version number from `version = 0.0.2`
    split=(${line//=/ })
    # remove double quotes from string
    versionNumber=${split[1]//\"/}
    break
  fi
  if [[ $line == *"qaware.com/terraform/ikea"* ]]; then
    found=true
  fi
done <main.tf

# remember the current directory name and mount the parent directory as volume into the container
# then, set the workspace directory to the original directory.
# mounting the parent directory into the container is necessary so that iac-modules can be referenced locally
# as they are mounted into the container too
docker run --rm -it -v "$(pwd)":/workspace -w /workspace qaware.com/terraform/ikea:"$versionNumber" "$@"
