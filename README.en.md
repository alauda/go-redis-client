# alauda/go-redis-client

## Introduction

alauda/go-redis-client is a secondary encapsulation on the official redis package. It is suitable for container environment. It can automatically obtain parameters from container environment variables and mounted configuration files and create redis-client instances. Convenient for developers

## Usage

### Installation and use in programs

#### 1. For Go developer,you need to install package

```shell
go get github.com/alauda/go-redis-client
```

#### 2. create redis-client in your programs

There are four ways to create redis-client based on security, operations and maintenance requirements

* Configuring parameters in programs and create redisclient

```go
//This method is mainly used in the test environment. It is not recommended in the production environment.
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

* Getting parameters from environment variables to create redisclient

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

* Get parameters from the configuration file to create redisclient

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

* Get parameters from configuration files and environment variables to create redisclient
Priority: Environment Variables > Configuration files

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

### Configuration customization

#### 1 For container maintainers, you need to configure Pod to inject the necessary environment variables

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
          mountPath: "/etc/pass" # it must be this value
          readOnly: true
      env:
        - name: REDIS_TYPE_READER # Injected  redis variables must have ENV_PREFIX when ENV_PREFIX is not null
          value: "cluster"
        - name: REDIS_HOST_READER
          value: "1.1.1.1 1.1.1.2 1.1.1.3"
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

#### 2 Configure secret and mount

* For Example:  
Secret Config info

```shell
apiVersion: v1
kind: Secret
metadata:
  name: redis-secret
type: Opaque
data:
  # only support toml
  redis.toml: `base64`
```

* config.toml text  
If you need to configure a WRITER cluster, you can change the READER suffix to WRITER

```shell
# normal and cluster
REDIS_TYPE_READER="cluster"

# ip，Separation with commas
REDIS_HOST_READER=["10.0.129.115","10.0.128.150","10.0.128.89"]

# port，Separation with commas，len(HOST)===len(PORT)
REDIS_PORT_READER=["26379","26379","26379"]

# redis's DB name，cluster mode is ignored.
REDIS_DB_NAME_READER=0

# passwd
REDIS_DB_PASSWORD_READER="alauda_redis_passwd"

# max connection , default:32
REDIS_MAX_CONNECTIONS_READER=32

# redis key's prefix
REDIS_KEY_PREFIX_READER="alauda_redis_passwd"

# Used to shield CONFIG commands,set true，Only the cluster mode has this variable.
REDIS_SKIP_FULL_COVER_CHECK_READER=false

# redis timeout default:5 sec
REDIS_TIMEOUT_READER=5
```

#### Refer to Detailed

[example](/example)  
