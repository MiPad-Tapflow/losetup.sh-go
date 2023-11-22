SKIPUNZIP=0

ui_print "#####################"
ui_print "TapFlow Project  -> WPS losetup_go install script "
ui_print "By ljlVink"
ui_print "#####################"

extract "$ZIPFILE" 'losetup.sh' "$MODPATH"

# 检测/data/Tapflow_project/文件夹是否存在，不存在则创建
tapflow_project_dir="/data/Tapflow_project"
if [ ! -d "$tapflow_project_dir" ]; then
    mkdir -p "$tapflow_project_dir"
    echo "已创建目录 $tapflow_project_dir"
fi

if [ ! -f "/vendor/etc/assets/mslgusrimg" ]; then
    ui_print "??mslgusrimg does not exist"
    
    # Check if the file exists in /sdcard/Downloads/
    if [ -f "/sdcard/Downloads/mslgusrimg" || ! -f "/data/Tapflow_project/mslgusrimg"]; then
        # Check if the file already exists in /data/Tapflow_project/
        if [ ! -f "/data/Tapflow_project/mslgusrimg" ]; then
            ui_print "Copying mslgusrimg from /sdcard/Downloads/"
            cp /sdcard/Downloads/mslgusrimg /data/Tapflow_project/
            resize2fs -f /data/Tapflow_project/mslgusrimg 3G 
            ui_print "Modified mslgusrimg size -> 3G"
        else
            ui_print "mslgusrimg already exists in /data/Tapflow_project/"
        fi
    else
        ui_print "Error: mslgusrimg not found in /sdcard/Downloads/"
        abort
    fi
fi

if [ ! -f "/vendor/etc/assets/mslgoptimg" ]; then
    ui_print "??mslgoptimg does not exist"
    
    # Check if the file exists in /sdcard/Downloads/
    if [ -f "/sdcard/Downloads/mslgoptimg" || ! -f "/data/Tapflow_project/mslgoptimg"]; then
        # Check if the file already exists in /data/Tapflow_project/
        if [ ! -f "/data/Tapflow_project/mslgoptimg" ]; then
            ui_print "Copying mslgoptimg from /sdcard/Downloads/"
            cp /sdcard/Downloads/mslgoptimg /data/Tapflow_project/
            resize2fs -f /data/Tapflow_project/mslgoptimg 3G 
            ui_print "Modified mslgoptimg size -> 3G"
        else
            ui_print "mslgoptimg already exists in /data/Tapflow_project/"
        fi
    else
        ui_print "Error: mslgoptimg not found in /sdcard/Downloads/"
        abort
    fi
fi

if [ -f "/vendor/etc/assets/mslgusrimg" ] && [ -f "/vendor/etc/assets/mslgoptimg" ]; then
    if [ ! -f "/data/Tapflow_project/mslgusrimg" ]; then
        ui_print "Copying mslgusrimg"
        cp /vendor/etc/assets/mslgusrimg /data/Tapflow_project/
        resize2fs -f /data/Tapflow_project/mslgusrimg 3G 
        ui_print "Modified mslgusrimg size -> 3G"
    else
        ui_print "mslgusrimg already exists in /data/Tapflow_project/"
    fi

    if [ ! -f "/data/Tapflow_project/mslgoptimg" ]; then
        ui_print "Copying mslgoptimg"
        cp /vendor/etc/assets/mslgoptimg /data/Tapflow_project/
        resize2fs -f /data/Tapflow_project/mslgoptimg 3G 
        ui_print "Modified mslgoptimg size -> 3G"
    else
        ui_print "mslgoptimg already exists in /data/Tapflow_project/"
    fi
fi


ui_print "Success."
ui_print "本模块将在设备重启后生效。"
ui_print "This module will take effect after the device is restarted."
ui_print "Этот модуль вступит в силу после перезагрузки устройства."
