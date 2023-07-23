:: starts the system under test
docker network create restaurant

call sut-build.cmd

docker compose -f sut/docker-compose.yml up --build --remove-orphans --exit-code-from restaurant-document-design-gateway

docker compose -f sut/docker-compose.yml down

docker image rm restaurant-document-design-gateway
pause