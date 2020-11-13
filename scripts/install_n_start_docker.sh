#!/bin/bash

sudo yum install zip -y
sudo yum install docker -y
sudo systemctl enable docker
sudo systemctl start docker
sudo mkdir -p /tinyurl 
sudo chown ec2-user:ec2-user /tinyurl
sudo unzip /deploy/gotinyurl.zip -d /tinyurl
sudo chown -R ec2-user:ec2-user /tinyurl
sudo chmod 777 gourlshortner*