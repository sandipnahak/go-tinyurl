#!/bin/bash

sudo yum install zip -y
sudo yum install docker -y
sudo systemctl enable docker
sudo systemctl start docker
sudo rm -rf  /tinyurl 
sudo mkdir -p /tinyurl 
sudo unzip  /opt/codedeploy-agent/deployment-root/$DEPLOYMENT_GROUP_ID/$DEPLOYMENT_ID/bundle.tar  -d /tinyurl
sudo cp -v /tinyurl/build/gourlshortner* /tinyurl
sudo chmod 777  /tinyurl/gourlshortner*