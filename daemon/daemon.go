package daemon

import (
	"fmt"
	"os"
	"syscall"
)

func Daemon(nochdir,noclose int,args []string)(int,error){
	if syscall.Getppid()==1{
		syscall.Umask(0)
		if nochdir==0{
			os.Chdir("/")
		}
		return 0,nil
	}
	
	files:=make([]*os.File,3,6)
	if noclose ==0 {
		nullDev,er:=os.OpenFile("/dev/null",0,0)
		if err!=nil{
			return 1,err
		}
		files[0],files[1],files[2]=nullDev,nullDev,nullDev
	}else{
		files[0],files[1],files[2]=os.Stdin,os.Stdout,os.Stderr
	}
	dir,_:=os.Getwd()
	sysattrs:=syscall.SysProcAttr{Setsid:true}
	attrs:=os.ProcAttr{Dir:dir,Env:os.Environ(),Files:files,Sys:&sysattrs}
	proc,err:=os.StartProcess(args[0],args,&attrs)
	if err!=nil{
		return -1,fmt.Errorf("can't create process %s:%s",args[0],err)
	}
	proc.Release()
	os.Exit(0)
	return 0,nil

}
