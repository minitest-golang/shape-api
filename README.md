# 1. Implementation Approach
This Golang assignment is straightforward, but I would like to use microservie approach to implement this shape API.

* One container for Golang RESTful service (User APIs and Shape APIs)
* One container for PostgresQL DB

# 2. Setups 
To bring-up the services, please install docker and docker-compose on your system.
* docker: 
https://docs.docker.com/engine/install/

* docker-compose: 
https://docs.docker.com/compose/install/

For Windows, Docker Desktop could be a choice.

# 3. Unit-test
Go to the project folder, then run following command:
`go test ./...`

# 4. Start containers for test
Please go to the project folder, then run following command:
`docker-compose up -d –-build`

# 5. Swagger API
* The Swagger API is available at: http://34.101.32.199:8081/app/v1/swagger/index.html 
* Or after starting docker containers, the Swagger API can be accessed at local: http://localhost:8081/app/v1/swagger/index.html 

# 6. How to test
* Using python test file: 

    `cd ./test`

    `python3 api_test.py`

* Using Postman

    Import Postman Requests at: `./test/Mini-Test-API.postman_collection.json`
