package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Write2Log(log string){
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println("["+currentTime+"] "+log)
	checkAndCreateDir("/dev/Tapflow")
	filePath := "/dev/Tapflow/losetup_logs"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	file.WriteString("["+currentTime+"] "+log + "\n"); 
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
func GetFreeLoop() string {
	cmd, _ := RunCMD("losetup", "-f")
	return cmd
}
func GetProperty(prop string) string {
	result, _ := RunCMD("getprop", prop)
	return result
}
func Resetprop() {
	//清除存在的prop
	Write2Log("running reset 2 props")
	RunCMD("setprop", "vendor.mslg.mslgoptimg", "")
	RunCMD("setprop", "vendor.mslg.mslgusrimg", "")
}
func ResetLoop()bool{
	Write2Log("trying reseting loop")
	RunCMD("losetup", "-D")
	cmd,_:=RunCMD("losetup","-a")
	if strings.Contains(cmd,"mslg"){
		//卸载失败->可能需要重启
		Write2Log("ERROR:reseting failed!! Maybe need reboot!!") 
		return false
	}
	return true
}
func MakePartitionRW(){
	Write2Log("remount partirion in rw.")
	RunCMD("mount","-o","remount,rw","/data/vendor/mslg/rootfs/opt")
	RunCMD("mount","-o","remount,rw","/data/vendor/mslg/rootfs/usr")
}
func Setprop(key string, value string) {
	Write2Log("Running setprop "+ key+" "+value)
	RunCMD("setprop", key, value)
}
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
func GetOptimgPath() (string, bool) { //位置，是否只读
	if fileExists("/data/Tapflow_project/mslgoptimg"){
		return "/data/Tapflow_project/mslgoptimg", false
	}else{
		return "/vendor/etc/assets/mslgoptimg",true
	}
}
func GetUsrimgPath() (string, bool) { //位置，是否只读
	if fileExists("/data/Tapflow_project/mslgusrimg"){
		return "/data/Tapflow_project/mslgusrimg", false
	}else{
		return "/vendor/etc/assets/mslgusrimg",true
	}
}
func SetupLoop(loop string, path string, ro bool) {
	Write2Log("going to setup loop to "+loop+" "+path)
	if ro {
		cmd, _ := RunCMD("losetup", "-r", loop, path)
		fmt.Println("setup loop result:", cmd)
	} else {
		cmd, _ := RunCMD("losetup", loop, path)
		fmt.Println("setup loop result:", cmd)
	}
}
func main() {
	Write2Log("-------------------")
	Write2Log("starting losetup.sh")
	//一般是33或者34 重新挂载多了不好，需要重启
	if !ResetLoop(){
		os.Exit(1)
	}
	Resetprop()
	var optimgloop string
	var usrimgloop string
	for !strings.HasPrefix(optimgloop, "/dev/block/loop") {
		time.Sleep(1 * time.Second)
		optimgloop = GetFreeLoop()
	}
	Path, isro := GetOptimgPath()
	SetupLoop(optimgloop, Path, isro)
	Setprop("vendor.mslg.mslgoptimg", optimgloop)
	for !strings.HasPrefix(usrimgloop, "/dev/block/loop") {
		time.Sleep(1 * time.Second)
		usrimgloop = GetFreeLoop()
	}
	Path, isro = GetUsrimgPath()
	SetupLoop(usrimgloop, Path, isro)
	Setprop("vendor.mslg.mslgusrimg", usrimgloop)
	MakePartitionRW()
	Write2Log("finish.")
	Write2Log("-------------------")
}