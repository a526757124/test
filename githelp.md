## 创建代码库

### git init 在当前目录新建一个 Git 代码库

### git init [project-name] 创建一个新的目录 git 代码库

### git clone [url] 下载一个项目和它的整个代码历史

## 配置

### git config --list 显示当前的 Git 配置

### git config -e [--global] 编辑 Git 配置文件

### 设置提交代码时的用户信息

#### git config [--global] user.name "[name]"

#### git config [--global] user.email "[email address]"

## 增加/删除文件

### git add [file1][file2] ... 添加指定文件到暂存区

### git add [dir] 添加指定目录到暂存区，包括子目录

### git add . 添加当前目录的所有文件到暂存区

### git rm [file1][file2] ... 删除工作区文件，并且将这次删除放入暂存区

## 代码提交

### git commit -m [message] 提交暂存区到仓库区

### git commit [file1][file2] ... -m [message] 提交暂存区的指定文件到仓库区

### git commit -v 提交时显示所有 diff 信息

## 分支

### git branch 列表所有本地分支

### git branch -r 列表所有远程分支

### git branch [branch-name] 创建一个新分支，但依然停留在当前分支

### git checkout -b [branch] 新建一个分支，并切换到该分支

### git checkout [branch-name] 切换到指定分支，并更新工作区

### git checkout - 切换到上一个分支

### git merge [branch] 合并指定分支到当前分支

### git branch -d [branch-name] 删除分支

### git push origin --delete [branch-name] 删除远程分支

### git branch -dr [remote/branch]

## 远程同步

### git remote update 更新远程仓库

### git fetch [remote] 下载远程仓库所有变动

### git remote -v 显示所有远程仓库

### git remote show [remote] 显示某个远程仓库信息

### git remote add [shortname][url] 增加一个新的远程仓库，并命名

### git pull [remote][branch] 取回远程仓库变化，并与本地分支合并

### git push [remote][branch] 上传本地到远程仓库

### git push [remote] --force 强行推送当前分支到远程仓库，即使有冲突

### git push [remote] --all 推送所有分支到远程仓库

## 标签

### git tag

### git tag [tag]

### git tag -d [tag]

### git push origin :refs/tags/[tagName]

### git show [tag]

### git push [remote][tag]

### gti push [remote] --tags

## Git 忽略规则匹配语法

在 .gitignore 文件中，每一行的忽略规则的语法如下：
空格不匹配任意文件，可作为分隔符，可用反斜杠转义

### #开头的文件标识注释，可以使用反斜杠进行转义

### ! 开头的模式标识否定，该文件将会再次被包含，如果排除了该文件的父级目录，则使用 ! 也不会再次被包含。可以使用反斜杠进行转义

### / 结束的模式只匹配文件夹以及在该文件夹路径下的内容，但是不匹配该文件

### / 开始的模式匹配项目跟目录

### 如果一个模式不包含斜杠，则它匹配相对于当前 .gitignore 文件路径的内容，如果该模式不在 .gitignore 文件中，则相对于项目根目录

### \*\* 匹配多级目录，可在开始，中间，结束

### ? 通用匹配单个字符

### [] 通用匹配单个字符列表

## .gitignore 规则不生效

### .gitignore 只能忽略那些原来没有被 track 的文件，如果某些文件已经被纳入了版本管理中，则修改.gitignore 是无效的。

### 解决方法就是先把本地缓存删除（改变成未 track 状态），然后再提交:

git rm -r --cached .
git add .
git commit -m 'update .gitignore'
