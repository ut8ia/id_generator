![Go](https://github.com/ut8ia/hookexec/workflows/Go/badge.svg?branch=master)
# Universally Unique Lexicographically Sortable Identifier (ULID) microservice 
## HTTP based microservice for generation and resolve ULID`s 

HTTP wrapper based on https://github.com/oklog/ulid ULID generator 


### Futures
- simple GET requests
- resolve your uid at millesecond via api 

### Typical use cases : 
- unique identifiers for distributed services
- unique id of request for tracing request inside your architecture

### why ULID ?
- sortable
- short
- url-safe
- simple
- has other implementations

JS version : https://github.com/ulid/javascript

PHP version : https://github.com/robinvdvleuten/php-ulid

Python : https://github.com/valohai/ulid2

### comparison
 - UUID OSF Realization by Google corp
  
    source : https://github.com/google/uuid
  
    example : c29ad4f5-ea3f-47d0–9526-b5d4e9029a45
  
    sortable : NO
    
 -  Sonyflake from SONY corp
  
    source : https://github.com/sony/sonyflake
  
    example : 20f8707d6000108
  
    sortable : YES
  
 -  Ksuid from segment.io
   
    source : https://github.com/segmentio/ksuid
   
    example : 0ujtsYcgvSTl8PAuAdqWYSMnLOv
   
    sortable : YES
   
 -  Ulid from OKlog
  
    source: https://github.com/oklog/ulid
  
    example : 01D78XYFJ1PRM1WPBCBT3VHMNV
  
    sortable : YES

### config example
```yaml
server:
  host: "0.0.0.0" # default address
  port: 8084 # port 
```

### examples
Generate new 
```bash
$  curl -X GET 0.0.0.0:8080/ 
```
Result : 01E5DHB0VWZ9P2HXN6E4JC3V3Q

Resolve your id as datetime
```bash
$  curl -X GET 0.0.0.0:8080/datetime?id=01E5DHB0VWZ9P2HXN6E4JC3V3Q 
```
Result : 2020-04-08 21:37:48 +0300 EEST

Resolve your id as timestamp
```bash
$ curl -X GET 0.0.0.0:8084/timestamp?id=01E5DHB0VWZ9P2HXN6E4JC3V3Q
```
Result : 1586371068796

### Build
For local run : 
```bash
$ make start
```
Build of container for Docker 
```bash
$ make build
```

### Running
Default config path is ./configs/config.yml so you can simply run
```bash
$ ./main
```
You can run N instances with different configs, specifies it in call argument :
```bash
$ ./main ./configs/first_config.yml 
....
$ ./hookexec ./configs/second_config.yml 
```

### Logging
Relay stdout to your logfile 
```bash
$ ./main ./config/config.yml &> ./logs/example.log
```