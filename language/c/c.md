# 交叉编译
```
# 下载对应内存项目，生成.config文件，执行下面语句
make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- -j$(nproc)
```

