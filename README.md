# Username Checker
API used to check if username exists in source.

Source list:
- Instagram
- Facebook
- Github

Check detail <a href="https://github.com/izzxx/username_checker/schema/source.go">source.go</a> to see avaible sources

## Check

### URL

* Url method GET
```
- http://localhost:{your_port}/api/check/instagram/{username}

- http://localhost:{your_port}/api/check/facebook/{username}

- http://localhost:{your_port}/api/check/github/izzxx
```

### Response

* User found
```json
{
    "status_code": 200,
    "message": "found",
    "url": "https://facebook.com/{username}"
}
```

* User not found
```json
{
    "status_code": 404,
    "message": "user not found",
    "url": ""
}
```

## Library

| Name | Links |
| --- | --- |
| go-chi | https://github.com/go-chi/chi |
| godotenv | https://github.com/joho/godotenv |

## Run Code

```
# move to directory
$ cd workspace

# clone into your workspace 
$ git clone https://github.com/izzxx/username_checker.git

# move to project
$ cd username_checker/

# make .env file, add this
# example:
PORT = ":9000"

# run code
$ go run main.go

# execute the call
$ curl localhost:{your_port}/api/check/{chose_source}/{username}

# stop 
Ctrl + c
```