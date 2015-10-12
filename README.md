# Gaussian Elimination

This program is a C implementation to improve performance of a Gaussian Elimination Code, using OpenMP, Pthreads and [Pool of Threads](https://github.com/Pithikos/C-Thread-Pool)

You can compile everything using the Makefile and run the objects created.
```sh
$ make
```
There is a script with alias for running the programs. You can type:
```sh
$ source script.sh
```
Then, type:
```sh
$ help
```
You will see all the alias to run the program. It should be run after ```sh make``` and should be with `source`

For run the serial version. you can type:

```sh
$ gcc gauss.c -o gauss.out
```


$ gcc threads_gauss.c -pthread -o threads.out

$ gcc chunk_threads_gauss.c -pthread -o chunkthreads.out

$ gcc openmp_gauss.c -fopenmp -o openmp.out

$ gcc chunk_thpool_gauss.c thpool.c -o chunkpool.out

$ gcc thpool_gauss.c thpool.c -o pool.out
```
