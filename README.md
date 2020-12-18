# Movies 

This is a rest api for aws lambda and api gateway with dynamodb

##### Update Lambda Function Source
```sh
$ cd index
$ ./build.sh
$ aws lambda update-function-code \
    --function-name MoviesShow \
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