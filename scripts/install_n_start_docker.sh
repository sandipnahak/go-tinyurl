#!/bin/bash

sudo yum install zip -y
sudo yum install docker -y
sudo systemctl enable docker
sudo systemctl start docker
sudo mkdir -p /tinyurl 
sudo unzip  $DEPLOYMENT_GROUP_ID/$DEPLOYMENT_ID/bundle.tar  -d /tinyurl
sudo chmod 777 gourlshortner*