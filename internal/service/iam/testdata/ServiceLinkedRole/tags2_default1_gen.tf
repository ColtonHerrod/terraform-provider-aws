# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "aws" {
  default_tags {
    tags = {
      (var.providerTagKey1) = var.providerTagValue1
    }
  }
}

resource "aws_iam_service_linked_role" "test" {
  aws_service_name = "autoscaling.amazonaws.com"
  custom_suffix    = var.rName

  tags = {
    (var.tagKey1) = var.tagValue1
    (var.tagKey2) = var.tagValue2
  }
}


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

variable "tagKey2" {
  type     = string
  nullable = false
}

variable "tagValue2" {
  type     = string
  nullable = false
}


variable "providerTagKey1" {
  type     = string
  nullable = false
}

variable "providerTagValue1" {
  type     = string
  nullable = false
}
