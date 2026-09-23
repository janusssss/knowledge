```bash
# 
|xargs -I {} bash -c 'ping -c 1 "{}" >/dev/null 2>&1 || echo "{}"'
```

