# registry-nacos (*这是一个由社区驱动的项目*)

使用 **nacos** 作为 **Kitex** 的注册中心

##  这个项目应当如何使用?

### 服务端

**registry-nacos/example/server/main.go**

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/registry"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    // ...
)

func main() {
    // ... 
    
    r, err := registry.NewDefaultNacosRegistry()
    if err != nil {
        panic(err)
    }
   
    svr := echo.NewServer(new(EchoImpl), server.WithRegistry(r))
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }
    // ...
}

```

### 客户端

**registry-nacos/example/client/main.go**

****

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/resolver"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    // ...
)

func main() {
    // ... 
   
    r, err := resolver.NewDefaultNacosResolver()
    if err != nil {
       panic(err) 
    }
   
    client, err := echo.NewClient("echo", client.WithResolver(r))
    if err != nil {
        log.Fatal(err)
    }
    // ...
}
```

## **环境变量**

| 变量名 | 变量默认值 | 作用 |
| ------------------------- | ---------------------------------- | --------------------------------- |
| serverAddr               | 127.0.0.1                          | nacos 服务器地址 |
| serverPort               | 8848                               | nacos 服务器端口            |
| namespace                 |                                    | nacos 中的 namespace Id |

## 兼容性

Nacos 2.0 和 1.X 版本的 nacos-sdk-go 是完全兼容的，[详情](https://nacos.io/en-us/docs/2.0.0-compatibility.html)

主要贡献者： [baiyutang](https://github.com/baiyutang)

