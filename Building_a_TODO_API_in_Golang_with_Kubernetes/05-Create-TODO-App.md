


#  Create App


##  pull the template
```
faas-cli template store pull golang-middleware
```


##  create the App
```
export PREFIX=ibackchina2018
export LANG=golang-middleware
export API_NAME=todo
faas-cli new --lang $LANG --prefix $PREFIX $API_NAME
```


##  edit code: ./todo/handler.go
**vim  ./todo/handler.go**


```
package function
import (
    "net/http"
    "encoding/json"
)
type Todo struct {
    Description string `json:"description"`
}
func Handle(w http.ResponseWriter, r *http.Request) {
    todos := []Todo{}
    todos = append(todos, Todo{Description: "Run faas-cli up"})
    res, _ := json.Marshal(todos)
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(res))
}
```

##  **code format**

```
gofmt -w -s ./todo/handler.go
```

##   deploy the code by building a new image, pushing it to the Docker Hub and deploying it to the cluster via the OpenFaaS API     


```

faas-cli up -f todo.yml --build-arg GO111MODULE=on


```



##   Test

```

curl http://127.0.0.1:8080/function/todo

{
 "description": "Run faas-cli up"
}


```

