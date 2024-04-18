# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_iam_policy" "test" {
  name = var.rName

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
    }
  ]
}
EOF

  tags = {
    (var.tagKey1) = var.tagValue1
  }
}

data "aws_partition" "current" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}


variable "rName" {
  type     = string
  nullable = false
}

variable "tagKey1" {
  type     = string
  nullable = false
}

variable "tagValue1" {
  type     = string
  nullable = false
}


