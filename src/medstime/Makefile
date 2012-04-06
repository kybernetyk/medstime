CC=go tool 6g
LL=go tool 6l
LIBS=
LNAME=medstime.6
CFLAGS=-o $(LNAME) -p medstime 
EXECNAME=medstime

.PHONY: all
all: medstime 

test: clean all run

.PHONY: medstime 
medstime:
	$(CC) $(CFLAGS) ../shared/*.go *.go $(LIBS)
	$(LL) -o $(EXECNAME) $(LNAME)
	rm $(LNAME)

clean:
	rm -rf *.o $(EXECNAME) core *.out

.PHONY: run
run: medstime 
	./medstime
