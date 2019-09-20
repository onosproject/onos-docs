#!/bin/bash
shopt -s extglob
git config user.name "adibrastegarnia"
git config user.email "arastega@purdue.edu"
git remote add gh-token "https://${GH_TOKEN}@github.com/onosproject/onos-docs.git";
git fetch gh-token && git fetch gh-token gh-pages:gh-pages;
CURRENT_PATH=$PWD
cd docs
git submodule add https://github.com/onosproject/onos-config.git
mv onos-config Configuration-Subsystem && cd Configuration-Subsystem && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-topo.git
mv onos-topo Topology-Subsystem && cd Topology-Subsystem && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-cli.git
mv onos-cli CLI-Subsystem && cd CLI-Subsystem && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-ztp.git
mv onos-ztp ZTP-Subsystem && cd ZTP-Subsystem && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-control.git
mv onos-control Control-Subsystem && cd Control-Subsystem && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-test.git
mv onos-test Test-Tools && cd Test-Tools && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-gui.git
mv onos-gui GUI-Subsystem && cd GUI-Subsystem && rm -rf !(docs)
cd $CURRENT_PATH
mkdocs gh-deploy -v --clean --remote-name gh-token;
