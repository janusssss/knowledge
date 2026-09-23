# 设置字符编码
:set fileencoding=utf-8
:set fileencoding=utf-8   " 转换为 UTF-8
:set fileencoding=cp936   " 转换为 GBK
:set fileencoding=gb18030 " 转换为 GB18030
:w                        " 保存生效

### 安装


```bash
sudo pacman -S vim
sudo pacman -S python-pynvim 
# python-pynvim 是一个 Python 库，它允许你在纯 Python 环境中与 Neovim（一种高性能的 Vim 替代品）进行交互
```

```bash
# 下载插件
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

```bash
# 编辑你的.vimrc，在它开头加入以下Vundle配置，并初始化：
call plug#begin('~/.vim/plugged')

" 插件列表
Plug 'fatih/vim-go'

call plug#end()
let g:go_use_gopls = 1

:PlugInstall # 安装 vim-go 及其他你可能添加的插件
:GoInstallBinaries # 安装该插件所需的 Go 工具
:PlugUpdate vim-go
```

```bash
# gopls 是官方提供的 Go 语言服务器，为现代编辑器提供高级功能，如代码补全、跳转到定义等
go install golang.org/x/tools/gopls@latest
```

```bash
# 打开你的 .vimrc 文件并添加以下内容来配置 Vim，使其更适合 Go 开发
" Set up sensible keybindings and go-specific settings
set nocompatible              " Required to get rid of some annoying warning messages later
filetype off                  " Close the default filetype setting at the start

" Load plugin configuration file for better management
let &runtimepath = append(&runtimepath, '~/.vim/bundle/')

" Set up sensible keybindings
let mapleader = ","            " Use comma as leader key
nnoremap <Leader>l :nohltablist<CR>  " Show only the current tab in terminal
nnoremap <Leader>e :enew<CR>        " Open a new buffer

" Load sensible settings (for better Vim experience)
call pathogen#runtime()

```

```bash
# 安装 vim-go 插件的配置
vim +GoInstallBinaries +qall
```



```bash
# 保存文件后运行
vim +BundleInstall +qall
```

### 快捷键

| 键位 | 作用             |
| ---- | ---------------- |
| `yy` | 复制             |
| `p`  | 粘贴             |
| `G`  | 光标跳到文件末尾 |
| `gg` | 光标跳到文件顶部 |
| `v`  | 进入可视行模式   |
| `d`  | 删除选中内容     |

### 设置

|                 |          |
| --------------- | -------- |
| `:set number`   | 显示行号 |
| `:set nonumber` | 隐藏行号 |
|                 |          |

