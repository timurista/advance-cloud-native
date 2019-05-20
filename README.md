# advance-cloud-native
Advanced cloud native using go

## Go microservice frameworks
revamped, conatinerization. Containerize the service, docker compose to build and run it locally. Run micro deploy in kubernetes.

## Anatomy
center --> microservice chassis
  inside is a service client to call others in resilient manner
  service discovery --> exposition of service endpoint
  bottom --> cluster like coordination, dont want to hardcode stage config
    - centralize to some store

### api gateway authorization and traffic managment
github.com/cncf/landscape
https://landscape.cncf.io/

## monkit
diagnosing and metrics, instrumenation, fault analysis etc.

## logrus
nice formatting

## Service Frameworks
- goals
- rpc, also can use json over http
- async like pub/sub, async os
### go micro
  - plugable, message encoping, client/server rpc

### gizmo
  - standardized config and logging
  - health check

### gin
  go, smashing performance
  gin-gonic/gin

## Kubernetes
- pods are the smallest unit in kubernetes
- 
```
$ minikube start 
$ minikube dashbaord // to see list of deployments
```

Troubleshooting
```
if you get access errors in minikube, may need to
downgrade minikube to v0.17.1.
 ```