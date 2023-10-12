cd .. && cd ..

sudo docker build -f ./build/dockerfile . -t restaurant-document-design-gateway:latest
sudo docker image ls
sudo docker tag restaurant-document-design-gateway localhost:32000/restaurant-document-design-gateway:latest
sudo docker push localhost:32000/restaurant-document-design-gateway:latest

kubectl apply --kustomize ./deployment/microk8s/overlays/dev