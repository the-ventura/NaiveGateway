terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  region  = "eu-west-1"
}

# AWS ECR repository
resource "aws_ecr_repository" "naive_gateway_repo" {
  name = "joaoventura/naive-gateway"
}
