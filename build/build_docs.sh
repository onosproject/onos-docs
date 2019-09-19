#!/bin/bash
shopt -s extglob
git config user.name "adibrastegarnia"
git config user.email "arastega@purdue.edu"
git remote add gh-token "https://${GH_TOKEN}@github.com/onosproject/onos-docs.git";
git fetch gh-token && git fetch gh-token gh-pages:gh-pages;
CURRENT_PATH=$PWD
cd docs
git submodule add https://github.com/onosproject/onos-config.git
cd onos-config && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-topo.git
cd onos-topo && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-cli.git
cd onos-cli && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-ztp.git
cd onos-ztp && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-control.git
cd onos-control && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-test.git
cd onos-test && rm -rf !(docs)
cd ..
git submodule add https://github.com/onosproject/onos-gui.git
cd onos-gui && rm -rf !(docs)
cd $CURRENT_PATH
mkdocs gh-deploy -v --clean --remote-name gh-token;
