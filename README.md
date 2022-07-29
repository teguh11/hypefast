# hypefast

for installation and running process.
1. clone this repository
2. run go mod tidy
3. run go main.go http

notes: this api use port 8080 for running. you can change this port in file cmd/http.go

while service running, you can use this endpoint for checking

1.  Generate short URL
    Request
    curl --location --request POST 'localhost:8080/generate' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "url": "https://detik.com"
    }'

    Response
    {
        "data": "http://localhost:8080/Xo9WxM",
        "err_msg": null
    }

2.  Redirect URL
    Request
    curl --location --request GET 'http://localhost:8080/Bp4rRG'
    
3.  Get Statistic for short URL
    Request
    curl --location --request GET 'localhost:8080/stat/Xo9WxM'
    
    Response
    {
        "data": {
            "url": "https://detik.com",
            "shorten_value": "http://localhost:8080/Xo9WxM",
            "created_at": "2022-07-29 09:46:26.8444878 +0700 +07 m=+15.721490601",
            "redirect_count": 2
        },
        "err_msg": null
    }
