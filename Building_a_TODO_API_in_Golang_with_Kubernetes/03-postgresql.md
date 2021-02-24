#  INSTALL  postgresql

## INSTALL

```
arkade install postgresql

```

##  CONNECT

```

export POSTGRES_PASSWORD=$(kubectl get secret --namespace default postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)

```


```

kubectl run postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:11.6.0-debian-9-r0 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host postgresql -U postgres -d postgres -p 5432

```


## edit database

```
CREATE TABLE todo (
  id              INT GENERATED ALWAYS AS IDENTITY,
  description     text NOT NULL,
  created_date    timestamp NOT NULL,
  completed_date  timestamp NOT NULL
);

```


```

 \dt
 
```


##  退出客户端
```

\q


```
