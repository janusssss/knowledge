# subernetes

## kubectl 命令

```bash
kubectl <command> [flag...]

Commands:
  get           Display one or many resources
  edit          Edit a resource on the server
  delete        Delete resources by filenames, stdin, resources and names, or by resources and label selector
  
flags:
	-o <wode|json|yaml>	选择输出格式
	-n <namespace>	命令空间
```

### get

```bash
kubectl get [command]

commands:
	svc <service-name>	查看指定服务的相关信息
```

#### pods

```bash
kubectl get pods [flag]

flags:
	-n <namespace>	命令空间
	-l --selector <selector>	根据上设置的键值对标签来选择特定的资源
```

```bash
kubectl exec -it -n az3 gostack-mongos-0 -- /bin/sh #进入pod内部
```

```bash
MongoConf:
  EndPoint: "gostack-mongos-0.gostack-mongos.az3.svc.cluster.net:27017"
  EndPoints:
    - "gostack-mongos-0.gostack-mongos.az3.svc.cluster.net:27017" 
  Username: "root"
  Password: "testaz3"
  Timeout: 300

```

```bash
mongo -u "root" -p "test" --authenticationDatabase "admin"
```

#### deployments

```bash
kubectl get deployments -n az2
```

#### services

```bash
```



### cp

```bash
kubectl cp gostack-mongos-0:/home/instance /path/to/local/directory -n az1
```

```bash
kubectl cp <source> <destination> [-n namespace]

<source> <destination>:
<namespace>/<pod-name>[:<container-name>]:<path-to-file-or-directory>
```

```bash
kubectl get deployments -n az2 | grep gostack-api
```



```bash
kubectl set image deployment/<deployment-name> <container-name>=<new-image>:<tag> -n az2
```



### rollout

```bash
kubectl rollout undo deployment/<deployment-name> -n az2
# 指定的 Deployment 回滚到之前的修订版本
```

```bash
1. 可以走单纯的打包流水线，手动去部署（部署前群里面同步一下，我看你还没有进，我明天催一下）
2. 星空镜像地址可以直接在内蒙环境拉取，电脑上拉取不了
3. 下面是部署的参考命令，你根据你要部署的组件去替换镜像地址；


### gostack_proxy
kubectl set image deployment/cn-nm-region1-az1-gostack-proxy gostack-proxy=docker.ctyun.cn:60001/gostack/gostack_proxy:202310292111-dev-amd64 -n az1

### gostack_api
repos-snapshot.ctyun.cn/compute/gostack/gostack:20250106131325-18ad53da92-amd64
kubectl set image deployment/cn-nm-region1-az1-gostack-api  api=docker.ctyun.cn:60001/gostack/gostack:202311022100-dev-amd64 -n az1
### gostack_scheduler
kubectl set image deployment/cn-nm-region1-az1-gostack-scheduler  scheduler=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1

### gostack_engine
kubectl set image deployment/cn-nm-region1-az1-gostack-engine  engine=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1
kubectl set image deployment/cn-nm-region1-az1-gostack-engine-2 engine=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1
kubectl set image deployment/cn-nm-region1-az1-gostack-engine-3 engine=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1

### gostack_其它
kubectl set image deployment/cn-nm-region1-az1-gostack-cron  cron=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1
kubectl set image deployment/cn-nm-region1-az1-gostack-event event=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1
kubectl set image deployment gostack-meta meta-api=docker.ctyun.cn:60001/gostack/gostack:202309131611-live-resize-amd64 -n az1


### gostack_agent，注意要升级哪些agent更新对应的daemon set
kubectl set image daemonset cn-nm-region1-az1-gostack-agent agent=docker.ctyun.cn:60001/gostack/agent:202311011819-dev-amd64 -n az1
kubectl set image daemonset gostack-agent agent=docker.ctyun.cn:60001/gostack/agent:202311011819-dev-amd64 -n az1

kubectl get pod -n az1 | grep cn-nm-region1-az1-gostack-agent | awk '{print $1}' | xargs kubectl delete pod -n az1
kubectl get pod -n az1 | awk '/^gostack-agent-[1-9,a-z]/' |grep -vE "aarch64|s8-|ir3-|-hx|haiguang"| awk '{print $1}' | xargs kubectl delete pod -n az1
```

