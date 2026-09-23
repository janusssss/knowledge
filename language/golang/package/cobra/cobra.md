	`obra` 是一个非常流行的 Go 包，用于创建强大的现代命令行应用。它由 Steve Francia（spf13）开发，广泛应用于许多知名的项目中，例如 Hugo、Kubernetes 和 Docker 等。Cobra 提供了简单易用的 API  来定义命令、参数和标志，并且支持子命令结构，使得构建复杂的 CLI 工具变得轻而易举。

# 命令

## cobra-cli

```bash
Usage:
  cobra-cli [command]

Available Commands:
  add         Add a command to a Cobra Application
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra-cli
  -l, --license string   name of license for the project
      --viper            use Viper for configuration
```

### init 
​	`cobra` 是一个非常流行的 Go  包，用于创建强大的现代命令行应用。它由 Steve Francia（spf13）开发，广泛应用于许多知名的项目中，例如  Hugo、Kubernetes 和 Docker 等。Cobra 提供了简单易用的 API  来定义命令、参数和标志，并且支持子命令结构，使得构建复杂的 CLI 工具变得轻而易举。

# 安装

```bash
go get -u github.com/spf13/cobra/cobra			 	#项目安装Cobra库
go install github.com/spf13/cobra-cli@latest 	#Cobra命令行工具cobra-cli
```

# cobra-cli

```bash
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-cli [command]

Available Commands:
  add         Add a command to a Cobra Application
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra-cli
  -l, --license string   name of license for the project
      --viper            use Viper for configuration

Use "cobra-cli [command] --help" for more information about a command.
```

## init

```bash
cobra-cli init ./ -a janus
```

```bash
Initialize (cobra init) will create a new application, with a license
and the appropriate structure for a Cobra-based CLI application.

Cobra init must be run inside of a go module (please run "go mod init <MODNAME>" first)

Usage:
  cobra-cli init [path] [flags]

Aliases:
  init, initialize, initialise, create

Flags:
  -h, --help   help for init

Global Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -l, --license string   name of license for the project
      --viper            use Viper for configuration
```

## add

```bash
cobra-cli add hello 
```

```ba
Add (cobra add) will create a new command, with a license and
the appropriate structure for a Cobra-based CLI application,
and register it to its parent (default rootCmd).

If you want your command to be public, pass in the command name
with an initial uppercase letter.

Example: cobra add server -> resulting in a new cmd/server.go

Usage:
  cobra-cli add [command name] [flags]

Aliases:
  add, command

Flags:
  -h, --help            help for add
  -p, --parent string   variable name of parent command for this command (default "rootCmd")

Global Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -l, --license string   name of license for the project
      --viper            use Viper for configuration
```

