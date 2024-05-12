SKIPUNZIP=0

ui_print "#####################"
ui_print "TapFlow Project  -> WPS losetup_go install script "
ui_print "By ljlVink"
ui_print "#####################"


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


ui_print "Success."
ui_print "本模块将在设备重启后生效。"
ui_print "This module will take effect after the device is restarted."
ui_print "Этот модуль вступит в силу после перезагрузки устройства."
