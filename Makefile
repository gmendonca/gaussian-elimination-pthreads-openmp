CC=gcc

program:

	 $(CC) gauss.c -o gauss.out

	 $(CC) threads_gauss.c -pthread -o threads.out

	 $(CC) chunk_threads_gauss.c -pthread -o chunkthreads.out

	 $(CC) openmp_gauss.c -fopenmp -o openmp.out

	 $(CC) chunk_thpool_gauss.c thpool.c -o chunkpool.out

	 $(CC) thpool_gauss.c thpool.c -o pool.out


clean:
	rm -rf *.out
