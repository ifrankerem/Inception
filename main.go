package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker 		 run image <cmd> <params>
// go run main.go run       <cmd> <params>


func main(){
	switch os.Args[1]{
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run(){
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin // binding the 0,1,2 fds to my child process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	if err := cmd.Run(); err != nil {
    fmt.Println("ERROR:", err)
    os.Exit(1)
	}
}

func child(){
	fmt.Printf("Running %v\n", os.Args[2:])

	syscall.Sethostname([]byte("container"))
	
	cmd := exec.Command(os.Args[2],os.Args[3:]...)
	cmd.Stdin = os.Stdin 
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	if err := cmd.Run(); err != nil {
    fmt.Println("ERROR:", err)
    os.Exit(1)
	}

}
