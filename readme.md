# losetup.sh

IN WIP for Tapflow Project.

### 使用

这是 Tapflow Project对应的模块后端,采用golang编写，目的是改变小米平板容器挂载/vendor/的两个img分区。采用的抢先挂载的方法，使其具有读写功能。

关于面具/ksu模块安装：如果你的/vendor/assets/etc/不存在mslgoptimg和mslgusrimg，那么就会去复制/sdcard/Downloads的mslgoptimg和mslgusrimg，如果没有请提前准备好。如果你的机器，如yudi，已经在/vendor/assets/etc/存在了，那么不用准备额外的两个镜像文件，安装即可。安装成功重启之后在面具/ksu中可以看到当前模块的状态。

关于修改分区大小：如果你要对某虚拟分区(usr,opt)修改大小，那么在/data/Tapflow_project/中创建need_resize_usr或need_resize_opt (不建议缩小，每次应该比上次要大)，文件格式如`1G`,`2G`等，重启后自动生效。在Tapflow中可以看到相关占用。

### 开发/编译

打包(linux)
```sh
scripts/build_pack2mod
```

直接导入到设备运行(linux)
```sh
scripts/build
```

直接导入到设备运行(windows)
```sh
scripts/build_to_android.bat

```