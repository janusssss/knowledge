```bash
sudo apt-get install mkvtoolnix # ubunut安装
mkvinfo input.mkv | grep -A 5 -B 3 subtitles # 查看信息
mkvextract tracks input.mkv 2:subtitle.srt # 导出

{"log_id":1900919657949378115,"error_msg":"missing required parameter from","error_code":282003}
```

