CC:=gcc

program: gauss.c

	 $(CC) gauss.c -o gauss.out


clean:
	rm -rf *.out
