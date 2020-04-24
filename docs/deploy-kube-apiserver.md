# deploy kube-apiserver

- Open the admission-controllers of ImagePolicyWebhook by adding --enable-admission-plugins
```
    kube-apiserver Increase startup parameters
    --enable-admission-plugins=ImagePolicyWebhook
    --admission-control-config-file=/kube-apiserver-image-admission.yaml
```

- Reference ImagePolicyWebhook in admission-control-config-file, configure kube-apiserver-image-admission.yaml
```
imagePolicy:
  kubeConfigFile: /kube-apiserver-image-admission-config.yaml
  allowTTL: 50
  denyTTL: 50
  retryBackoff: 500
  defaultAllow: true
```
> Note: If you cannot connect to the webhook, all images will be allowed by default

- kubeconfig content kube-apiserver-image-admission-config.yaml
```
apiVersion: v1
kind: Config
clusters:
  - name: image-admission-webhook
    cluster:
      certificate-authority: /etc/kubernetes/ssl/domain-ca.pem
      server: https://kube-admission-image.kube-system/images_admission
contexts:
  - context:
      cluster: image-admission-webhook
      user: apiserver-client
    name: admission_validator
current-context: admission_validator
preferences: {}
users:
  - name: apiserver
    user:
      client-certificate: /etc/kubernetes/ssl/apiserver-client.pem
      client-key: /etc/kubernetes/ssl/apiserver-client-key.pem
```
