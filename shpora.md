# crypto project

swagger generate server -f ./swagger.yaml --exclude-main --server-package ./api/restapi --model-package ./internal/models

docker-compose up --build -d

docker-compose up -d

docker-compose ps

docker exec -it crypto_app sh