### logs

```bash
kubectl logs -l <label-selector> -n <namespace>
# 你可以结合 -l 参数和 kubectl logs 来查看所有符合标签选择器的 Pods 的日志：

--previous
```

```bash
kubectl exec -it gostack-mongos-0 -n az2 -- /bin/bash
```

### describe

```bash
kubectl describe 
```

### exec

```bash
kubectl exec -it cn-nm-region1-az2-gostack-agent-mmfw4 -n az2 -- /bin/sh
```



# docker

## 命令

| 命令                        | 作用                             |
| --------------------------- | -------------------------------- |
| docker ps                   | 列出当前正在运行的容器           |
| docker images               | 列出本地机器上的所有 Docker 镜像 |
| docker rm <NAMES>           | 删除一个或多个已停止的容器       |
| docker rmi                  | 删除一个或多个镜像               |
| docker build  -t name:tag . | 利用Dockerfile文件打包镜像       |
| docker start <ID or NAME>   | 启动容器                         |
| docker stop <containerName> | 停止容器运行                     |

```bash
Usage:	docker [OPTIONS] COMMAND

A self-sufficient runtime for containers

Options:
      --config string      Location of client config files (default "/root/.docker")
  -D, --debug              Enable debug mode
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
      --tls                Use TLS; implied by --tlsverify
      --tlscacert string   Trust certs signed only by this CA (default "/root/.docker/ca.pem")
      --tlscert string     Path to TLS certificate file (default "/root/.docker/cert.pem")
      --tlskey string      Path to TLS key file (default "/root/.docker/key.pem")
      --tlsverify          Use TLS and verify the remote
  -v, --version            Print version information and quit

Management Commands:
  builder     Manage builds
  config      Manage Docker configs
  container   Manage containers
  engine      Manage the docker engine
  image       Manage images
  network     Manage networks
  node        Manage Swarm nodes
  plugin      Manage plugins
  secret      Manage Docker secrets
  service     Manage services
  stack       Manage Docker stacks
  swarm       Manage Swarm
  system      Manage Docker
  trust       Manage trust on Docker images
  volume      Manage volumes

Commands:
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  events      Get real time events from the server
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  history     Show the history of an image
  images      List images
  import      Import the contents from a tarball to create a filesystem image
  info        Display system-wide information
  inspect     Return low-level information on Docker objects
  kill        Kill one or more running containers
  load        Load an image from a tar archive or STDIN
  login       Log in to a Docker registry
  logout      Log out from a Docker registry
  logs        Fetch the logs of a container
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  ps          List containers
  pull        Pull an image or a repository from a registry
  push        Push an image or a repository to a registry
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  rmi         Remove one or more images
  run         Run a command in a new container
  save        Save one or more images to a tar archive (streamed to STDOUT by default)
  search      Search the Docker Hub for images
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  tag         Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  version     Show the Docker version information
  wait        Block until one or more containers stop, then print their exit codes

Run 'docker COMMAND --help' for more information on a command.
```



### build

```bash
docker build -f MyDockerfile -t myapp:latest .
```

```bash
docker build [OPTIONS] PATH | URL | -
# PATH：指定包含Dockerfile和所有需要的文件的目录路径。
# URL：可以是一个Git仓库地址，Docker会克隆该仓库（以及子模块）并使用其作为上下文。
# -：通过标准输入(stdin)发送一个tar压缩包或者一个文本文件。
```



### docker exec

```bash
docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
# 在已经运行的 Docker 容器中执行命令
docker exec -it CONTAINER_NAME_OR_ID /bin/bash
#  进入容器并启动一个交互式的 Bash shell
```

