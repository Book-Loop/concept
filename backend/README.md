# Book Loop Go Server with an Nginx Proxy and a Postgres Database

This project contains the source code and configuration files for deploying a Book Loop Go server with an Nginx proxy and a Postgres database using Docker Compose.

The project structure is as follows:

```bash
├── backend
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── db
│   └── password.txt
├── compose.yaml
├── proxy
│   └── nginx.conf
└── README.md
```

## Prerequisites

Before deploying the application, make sure you have the following software installed on your system:

* Docker Engine
* Docker Compose

## Deployment

To deploy the application, follow these steps:

1. Clone this repository to your local machine.

2. Open a terminal and navigate to the project directory.

3. Run the following command to start the application:

    `docker-compose up -d`

4. Docker Compose will build the required Docker images and start the containers. Wait for the process to complete.

5. Once the application starts, navigate to `http://localhost:85` in your web browser or run the following command to test the server:

    `curl localhost:85`

    You should see a response with a list of book posts.

6. When you're done testing the application, stop and remove the containers by running the following command:

    `docker-compose down`

## Services

The application consists of the following services:

* `backend`: A Go server that serves a list of book posts.
* `db`: A Postgres database for storing the book posts.
* `proxy`: An Nginx reverse proxy that routes requests to the backend server.

The Docker Compose file (`compose.yaml`) defines the services and their configurations. The backend service is built using the Dockerfile in the `backend` directory. The proxy service uses the `nginx.conf` file in the `proxy` directory to configure Nginx.

## Configuration

The `compose.yaml` file contains the configuration for the services. You can modify the file to change the port mappings or the database credentials. The database password is stored in the `password.txt` file in the `db` directory.
