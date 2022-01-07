# my Http Tool 
It is a simple Http tool to send request in parallel to the specified Url that would be given in
command line.



## How to Run:
you can run tests by :
```
make test
```
you can build the project by:
```
build myhttp
```
you can install dependencies by :
```
make vendor
```
to specify the number of parallel requests  you can use the below command :
```
./myhttp -parallel 3 http://google.com http://adjust.com http://facebook.com
```
you can specify the concurrency limit like above and if you don't use parallel flag the default 
number is  10.


