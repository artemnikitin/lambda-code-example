provider "aws" {
  region     = "us-east-1"
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "permissions" {
  name = "policy_for_lambda"
  role = "${aws_iam_role.iam_for_lambda.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "xray:PutTraceSegments",
        "xray:PutTelemetryRecords",
        "dynamodb:GetRecords",
        "dynamodb:GetShardIterator",
        "dynamodb:DescribeStream",
        "dynamodb:ListStreams",
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_lambda_function" "lambda_deployment" {
  function_name = "lambda_deployment_terraform"
  description   = "Example of Lambda deployment with Terraform"

  filename         = "build/lambda_event.zip"
  source_code_hash = "${base64sha256(file("build/lambda_event.zip"))}"

  role    = "${aws_iam_role.iam_for_lambda.arn}"
  handler = "lambda_event"
  runtime = "go1.x"
  timeout = 300

  tracing_config {
    mode = "Active"
  }
}

resource "aws_lambda_event_source_mapping" "event_source_mapping" {
  batch_size        = 100
  event_source_arn  = "xyz" # put here actual ARN of existed DynamoDB stream
  enabled           = true
  function_name     = "${aws_lambda_function.lambda_deployment.function_name}"
  starting_position = "TRIM_HORIZON"
}

resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_deployment.function_name}"
  retention_in_days = 7
}
