# Remote request



    curl http://checkip.amazonaws.com/



🧪 1. Healthcheck

    curl http://<IP-PUBLICA-DE-LA-EC2>:9090/v1/health


📝 2. Signup en Cognito (/v1/signup)

    curl -X POST http://54.161.41.246:9090/v1/signup \
    -H "Content-Type: application/json" \
    -d '{
        "email": "testuser@example.com",
        "password": "TestPassword123!"
    }'


3. Confirmar el código (/v1/confirm)

        curl -X POST http://54.161.41.246:9090/v1/confirm \
        -H "Content-Type: application/json" \
        -d '{
            "email": "testuser@example.com",
            "confirmation_code": "123456"
        }'


4. Hacer login y obtener el token JWT (/v1/login)


        curl -X POST http://54.161.41.246:9090/v1/login \
        -H "Content-Type: application/json" \
        -d '{
            "email": "testuser@example.com",
            "password": "TestPassword123!"
        }'


5. Usar el token en rutas protegidas (como crear un producto)

        curl -X POST http://54.161.41.246:9090/v1/products \
        -H "Authorization: Bearer <ACCESS_TOKEN>" \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Coca Cola",
            "price": 100,
            "description": "Random drink"
            "stock": 50
        }'


        curl -X POST http://54.161.41.246:9090/v1/products \
        -H "Authorization: Bearer eyJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1NDY4ZjQ0OC0wMGUxLTcwNzctMTJlZi0wNTQyYWI3OGM2MTUiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9aVHpTbmxHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4MmE3ZDc1YS03MzdlLTRhNDctYjAzMy03MDg5ZWFiMzc3OTEiLCJldmVudF9pZCI6ImZiOGY0YzZlLTM4NWQtNGYxNC1iZjdlLWUwMDljYWRkZjAwMiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM0NzQzMTMsImV4cCI6MTc1MzQ3NzkxMywiaWF0IjoxNzUzNDc0MzEzLCJqdGkiOiJkN2Y1MjczYi1mODdlLTQwZTEtOWIwNS1kNDgyMTUyYzY3MzciLCJ1c2VybmFtZSI6InRlc3R1c2VyQGV4YW1wbGUuY29tIn0.khtJLSIj1_4u3ZTX5hS7CQYWKGYL0xXl6EFDh3tawSgLGKhmqDl_mxyh__7oJC2LIDw-iHFcPmLGI9tM1YxGgAwWrow4WSnxfoqrAYezQbjNzPizhASlW1mH4VtlFpCbi4g6dx8XOvSOfUy0_J7JDERitPfKG6wejFHXzowVy6eSHLKbev9Z-IWaUowISlpmamO3aOs0oY_ay8OBVdZBI41Vwda-2vHapuW-KalHLI380v9IL61X5Pyf-7uR7bXglBELX2C9YSrbTLHnkEuJUE1rQvUvR-sQ1RnVCHxRUlSTql130AWNYK_ZnStZPKF-B99zPubFgYB0th2Arl0_nA" \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Coca Cola",
            "description": "Random drink"
            "price": 100,
            "stock": 50
        }'


curl -X POST http://54.161.41.246:9090/v1/products \
  -H "Authorization: Bearer eyJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1NDY4ZjQ0OC0wMGUxLTcwNzctMTJlZi0wNTQyYWI3OGM2MTUiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9aVHpTbmxHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4MmE3ZDc1YS03MzdlLTRhNDctYjAzMy03MDg5ZWFiMzc3OTEiLCJldmVudF9pZCI6ImZiOGY0YzZlLTM4NWQtNGYxNC1iZjdlLWUwMDljYWRkZjAwMiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM0NzQzMTMsImV4cCI6MTc1MzQ3NzkxMywiaWF0IjoxNzUzNDc0MzEzLCJqdGkiOiJkN2Y1MjczYi1mODdlLTQwZTEtOWIwNS1kNDgyMTUyYzY3MzciLCJ1c2VybmFtZSI6InRlc3R1c2VyQGV4YW1wbGUuY29tIn0.khtJLSIj1_4u3ZTX5hS7CQYWKGYL0xXl6EFDh3tawSgLGKhmqDl_mxyh__7oJC2LIDw-iHFcPmLGI9tM1YxGgAwWrow4WSnxfoqrAYezQbjNzPizhASlW1mH4VtlFpCbi4g6dx8XOvSOfUy0_J7JDERitPfKG6wejFHXzowVy6eSHLKbev9Z-IWaUowISlpmamO3aOs0oY_ay8OBVdZBI41Vwda-2vHapuW-KalHLI380v9IL61X5Pyf-7uR7bXglBELX2C9YSrbTLHnkEuJUE1rQvUvR-sQ1RnVCHxRUlSTql130AWNYK_ZnStZPKF-B99zPubFgYB0th2Arl0_nA" \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Coca Cola",
        "description": "Random drink",
        "price": 100,
        "stock": 50
    }'
