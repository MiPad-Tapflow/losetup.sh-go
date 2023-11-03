package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

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
func getProperty(prop string) string {
	result, _ := RunCMD("getprop", prop)
	return result
}
func Resetprop() {
	//清除存在的prop
	fmt.Println("running reset 2 props")
	RunCMD("setprop", "vendor.mslg.mslgoptimg", "")
	RunCMD("setprop", "vendor.mslg.mslgusrimg", "")
}
func ResetLoop() {
	fmt.Println("reseting loop")
	RunCMD("losetup", "-D")
}
func Setprop(key string, value string) {
	fmt.Println("Running setprop", key, value, ".")
	RunCMD("setprop", key, value)
}
func GetOptimgPath() (string, bool) { //位置，是否只读
	return "/data/Tapflow_project/mslgoptimg", false
}
func GetUsrimgPath() (string, bool) { //位置，是否只读
	return "/data/Tapflow_project/mslgusrimg", false
}
func SetupLoop(loop string, path string, ro bool) {
	fmt.Println("going to setup loop to ", loop, path)
	if ro {
		cmd, _ := RunCMD("losetup", "-r", loop, path)
		fmt.Println("setup loop result:", cmd)
	} else {
		cmd, _ := RunCMD("losetup", loop, path)
		fmt.Println("setup loop result:", cmd)
	}
}
func main() {
	ResetLoop()
	Resetprop()
	var optimgloop string
	var usrimgloop string
	//一般是33或者34 重新挂载多了不好，需要重启
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
}
