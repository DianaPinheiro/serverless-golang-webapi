# Bulding a Web API with AWS Lambda, Golang and Serverless Framework

In this project you will find an example of how to construct a Go applications - Web API to manage candidates for a job in recruitment process - following the Serverless approach using AWS Lambda, API Gateway, DynamoDB, and the Serverless Framework. 
AWS Lambda is the third compute service from Amazon. It's very different from the existing two compute services EC2 (Elastic Compute Cloud) and ECS (Elastic Container Service). AWS Lambda is an event-driven, serverless computing platform that executes your code in response to events. It manages the underlying infrastructure scaling it up or down to meet the event rate. You're only charged for the time your code is executed.

## Pre requisites
* AWS account
* Go installed on your machine
* AWS CLI and configure it

## Serverless Framework, what is it?

The Serverless Framework enables developers to deploy backend applications as independent functions that will be deployed to AWS Lambda. It also configures AWS Lambda to run your code in response to HTTP requests using Amazon API Gateway.

To install Serverless on your machine, run the below mentioned npm command.
```
$ npm install serverless -g
```

## Install and Run Web API

After clone the project, to gather your dependencies and build the proper binaries for your functions you should run the bellow command
```
$ make
```

After install the project and all dependencies you can deploy all functions as shown bellow
```
$ serverless deploy
```
If you have a custom profile, you must specify this profile on the command:
```
$ serverless deploy --aws-profile myCustomProfile
```

After run deploy command, you will see in the console the endpoints as shown bellow
```
POST - https://8k3me6szzf.execute-api.us-east-1.amazonaws.com/dev/candidates
GET - https://8k3me6szzf.execute-api.us-east-1.amazonaws.com/dev/candidates
GET - https://8k3me6szzf.execute-api.us-east-1.amazonaws.com/dev/candidates/{id}
DELETE - https://8k3me6szzf.execute-api.us-east-1.amazonaws.com/dev/candidates/{id}
```

## Why use Go for your Lambdas?

The combination of safety + speed is the big reason why people want so much to care about Golang. 
Golang is a compiled, statically-typed language. This can help catch simple errors and maintain correctness as your application grows. 
This safety is really useful for production environments.

However, we've had Java and C# support in Lambda for years. These are both compiled, static languages as well. What's the difference?

Java and C# both have notoriously slow cold-start time, in terms of multiple seconds. With Go, the cold-start time is much lower. In my haphazard testing, I was seeing cold-starts in the 200-400ms range, which is much closer to Python and Javascript.

Speed and safety. A pretty nice combo.

