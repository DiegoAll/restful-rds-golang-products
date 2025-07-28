

1. Healthcheck

    https://{API_Gateway_URL/}{stage}/{API_versioning}/{endpoint}

    https://kj85iaik52.execute-api.us-east-1.amazonaws.com/prod/v1/health


2. Create Product (Without token)


curl -X POST https://kj85iaik52.execute-api.us-east-1.amazonaws.com/prod/v1/products \
    -H "Content-Type: application/json" \
    -d '{ "name": "Unauthorized Product", "price": 10, "description": "Should be rejected", "stock": 1 }'


3. Create Product (With token)


curl -X POST "https://kj85iaik52.execute-api.us-east-1.amazonaws.com/prod/v1/products" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer <TU_ACCESS_TOKEN_VALIDO_Y_LIMPIO>" \
    -d '{ "name": "Authorized Product", "price": 100, "description": "This should work now", "stock": 50 }'

curl -X POST "https://kj85iaik52.execute-api.us-east-1.amazonaws.com/prod/v1/products" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer <TU_ACCESS_TOKEN_VALIDO_Y_LIMPIO>" \
    -d '{ "name": "Authorized Product", "price": 100, "description": "This should work now", "stock": 50 }'