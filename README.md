# Go LDAP Webservice

## General Information
POST http://host:44444/api/ldap-login 

Request:
{
    "username": "username",
    "password": "password
}"

Response Success:
{
    "result": "OK",
    "user_dn": "cn=user,ou=dept,ou=city,o=company"
}

Response Fail:
{
    "result": "Invalid username or password (n)"
}

## Technologies Used
- Go 1.19

## Setup
1. Follow setup folder instructions
2. Use Postman collection for testing