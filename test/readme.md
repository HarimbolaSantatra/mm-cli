# TEST
This directory contains scripts for all sort of tests.

To run all the test, navigate to the root folder and run:
```
make test
```

Files:
- `test-python-env`: test python dependencies
- `test-request.sh`: make a test request to server
- `test.py`: test the HTML parser. Test if the scraper needs an update. That's the case only if the website has been updated by the admin and it's not in the same structure anymore.
- `website-change.sh`: test if website HTML structure have changed. Test that by making a request and searching the word `lakozia`.
