# 拉取子模块
git submodule update --init --recursive

# 清除保存的账号信息
```bash
git credential-manager reject https://github.com
git credential-manager reject https://gitlab.com
```

Git提交标题格式

- feat: 新增功能
- fix: 修复 bug
- docs: 文档更新
- style: 格式化代码，不影响功能
- refactor: 重构代码，不添加新功能也不修复 bug
- test: 添加或修改测试用例
- chore: 其他维护工作，如更新依赖包

### 修改最近一次commit的message

```bash
git commit --amend -m "fix# 快照统计功能调试"
```

### fatal: Cannot prompt because user interactivity has been disabled

- JetBrains

```bash
# 新版git默认使用token拉取，需要设置用户账号认证，勾选Use Credential helper
JetBrains --> file --> setting --> Version Control --> Git --> Update --> Use Credential helper
# 更新~/.gitconfig文件内容
[credential]
	username = username
	helper = manager-core
```



### 改变当前项目的put的用户信息
```bash
git config user.name "你的新用户名"
git config user.email "你的新邮箱"
```

在 Git 项目中查看当前配置的用户名有几种方法，你可以根据自己的需求选择最合适的一种。以下是常用的几种方式：
方法一：使用命令行

    查看全局设置的用户名：
    如果你想查看为所有 Git 操作设置的用户名（即全局配置），可以在命令行输入以下命令：
    
    bash
    深色版本

git config --global user.name

查看当前项目设置的用户名：
如果你在特定项目里设置了不同于全局设置的用户名，可以使用如下命令查看该项目级别的用户名配置：

bash
深色版本

    git config user.name
    
    或者，如果你想直接查看项目配置文件中的设置，可以直接打开项目根目录下的 .git/config 文件查找 [user] 部分。

方法二：查看项目配置文件

    直接打开你的 Git 项目的 .git 文件夹，找到并打开 config 文件。在这个文件中，你可以看到类似于下面的配置段落，它显示了当前项目的用户信息。
    
    深色版本
    
    [user]
        name = Your Name
        email = you@example.com

通过上述任意一种方法，你都可以轻松地查看当前 Git 项目所配置的用户名。如果你发现没有为特定项目设置用户名，那么 Git 会默认使用全局设置的用户名。

### git-credential-cache

```bash
# linux
git config --global credential.helper store # 凭证以明文形式永久存储在磁盘文件中, ~/.git-credentials
git config --global credential.helper cache #配置Git使用内存缓存来临时存储凭证信息
git config --global credential.helper 'cache --timeout=3600' # 设置过期时间
```

### manager-core

```bash
 # linux
 curl -LO https://github.com/git-ecosystem/git-credential-manager/releases/download/v2.6.1/gcm-linux_amd64.2.6.1.deb
 sudo apt install ./gcm-linux_amd64.2.6.1.deb
 git config --global credential.helper manager
 # 明文存储
 git config --global credential.credentialStore plaintext
```

```bash

### 仓库 URL 为 SSH

```bash
git remote set-url origin git@github.com:Jauns27149/bookkeeper.git
git remote set-url origin http://janus@10.19.79.176:8190/DataLake/ctc-datalake-dcp-repository.git
```


# token拉代码
```bash
git clone http://janus@10.19.79.176:8190/DL-GLC/ctc-dl-glc-safe-log-service-repository.git
git clone http://janus@token10.19.79.176:8190/DL-GLC/ctc-dl-glc-safe-log-service-repository.git
```
