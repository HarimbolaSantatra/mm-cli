# INTERNALS
## mm-parsing
To show usage of `mm-parsing`:
```
./mm-parsing --help
```

There are two way to use the `mm-parsing` script:
```bash
./mm-parsing -c [HTML_STRING] # Default
./mm-parsing -f file.html
```

__Example__:
```
cat result.html | xargs -0 ./mm-parsing -c
```
