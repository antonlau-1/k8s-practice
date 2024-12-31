This is an experimental project to learn about the basic of Kubernetes.

The main objectives are to:
- learn about different kinds of deployment - Pods, Deployment, Service
- learn about Kubernetes orchestration capabilities - load-balancing, telemetry
- learn about deployment customisation using YAML files

k8s-hello
This is a basic service that returns Hello World when a request is made.

k8s-to-nginx
This service allows user to communicate between two deployments - default nginx and a basic hello endpoint.

deployment.yaml
This is the deployment configuration file associated with the k8s-hello deployment.

service.yaml
This is the service configuration file asssociated with the k8s-hello deployment.

k8s-to-nginx.yaml
This is the dual-purpose configuration file for the k8s-to-nginx deployment/ service.

nginx.yaml
This is the deployment configuration file associated with the nginx deployment
