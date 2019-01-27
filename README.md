![Build status](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoicGU4OGU1UmNNbXpJZ3QwYnBZSXpmSDVRUTJhUkh5TUN5UmtQWmRpQTZ0dUJVMzFvSXR3WU0yd3hxYnFrN0ltVUtSTTN4TmZja3lCaVhEa3dvOTl5U0VFPSIsIml2UGFyYW1ldGVyU3BlYyI6IkRQdUFadlk2cElJZlBaY0giLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)
# delex-code-example
Examples for writing AWS Lambda code for DelEx conference

### Setup
You need to install Go 1.11 or higher.    
On MacOS simply run `brew install go`. To install Go on other platforms check [official docs](https://golang.org/doc/install).

### Run tests
To run tests simply invoke `make`. It will run linters and unit tests.

### Run Lambda locally
1) Install [AWS SAM CLI](https://github.com/awslabs/aws-sam-cli)
2) Run:
```bash
sam local generate-event dynamodb update | sam local invoke
```