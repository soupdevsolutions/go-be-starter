go-get:
    @echo "Getting dependencies..."
    @cd src && go get -v -t ./...

init name:
    @echo "Initializing project..."
    @chmod +x ./scripts/init_project.sh
    @./scripts/create.sh {{name}} 
    @rm -f ./scripts/init_project.sh

create-migration name:
    @echo "Creating migration..."
    @migrate create -ext sql -dir ./migrations -seq {{name}} 

start-local:
    @echo "Initializing local database..."
    @chmod +x ./scripts/init_db.sh
    @./scripts/init_db.sh
    @cd src && go run .
