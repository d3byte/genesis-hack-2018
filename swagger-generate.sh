BLUE='\033[0;34m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "${BLUE}Starting generation of swagger.json${NC}"

swagger-codegen generate \
   -i ./swagger-specs/swagger.json \
   -l typescript-angular \
   -o ./src/app/swagger/

echo "${GREEN}Generated swagger.json${NC}"
