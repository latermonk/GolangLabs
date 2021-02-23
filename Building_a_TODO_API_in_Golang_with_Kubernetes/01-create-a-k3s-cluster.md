#  k3s


https://k3d.io/

```
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
```

k3d cluster create mycluster
```
k3d cluster create mycluster
```

##   You can now use it like this:
```
kubectl config use-context k3d-mycluster
kubectl cluster-info

```




```

kubectl get cs


kubectl get nodes
```

