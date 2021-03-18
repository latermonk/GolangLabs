#  arkade install openfaas

```
arkade install openfaas
```


```
arkade info openfaas
```



-----

Info for app: openfaas
# Get the faas-cli
curl -SLsf https://cli.openfaas.com | sudo sh

# Forward the gateway to your machine
kubectl rollout status -n openfaas deploy/gateway
kubectl port-forward -n openfaas svc/gateway 8080:8080 &

# If basic auth is enabled, you can now log into your gateway:
PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
echo -n $PASSWORD | faas-cli login --username admin --password-stdin



#  Open-faas  UI with admin & $PASSWORD


![openfaas-ui](_image/openfaas-ui.jpg)


# Deploy an App

```
faas-cli store deploy figlet

faas-cli list

```



# Find out more at:
# https://github.com/openfaas/faas
