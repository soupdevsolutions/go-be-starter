MOD_NAME = soup.dev/go-be-starter
NEW_MOD_NAME = $0
find . -name '*' -exec sed -i -e 's/$MOD_NAME/$NEW_MOD_NAME' {} \; 
