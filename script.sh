#!/bin/bash

gauss()
{
  ./gauss.out
}

threads_gauss()
{
  ./threads.out
}

chunk_threads()
{
  ./chunkthreads.out
}

openmp()
{
  ./openmp.out
}

pool_threads()
{
  ./pool.out
}

chunk_pool_threads()
{
  ./chunkpool.out
}

help()
{
    echo "        "
    echo "        "
    echo "        gauss - serial version"
    echo "        threads_gauss - pthread implementation"
    echo "        chunk_threads - pthread implementation with chunks"
    echo "        openmp - OpenMP implementation"
    echo "        pool_threads - Pool of threads implementation"
    echo "        chunk_pool_thread - Pool of threads implementation with chunks"
    echo "        "
    echo "        "
}
