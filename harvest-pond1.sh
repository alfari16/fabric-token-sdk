#!/bin/bash

curl -X POST http://vmi836509.contaboserver.net:9100/api/v1/issuer/issue -H 'Content-Type: application/json' -d '{
    "amount": {"value": 1000, "code":"IDR"},
    "counterparty": {"node": "owner1","account": "bob"},
    "message": "harvest pond 1"    
}'

curl -X POST http://vmi836509.contaboserver.net:9200/api/v1/owner/accounts/bob/redeem -H 'Content-Type: application/json' -d '{
  "amount": {
    "code": "KBY",
    "value": 1000
  },
  "message": "Kabayan Payment from IDR"
}'

curl -X POST http://vmi836509.contaboserver.net:9200/api/v1/owner/accounts/bob/redeem -H 'Content-Type: application/json' -d '{
  "amount": {
    "code": "IDR",
    "value": 1000
  },
  "message": "Kabayan Payment to KBY"
}'
