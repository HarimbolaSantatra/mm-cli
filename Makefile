OUT = mm
mm: 
	go build .

.PHONY: clean
clean:
	rm $(OUT)

.PHONY: test
test:
	./test/test-python-env
	python3 test/test.py
	go test ./client
