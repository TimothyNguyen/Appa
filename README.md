# Appa
An innovate platform to apply for job apps more easily, get noticed, and land your dream job!

# Dependency
1. Docker
2. Docker-Compose

# How to run
In the current directory, run:
> docker-compose up --build
To start the project.

# Explanation
The mongoDB is working on mongo:27017 and you can monitor the content inside using mongo_express.
Mongo_Express is accessible on port [8081](localhost:8081).
Nginx is accessible on port [8080](localhost:8080).
When develope the project, you can use the following url: "mongodb://root:password@mongo:27017" to connect to the mongoDB.