#!/bin/bash
version=$1
imageName=qaware.com/terraform/ikea

if [ "$1" = "" ]; then
  echo "Version number must not be empty. Please provide version number in format 1.2.3 as argument."
fi

docker buildx build --platform linux/amd64 -t $imageName:"$version" -t $imageName:latest --build-arg VERSION="$version" .
