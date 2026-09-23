### 命令
```bash
ollama run deepseek-r1:14b
```

```bash
ollama serve --host 0.0.0.0
```

```bash
curl http://localhost:11434/api/chat -d '{
  "model": "qwen2.5:0.5b",
  "messages": [
    {"role": "user", "content": "你好，你是谁？"}
  ],
  "stream": false
}'


curl http://192.168.169.186:11434/api/chat -d '{
  "model": "qwen2.5:0.5b",
  "messages": [
    {"role": "user", "content": "你好，你是谁？"}
  ],
  "stream": false
}'
```

```bash
docker run -d -v ollama:/root/.ollama -p 11444:11434 --name ollama ollama/ollama
docker exec -it ollama ollama run qwen2.5:0.5b

curl http://172.20.70.253:11444/api/chat -d '{
  "model": "qwen2.5:0.5b",
  "messages": [
    {"role": "user", "content": "你好，你是谁？"}
  ],
  "stream": false
}'

```

