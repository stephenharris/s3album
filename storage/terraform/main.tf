data "aws_caller_identity" "current" {}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "3.58.0"
    }
  }
}

provider "aws" {
  region = var.region
}