- options
  - -d 或 --detach: 在后台运行命令
  - -i 或 --interactive: 保持 STDIN 打开，即使没有附加
  - -t 或 --tty: 分配一个伪TTY
  - -u 或 --user: 指定运行命令的用户
  - -e 或 --env: 设置环境变量

### logs

```bash
1f70364ba09eFlag shorthand -h has been deprecated, please use --help

Usage:	docker logs [OPTIONS] CONTAINER

Fetch the logs of a container

Options:
      --details        Show extra details provided to logs
  -f, --follow         Follow log output
      --since string   Show logs since timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g. 42m for 42 minutes)
      --tail string    Number of lines to show from the end of the logs (default "all")
  -t, --timestamps     Show timestamps
      --until string   Show logs before a timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g. 42m for 42 minutes)
```

### inspect

```bash
Flag shorthand -h has been deprecated, please use --help

Usage:	docker inspect [OPTIONS] NAME|ID [NAME|ID...]

Return low-level information on Docker objects

Options:
  -f, --format string   Format the output using the given Go template
  -s, --size            Display total file sizes if the type is container
  -t, --time int        Seconds to wait for inspect timeout (default 120)
      --type string     Return JSON for specified type
```

### run

```bash
docker run -d \
 --name janus \
 --network host \
 -v /etc/nova/nova.conf:/etc/nova/nova.conf \
 36a6b0a1416b
```



