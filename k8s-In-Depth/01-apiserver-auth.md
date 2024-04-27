# 01 使用 `curl` 访问 Kubernetes API  

```
minikube start
```

```
cat ~/.kube/config
```



```bash
export API_SERVER_URL=https://127.0.0.1:32784
```

访问：  

```bash
curl $API_SERVER_URL/api/v1/namespaces
```

输出：   

```bash
curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

SSL问题，加上参数 -k 

```bash
curl $API_SERVER_URL/api/v1/namespaces
```

输出：  

```json
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "namespaces is forbidden: User \"system:anonymous\" cannot list resource \"namespaces\" in API group \"\" at the cluster scope",
  "reason": "Forbidden",
  "details": {
    "kind": "namespaces"
  },
  "code": 403
}%
```

```
minikube delete
```

# 02 auth with static token 

Create dir:  

```
mkdir -p ~/.minikube/files/etc/ca-certificates
cd ~/.minikube/files/etc/ca-certificates

```



create file :    

```
cat <<EOF > tokens.csv
token1,arthur,1,"admin,dev,qa"
token2,daniele,2,dev
token3,errge,3,qa
EOF
```



Start minikube with:   

```
ca-certificates minikube start \
  --extra-config=apiserver.token-auth-file=/etc/ca-certificates/tokens.csv
```

再次查看：   

```
cat ~/.kube/config
```

设置：  

```
export APISERVER=https://127.0.0.1:57761
```

```
export CACERT=/Users/learnk8s/.minikube/ca.crt
```

```bash
curl --cacert ${CACERT} -X GET ${APISERVER}/api/v1/namespaces
```

结果:

```json
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "namespaces is forbidden: User \"system:anonymous\" cannot list resource \"namespaces\" in API group \"\" at the cluster scope",
  "reason": "Forbidden",
  "details": {
    "kind": "namespaces"
  },
  "code": 403
}%
```



```bash
curl --cacert ${CACERT} --header "Authorization: Bearer token1" -X GET ${APISERVER}/api/v1/namespaces
```

```json
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "namespaces is forbidden: User \"arthur\" cannot list resource \"namespaces\" in API group \"\" at the cluster scope",
  "reason": "Forbidden",
  "details": {
    "kind": "namespaces"
  },
  "code": 403
}%
```

说明用户已经识别出来了，但是权限不够， 怎么办？   

怎么办？  ->   加role    

```bash
cat <<EOF > rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin
subjects:
- kind: Group
  name: admin
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF
```

