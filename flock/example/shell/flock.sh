#!/bin/bash

$(cd `dirname $0`;)


#linux
/usr/bin/flock -xn flock.lock  -c 'go run main.go'



