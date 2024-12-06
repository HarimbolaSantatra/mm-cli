DIST		= dist
INSTALL		= $(HOME)/.local/bin
scrap_script 	= mm-parsing
installed	= $(INSTALL)/$(scrap_script) $(INSTALL)/$(OUT)
OUT 		= mm
TMPDIR 		= test/

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
	./test/test-python-env --verbose
	python3 test/test.py
	./test/website-change.sh
	goreleaser healthcheck

.PHONY: uninstall
uninstall: $(installed)
	rm $(installed)
