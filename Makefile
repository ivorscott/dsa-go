#!make

time:
	go test ./time -bench=.

binarysearch: 
	go test ./arrays/binarysearch -bench=.

binarysearchr: 
	go test ./arrays/binarysearchr -bench=.

commondivosor:
	go test ./arrays/commondivisor -bench=.

factorial: 
	go test ./arrays/factorial -bench=.

findone: 
	go test ./arrays/findone -bench=.
	
sort: 
	go test ./arrays/sort -bench=.
	
sumarray: 
	go test ./arrays/sumarray -bench=.

.PHONY: time
.PHONY: sumarray

