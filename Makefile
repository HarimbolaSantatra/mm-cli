OUT = mm
TMPDIR = test/

mm: 
	go build .

.PHONY: clean
clean:
	-@rm $(OUT) 2>/dev/null
	-@rm $(TMPDIR)/tmp.* 2>/dev/null

.PHONY: test
test:
	./test/test-python-env
	python3 test/test.py
	./test/website-change.sh
