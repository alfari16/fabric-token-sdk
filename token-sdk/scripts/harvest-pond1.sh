curl -X POST http://localhost:9100/api/v1/issuer/issue -H 'Content-Type: application/json' -d '{"amount": {"code": "IDR","value": 800},"counterparty": {"node": "owner1","account": "bob"},"message": "harvest pond 1"    }'

curl -X POST http://localhost:9200/api/v1/owner/accounts/bob/redeem -H 'Content-Type: application/json' -d '{
  "amount": {
    "code": "KBY",
    "value": 800
  },
  "message": "[AUTO A/R] Kabayan Payment from IDR"
}'

curl -X POST http://localhost:9200/api/v1/owner/accounts/bob/redeem -H 'Content-Type: application/json' -d '{
  "amount": {
    "code": "IDR",
    "value": 800
  },
  "message": "[AUTO A/R] Kabayan Payment to KBY"
}'
