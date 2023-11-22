mkdir /dev/msl
mkdir /dev/msl/rdp
mount --b /dev /data/vendor/mslg/rootfs/dev
mount --b /sys /data/vendor/mslg/rootfs/sys
mount --b /proc /data/vendor/mslg/rootfs/proc
mount --b /dev/msl/rdp /data/vendor/mslg/rootfs/tmp/msl/rdp
mount --b /storage/self/primary /data/vendor/mslg/rootfs/tablet