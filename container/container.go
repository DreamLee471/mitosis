package container


/*
#define _GNU_SOURCE
#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sched.h>
#include <signal.h>
#include <sys/mount.h>
#define STACK_SIZE (1024*1024)
static char container_stack[STACK_SIZE];
char* const container_args[]={"/bin/bash",NULL};

struct container{
        int pid;
        int ppid;
        char* name;
        char* rootfs;
        char* writefs;
        char* initcmd;
};

static void unmount_resource(int sig){
	umount("/home/dreamlee/go_workspace/miroot/proc");
	umount("/home/dreamlee/go_workspace/miroot/sysfs");
	umount("/home/dreamlee/go_workspace/miroot/tmp");
	umount("/home/dreamlee/go_workspace/miroot/dev");
	umount("/home/dreamlee/go_workspace/miroot/dev/pts");
	umount("/home/dreamlee/go_workspace/miroot/shm");
	umount("/home/dreamlee/go_workspace/miroot/run");
}
int container_main(void* arg){
	struct container *c;
	c=(struct container*)arg;
	sethostname((*c).name,20);
	struct sigaction sa;
	if(chdir((*c).rootfs)!=0 || chroot("./")!=0){
		perror("chdir/chroot");
	}
	system("mount -t proc proc /proc");
	
	sigemptyset(&sa.sa_mask);
	sa.sa_flags=0;
	sa.sa_handler=unmount_resource;
	sigaction(SIGINT,&sa,NULL);
	sigaction(SIGQUIT,&sa,NULL);
	execv(container_args[0],container_args);
	return 1;
}

void gosystem(char* cmd){
	system(cmd);
}

static void chld_handler(int sig){
	int status;
	waitpid(-1,&status,WNOHANG);
}

void startContainer(int* container_id,struct container *c){
	sigset_t blockMask,emptyMask;
	struct sigaction sa;
	int status;
	setbuf(stdout,NULL);
	*container_id=clone(container_main,container_stack+STACK_SIZE,CLONE_NEWUTS  | CLONE_NEWPID | CLONE_NEWUTS | SIGCHLD ,c);
	
	sigemptyset(&sa.sa_mask);
	sa.sa_flags=0;
	sa.sa_handler=chld_handler;
	if(sigaction(SIGCHLD,&sa,NULL)==-1){
		perror("sigaction");
	}
	sigemptyset(&blockMask);
	sigaddset(&blockMask,SIGCHLD);
	if(sigprocmask(SIG_SETMASK,&blockMask,NULL)==-1){
		perror("sigprocmask");
	}
	c->pid=*container_id;
	waitpid(c->pid,NULL,0);
	_exit(0);
}
*/
import "C"

import (
	"fmt"
	"syscall"
)

type Container struct {
	Id	int
	Pid	int
	Alias	string
	RootFs	string "/home/dreamlee/go_workspace/miroot"
	ParentFs	string
	WriteFs	string
	Hostname	string
	InitCmd	string
}


func (c *Container) Start(){
	if len(c.ParentFs)==0{
		c.ParentFs="/home/dreamlee/go_workspace/miroot"
	}
	if len(c.WriteFs)==0{
                c.WriteFs="/home/dreamlee/go_workspace/microot/layer/"+c.Hostname
        }
	c.RootFs="/home/dreamlee/go_workspace/microot/"+c.Hostname
	syscall.Mkdir(c.WriteFs,0755)
	syscall.Mkdir("/home/dreamlee/go_workspace/microot/"+c.Hostname,0755)
	C.gosystem(C.CString(fmt.Sprintf("mount -t aufs -o br=%s:%s none /home/dreamlee/go_workspace/microot/%s",c.WriteFs,c.ParentFs,c.Hostname)))
	var container_id C.int
	c_container := &C.struct_container{name:C.CString(c.Hostname),rootfs:C.CString(c.RootFs)}
	C.startContainer(&container_id,c_container)
}
