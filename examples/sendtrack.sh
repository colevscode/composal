#!/bin/bash

# test running the server, then sending a test message
# requires server to be built first with `go build`
# requires curl 
# expects that you're gonna cd into the examples folder
# run as `./sendtrack.sh track1.json out.wav`

trap "kill 0" EXIT

../composal -port 8088 & curl --data "@$1" http://localhost:8088/play > $2 
sleep 1
play $2