# alauda/go-redis-client

## 简介

alauda/go-redis-client 是在redis官方包上进行的二次封装。它适用于容器环境,能自动的从容器的环境变量、挂载的配置文件获取参数并创建redis-client实例。方便开发人员

## 使用

### 安装以及使用

#### 1. 对于Go开发人员，你需要安装package

```shell
go get github.com/alauda/go-redis-client
```

#### 2. 创建redis-client

根据不同的安全需求、运维定制,您有四种方式创建redis-client

* 自配置参数,创建redisclient

```go
//此方法主要应用与测试环境,生产环境不建议使用
import "redis" github.com/alauda/go-redis-client

func main() {
  // check options.go for more details
  opts := redis.RedisClientOptions{
    Type: 	  redis.ClientNormal,
    Hosts:    []string{"localhost:6379"},
    Password: "123456",
    Database: 0,
  }
  client := redis.NewRedisClient(opts)
  if err := client.Ping().Err(); err != nil {
    panic(err)
  }
  
  // Using cluster mode
  clusterOpts := redis.RedisClientOptions{
    Type:      redis.ClientCluster,
    Hosts:     []string{"localhost:7000","localhost:7001","localhost:7002"},
    Password:  "123456",
    Database:  0,
    // all keys with a prefix
    KeyPrefix: "my-app:",
  }
  clusterClient := redis.NewRedisClient(clusterOpts)
  if err := clusterClient.Ping().Err(); err != nil {
    panic(err)
  }
}
```

* 从环境变量中获取参数来创建redisclient

```go
import redis "github.com/alauda/go-redis-client"
func main(){
    //RWType:	
    //    OnlyWrite
    //    OnlyWrite
    //    ReadAndWrite
    r:= redis.OnlyRead
    client,err:= redis.AutoConfigRedisClientInEnv(r)
    if err!{
        panic(err)
    }
}
```

* 从配置文件中获取参数来创建redisclient

```go
import redis "github.com/alauda/go-redis-client"
func main(){
    //RWType:	
    //    OnlyWrite
    //    OnlyWrite
    //    ReadAndWrite
    r:= redis.OnlyRead
    client,err:= redis.AutoConfigRedisClientInVolume(r)
    if err!{
        panic(err)
    }
}
```

* 从配置文件和环境变量中获取参数来创建redisclient
优先级:环境变量>配置文件

```go
import redis "github.com/alauda/go-redis-client"
func main(){
    //RWType:	
    //    OnlyWrite
    //    OnlyWrite
    //    ReadAndWrite
    r:= redis.OnlyRead
    client,err:= redis.AutoConfigRedisClient(r)
    if err!{
        panic(err)
    }
}
```

### 配置定制

#### 1 对于容器维护人员,你需要配置Pod注入必要的环境变量

```shell
apiVersion: v1
...
spec:
  volumes:
    - name: redis-secret
      secret:
        secretName: redis-secret
  containers:
      ....
      volumeMounts:
        - name: redis-secret
          mountPath: "/etc/paas" # it must be this value
          readOnly: true
      env:
        - name: REDIS_TYPE_READER # Injected  redis variables must have ENV_PREFIX when ENV_PREFIX is not null
          value: "cluster"
        - name: REDIS_HOST_READER
          value: "1.1.1.1,1.1.1.2,1.1.1.3"
        - name: REDIS_PORT_READER
          value: "26379 26379 26379"
        - name: REDIS_DB_NAME_READER
          value: '0'
        - name: REDIS_DB_PASSWORD_READER
          value: "aiyijing"         # redis:passwd
        - name: REDIS_MAX_CONNECTIONS_READER
          value: "32"
        - name: REDIS_KEY_PREFIX_READER
          value: "aiyijing_"
        - name: REDIS_SKIP_FULL_COVER_CHECK_READER
          value: "false"
        - name: REDIS_TIMEOUT_READER
          value: '5'
```

#### 2 配置挂载的secret

* For Example:Secret 配置信息

```shell
apiVersion: v1
kind: Secret
metadata:
  name: redis-secret
type: Opaque
data:
  # 目前仅支持toml文件
  redis.toml: `base64`
```

* config.toml 明文  
如果需要配置WRITER集群,可以将READER后缀更改为WRITER

```shell
# normal为单例，cluster为集群
REDIS_TYPE_READER="cluster"

# ip，集群模式下IP会有多个，以逗号分割，IP数量和端口数量一致
REDIS_HOST_READER=["10.0.129.115","10.0.128.150","10.0.128.89"]

# 端口，集群模式下端口会有多个，以逗号分割，IP数量和端口数量一致
REDIS_PORT_READER=["26379","26379","26379"]

# redis的DB name，只有在单例模式下才管用，集群模式会被忽略
REDIS_DB_NAME_READER=0

# 密码,如果没有密码则传空
REDIS_DB_PASSWORD_READER="alauda_redis_passwd"

# 最大连接数, 默认为32
REDIS_MAX_CONNECTIONS_READER=32

# redis key的前缀
REDIS_KEY_PREFIX_READER="alauda_redis_passwd"

# 当用户屏蔽了CONFIG命令时,需要把这个值改为true，只有集群模式有这个变量
REDIS_SKIP_FULL_COVER_CHECK_READER=false

# redis连接和操作的超时时间为5秒
REDIS_TIMEOUT_READER=5
```

#### 详细配置请参考

[example](/example)  
