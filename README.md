# GO Transfer
GO Transfer is a simplified banking back-end API that create users account, authenticate the user and make transactions between accounts.

This API is meant only to be an ability assessment of my skills as a developer, it has no meaningful use or contribution.

## Features
This API has three Endpoints as follow

+ /accounts

    + Create user account
    + Return the list of all accounts in the database
    + Return the balance for a specific user

+ /login
    + Authenticate the user and return a token
    + Support Json Web Tokens

+ /transfer
    + Allows transactions between accounts (Requires authentication)
    + List all the transaction for a specific user (Requires authentication)

## Requirements
	github.com/go-playground/validator v9.31.0
	github.com/golang-jwt/jwt v3.2.2
	github.com/labstack/echo/v4 v4.9.0
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/stretchr/testify v1.7.0

# Quickstart

## Docker

```
docker pull geronaso/go-transfer
docker build --tag geronaso/go-transfer . 
docker run -p 1323:1323 -it geronaso/go-transfer
```

## Local

```
git clone https://github.com/Geronaso/Go-transfer.git
cd go-transfer
go mod init
go mod tidy
go run main.go
```

## Google Cloud Provider

https://go-transfer-geronaso-bcaeeczeha-uc.a.run.app

I made a CI/CD pipeline that run the written tests and deploy

# Database

The database.db file is pre loaded with two tables, account and transfers. These tables are also pre loaded with some values for easy testing:

id | user | cpf | balance | date_joined | password
--- | --- | --- | --- | --- | ---
1 | Paulo | 12345678905 | 1200 | 2022-10-06T01:06:53-03:00 | hash(123456)
2 | Carina | 12345678903 | 1300 | 2022-10-06T01:06:55-03:00 | hash(123456)
3 | Juliana | 12345678904 | 500 | 2022-10-06T01:07:29-03:00 | hash(123456)


id | account_origin_id | account_destination_id | amount | created_at
--- | --- | --- | --- | ---
1 | 3 | 2 | 300.0 | 2022-10-06T01:08:23-03:00
2 | 3 | 2 | 200.0 | 2022-10-06T01:08:35-03:00



# Endpoints


## /accounts

+ Request (application/json)
+ Method [POST]

    + Body
```    
        {
            "name": "Carina",
            "cpf": "12345678903",
            "secret": "123456"
        }
```       

+ Response 200 (application/json)
    + Body
```    
        {
            "Account Created"
        }
```

## /accounts
+ Request (application/json)
+ Method [GET]

+ Response 200 (application/json)
    + Body
```    
        [
            {
                "Id": 1,
                "Name": "Paulo",
                "Cpf": "12345678905",
                "Balance": 1200,
                "Date_Joined": "2022-10-06T01:06:53-03:00"
            },
            {
                "Id": 2,
                "Name": "Carina",
                "Cpf": "12345678903",
                "Balance": 1300,
                "Date_Joined": "2022-10-06T01:06:55-03:00"
            },
            {
                "Id": 3,
                "Name": "Juliana",
                "Cpf": "12345678904",
                "Balance": 500,
                "Date_Joined": "2022-10-06T01:07:29-03:00"
            }
        ]
```   

## /accounts/:id/balance


+ Request (application/json)
+ Method [GET]

+ Response 200 (application/json)
    + Body
```    
        {
            "1200"
        }
```   

## /login

+ Request (application/json)
    + Method [POST]

    + Body
```    
        {
            "cpf": "12345678903",
            "secret": "123456"
        }
```       

+ Response 200 (application/json)
    + Body
```    
        {
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiIxMjM0NTY3ODkwNCIsImV4cCI6MTY2NTExNjU3Nn0.nyHJ-sOm8vk3fusByDgw1opeMlu2zA8NXFdbvbbDe0Y"
        }
```
### JWT

The token provided by the API expires after 24 hours.

## /transfer

+ Request (application/json)
    + Method [GET]

    + Headers
```
            Authorization: Bearer [access_token]
```
+ Response 200 (application/json)
    + Body
```    
        [
            {
                "Id": 1,
                "Account_origin_id": 3,
                "Account_destination_id": 2,
                "Amount": 300,
                "Created_at": "2022-10-06T01:08:23-03:00"
            },
            {
                "Id": 2,
                "Account_origin_id": 3,
                "Account_destination_id": 1,
                "Amount": 200,
                "Created_at": "2022-10-06T01:08:35-03:00"
            }
        ]
```

## /transfer

+ Request (application/json)
    + Method [POST]

    + Headers
``` 
            Authorization: Bearer [access_token]
```
    + Body
```    
        {
            "destination_cpf": "12345678902",
            "amount": 700
        }
```       

+ Response 200 (application/json)
    + Body
```    
        {
            "Transfer Completed"
        }
```



## Testing Endpoints

[POSTMAN](https://www.postman.com/foxmc/workspace/go-transfer)

I have made a workspace at postman for easy testing.


The req folder also has some http templates to do some easy testing, but you can also use your preferred tool to test the endpoints!

# Tests

The folder tests has some written tests for the endpoint /accounts. The database.db localized in the tests folder is initialized with empty tables for testings, after running the tests it is necessary to re-initialize the table with the init.sql script. So if you want to run the tests again, please run the script before testing.

```
go test .\tests\
```

# Pipeline

The pipeline is ilustrated in the GCP-Deploy.yml file at .github/workflows

It runs the following steps on push.

```mermaid
graph TD;
    Authenticate GCP-->Docker Build;
    Docker Build-->Go test;
    Go test --> GCP Deploy;

```

Authenticate GCP authenticates at my account at GCP.

Docker Build makes the docker file for deployment.

Go test tests the docker image with written tests.

GCP Deploy finally deploy the application to production at GCP.


## TO-DO

* Write better error messages.
    * It is necessary to write better error messages specially for database related errors, so it becomes easier to debug and also to be more of a RESTful API.

* Write more tests
    * Currently only the endpoint /accounts has written tests, it is necessary to write tests for the others endpoints as well.


## Author
Cézar Murilo (cezarmgt@gmail.com)

## License

GNU General Public License v3.0 or later.

See [LICENSE](LICENSE) for the full text.
