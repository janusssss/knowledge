# 终端配置PIN
```fish
curl -k -X POST https://localhost:47990/api/pin \
        -u "janus":"27149" \
        -H "Content-Type: application/json" \
        -d '{"pin":"2643", "name":"surface"}'
{"status":true}⏎
```
