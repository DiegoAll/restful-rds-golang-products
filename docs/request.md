# Request

Por el browser

    http://localhost:9090/v1/login


root@ph03nix:/home/diegoall/Projects/restful-rds-golang-products# curl -k http://localhost:9090/v1/login
<a href="https://accounts.google.com/o/oauth2/auth?access_type=offline&amp;client_id=89960475367-37h7e26id256t33v5b7p5aho4kbv0gio.apps.googleusercontent.com&amp;redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&amp;response_type=code&amp;scope=openid+profile+email&amp;state=state-token">Temporary Redirect</a>.


http://localhost:8080/callback?state=state-token&code=4%2F0AVMBsJhYgkJXGpePJKOrjvCMPgNxfaSiz2_PUp4dVh9MWjO-h6WLL0auq-O4BtBf-3IiPQ&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=consent


**Submit**

    curl -X POST http://localhost:9090/v1/signup -H "Content-Type: application/json" -d '{
        "email": "dposadallano@hotmail.com",
        "password": "TuPasswordSeguro123!"
    }'

**Confirm**

    curl -X POST http://localhost:9090/v1/confirm \
    -H "Content-Type: application/json" \
    -d '{
        "email": "dposadallano@hotmail.com",
        "code": "677769"
    }'

**Login - Returns a AccessToken, IdToken and RefreshToken**

    curl -X POST http://localhost:9090/v1/login \
    -H "Content-Type: application/json" \
    -d '{
        "email": "dposadallano@hotmail.com",
        "password": "TuPasswordSeguro123!"
    }'


**¿Qué Token Utilizar para Endpoints Protegidos?**

De los tokens que te ha devuelto Cognito, el que debes utilizar para acceder a los endpoints protegidos (como /v1/products si lo configuraras con un middleware de autenticación JWT) es el AccessToken.


    curl -X POST http://localhost:9090/v1/products \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer EYJraWQiOiJTR1pKb0dLY3BcL1lVb1NoVkYwdWl6RzdRZnNmUjBlNVVKdDN2QmtKdCtcLzg9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIxNGQ4NDQ2OC0wMDMxLTcwMWUtZTU2Mi1lZmE0YzhlNjA0ZmUiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9aVHpTbmxHODEiLCJjbGllbnRfaWQiOiIyb2E0NXJjcmw2NnFvcGh2Y2NhZWVzZHRsOSIsIm9yaWdpbl9qdGkiOiI4M2QyNDcyMC1kZmQ3LTRmYTQtODVlOC04ZGE1MWI5MDI0ODAiLCJldmVudF9pZCI6IjUyNDRmZDQ1LTZkYzctNDc3Mi05N2YxLTRjOTQzODE1YTBiNSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NTMzMDk1MDcsImV4cCI6MTc1MzMxMzEwNywiaWF0IjoxNzUzMzA5NTA3LCJqdGkiOiJkZmUxMTJiOS0wOGY4LTQ3OWItOWNiMi0xOTE1NmI5MmYwN2UiLCJ1c2VybmFtZSI6ImRwb3NhZGFsbGFub0Bob3RtYWlsLmNvbSJ9.HdoyeDcioSlUA329mI4vTphlBFT9K2EMbDXru4RvkYbJuu5tDlO_2G_XDPdoFhgX6uMOIplu5CvmRlsYpG7a6dZEmc5qCFvTrZIyorK26hAogU_w6zIUTy-U15Nkao7mCWcSCkG4C4KwD_6mxCWATe3Bu1vmnT6cJjTOH6wsMIuGYWCR4DaSF1GrdzaMrm_xF3wkvQb-acHl7cS29wk3x-8Xq8A3v048p-yU43WaMYHpPEMBZ0xKaWnUkq_3Dn5Eec-P_5CLXkok3NCfqCsPULYxiQHMrTkX9yVcyIxKtnSCG8P1adWpIo65kL4vMleS0QNWLWBuTwOPPy35g-7mXg" \
    -d '{
        "name": "Producto Autenticado",
        "description": "Este producto requiere autenticación.",
        "price": 50.99
    }'