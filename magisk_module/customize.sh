SKIPUNZIP=0

ui_print "TapFlow Project -> MSLG losetup_go install script "

rootfs_dir="/data/rootfs"
work_dir="/data/rootfs/losetup.sh-go"
usr_img="/data/rootfs/losetup.sh-go/usr.img"

if [ ! -d "$rootfs_dir" ]; then
    abort "You're trying to install this module on unsupported devices! Abort."

fi

if [ ! -d "$work_dir" ]; then
    echo "$work_dir not exist,creating.."
    mkdir $work_dir 
fi

if [ ! -f "$usr_img" ]; then
    echo "making new usr.img"
    truncate -s 1099511627776 /data/rootfs/losetup.sh-go/usr.img # 1T
    mkfs.ext4 /data/rootfs/losetup.sh-go/usr.img
fi


ui_print "安装成功,本模块将在设备重启后生效。"
ui_print "Success! This module will take effect after the device is restarted."
