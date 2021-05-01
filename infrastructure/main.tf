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

resource "aws_ecr_lifecycle_policy" "naive_gateway_repo" {
  repository = aws_ecr_repository.naive_gateway_repo.name

  policy = <<EOF
{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Only have a version history of 5 images",
            "selection": {
                "tagStatus": "any",
                "countType": "imageCountMoreThan",
                "countNumber": 5
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
EOF
}
