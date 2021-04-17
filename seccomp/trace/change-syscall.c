#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <err.h>
#include <errno.h>
#include <seccomp.h>
#include <ctype.h>
#include <sys/types.h>
#include <sys/utsname.h>
#include <sys/signal.h>
#include <sys/wait.h>
#include <sys/user.h>
#include <sys/ptrace.h>
#include <sys/socket.h>
#include <syscall.h>

static void sandbox(void)
{
    scmp_filter_ctx seccomp_ctx = seccomp_init(SCMP_ACT_ALLOW);
    if (!seccomp_ctx)
        err(1, "seccomp_init failed");

    if (seccomp_rule_add_exact(seccomp_ctx, SCMP_ACT_TRACE(0), seccomp_syscall_resolve_name("uname"), 0)) {
        perror("seccomp_rule_add_exact failed");
        exit(1);
    }

    if (seccomp_load(seccomp_ctx)) {
        perror("seccomp_load failed");
        exit(1);
    }

    seccomp_release(seccomp_ctx);
}

int call_getpid()
{
    struct utsname name;
   
    syscall(SYS_uname, SYS_mkdir, "dir", 0777);

    return 0;
}

int main(void)
{
    int pid;
    int status;
    struct user_regs_struct regs;

    switch( (pid = fork() )) {
        case -1: perror("fork failed"); return 1;
        case 0:
            ptrace(PTRACE_TRACEME, 0, NULL, NULL);
            sandbox();
            kill(getpid(), SIGSTOP);
            call_getpid();
            return 0;
    }

    waitpid(pid, 0, 0);

    ptrace(PTRACE_SETOPTIONS, pid, NULL, PTRACE_O_TRACESECCOMP);
    ptrace(PTRACE_CONT, pid, NULL, NULL);

    while(1) {
        if (waitpid(pid, &status, __WALL) == -1) {
            break;
        }

        if (WIFEXITED(status) || WIFSIGNALED(status)) {
            break;
        }

        if (WIFSTOPPED(status) && WSTOPSIG(status) == SIGTRAP && status >> 16 == PTRACE_EVENT_SECCOMP) {
            ptrace(PTRACE_GETREGS, pid, NULL, &regs);
            printf("orig_rax = %lld\n", regs.orig_rax);
            regs.orig_rax = regs.rdi;
            regs.rdi = regs.rsi;
            regs.rsi = regs.rdx;
            regs.rdx = regs.r10;
            regs.r10 = regs.r8;
            regs.r8 = regs.r9;
            regs.r9 = 0;
            ptrace(PTRACE_SETREGS, pid, NULL, &regs);
            printf("orig_rax = %lld\n", regs.orig_rax);
        }

        ptrace(PTRACE_CONT, pid, NULL, NULL);
    }

    return 0;
}

