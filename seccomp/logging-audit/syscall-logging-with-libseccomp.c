#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/utsname.h>
#include <seccomp.h>
#include <err.h>

static void sandbox(void)
{
    scmp_filter_ctx seccomp_ctx = seccomp_init(SCMP_ACT_LOG);
    if (!seccomp_ctx)
        err(1, "seccomp_init failed");

    if (seccomp_load(seccomp_ctx)) {
        perror("seccomp_load failed");
        exit(1);
    }

    seccomp_release(seccomp_ctx);
}

int main(void)
{
    struct utsname name;

    sandbox();

    getpid();
    if (uname(&name)) {
        perror("uname failed");
        return 1;
    }

    printf("uname: %s\n", name.sysname);

    return 0;
}
