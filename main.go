package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var err error
var ModulePath string

func Write2Log(log string) {
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println("[" + currentTime + "] " + log)
	filePath := "/dev/Tapflow/losetup_logs"
	writestring(filePath, "["+currentTime+"] "+log+"\n", true)
}

// mode true:追加
// false:覆盖
func writestring(filepath_underdev string, text string, mode bool) {
	if mode {
		file, _ := os.OpenFile(filepath_underdev, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()
		file.WriteString(text)
	} else {
		file, _ := os.OpenFile(filepath_underdev, os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()
		file.WriteString(text)
	}
}
func checkAndCreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
func checkerr(err error, step string) {
	if err != nil {
		Write2Log("error occured :(" + step + ") " + err.Error())
		SetCurnetPropMode("error occured :("+step+") ", 2)
		os.Exit(1)
	}
}
func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return ex
}
func RunCMD(Name string, ar ...string) (string, error) {
	cmd := exec.Command(Name, ar...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// 挂载普通img镜像
func MountLegacyImg(Type string, imgpath string, Dest string, isRo bool) error {
	args := []string{"-t", Type}
	if isRo {
		args = append(args, "-r")
	}
	args = append(args, imgpath, Dest)
	_, err := RunCMD("mount", args...)
	if err != nil {
		return err
	}
	return nil
}

// 挂载overlay镜像
func MountOverlayImg(lowerdir string, upperdir string, workdir string, Dst string) error {
	_, err = RunCMD("mount", "-t", "overlay", "overlay", "-o", "lowerdir="+lowerdir+",upperdir="+upperdir+",workdir="+workdir, Dst)
	if err != nil {
		return err
	}
	return nil
}

func modifyMagiskDescription(newDescription string) error {
	file, err := os.OpenFile(ModulePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// 读取文件内容
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	content := make([]byte, fileSize)
	_, err = file.Read(content)
	if err != nil {
		return err
	}
	contentStr := string(content)
	newContent := strings.Replace(contentStr, "description=", fmt.Sprintf("description=%s\n", newDescription), -1)
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(newContent))
	if err != nil {
		return err
	}
	return nil
}

// current:定义
// 0 :😋 正常
// 1 :🤔 等待
// 2 :😰 错误
func SetCurnetPropMode(msg string, current int) {
	if current == 0 {
		err := modifyMagiskDescription("[😋 losetup_go]:" + msg)
		UpdateCurrentMode("0")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if current == 1 {
		err := modifyMagiskDescription("[🤔 losetup_go]:" + msg)
		if err != nil {
			fmt.Println(err)
		}
		UpdateCurrentMode("1")
		return
	}
	if current == 2 {
		err := modifyMagiskDescription("[😰 losetup_go]:" + msg)
		UpdateCurrentMode("2")
		if err != nil {
			fmt.Println(err)
		}
	}
}
func Setprop(key string, value string) {
	Write2Log("Running setprop " + key + " " + value)
	RunCMD("setprop", key, value)
}

// current:定义
// 0 :😃 正常
// 1 :🤔 等待
// 2 :😰 错误
func UpdateCurrentMode(current string) {
	writestring("/dev/Tapflow/current", current, false)
}
func init() {
	checkAndCreateDir("/dev/Tapflow")
	writestring("/dev/Tapflow/version", "V2.0_20240512_Release", false)
	ModulePath = filepath.Dir(getExecutablePath()) + "/module.prop"
}
func chcon_folder(label, folderpath string) {
	str, _ := RunCMD("chcon", label, folderpath)
	Write2Log("restorerecon log:" + str)
}
func main() {
	work_path := "/data/rootfs/losetup.sh-go"
	Write2Log("-------------------")
	Write2Log("starting losetup for Tapflow project")
	//1.mount usr.img
	checkAndCreateDir(work_path)
	checkAndCreateDir(filepath.Join(work_path, "usr"))
	checkAndCreateDir(filepath.Join(work_path, "partition_ro"))
	checkAndCreateDir(filepath.Join(work_path, "partition_ro", "usr"))
	err = MountLegacyImg("ext4", filepath.Join(work_path, "usr.img"), filepath.Join(work_path, "usr"), false)
	checkerr(err, "mount legacy img")
	//create workdir and upperdir
	checkAndCreateDir(filepath.Join(work_path, "usr", "upper"))
	checkAndCreateDir(filepath.Join(work_path, "usr", "work"))
	//mount erofs mslgusrimg
	err = MountLegacyImg("erofs", "/odm/etc/assets/mslgusrimg", filepath.Join(work_path, "partition_ro", "usr"), true)
	checkerr(err, "mount(ro) usr from odm")
	//mount overlay usr.img
	err = MountOverlayImg(filepath.Join(work_path, "partition_ro", "usr"), filepath.Join(work_path, "usr", "upper"), filepath.Join(work_path, "usr", "work"), "/data/rootfs/usr")
	//no need to mount mslgkingsoftimg and mslgappsimg ,because /odm/bin/losetup.sh loaded
	//wait 5 secs and override system prop
	SetCurnetPropMode("Wait For 5 secs ", 1)
	time.Sleep(time.Duration(5) * time.Second)
	RunCMD("setprop", "vendor.mslg.mslgusrimg", "null")
	//set usr sec label
	chcon_folder("u:object_r:mslg_rootfs_file:s0", "/data/rootfs/usr/")
	Write2Log("finish.")
	SetCurnetPropMode("Finished! ", 0)
	Write2Log("-------------------")
}
