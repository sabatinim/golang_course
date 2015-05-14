dynamic:
	rm -f libkeyboard.dylib
	gcc -c -fPIC keyboard.c
	gcc -dynamiclib -lncurses -o libkeyboard.dylib keyboard.o
	rm -f keyboard.o

shared:
	rm -f libkeyboard.so
	gcc -c -fPIC keyboard.c
	gcc -shared -W1 -lncurses -soname,libkeyboard.so -o libkeyboard.so keyboard.o
	rm -f keyboard.o
