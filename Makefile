DIST		= dist
INSTALL		= $(HOME)/.local/bin
scrap_script 	= mm-parsing
OUT 		= mm
installed	= $(INSTALL)/$(scrap_script) $(INSTALL)/$(OUT)
TMPDIR 		= tmp/

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
	./test/website-change.sh
	./test.sh
	go test ./utils/ ./client/
	goreleaser healthcheck

.PHONY: uninstall remove
remove: uninstall
uninstall: $(installed)
	rm -v $(installed)
