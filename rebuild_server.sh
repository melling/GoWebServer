#!/bin/sh

export GOPATH=`pwd`:$HOME

DATE=`date "+%Y%m%d"`

rm bin/webserver
cd src/webserver
go install
cd -
cp -p bin/webserver bin/webserver.${DATE}
ls -l bin/webserver
