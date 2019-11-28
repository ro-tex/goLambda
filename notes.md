## Non-code configurations needed for setup
IAM user needs the following permissions:
  - `cloudformation:CreateStack` (I used `AWSCloudFormationFullAccess`)
  - `apigateway:POST` (I used `AmazonAPIGatewayAdministrator`)
  - `iam:CreateRole` (I used `IAMFullAccess`)
  - My user also has `AWSLambdaFullAccess`

## TODO
 - how to reuse code?
 - JWT protection
 - write to DynamoDB table
 - read from DynamoDB table

## Examples
https://github.com/serverless/examples

## Notes

Invoke locally:
`sls invoke local -f functionName`
