# 程序功能

该程序实现了以下三个功能：

- 安装软件包
- 初始化gdb
- 文件传输

# 安装软件包

通过调用 App/src/packages 中的 InstallAllPkgs 函数实现安装所有需要的软件包。在实现该功能之前，需要先调用 GetAllPkgs 函数获得需要安装的软件包列表。
# 初始化gdb

通过调用 gdbinit 函数初始化gdb，其具体实现为在C语言代码中调用 gdb() 函数，该函数执行以下操作：

    调用 popen 函数打开一个进程，执行 pwd 命令获取当前路径。
    构造要执行的命令行命令，其中包括调用Python脚本 gdbinit.py 并将其路径添加到命令行命令中。
    使用 system 函数执行命令行命令。

# 文件传输

通过调用 ScpMenu 函数实现文件传输功能。该函数通过交互式命令行菜单提供以下选项：

    上传文件
    下载文件
    退出

在选择上传或下载文件后，会提示用户输入相应的源文件路径和目标文件路径。然后，调用 src/packages 中的 ScpFile 函数完成文件传输。
主程序

主程序通过交互式命令行菜单提供了以上三个功能的选项。用户可以输入相应的选项编号来选择需要执行的功能。
C语言代码

程序中使用了一些C语言代码，包括在 gdbinit 函数中调用的 gdb 函数，以及在 gdb 函数中使用的 popen 和 system 函数。其中，popen 函数用于打开进程，system 函数用于执行命令行命令。这些C语言代码通过在代码文件中添加注释 // #include<stdio.h> 和 import "C" 来实现与Go语言的集成。
