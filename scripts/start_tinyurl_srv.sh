#!/bin/bash

cd /tinyurl 
nohup sudo ./gourlshortner-linux > /dev/null 2>&1 &
exit 0