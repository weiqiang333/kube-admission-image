# Test acceptance

### check latest tag
- curl
```shell script
# 请求:
    curl -X POST --data \
     '{"kind":"ImageReview","apiVersion":"imagepolicy.k8s.io/v1alpha1","metadata":{"creationTimestamp":null},"spec":{"containers":[{"image":"registry.cn-hangzhou.aliyuncs.com/imagewebhook:3.14.0-20200421-21843c7"},{"image":"registry.cn-hangzhou.aliyuncs.com/action:latest"}],"namespace":"pipeline-10009687"},"status":{"allowed":false}}' \
     https://kube-admission-image.kube-system/images_admission
# 响应：
    {"kind":"ImageReview","apiVersion":"imagepolicy.k8s.io/v1alpha1","metadata":{"creationTimestamp":null},"spec":{},"status":{"allowed":false,"reason":"Images using latest tag are not allowed"}}
```

- deploy check
```shell script
# create pod: 
    cat >> test-latest.yaml << EOF
apiVersion: v1
kind: ReplicationController
metadata:
  name: test-latest
spec:
  selector:
    app: test-latest
  template:
    metadata:
      name: test-latest
      labels:
        app: test-latest
    spec:
      containers:
      - name: test-latest
        image: nginx:latest
        ports:
        - containerPort: 80
EOF

  kubectl apply -f test-latest.yaml
  kubectl describe rc nginx-latest

  # Warning  FailedCreate  22s (x5 over 60s)  replication-controller  (combined from similar events): Error creating: pods "test-latest-bd2qs" is forbidden: image policy webhook backend denied one or more images: Images using latest tag are not allowed

# 成功拒绝最新镜像
```
