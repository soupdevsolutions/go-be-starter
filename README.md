# go-be-starter
Template for starting a BE application in Go.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker)
- [Just](https://github.com/casey/just) (optional, but recommended)

## Getting Started

To get started with the project, run the following commands (if you do not want to use `just`, you can run the commands from the `.justfile` directly).

1. Clone the repository:

```sh
git clone https://github.com/Mirch/go-be-starter 
mv go-be-starter <project-name>
cd <project-name> 
```

2. Initialize the project:

The `<project-name>` will change the Go module name (feel free to change this with a find-and-replace), while the `<database-name>` will change the local database name in the local config and scripts.

```sh  
just init <project-name> <database-name>
```

3. Run the project

This will start the local database container and run the Go project. 

```sh
just start-local
```

You can test that everything is working by calling the `/api/hello` endpoint:

```sh 
curl http://127.0.0.1:8080/api/hello
```

## Working with the project 

1. Creating migrations

To create a new migration, run the following command:

```sh
just create-migration <migration-name>
```

This will create a `.up` and `.down` file in the `migrations` directory, where you can write your SQL queries.
