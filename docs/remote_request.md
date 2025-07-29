# Remote request



    curl http://checkip.amazonaws.com/



🧪 1. Healthcheck

    curl http://<IP-PUBLICA-DE-LA-EC2>:9090/v1/health

    curl http://3.92.73.118:9090/v1/health




📝 2. Signup en Cognito (/v1/signup)

    curl -X POST http://54.211.39.213:9090/v1/signup \
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


        curl -X POST http://54.211.39.213:9090/v1/login \
        -H "Content-Type: application/json" \
        -d '{
            "email": "testuser@example.com",
            "password": "TestPassword123!"
        }'


5. Usar el token en rutas protegidas (como crear un producto)

        curl -X POST http://54.211.39.213:9090/v1/products \
        -H "Authorization: Bearer <ACCESS_TOKEN>" \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Coca Cola",
            "price": 100,
            "description": "Random drink",
            "stock": 50
        }'


        curl -X POST http://54.211.39.213:9090/v1/products \
        -H "Authorization: Bearer eyJraWQiOiJIdGdBN0hzZXo2clRVZWE2eVAwN0VNdTNrcDBRZmxjUHRYeEJOWWxDeGVBPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJlNDc4YTQ1OC1iMDIxLTcwZWMtMDc3ZS03ZTk5OWQyMjNhNTYiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV94cjA1NFczWWkiLCJjbGllbnRfaWQiOiI1Ymo1M25kMXQwdjcyNzlxa3BhajBib25qOCIsIm9yaWdpbl9qdGkiOiJmNGZmMWMzOC1kZjQxLTRkMjMtYTViMC05OGUxZDgyMzQ2ZWYiLCJldmVudF9pZCI6Ijk5MDE4M2EwLTc4NDUtNDU4Zi1hOWIyLWM3ZDljZmViM2VjYiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM2ODA2ODAsImV4cCI6MTc1MzY4NDI4MCwiaWF0IjoxNzUzNjgwNjgwLCJqdGkiOiIxM2RlZGU2My1iZDlmLTRlYTAtYjQ1Mi1iZmFjY2NjODkxMDMiLCJ1c2VybmFtZSI6ImU0NzhhNDU4LWIwMjEtNzBlYy0wNzdlLTdlOTk5ZDIyM2E1NiJ9.AGBO-M1CfN6sqGoHj0HOMf1UtFabikRwt6qgYbcniuaveY0-n1fBL92UbqQVxbyUf5iqnh0vy5YQ7_aDynSqAyeh8nCUXtIMmNB9BrC8pSgS_LsP0GxQyC7OCL98IiYuU9XUOSiah0DBc0fQ4r1AA_eTU8cgnHyFtqPrRHy_z5qJvpdE7NBjaN85jUNHMshU83SjZWIJd7yJuycs8NDnQlMQwT5H_p0RJo_fQKTXX3ApA06-4YFH1AfEXIunRu4vr6geyBofGfD712Ap1RNZr2Y3WJJZQC3XZ6C0e8nnax0W-PDEWOqIioJWPQRVfd1UUTmvFfzZJSPOb9SF-yiUyA" \
        -H "Content-Type: application/json" \
        -d '{
            "name": "Coca Cola",
            "description": "Random drink",
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


curl -X POST https://kj85iaik52.execute-api.us-east-1.amazonaws.com/v1/products \
    -H "Authorization: Bearer eyJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIzNDM4MDQwOC01MDIxLTcwNTMtN2ZkMy0xOGNmNDBhMjkxNWQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9aVHpTbmxHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4ZjlhMDE1MC0zY2E3LTQ0NGItYTQwYi01ZTkwNGI0ODg5ZGMiLCJldmVudF9pZCI6ImNmNzY2YzNjLWM3YTctNGExNC1hMzUxLWJkNmZiZTE0MmZkYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM1NjM3NTgsImV4cCI6MTc1MzU2NzM1OCwiaWF0IjoxNzUzNTYzNzU4LCJqdGkiOiI2OGVjMjRmZC05ZTBmLTRmNWYtYTNkOS03YzU2NGMwYjVlNDciLCJ1c2VybmFtZSI6InRlc3R1c2VyQGV4YW1wbGUuY29tIn0.xA_gfCX-MCvIxa8pygyR7BQFBjwe7sKGjZszz5qdWdyUbJ39uVShkTEA2iIxzEEPhRL74qi_h67XaU1BROBK1ZxrDQhLYhIAfsZXZ3OucIdRLZJhkSfNMAObpYqZojOjuNFZSy_oJhxurHTrTe_zEprbma1l2IrL1kK8IUs4Q0s_3KgxjyCPgOJnWnBW8ahVLg-6gNfpSj6msD--BmTsSHmkA7QejpzRYB0d_7IErmALHLUu9VqxBYHLKlLXM8aMbKvLMHZDGLuj7swRJfcy8qj8tUdu8RzSlNBGyQgxihfv4zfic_ZGI_q9WYOYBbwBUi4E62-l0orMyRUHcSdn_A" \
    -H "Content-Type: application/json" \
    -d '{ "name": "Authorized Product", "price": 100, "description": "This should work now", "stock": 50 }'

eyJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIzNDM4MDQwOC01MDIxLTcwNTMtN2ZkMy0xOGNmNDBhMjkxNWQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9aVHpTbmxHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4ZjlhMDE1MC0zY2E3LTQ0NGItYTQwYi01ZTkwNGI0ODg5ZGMiLCJldmVudF9pZCI6ImNmNzY2YzNjLWM3YTctNGExNC1hMzUxLWJkNmZiZTE0MmZkYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM1NjM3NTgsImV4cCI6MTc1MzU2NzM1OCwiaWF0IjoxNzUzNTYzNzU4LCJqdGkiOiI2OGVjMjRmZC05ZTBmLTRmNWYtYTNkOS03YzU2NGMwYjVlNDciLCJ1c2VybmFtZSI6InRlc3R1c2VyQGV4YW1wbGUuY29tIn0.xA_gfCX-MCvIxa8pygyR7BQFBjwe7sKGjZszz5qdWdyUbJ39uVShkTEA2iIxzEEPhRL74qi_h67XaU1BROBK1ZxrDQhLYhIAfsZXZ3OucIdRLZJhkSfNMAObpYqZojOjuNFZSy_oJhxurHTrTe_zEprbma1l2IrL1kK8IUs4Q0s_3KgxjyCPgOJnWnBW8ahVLg-6gNfpSj6msD--BmTsSHmkA7QejpzRYB0d_7IErmALHLUu9VqxBYHLKlLXM8aMbKvLMHZDGLuj7swRJfcy8qj8tUdu8RzSlNBGyQgxihfv4zfic_ZGI_q9WYOYBbwBUi4E62-l0orMyRUHcSdn_A







curl -X POST https://kj85iaik52.execute-api.us-east-1.amazonaws.com/v1/products \
    -H 'Authorization: Bearer eyJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIzNDM4MDQwOC01MDIxLTcwNTMtN2ZkMy0xOGNmNDBhMjkxNWQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9ZVHpOblpHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4ZjlhMDE1MC0zY2E3LTQ0NGItYTQwYi01ZTkwNGI0ODg5ZGMiLCJldmVudF9pZCI6ImNmNzY2YzNjLWM3YTctNGExNC1hMzUxLWJkNmZiZTE0MmZkYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTM1NjM3NTgsImV4cCI6MTc1MzU2NzM1OCwiaWF0IjoxNzUzNTYzNzU4LCJqdGkiOiI2OGVjMjRmZC05ZTBmLTRmNWYtYTNkOS03YzU2NGMwYjVlNDciLCJ1c2VybmFtZSI6InRlc3R1c2VyQGV4YW1wbGUuY29tIn0.xA_gfCX-MCvIxa8pygyR7BQFBjwe7sKGjZszz5qdWdyUbJ39uVShkTEA2iIxzEEPhRL74qi_h67XaUjOjuNFZSy_oJhxurHTrTe_zEprbma1l2IrL1kK8IUs4Q0s_3KgxjyCPgOJnWnBW8ahVLg-6gNfpSj6msD--BmTsSHmkA7QejpzRYB0d_7IErmALHLUu9VqxBYHLKlLXM8aMbKvLMHZDGLuj7swRJfcy8qj8tUdu8RzSlNBGyQgxihfv4zfic_ZGI_q9WYOYBbwBUi4E62-l0orMyRUHcSdn_A' \
    -H "Content-Type: application/json" \
    -d '{ "name": "Authorized Product", "price": 100, "description": "This should work now", "stock": 50 }'