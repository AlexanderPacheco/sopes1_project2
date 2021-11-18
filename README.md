# USAC SQUID GAME

## Installing gcloud command
```
echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
sudo apt-get install apt-transport-https ca-certificates gnupg
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
sudo apt-get update && sudo apt-get install google-cloud-sdk
gcloud init
```  

## Create cluster
```
gcloud container clusters create k8s-demo --num-nodes=1 --tags=allin,allout --machine-type=n1-standard-2 --no-enable-network-policy
```

## Installing helm
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

## Install kubeclt stable version
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```

## Install a kubectl specific version
```
curl -LO https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```
> *Nota* Las versiones de kubectl server al client deben ser al menos una version de diferencia, sino no funciona chaos-mesh. El siguiente comando es para ver las versiones

```
kubectl version
```

## Configure kubectl to connect to your cluster
```
gcloud container clusters get-credentials k8s-demo --zone us-central1-c --project sopes1-323620
kubectl get nodes
```

## Installing NGINX ingress controller
```
kubectl create ns nginx-ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx 
helm repo update 
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress
kubectl get services -n nginx-ingress
```

## Installing Docker
```
sudo apt-get install docker.io
sudo usermod -aG docker developer
```

## Building myapp image
```
/bin/bash build.sh 
```

## Install Linkerd
```
curl -fsL https://run.linkerd.io/install | sh

**Agregar el usuario whoami en el export
nano ~/.bashrc <- export PATH=$PATH:/home/YOUR_USER/.linkerd2/bin
o bien 
nano ~/.bashrc <- export PATH=$PATH:$HOME/.linkerd2/bin
**Si no reconoce el comando linkerd, hay que reiniciar porque tendria que jalar :V

linkerd install | kubectl apply -f -
linkerd check
linkerd viz install | kubectl apply -f -
linkerd check
```

## Open Linkerd dashboard
```
linkerd viz dashboard
```

## Inject NGINX ingress-controller
```
kubectl get deployment nginx-ingress-ingress-nginx-controller -n nginx-ingress  -o yaml | linkerd inject - | kubectl apply -f -
```
Check if you can find 2 pods for the NGINX ingress-controller
```
kubectl get pods -n nginx-ingress
```

### Get Load Balancer IP
```
kubectl get svc -n nginx-ingress
```

### Get info squidgame
```
kubectl get all -n squidgame
```

### Get routes for squidgame
```
kubectl describe ingresses -n squidgame
```


## Create the project
*Note:* You have to change the IP in the ingress definition
```
kubectl  apply -f squidgame.yaml 
kubectl delete -f squidgame.yaml 
```

## Test traffic split
```
export LOADBALANCER_IP=X.X.X.X
for i in {0..100000}; do  curl http://${LOADBALANCER_IP}.nip.io/myapp; done
```

## Chaos Mesh Installation
```
curl -sSL https://mirrors.chaos-mesh.org/v2.0.3/install.sh | bash
```
Opening the dashboard
```
kubectl port-forward -n chaos-testing svc/chaos-dashboard 2333:2333
```
```
kubectl apply -f chaos_mesh/pod-experiments.yaml
```
```
kubectl get pods -n squidgame -w
```
## Install and Configure gRPC on Kubernetes

#### Iniciamos el proyecto

vidio
https://www.youtube.com/watch?v=_eTR0if-KYc&ab_channel=CarlosDavid

`mkdir gRPC-Client-api`

`cd gRPC-Client-api`

`go mod init github.com/racarlosdavid/demo-gRPC-kubernetes/gRPC-Client-api`

`mkdir gRPC-Server`

`cd gRPC-Server`

`go mod init github.com/racarlosdavid/demo-gRPC-kubernetes/gRPC-Server`


#### Instalar dependencias gRPC

`go get -u google.golang.org/grpc`

`go get github.com/golang/protobuf/proto@v1.5.2`

`go get google.golang.org/protobuf/reflect/protoreflect@v1.27.1`

`go get google.golang.org/protobuf/runtime/protoimpl@v1.27.1`

#### Instalar dependencias para compilar el .proto

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1`

`export PATH="$PATH:$(go env GOPATH)/bin"`

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/demo.proto`

#### API - Instalamos gorilla mux para el server y go-randomdata para nombres random
`go get -u github.com/gorilla/mux`

`go get github.com/Pallinder/go-randomdata`

#### Creacion de las imagenes de los contenedores
`docker build -t "racarlosdavid/grpc_client_api" .`

`docker build -t "racarlosdavid/grpc_server" .`

#### Prueba de los contenedores
`docker run -it -d -p 2000:2000 -e HOST=192.168.1.4 --name grpc_client_api_ racarlosdavid/grpc_client_api`

`docker run -it -d -p 50051:50051 --name grpc_server_ racarlosdavid/grpc_server`

#### Subir contenedores a dockerhub
`docker login`

`docker push racarlosdavid/grpc_client_api`

`docker push racarlosdavid/grpc_server`

#### Creacion del cluster
`gcloud config get-value project`

`gcloud config set project <NOMBRE DEL PROYECTO>`

`gcloud config set compute/zone us-central1-a`

`gcloud container clusters create <NOMBRE DEL CLUSTER> --num-nodes=1 --tags=all-in,all-out --machine-type=n1-standard-2 --no-enable-network-policy`

#### Instalacion de Ingress
`helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx`

`helm repo update`

`helm install ingress-nginx ingress-nginx/ingress-nginx -n ayd2-backend`

#### Kubernetes

`kubectl apply -f deployment-grpc-kubernetes.yml`

`kubectl apply -f ingress-grpc-kubernetes.yml`

`kubectl get services`

`kubectl get pods`

`kubectl get pods -o wide`

`kubectl describe pod <NOMBRE DEL POD>`

### Instalacion RabbitMQ

`docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management`

> En mi api de go instalar

`go get github.com/rabbitmq/amqp091-go`