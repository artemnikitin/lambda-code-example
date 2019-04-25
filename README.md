![Build status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoicGU4OGU1UmNNbXpJZ3QwYnBZSXpmSDVRUTJhUkh5TUN5UmtQWmRpQTZ0dUJVMzFvSXR3WU0yd3hxYnFrN0ltVUtSTTN4TmZja3lCaVhEa3dvOTl5U0VFPSIsIml2UGFyYW1ldGVyU3BlYyI6IkRQdUFadlk2cElJZlBaY0giLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)
# lambda-code-example
Examples of AWS Lambda code in Go

### Setup
You need to install Go 1.11 or higher.    
On MacOS simply run `brew install go`. To install Go on other platforms check [official docs](https://golang.org/doc/install).

### Run tests
To run tests simply invoke `make`. It will run linters and unit tests.

### Run Lambda locally
1) Install [AWS SAM CLI](https://github.com/awslabs/aws-sam-cli)
2) Run it:     

via SAM CLI
```bash
make build 
sam local generate-event dynamodb update | sam local invoke
```
or via make: 
```bash
make run-local
```

### Deploy Lambda with Terraform
1) Install [Terraform](https://learn.hashicorp.com/terraform/getting-started/install.html)
2) Run it:

via command line
```bash
make build & terraform init & terraform apply -auto-approve
```
or via make: 
```bash
make tf-deploy
```

### Deploy Lambda with AWS SAM CLI
1) Install [AWS SAM CLI](https://github.com/awslabs/aws-sam-cli)
2) Run it:     

via SAM CLI
```bash
make build

sam package \
    --template-file template.yaml \
    --output-template-file serverless-output.yaml \
    --s3-bucket your-s3-bucket # put name of your S3 bucket here
    
sam deploy \
    --template-file serverless-output.yaml \
    --stack-name my-lambda-deployment \
    --capabilities CAPABILITY_IAM
```
or via make: 
```bash
make sam-deploy
```
