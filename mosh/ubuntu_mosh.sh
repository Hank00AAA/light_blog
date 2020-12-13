#!/bin/bash

sudo apt-get install python-software-properties
sudo add-apt-repository ppa:keithw/mosh
sudo apt-get update
sudo apt-get install mosh

sudo iptables -I INPUT -p udp --dport 60001 -j ACCEPT 

mosh_server