# CI User
resource "aws_iam_user" "ci_user" {
  name = "ci"
}

# ECR
resource "aws_iam_group" "ecr_push_pull" {
  name = "ecr_push_pull"
}

resource "aws_iam_group_membership" "ecr_push_pull" {
  name = "ecr_push_pull_membership"

  users = [
    aws_iam_user.ci_user.name,
  ]

  group = aws_iam_group.ecr_push_pull.name
}


resource "aws_iam_policy" "ecr_push_pull" {
  name        = "ecr-push-pull"
  description = "Policy allowing push and pull to ecr"
  policy      =  <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PushPull",
            "Effect": "Allow",
            "Action": [
                "ecr:PutImageTagMutability",
                "ecr:StartImageScan",
                "ecr:ListTagsForResource",
                "ecr:UploadLayerPart",
                "ecr:BatchDeleteImage",
                "ecr:ListImages",
                "ecr:DeleteRepository",
                "ecr:CompleteLayerUpload",
                "ecr:TagResource",
                "ecr:DescribeRepositories",
                "ecr:DeleteRepositoryPolicy",
                "ecr:BatchCheckLayerAvailability",
                "ecr:ReplicateImage",
                "ecr:GetLifecyclePolicy",
                "ecr:PutLifecyclePolicy",
                "ecr:DescribeImageScanFindings",
                "ecr:GetLifecyclePolicyPreview",
                "ecr:CreateRepository",
                "ecr:PutImageScanningConfiguration",
                "ecr:GetDownloadUrlForLayer",
                "ecr:DeleteLifecyclePolicy",
                "ecr:PutImage",
                "ecr:UntagResource",
                "ecr:BatchGetImage",
                "ecr:DescribeImages",
                "ecr:StartLifecyclePolicyPreview",
                "ecr:InitiateLayerUpload",
                "ecr:GetRepositoryPolicy"
            ],
            "Resource": "${aws_ecr_repository.naive_gateway_repo.arn}"
        },
        {
            "Sid": "EcrAccess",
            "Effect": "Allow",
            "Action": [
                "ecr:GetRegistryPolicy",
                "ecr:DescribeRegistry",
                "ecr:GetAuthorizationToken",
                "ecr:DeleteRegistryPolicy",
                "ecr:PutRegistryPolicy",
                "ecr:PutReplicationConfiguration"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_group_policy_attachment" "ecr_push_pull" {
  group      = aws_iam_group.ecr_push_pull.name
  policy_arn = aws_iam_policy.ecr_push_pull.arn
}
