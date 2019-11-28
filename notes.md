## Non-code configurations needed for setup
IAM user needs the following permissions:
  - `cloudformation:CreateStack` (I used `AWSCloudFormationFullAccess`)
  - `apigateway:POST` (I used `AmazonAPIGatewayAdministrator`)
  - `iam:CreateRole` (I used `IAMFullAccess`)
  - My user also has `AWSLambdaFullAccess`

## TODO
 - JWT protection
 - write to DynamoDB table
 - read from DynamoDB table
