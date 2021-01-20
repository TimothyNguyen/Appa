# Appa
An innovate platform to apply for job apps more easily, get noticed, and land your dream job!

# Dependency
1. Docker
2. Docker-Compose

# Before you run
First you need to config the URL to the mongodb. This determines which database you are using.

Second, choose a pair of secret keys to encrypt the access token and refresh token. They should be different and random.

All these configuration should be set in the .env file.
# How to run
In the current directory, run:
> docker-compose up --build
To start the project.

# Explanation
The jwt-auth is working at port 8000.

It has two routers:

1. Register
URL: localhost:8000/register
Accept:
{
    email: string,
    password: string
}

Return:
On success
{
    status: "success"
    msgs: null,
}
On failure
{
    status: "unsuccess",
    msgs: [ string, ... ] // error message
}
2. Login Authentication
URL: localhost:8000/login
Accept:
{
    email: string,
    password: string
}

Return:
On success:
{
    status: "success",
    msgs: null,
    data: {
        access_token: string,
        refresh_token: string,
    }
}
On failure:
{
    status: "unsuccess",
    msgs: [ string, ... ], // error msg
    data: null
}

To build the data:
docker-compose up --build

