to-gist
===========


STDIN to Gist.



Installation
---------

```sh
go get github.com/pocke/to-gist
```




Usage
------


```sh
# Create Gist by ifconfig output
$ ifconfig | to-gist ifconfig.txt

# Create Gist as private
$ ifconfig | to-gist ifconfig.txt -p
```
