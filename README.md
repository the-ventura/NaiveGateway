# NaiveGateway
A naive implementation of a payment gateway. It contains no security or authorization, every user can deposit and withdraw from every account.

## Running

### The easy way (recommended)
There is a docker compose file already configured which will build and run everything you need.
First build everything with
```
docker compose build -f docker-compose-demo.yaml
```
Then run the stack
```
docker compose up -f docker-compose-demo.yaml
```
Visit localhost:3000 in your machine and you should see the landing page.

### Locally
You will need a postgres instance running, you can launch one directly in your machine using the included docker-compose-dev.yaml file by running
```
docker compose up -f docker-compose-dev.yaml
```
You may want to copy the `docker-compose.yaml.example` file and drop the example extention and customize the file.

Next you should compile the gateway if you have not downloaded the binary from the releases page by running
```
make
```
This will create a file called `naivegateway` inside the `./bin` folder

Run database migrations with
```
./bin/naivegateway database migrate -t `cat MIGRATION_VERSION`
```

Run the api server with
```
./bin/naivegateway api
```

### Using docker
You will need a postgres instance running in the same network as the docker container. You can build the docker container from the dockerfile included in this repo by running
```
docker build -t naivegateway .
```

and then run it with
```
docker run naivegateway
```
You may want to set a volume with the configurations or use environment variables. More on that down below.
You may also want run the migrations as in the above step

## Configuring
This project comes with two complimentary methods of configuration; A config file and environment variables
Check `configs/config.yaml.example` for the standard configuration file.

### Environment Variables
| Name                        | Default             | Description                                                          |
|-----------------------------|---------------------|----------------------------------------------------------------------|
| API_URL                     |                     | The api service's public url                                         |
| GATEWAY_API_ALLOWED_ORIGINS |                     | Cors domains, comma separated. e.g: "app.domain.com, api.domain.com" |
| GATEWAY_API_PORT            | 5009                | Default port to listen on                                            |
| GATEWAY_CONFIG_PATH         | configs/config.yaml | Default path for the configuration file                              |
| GATEWAY_DB_DB_NAME          | gateway             | Database to connect to                                               |
| GATEWAY_DB_HOST             | localhost           | Database host                                                        |
| GATEWAY_DB_PASSWORD         |                     | Database password                                                    |
| GATEWAY_DB_PORT             | 5432                | Database port                                                        |
| GATEWAY_DB_USER             |                     | Database user                                                        |
| GATEWAY_FRONTEND_PORT       | 3000                | Default port for the frontend service                                |
| GATEWAY_LOG_LEVEL           | info                | Log level                                                            |
