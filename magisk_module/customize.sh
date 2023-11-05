SKIPUNZIP=0

ui_print "#####################"
ui_print "TapFlow Project  -> WPS losetup_go install script "
ui_print "By ljlVink"
ui_print "本模块将在设备重启后生效。"
ui_print "This module will take effect after the device is restarted."
ui_print "Этот модуль вступит в силу после перезагрузки устройства."
ui_print "#####################"

extract "$ZIPFILE" 'losetup.sh' "$MODPATH"
