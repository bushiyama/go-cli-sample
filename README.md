# go-cli-sample

CLIとしてGo書いたことなかったので、書いてみた。  
せっかくなので`urfave/cli`を使ってみた。  
※[urfave/cli/v2 usage](https://github.com/urfave/cli/blob/master/docs/v2/manual.md#getting-started)

## usage

CLIとして叩けるようにする手順。

1. go get.  
```$ GO111MODULE=on go get github.com/urfave/cli/v2```

1. make [hoge]/[hoge.go]

1. Install our command to the $GOPATH/bin directory:  
```$ go install```

1. Finally run our new command:  

```text
$ hoge
Hello!
```

1. cli also generates neat help text:

```text
$ hoge help
NAME:
    hoge - fuga!

USAGE:
    hoge [global options] command [command options] [arguments...]

COMMANDS:
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS
    --help, -h  show help (default: false)
```
