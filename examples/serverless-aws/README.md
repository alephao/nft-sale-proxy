# Serverless AWS Lambda

This example deploys [alephao/nft-sale-proxy](https://github.com/alephao/nft-sale-proxy) to [AWS Lambda](https://aws.amazon.com/lambda/) using the [Serverless framework](https://www.serverless.com).

You can see the configuration for this proxy on [serverless.yml](serverless.yml), under the `environment` key. It is simulating the BAYC sale, but only showing the first 1000 BAYCs and returning the image of BAYC 0 as the incognito image.

### Deploying this example

1. Create an AWS account and get the aws keys following the [instructions here](https://www.serverless.com/framework/docs/providers/aws/guide/credentials)
2. Install dependencies `go mod tidy && npm install`
3. `cd serverless-aws && make deploy`

### CI/CD

There is an example configuration for deploying the proxy on every push to master using GitHub Actions: [.github/workflows/deploy.yml].