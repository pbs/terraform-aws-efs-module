data "aws_vpcs" "vpc" {}

resource "aws_security_group" "sg" {
  vpc_id = data.aws_vpcs.vpc.ids[0]
}

module "efs" {
  source = "../.."

  security_groups = [aws_security_group.sg.id]

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
