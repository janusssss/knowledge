### dnf

```bash
dnf <command> [flag]

commands:
	download                 Download package to current directory
		--resolve             resolve and download needed dependencies
  config-manager
  	--add-repo					添加仓库源
  repolist							展示配置的仓库源
	
	
```

### 添加仓库

```bash
sudo dnf config-manager --add-repo http://repo.mysql.com/yum/mysql-5.7-community/el/6/x86_64/
```

