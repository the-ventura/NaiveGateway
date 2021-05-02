# Infrastructure

## Description

All required infrastructure is described as code as terraform files in the `infrastructure` directory.

It assumes you have an AWS account and sufficient priviledges to create infrastructure.

State is stored locally for convenience but it really should be stored remotely if working in a team to avoid state corruption.

Running `terraform apply` on a blank state will create:

* An ecr repository which only allows 5 images at a time, deleting the oldest ones.
* A user group with full permissions to the ecr repository
* A user called `ci` belonging to the group.
