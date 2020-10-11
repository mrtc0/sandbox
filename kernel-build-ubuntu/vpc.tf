module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  # version = "2.47.0"
  # https://github.com/terraform-aws-modules/terraform-aws-vpc/issues/267
  version = "~> 1.66.0"

  name = "test-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["ap-northeast-1a"]
  public_subnets  = ["10.0.10.0/24"]
  private_subnets = ["10.0.1.0/24"]

  map_public_ip_on_launch = true

  enable_nat_gateway = true
  single_nat_gateway = true
}
