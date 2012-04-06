CC=go tool 6g
LL=go tool 6l
LIBS=
LNAME=smsdemon.6
CFLAGS=-o $(LNAME) -p smsdemon
EXECNAME=smsdemon

.PHONY: all
all: smsdemon	

test: clean all run

.PHONY: smsdemon 
smsdemon:
	$(CC) $(CFLAGS) ../shared/*.go *.go $(LIBS)
	$(LL) -o $(EXECNAME) $(LNAME)
	rm $(LNAME)

clean:
	rm -rf *.o $(EXECNAME) core *.out

.PHONY: run
run: smsdemon 
	./smsdemon
