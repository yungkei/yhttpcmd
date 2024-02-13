# yhttpcmd
[![Software License](https://img.shields.io/github/license/yungkei/yhttpcmd)](LICENSE)
![Release](https://img.shields.io/github/v/release/yungkei/yhttpcmd?logo=github)
![Docker](https://img.shields.io/docker/v/yungkei/yhttpcmd?logo=docker)

yhttpcmd is a tiny cmd2http(convert command as http service) server, used to execute local commands via http.

## Install
There are various ways of installing yhttpcmd.

### Go 
```bash
go install github.com/yungkei/yhttpcmd@master
```

### Docker Images
Docker images are available on [Docker Hub](https://hub.docker.com/r/yungkei/yhttpcmd/).

You can launch a yhttpcmd container for trying it out with

```bash
docker run --name yhttpcmd -d -p 8080:8080 yungkei/yhttpcmd
```

yhttpcmd will now be reachable at <http://localhost:8080/>.

It is worth noting that, yhttpcmd only converts the commands set in the configuration into HTTP services.


## Usage

### Linux/Windows/MacOS
```bash
yhttpcmd start --config=yhttpcmd.yaml
```

### Docker
```dockerfile
FROM yungkei/yhttpcmd
# overwrite default configuration file
COPY yhttpcmd.yaml /yhttpcmd/yhttpcmd.yaml
CMD [ "./yhttpcmd","start" ]
```

### Configuration
```yaml
server:
  port: 8080
command_configs:
  - command: echo
    route: echo
  - command: <cmd>
    route: <route>
```

### Execute commands via HTTP
``` bash
curl -H "Content-Type:application/json" -d "{\"args\":\"hello\"}" -X POST 127.0.0.1:8080/echo
{"Command":"/bin/echo hello","Message":"hello\n"}
```

## License

Apache License 2.0, see [LICENSE](https://github.com/yungkei/yhttpcmd/blob/main/LICENSE).