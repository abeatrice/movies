# Movies 

This is a rest api for aws lambda and api gateway with dynamodb

##### Create Lambda Function
```sh
$ cd delete
$ ./build.sh
$ aws lambda create-function \
    --function-name MoviesDelete \
    --zip-file fileb://./deployment.zip \
    --runtime go1.x \
    --handler main \
    --role arn:aws:iam::803551335240:role/lambda-cloud-watch-full \
    --environment Variables={TABLE_NAME=movies} \
    --region us-east-1
```

##### Create API Gateway method
 - Get API_ID
```sh
$ aws apigateway get-rest-apis --query "items[?name==\`Movies\`].id" --output text
```
*ta6y8fu9j0*

 - Get Resource_ID
```sh
$ aws apigateway get-resources --rest-api-id API_ID --query "items[?path==\`/movies\`].id" --output text
```
*0f4dix*

 - Create API Gateway Method
```sh
$ aws apigateway put-method \
    --rest-api-id API_ID \
    --resource-id RESOURCE_ID \
    --http-method DELETE \
    --authorization-type "NONE" \
    --region us-east-1
```

 - Set Lambda Function as target of DELETE method in API Gateway
```sh
$ aws apigateway put-integration \
    --rest-api-id API_ID \
    --resource-id RESOURCE_ID \
    --http-method DELETE \
    --type AWS_PROXY \
    --integration-http-method DELETE \
    --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:803551335240:function:MoviesDelete/invocations \
    --region us-east-1
```

 - Set put method response in API Gateway
```sh
$ aws apigateway put-method-response \
    --rest-api-id API_ID \
    --resource-id RESOURCE_ID \
    --http-method DELETE \
    --status-code 200 \
    --response-models '{"application/json":"Empty"}' \
    --region us-east-1
```

 - Redeploy API
```sh
$ aws apigateway create-deployment \ 
    --rest-api-id API_ID \
    --stage-name staging \
    --region us-east-1
```

##### Update Lambda Function Source
```sh
$ cd index
$ ./build.sh
$ aws lambda update-function-code \
    --function-name MoviesIndex \
    --zip-file fileb://./deployment.zip \
    --region us-east-1
```

##### Create Env Var
```sh
$ aws lambda update-function-configuration \
    --function-name MoviesShow \
    --environment Variables={TABLE_NAME=movies} \
    --region us-east-1
```
