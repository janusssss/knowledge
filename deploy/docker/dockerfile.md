# 概述
Dockerfile 是一个文本文件，其中包含了构建 Docker 镜像所需的指令。
通过Dockerfile，你可以定义应用程序的运行环境、依赖项以及如何运行你的应用程序。
Docker 可以读取这个文件并自动执行其中的指令来构建一个镜像。 

- 基础镜像（Base Image）：这是你将要构建的新镜像的基础。通常是一个官方的、轻量级的操作系统镜像，比如 ubuntu、alpine 或 node 等。 
- 指令（Instructions）：这些是用于构建镜像的命令，例如 FROM、RUN、COPY、ADD、CMD、EXPOSE 等。每个指令都会创建一个新的镜像层。 
- 标签（Labels）：可以为镜像添加元数据，比如版本、维护者等信息。 
- 构建上下文（Build Context）：Docker 构建镜像时会将当前目录下的文件和目录作为上下文发送给 Docker 守护进程。你可以使用 .dockerignore 文件来排除不需要的文件。 


