# Request

Por el browser

    http://localhost:9090/v1/login


root@ph03nix:/home/diegoall/Projects/restful-rds-golang-products# curl -k http://localhost:9090/v1/login
<a href="https://accounts.google.com/o/oauth2/auth?access_type=offline&amp;client_id=89960475367-37h7e26id256t33v5b7p5aho4kbv0gio.apps.googleusercontent.com&amp;redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&amp;response_type=code&amp;scope=openid+profile+email&amp;state=state-token">Temporary Redirect</a>.


http://localhost:8080/callback?state=state-token&code=4%2F0AVMBsJhYgkJXGpePJKOrjvCMPgNxfaSiz2_PUp4dVh9MWjO-h6WLL0auq-O4BtBf-3IiPQ&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=consent




    curl -X POST http://localhost:9090/v1/signup -H "Content-Type: application/json" -d '{
        "email": "dposadallano@hotmail.com",
        "password": "TuPasswordSeguro123!"
    }'


    curl -X POST http://localhost:9090/v1/confirm \
    -H "Content-Type: application/json" \
    -d '{
        "email": "dposadallano@hotmail.com",
        "code": "677769"
    }'

**Returns a AccessToken, IdToken and RefreshToken**

    curl -X POST http://localhost:9090/v1/login \
    -H "Content-Type: application/json" \
    -d '{
        "email": "dposadallano@hotmail.com",
        "password": "TuPasswordSeguro123!"
    }'


