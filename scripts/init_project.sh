MOD_NAME=soup.dev/go-be-starter
NEW_MOD_NAME=$1

DATABASE_NAME=main
NEW_DATABASE_NAME=$1

find . \( -name \*.go -o -name \*.mod \) -exec sed -i '' "s#$MOD_NAME#$NEW_MOD_NAME#g" {} \;
find . \( -name \*.sh -o -name \*.yaml \) -exec sed -i '' "s#$DATABASE_NAME#$NEW_DATABASE_NAME#g" {} \;
