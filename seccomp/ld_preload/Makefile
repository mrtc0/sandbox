CC = gcc

build:
	$(CC) -shared -fPIC -o seccomp.so seccomp.c -lseccomp
	$(CC) uname.c -o uname
	$(CC) uname.c -static -o uname-static

run:
	LD_PRELOAD=./seccomp.so ./uname 
	LD_PRELOAD=./seccomp.so ./uname-static
