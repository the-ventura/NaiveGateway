# CI/CD

## Design

The cicd pipeline is as simple as possible containiing only a build step and an upload step.
It can easily be expanded into testing the services and deployiing to a running environment.

## Technology

Github actions was chosen because it is simple, fast and free. It provides an easy and integrated way of running jobs.

## Requirements

Currently the cicd pipeline requires secrets to be set in the repo's secrets tab:
`AWS_ACCESS_KEY` and `AWS_SECRED_KEY_ID`. These keys should belong to a user with read and write permissions to AWS ECR.

## Current steps

### Build

Builds an docker image according to the Dockerfile.
The image gets tagged with the git sha and `latest`

### Upload

Uploads the built image to a private repo on AWS ECR
