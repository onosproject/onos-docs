#!/bin/bash
shopt -s extglob
CURDIR=$(pwd)
git clone https://github.com/onosproject/onos-config.git ./content/onos-config
cd ./content/onos-config && rm -rf !(docs)
cd $CURDIR
git clone https://github.com/onosproject/onos-topo.git ./content/onos-topo
cd ./content/onos-topo && rm -rf !(docs)
cd $CURDIR
git clone https://github.com/onosproject/onos-cli.git ./content/onos-cli
cd ./content/onos-cli && rm -rf !(docs)
cd $CURDIR
git clone https://github.com/onosproject/onos-ztp.git ./content/onos-ztp
cd ./content/onos-ztp && rm -rf !(docs)
cd $CURDIR
git clone https://github.com/onosproject/onos-gui.git ./content/onos-gui
cd ./content/onos-gui && rm -rf !(docs)
cd $CURDIR
git clone https://github.com/onosproject/onos-test.git ./content/onos-test
cd ./content/onos-test && rm -rf !(docs)
cd $CURDIR
