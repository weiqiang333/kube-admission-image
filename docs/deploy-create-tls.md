# create TLS

#### domain cert ca
```
cat > domain-ca-csr.json <<EOF
{
  "CN": "domain-ca",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "ca":{
    "expiry": "47520h"
  }
}
EOF

cfssl -loglevel=4 gencert -initca domain-ca-csr.json | cfssljson -bare domain-ca
```

- domain-ca.pem Is the root certificate of your domain name certificate, as a verification

#### domain cert: kube-admission-image
```
cat > kube-admission-image-csr.json <<EOF
{
  "CN": "kube-admission-image.kube-system",
  "key": {
    "algo": "rsa",
    "size": 2048
  }
}
EOF

cfssl -loglevel=4 gencert \
  -ca=domain-ca.pem \
  -ca-key=domain-ca-key.pem \
  kube-admission-image-csr.json | cfssljson -bare kube-admission-image
```

- CN is your webhook access domain name
- kube-admission-image.pem and kube-admission-image-key.pem is webhook TLS certificate
