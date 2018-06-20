## 创建代码库
### git init 在当前目录新建一个Git代码库
### git init [project-name] 创建一个新的目录git代码库
### git clone [url] 下载一个项目和它的整个代码历史

## 配置
### git config --list 显示当前的Git配置
### git config -e [--global] 编辑Git配置文件

### 设置提交代码时的用户信息
#### git config [--global] user.name "[name]"
#### git config [--global] user.email "[email address]"

## 增加/删除文件
### git add [file1] [file2] ... 添加指定文件到暂存区
### git add [dir] 添加指定目录到暂存区，包括子目录
### git add . 添加当前目录的所有文件到暂存区
### git rm [file1] [file2] ... 删除工作区文件，并且将这次删除放入暂存区