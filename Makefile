DIST		= dist
INSTALL		= $(HOME)/.local/bin
scrap_script 	= mm-parsing
installed	= $(INSTALL)/$(scrap_script) $(INSTALL)/$(OUT)
OUT 		= mm
TMPDIR 		= test/

endpoint 	= https://malagasyword.org

mm: 
	go build .

.PHONY: clean
clean:
	-@rm $(OUT) 2>/dev/null
	-@rm $(TMPDIR)/tmp.* 2>/dev/null
	-@rm -r $(DIST)/* 2>/dev/null

.PHONY: install
install: $(OUT)
	cp $(OUT) $(INSTALL)
	cp $(scrap_script) $(INSTALL)

.PHONY: test
test:
	wget -q --spider $(endpoint)
	./test/test-python-env --verbose
	python3 test/test.py
	./test/website-change.sh
	go test ./utils/ ./client/
	goreleaser healthcheck

.PHONY: uninstall
uninstall: $(installed)
	rm $(installed)
