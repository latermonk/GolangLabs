#  k3s


https://k3d.io/

```
wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
```

```
curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
```

----




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



##  kubectl sheetcheet

```
apt install  bash-completion   

source <(kubectl completion bash) 
echo "source <(kubectl completion bash)" >> ~/.bashrc 
```

**vim  ~/.bashrc**

```
alias k=kubectl   
complete -F __start_kubectl k

```




**REFERENCE**

https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d           


