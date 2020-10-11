resource "aws_instance" "private_instance" {
  ami                  = "ami-01c36f3329957b16a"                            // ubuntyu 18.04 ap-northeast-1 hvm:ebs-ssd
  instance_type        = "t2.micro"
  subnet_id            = "${module.vpc.private_subnets[0]}"
  iam_instance_profile = "${aws_iam_instance_profile.systems_manager.name}"

  tags {
    Name = "private_instance"
  }
}

output "instance id" {
  value = "${aws_instance.private_instance.id}"
}
