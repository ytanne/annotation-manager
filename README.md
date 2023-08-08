# annotation-manager

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Go Migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- [Sqlite driver](https://github.com/golang-migrate/migrate/issues/670#issuecomment-1118029997) 

## Setup

1. Clone the repository
2. Setup the database
    - `migrate -database sqlite3://database/sqlite.db -path ./migrations up`
3. Create container
    - `docker build -t annotation-manager .`
4. Run container
    - `docker run -p 8080:8080 -v ./config/:/app/config/ -v ./database/:/app/database annotation-manager`
5. Use Postman documentation (from ./documents/postman) to test the application (*Optional*)

## Thoughts and considerations
The application is far from being perfect and production ready. I would have liked to spend more time on it but unfortunately I didn't have time to do so. I have listed some of the things that I would have liked to do if I had more time:
 - Right now the application doesn't have almost any user management. It has hardcoded user credentials (i.e. 'admin:admin' and 'test:test'). Admin has all rights and test user doesn't have any.
 - The configuration (i.e. obtaining values from config file) could contain bugs since I didn't have time to test it properly.
 - The application doesn't have any tests. I would have liked to write some unit tests for all layers (i.e. controller, usecase, repository) and packages.
 - The application doesn't have any logging.
 - The application doesn't have any tracing.
 - The application could have Swagger documentation.
 - The application could have better error handling and http responses.
 - The application could have graceful shutdown.
 - The application could have better input validation.
 - I have used gin as a web framework. I would have liked to use something more lightweight like [fiber](https://gofiber.io/).
 - I have implemented single stage build in Dockerfile. I would have liked to implement multi stage build to reduce the final size of the image.
 - There are many things to do yet. But I think that this is a good start.
