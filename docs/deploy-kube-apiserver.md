# deploy kube-apiserver

- 通过增加 --enable-admission-plugins 来开启 ImagePolicyWebhook 的 admission-controllers
```
    kube-apiserver 增加启动参数
    --enable-admission-plugins=ImagePolicyWebhook
    --admission-control-config-file=/kube-apiserver-image-admission.yaml
```

- admission-control-config-file 中引用 ImagePolicyWebhook 配置 kube-apiserver-image-admission.yaml
```
imagePolicy:
  kubeConfigFile: /kube-apiserver-image-admission-config.yaml
  allowTTL: 50
  denyTTL: 50
  retryBackoff: 500
  defaultAllow: true
```
> 注意：如果连接不到 webhook 将默认允许所有 images

- kubeconfig 内容 kube-apiserver-image-admission-config.yaml
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
