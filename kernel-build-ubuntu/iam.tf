data "aws_iam_policy_document" "ec2_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

data "aws_iam_policy" "systems_manager" {
  arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_iam_role" "role" {
  name               = "ec2-role"
  assume_role_policy = "${data.aws_iam_policy_document.ec2_assume_role.json}"
}

resource "aws_iam_role_policy_attachment" "default" {
  role       = "${aws_iam_role.role.name}"
  policy_arn = "${data.aws_iam_policy.systems_manager.arn}"
}

resource "aws_iam_instance_profile" "systems_manager" {
  name = "InstanceProfile"
  role = "${aws_iam_role.role.name}"
}