```bash

Usage:	docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container

Options:
      --add-host list                  Add a custom host-to-IP mapping (host:ip)
      --annotation list                Set annotations on a container
  -a, --attach list                    Attach to STDIN, STDOUT or STDERR
      --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
      --blkio-weight-device list       Block IO weight (relative device weight) (default [])
      --cap-add list                   Add Linux capabilities
      --cap-drop list                  Drop Linux capabilities
      --cgroup-parent string           Optional parent cgroup for the container
      --cidfile string                 Write the container ID to the file
      --cpu-period int                 Limit CPU CFS (Completely Fair Scheduler) period
      --cpu-quota int                  Limit CPU CFS (Completely Fair Scheduler) quota
      --cpu-rt-period int              Limit CPU real-time period in microseconds
      --cpu-rt-runtime int             Limit CPU real-time runtime in microseconds
  -c, --cpu-shares int                 CPU shares (relative weight)
      --cpus decimal                   Number of CPUs
      --cpuset-cpus string             CPUs in which to allow execution (0-3, 0,1)
      --cpuset-mems string             MEMs in which to allow execution (0-3, 0,1)
  -d, --detach                         Run container in background and print container ID
      --detach-keys string             Override the key sequence for detaching a container
      --device list                    Add a host device to the container
      --device-cgroup-rule list        Add a rule to the cgroup allowed devices list
      --device-read-bps list           Limit read rate (bytes per second) from a device (default [])
      --device-read-iops list          Limit read rate (IO per second) from a device (default [])
      --device-write-bps list          Limit write rate (bytes per second) to a device (default [])
      --device-write-iops list         Limit write rate (IO per second) to a device (default [])
      --disable-content-trust          Skip image verification (default true)
      --dns list                       Set custom DNS servers
      --dns-option list                Set DNS options
      --dns-search list                Set custom DNS search domains
      --entrypoint string              Overwrite the default ENTRYPOINT of the image
  -e, --env list                       Set environment variables
      --env-file list                  Read in a file of environment variables
      --expose list                    Expose a port or a range of ports
      --files-limit int                Tune container files limit (set -1 for unlimited)
      --group-add list                 Add additional groups to join
      --health-cmd string              Command to run to check health
      --health-exit-on-unhealthy       Shut down a container if it becomes Unhealthy
      --health-interval duration       Time between running the check (ms|s|m|h) (default 0s)
      --health-retries int             Consecutive failures needed to report unhealthy
      --health-start-period duration   Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
      --health-timeout duration        Maximum time to allow one check to run (ms|s|m|h) (default 0s)
      --help                           Print usage
      --hook-spec string               file containing hook definition(prestart, poststart, poststop)
  -h, --hostname string                Container host name
      --hugetlb-limit hugetlb          Huge page limit (format: [size:]<limit>, e.g. --hugetlb-limit 2MB:32MB) (default [])
      --init                           Run an init inside the container that forwards signals and reaps processes
  -i, --interactive                    Keep STDIN open even if not attached
      --ip string                      IPv4 address (e.g., 172.30.100.104)
      --ip6 string                     IPv6 address (e.g., 2001:db8::33)
      --ipc string                     IPC mode to use
      --isolation string               Container isolation technology
      --kernel-memory bytes            Kernel memory limit
  -l, --label list                     Set meta data on a container
      --label-file list                Read in a line delimited file of labels
      --link list                      Add link to another container
      --link-local-ip list             Container IPv4/IPv6 link-local addresses
      --log-driver string              Logging driver for the container
      --log-opt list                   Log driver options
      --mac-address string             Container MAC address (e.g., 92:d0:c6:0a:29:33)
  -m, --memory bytes                   Memory limit
      --memory-reservation bytes       Memory soft limit
      --memory-swap bytes              Swap limit equal to memory plus swap: '-1' to enable unlimited swap
      --memory-swappiness int          Tune container memory swappiness (0 to 100) (default -1)
      --mount mount                    Attach a filesystem mount to the container
      --name string                    Assign a name to the container
      --network string                 Connect a container to a network (default "default")
      --network-alias list             Add network-scoped alias for the container
      --no-healthcheck                 Disable any container-specified HEALTHCHECK
      --oom-kill-disable               Disable OOM Killer
      --oom-score-adj int              Tune host's OOM preferences (-1000 to 1000)
      --pid string                     PID namespace to use
      --pids-limit int                 Tune container pids limit (set -1 for unlimited)
      --privileged                     Give extended privileges to this container
  -p, --publish list                   Publish a container's port(s) to the host
  -P, --publish-all                    Publish all exposed ports to random ports
      --read-only                      Mount the container's root filesystem as read only
      --restart string                 Restart policy to apply when a container exits (default "no")
      --rm                             Automatically remove the container when it exits
      --runtime string                 Runtime to use for this container
      --security-opt list              Security Options
      --shm-size bytes                 Size of /dev/shm
      --sig-proxy                      Proxy received signals to the process (default true)
      --stop-signal string             Signal to stop a container (default "SIGTERM")
      --stop-timeout int               Timeout (in seconds) to stop a container
      --storage-opt list               Storage driver options for the container
      --sysctl map                     Sysctl options (default map[])
      --tmpfs list                     Mount a tmpfs directory
  -t, --tty                            Allocate a pseudo-TTY
      --ulimit ulimit                  Ulimit options (default [])
  -u, --user string                    Username or UID (format: <name|uid>[:<group|gid>])
      --userns string                  User namespace to use
      --uts string                     UTS namespace to use
  -v, --volume list                    Bind mount a volume
      --volume-driver string           Optional volume driver for the container
      --volumes-from list              Mount volumes from the specified container(s)
  -w, --workdir string                 Working directory inside the container
```

### image

#### list

```bash

List images

Aliases:
  ls, images, list

Options:
  -a, --all             Show all images (default hides intermediate images)
      --digests         Show digests
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print images using a Go template
      --no-trunc        Don't truncate output
  -q, --quiet           Only show numeric IDs
```

### ps

```bash
Usage:	docker ps [OPTIONS]

List containers

Options:
  -a, --all             Show all containers (default shows just running)
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print containers using a Go template
  -n, --last int        Show n last created containers (includes all states) (default -1)
  -l, --latest          Show the latest created container (includes all states)
      --no-trunc        Don't truncate output
  -q, --quiet           Only display numeric IDs
  -s, --size            Display total file sizes
```

### start

```bash
docker start [OPTIONS] CONTAINER [CONTAINER...]

Start one or more stopped containers

Options:
  -a, --attach               Attach STDOUT/STDERR and forward signals
      --detach-keys string   Override the key sequence for detaching a container
  -i, --interactive          Attach container's STDIN
```

### restart

