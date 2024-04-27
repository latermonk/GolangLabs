# 使用 `curl` 访问 Kubernetes API  

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

```jason
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

