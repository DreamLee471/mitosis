package container


/*
#define _GNU_SOURCE
#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sched.h>
#include <signal.h>
#include <sys/mount.h>
#define STACK_SIZE (1024*1024)
static char container_stack[STACK_SIZE];

char* const container_args[]={"/bin/bash",NULL};

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
	sethostname((char*)arg,10);
	struct sigaction sa;
	if(mount("proc","/home/dreamlee/go_workspace/miroot/proc","proc",0,NULL)!=0){
		perror("proc");
	}
	
	if(mount("sysfs","/home/dreamlee/go_workspace/miroot/sysfs","sysfs",0,NULL)!=0){
                perror("sys");
        }

	if(mount("none","/home/dreamlee/go_workspace/miroot/tmp","tmpfs",0,NULL)!=0){
                perror("tmp");
        }

	if(mount("udev","/home/dreamlee/go_workspace/miroot/dev","devtmpfs",0,NULL)!=0){
                perror("dev");
        }

	if(mount("devpts","/home/dreamlee/go_workspace/miroot/dev/pts","devpts",0,NULL)!=0){
                perror("dev/pts");
        }

	if(mount("shm","/home/dreamlee/go_workspace/miroot/dev/shm","tmpfs",0,NULL)!=0){
                perror("dev/shm");
        }

	if(mount("tmpfs","/home/dreamlee/go_workspace/miroot/run","tmpfs",0,NULL)!=0){
                perror("proc");
        }

	if(chdir("/home/dreamlee/go_workspace/miroot")!=0 || chroot("./")!=0){
		perror("chdir/chroot");
	}
	sigemptyset(&sa.sa_mask);
	sa.sa_flags=0;
	sa.sa_handler=unmount_resource;
	sigaction(SIGINT,&sa,NULL);
	sigaction(SIGQUIT,&sa,NULL);
	execv(container_args[0],container_args);
	return 1;
}

static void chld_handler(int sig){
	int status;
	waitpid(-1,&status,WNOHANG);
}


struct container{
	int pid;
	int ppid;
	char* name;
	char* rootfs;
	char* writefs;
};


struct container test(){
	struct container c;
	return c;
}

void startContainer(int* container_id,struct container *c){
	sigset_t blockMask,emptyMask;
	struct sigaction sa;
	int status;
	setbuf(stdout,NULL);
	*container_id=clone(container_main,container_stack+STACK_SIZE,CLONE_NEWUTS  | CLONE_NEWPID | CLONE_NEWUTS | SIGCHLD,(*c).name);
	
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
	waitpid(*container_id,NULL,0);
}

*/
import "C"

import (
	"fmt"
)


type Container struct {
	Id	int
	Pid	int
	Alias	string
	RootFs	string
	WriteFs	string
	Hostname	string
}


func (c *Container) Start(){
	var container_id C.int
	c_container := &C.struct_container{name:C.CString(c.Hostname)}
	C.startContainer(&container_id,c_container)
	fmt.Println("pid:",container_id)
}
