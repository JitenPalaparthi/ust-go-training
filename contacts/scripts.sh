#!/bin/sh

# To insert contact
curl --location --request POST 'localhost:58080/v1/contact/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Jiten",
    "email":"jitenp@outlook.Com",
    "contactNo":"9618558500",
    "address":"Yeshvantpur, Bangalore",
    "moreInfo":"Demo purpose"
}'

curl --location --request GET 'localhost:58080/v1/contact/get/1'