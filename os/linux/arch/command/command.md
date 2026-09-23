### pacman || yay

```bash
usage:  pacman <operation> [...]
operations:
    pacman {-h --help}
    pacman {-V --version}
    pacman {-D --database} <options> <package(s)>
    pacman {-F --files}    [options] [file(s)]
    pacman {-Q --query}    [options] [package(s)]
    pacman {-R --remove}   [options] <package(s)>
    pacman {-S --sync}     [options] [package(s)]
    pacman {-T --deptest}  [options] [package(s)]
    pacman {-U --upgrade}  [options] <file(s)>
```



```bash
pacman <operation> [package...]
operations:
  -h --help
  -V --version
  -R --remove  <package...)>
  -S --sync    [package...]
    -b, --dbpath <path>  set an alternate database location
    -c, --clean          remove old packages from cache directory (-cc for all)
    -d, --nodeps         skip dependency version checks (-dd to skip all checks)
    -g, --groups         view all members of a package group
    (-gg to view all groups and members)
    -i, --info           view package information (-ii for extended information)
    -l, --list <repo>    view a list of packages in a repo
    -p, --print          print the targets instead of performing the operation
    -q, --quiet          show less information for query and search
    -r, --root <path>    set an alternate installation root
    -s, --search <regex> search remote repositories for matching strings
    -u, --sysupgrade     upgrade installed packages (-uu enables downgrades)
    -v, --verbose        be verbose
    -w, --downloadonly   download packages but do not install/upgrade anything
    -y, --refresh        download fresh package databases from the server

    pacman {-D --database} <options> <package(s)>
    pacman {-F --files}    [options] [file(s)]
    pacman {-Q --query}    [options] [package(s)]
    pacman {-T --deptest}  [options] [package(s)]
    pacman {-U --upgrade}  [options] <
```