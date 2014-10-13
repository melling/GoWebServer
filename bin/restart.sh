#!/bin/sh

host0=`hostname`
echo $host0

if test "$host0" = "mb2" ; then
	PID1=`ps  | grep "webserver" | grep -v grep | awk '{print $1}'`
else
	PID1=`ps -elf | grep "webserver" | grep -v grep | awk '{print $4}'`
fi
# PID2=`ps  | grep "go-build" | grep -v grep | awk '{print $1}'`


CMD="kill $PID1"

echo $CMD
out=`$CMD`
echo $out

# ps -elf | grep "go run spanishdb.go" | grep -v grep
# ps -elf | grep "/tmp/go-" | grep -v grep

ps  | grep "webserver" | grep -v grep
# ps  | grep "/tmp/go-" | grep -v grep

export PORT=8080
# Must start one dir up because of resources/ dir
export GOPATH=go:$GOPATH
#git pull
./rebuild_server.sh
mv webserver bin
bin/webserver &
# go run go/spanishdb.go go/Html.go go/Menu.go go/SummaryView.go \
#  go/DictionaryDAO.go go/TranslatorComment.go go/QuestionView.go \
#   go/QuestionModel.go go/PictureView.go go/OppositesView.go \
#   go/OppositesModel.go go/CategoryModel.go go/CategoryView.go \
#   go/MajorGroupModel.go \
#   go/HomeView.go \
#   go/MajorGroupView.go go/LocalizationModel.go go/LocalizationView.go &
