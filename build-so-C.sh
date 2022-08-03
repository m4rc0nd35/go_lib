gcc -c -fPIC hello.c -o hello.o && \
gcc hello.o -shared -o libhello.so

OR

gcc -shared -o libhello.so -fPIC hello.c