provider "aws" {
  region  = "us-west-1"
  profile = "aws_account_profile"
}

locals {
  name   = "example-ec2-complete"
  region = "us-west-1"

  user_data = <<-EOT
  #!/bin/bash
  echo "Hello Terraform!"
  EOT

  tags = {
    Owner       = "user"
    Environment = "dev"
  }
}


data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-hvm-*-x86_64-gp2"]
  }
}

resource "aws_instance" "web" {
  ami           = "${data.aws_ami.amazon_linux.id}"
  instance_type = "t2.micro"

  tags = {
    Name = "terraform-ec2"
  }
  availability_zone           = element(module.vpc.azs, 0)
  subnet_id                   = element(module.vpc.private_subnets, 0)
  vpc_security_group_ids      = [module.security_group.security_group_id]



}

module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = local.name
  description = "Security group for example usage with EC2 instance"
  vpc_id      = module.vpc.vpc_id

  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["http-80-tcp", "all-icmp"]
  egress_rules        = ["all-all"]

  tags = local.tags
}
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = local.name
  cidr = "10.99.0.0/18"

  azs              = ["${local.region}a", "${local.region}b"]
  public_subnets   = ["10.99.0.0/24", "10.99.1.0/24"]
  private_subnets  = ["10.99.3.0/24", "10.99.4.0/24"]
  database_subnets = ["10.99.7.0/24", "10.99.8.0/24"]

  tags = local.tags
}



module "cluster" {
  source  = "terraform-aws-modules/rds-aurora/aws"

  name           = "test-aurora-db-postgres96"
  engine         = "aurora-postgresql"
  engine_version = "11.12"
  instance_class = "db.t2.micro"
  instances = {
    1 = {
      instance_class = "db.t2.micro"
    }
  }

  vpc_id  = module.vpc.vpc_id
  subnets = [module.security_group.security_group_id]

#   allowed_security_groups = [module.security_group.security_group_id]
  allowed_cidr_blocks     = ["10.20.0.0/20"]

  storage_encrypted   = true
  apply_immediately   = true
  monitoring_interval = 10
  db_parameter_group_name         = "default"
  db_cluster_parameter_group_name = "default"
  enabled_cloudwatch_logs_exports = ["postgresql"]
  tags = {
    Environment = "dev"
    Terraform   = "true"
  }
}