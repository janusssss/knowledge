# 手动启动lsp
```bash
:lua vim.lsp.start({ name = 'gopls', cmd = { 'gopls' }, root_dir = vim.fn.getcwd() })

lsp enable
```
