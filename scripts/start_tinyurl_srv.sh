#!/bin/bash

cd /tinyurl 
nohup sudo ./gourlshortner-linux > /tmp/url.log &
exit 0