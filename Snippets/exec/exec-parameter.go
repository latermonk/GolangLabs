// 参数构建有讲究  每一段只能有一个参数 ， 不然的话 在 wait 阶段会报错  error: exit status 2


cmd := exec.Command("ansible-playbook", "-i /root/install/go-ansible/example/golang-lab/inventorydir", "/root/install/exec-go/redis.yaml")


//错误写法
cmd := exec.Command("ansible-playbook", "-i /root/install/go-ansible/example/golang-lab/inventorydir /root/install/exec-go/redis.yaml")
