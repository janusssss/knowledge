### 打包命令

```bash
CGO_ENABLED=1 GOOS=windows GOARCH=ajanusjanusmd64 CC=x86_64-w64-mingw32-gcc fyne package -os windows

# 安卓
fyne package -os android --app-id com.janus.notes
fyne-cross android -app-id com.janus.notes
```