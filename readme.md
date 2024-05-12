# losetup.sh

IN WIP for Tapflow Project.

### 使用

这是`Tapflow Project` 对应的模块后端,使小米平板`mslg`新方案的`usr`(EROFS)分区可读写

与[KernelSU](https://github.com/tiann/KernelSU)采用类似的overlay方案和稀疏文件系统方案，使得usr分区可读写，并且没有分区容量限制。

### 开发/编译

打包(linux)

```sh
scripts/build_pack2mod
```