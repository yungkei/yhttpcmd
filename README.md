# yhttpcmd
[![Software License](https://img.shields.io/github/license/yungkei/yhttpcmd
)](LICENSE)

yhttpcmd is a tiny cmd2http(convert command as http service) server, used to execute local commands via http.

## Installation

```bash
go install github.com/yungkei/yhttpcmd@master
```

## Usage

### Linux/Windows/MacOS
```bash
yhttpcmd start --config=yhttpcmd.yaml
```

### Configuration
```yaml
server:
  port: 8080
command_configs:
  - command: echo
    route: echo
```

## License

yhttpcmd is released under the Apache 2.0 license.