```bash
docker restart [OPTIONS] CONTAINER [CONTAINER...]

Restart one or more containers

Options:
  -t, --time int   Seconds to wait for stop before killing the container (default 10)
```

### rm

```bash
Usage:	docker rm [OPTIONS] CONTAINER [CONTAINER...]

Remove one or more containers

Options:
  -f, --force     Force the removal of a running container (uses SIGKILL)
  -l, --link      Remove the specified link
  -v, --volumes   Remove the volumes associated with the container
```

### cp

```bash
docker cp manager.py nova_compute:/usr/lib/python2.7/site-packages/nova/compute/manager.py
```

```bash
docker cp [option...] source destination	# 从宿主机复制文件到容器内部，或相反路径
container_path: <container_name:path>
Options:
  -a, --archive       Archive mode (copy all uid/gid information)
  -L, --follow-link   Always follow symbol link in SRC_PATH
```

### tag

```bash
docker tag emulator:1.0.0 repos-devops.ctyun.cn/starsky-base/compute/go1.24-emulator:1.2.0
```

```bash
docker tag <源镜像>[:<源标签>] <目标仓库>/<目标镜像>:<目标标签>
```

### save

```bash
# # 将镜像保存为tar文件
docker save -o myimage.tar myimage:latest
```

### load

```bash
docker load -i myimage.tar
# -i: 选项用于指定输入文件的路径
```

## Dockerfile

```bash
 "Mounts": [
            {
                "Type": "volume",
                "Name": "libvirtd",
                "Source": "/var/lib/docker/volumes/libvirtd/_data",
                "Destination": "/var/lib/libvirt",
                "Driver": "local",
                "Mode": "rw",
                "RW": true,
                "Propagation": ""
            },
            {
                "Type": "bind",
                "Source": "/lib/modules",
                "Destination": "/lib/modules",
                "Mode": "ro",
                "RW": false,
                "Propagation": "rprivate"
            },
            {
                "Type": "volume",
                "Name": "iscsi_info",
                "Source": "/var/lib/docker/volumes/iscsi_info/_data",
                "Destination": "/etc/iscsi",
                "Driver": "local",
                "Mode": "rw",
                "RW": true,
                "Propagation": ""
            },
            {
                "Type": "bind",
                "Source": "/dev",
                "Destination": "/dev",
                "Mode": "rw",
                "RW": true,
                "Propagation": "rprivate"
            },
            {
                "Type": "bind",
                "Source": "/usr/bin/vhost_disk",
                "Destination": "/usr/bin/vhost_disk",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            },
            {
                "Type": "volume",
                "Name": "kolla_logs",
                "Source": "/var/lib/docker/volumes/kolla_logs/_data",
                "Destination": "/var/log",
                "Driver": "local",
                "Mode": "rw",
                "RW": true,
                "Propagation": ""
            },
            {
                "Type": "volume",
                "Name": "nova_compute",
                "Source": "/var/lib/docker/volumes/nova_compute/_data",
                "Destination": "/var/lib/nova",
                "Driver": "local",
                "Mode": "rw",
                "RW": true,
                "Propagation": ""
            },
            {
                "Type": "bind",
                "Source": "/etc/localtime",
                "Destination": "/etc/localtime",
                "Mode": "ro",
                "RW": false,
                "Propagation": "rprivate"
            },
            {
                "Type": "bind",
                "Source": "/run",
                "Destination": "/run",
                "Mode": "shared",
                "RW": true,
                "Propagation": "shared"
            },
            {
                "Type": "bind",
                "Source": "/etc/kolla/nova-compute",
                "Destination": "/var/lib/kolla/config_files",
                "Mode": "ro",
                "RW": false,
                "Propagation": "rprivate"
            },
            {
                "Type": "bind",
                "Source": "/usr/etc/vhost_disk",
                "Destination": "/usr/etc/vhost_disk",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
        ],

```

### docker-compose
#### up
```bash
docker-compose up -d
```
```bash
docker-compose up [options] [SERVICE...]
-d 后台运行
```