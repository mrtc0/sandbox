
CC = gcc

uname:
	$(CC) uname.c -o uname
	./uname

uname_with_raw_seccomp:
	$(CC) uname_with_raw_seccomp.c -o uname_with_raw_seccomp
	./uname_with_raw_seccomp

uname_with_libseccomp:
	$(CC) uname_with_libseccomp.c -lseccomp -o uname_with_libseccomp
	./uname_with_libseccomp
