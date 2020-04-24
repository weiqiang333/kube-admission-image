# kube-admission-image

[ImagePolicyWebhook](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#imagepolicywebhook) on admission-controllers for kubernetes

Use kube-admission-image for customized image verification or scanning


## admission category

This will define the game rules for admission

 - [images name specification](docs/name-policy.md)
 - [images unauthorized source Specification](docs/source-policy.md)
 - images size specification
 - images startup specifications
 - images known vulnerability scan


## building
- go get github.com/weiqiang333/kube-admission-image

#### dev
- go run kube-admission-image.go

#### container
- docker pull weiqiang333/kube-admission-image

#### docker hub
- [docker hub](https://hub.docker.com/repository/docker/weiqiang333/kube-admission-image)


## open ImagePolicyWebhook

#### TLS
 - Create a TLS certificate to protect the webhook service
 - kubernetes will be accessed via the TLS protocol
 
    [View detailed process](docs/deploy-create-tls.md)

#### kube-apiserver
 - kube-apiserver admission-control add ImagePolicyWebhook Control plugin
 - config admission-control-config-file
 - load kube-admission-image kubeconfig

    [View detailed process](docs/deploy-kube-apiserver.md)

#### deploy kube-admission-image
- First, create the TLS secret required by the webhook:
```
kubectl -n kube-system create secret tls tls-kube-admission-image \
  --key kube-admission-image-key.pem \
  --cert kube-admission-image.pem
```
- deploy kube-admission-image
```
kubectl apply -f configs/kubernetes/kube-admission-image-deployment.yaml
```


## FAQ
- Pay attention to using strategy logic to avoid chicken and egg problems
