# Auth with oidc 

## 01 start minikube 

```bash
minikube start --addons=ingress --vm=true --memory=8192 --cpus=4
```



```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.5/cert-manager.yaml	
```

https://cert-manager.io/docs/installation/        

```bash
kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=controller -n cert-manager --timeout=60s && \
kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=cainjector -n cert-manager --timeout=60s && \
kubectl wait --for=condition=ready pod -l app.kubernetes.io/component=webhook -n cert-manager --timeout=60s
```



```bash
kubectl apply -f - <<EOF
---
apiVersion: v1
kind: Namespace
metadata:
  name: keycloak
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: keycloak-selfsigned
  namespace: keycloak
  labels:
    app: keycloak
spec:
  selfSigned: {}
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: keycloak-selfsigned
  namespace: keycloak
  labels:
    app: keycloak
spec:
  isCA: true
  commonName: keycloak-selfsigned-ca
  privateKey:
    algorithm: ECDSA
    size: 256
  issuerRef:
    name: keycloak-selfsigned
    kind: Issuer
    group: cert-manager.io
  secretName: ca.crt
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: keycloak
  namespace: keycloak
  labels:
    app: keycloak
spec:
  ca:
    secretName: ca.crt
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: keycloak
  namespace: keycloak
  labels:
    app: keycloak
spec:
  isCA: false
  commonName: keycloak
  dnsNames:
    - keycloak.$(minikube ip).nip.io
  privateKey:
    algorithm: RSA
    encoding: PKCS1
    size: 4096
  issuerRef:
    kind: Issuer
    name: keycloak
    group: cert-manager.io
  secretName: keycloak.tls
  subject:
    organizations:
      - Local Eclipse Che
  usages:
    - server auth
    - digital signature
    - key encipherment
    - key agreement
    - data encipherment
---
apiVersion: v1
kind: Service
metadata:
  name: keycloak
  namespace: keycloak
  labels:
    app: keycloak
spec:
  ports:
  - name: http
    port: 8080
    targetPort: 8080
  selector:
    app: keycloak
  type: LoadBalancer
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: keycloak
  namespace: keycloak
  labels:
    app: keycloak
spec:
  replicas: 1
  selector:
    matchLabels:
      app: keycloak
  template:
    metadata:
      labels:
        app: keycloak
    spec:
      containers:
      - name: keycloak
        image: quay.io/keycloak/keycloak:18.0.2
        args: ["start-dev"]
        env:
        - name: KEYCLOAK_ADMIN
          value: "admin"
        - name: KEYCLOAK_ADMIN_PASSWORD
          value: "admin"
        - name: KC_PROXY
          value: "edge"
        ports:
        - name: http
          containerPort: 8080
        readinessProbe:
          httpGet:
            path: /realms/master
            port: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: keycloak
  namespace: keycloak
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/proxy-connect-timeout: '3600'
    nginx.ingress.kubernetes.io/proxy-read-timeout: '3600'
    nginx.ingress.kubernetes.io/ssl-redirect: 'true'
spec:
  tls:
    - hosts:
        - keycloak.$(minikube ip).nip.io
      secretName: keycloak.tls
  rules:
  - host: keycloak.$(minikube ip).nip.io
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: keycloak
            port:
              number: 8080
EOF
```



```bash
kubectl get secret ca.crt -o "jsonpath={.data['ca\.crt']}" -n keycloak | base64 -d > keycloak-ca.crt
```



```bash
minikube ssh sudo "mkdir -p /etc/ca-certificates" && \
  minikube cp keycloak-ca.crt /etc/ca-certificates/keycloak-ca.crt
```

##   02 minikube with oidc  

```bash
minikube start \
    --extra-config=apiserver.oidc-issuer-url=https://keycloak.$(minikube ip).nip.io/realms/che \
    --extra-config=apiserver.oidc-username-claim=email \
    --extra-config=apiserver.oidc-client-id=k8s-client \
    --extra-config=apiserver.oidc-ca-file=/etc/ca-certificates/keycloak-ca.crt
```



```bash
kubectl exec deploy/keycloak -n keycloak -- bash -c \
    "/opt/keycloak/bin/kcadm.sh config credentials \
        --server http://localhost:8080 \
        --realm master \
        --user admin  \
        --password admin && \
    /opt/keycloak/bin/kcadm.sh create realms \
        -s realm='che' \
        -s displayName='che' \
        -s enabled=true \
        -s registrationAllowed=false \
        -s resetPasswordAllowed=true && \
    /opt/keycloak/bin/kcadm.sh create clients \
        -r 'che' \
        -s clientId=k8s-client \
        -s id=k8s-client \
        -s redirectUris='[\"*\"]' \
        -s directAccessGrantsEnabled=true \
        -s secret=eclipse-che && \
    /opt/keycloak/bin/kcadm.sh create users \
        -r 'che' \
        -s username=test \
        -s email=\"test@test.com\" \
        -s enabled=true \
        -s emailVerified=true &&  \
    /opt/keycloak/bin/kcadm.sh set-password \
        -r 'che' \
        --username test \
        --new-password test"
```

## 03 get token from oidc

```bash
curl -ks -X POST https://${Keycloak服务域名}/realms/myrealm/protocol/openid-connect/token \
-d grant_type=password -d client_id=ack \
-d username=myuser -d password=${已设置的用于登录Keycloak的密码} -d scope=openid \
-d client_secret=${client credential}
```



## 04 access apiserver with token  



```bash
curl -k https://${API server地址}/api/v1/namespaces -H "Authorization: Bearer ${id token}"
```



---

#  Reference : 

https://www.keycloak.org/getting-started/getting-started-kube   

https://eclipse.dev/che/docs/stable/administration-guide/installing-che-on-minikube-keycloak-oidc/   

https://help.aliyun.com/zh/ack/serverless-kubernetes/use-keycloak-oidc-provider-for-identity-authentication  
