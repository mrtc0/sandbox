#include <stdio.h>
#include <seccomp.h>

__attribute__((constructor)) void configure_seccomp(void)
{
    scmp_filter_ctx seccomp_ctx = seccomp_init(SCMP_ACT_ALLOW);
    if (!seccomp_ctx)
    {
        perror("seccomp_init failed");
    }

    if (seccomp_rule_add_exact(seccomp_ctx, SCMP_ACT_KILL, seccomp_syscall_resolve_name("uname"), 0))
    {
        perror("seccomp_rule_add_exact failed");
    }

    if (seccomp_load(seccomp_ctx))
    {
        perror("seccomp_load failed");
    }

    seccomp_release(seccomp_ctx);
    printf("Configuring seccomp\n");
}
