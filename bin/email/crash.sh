#!/bin/bash

emails="toutiao_monitor@truxing.com"
sendmail="gomail  -q smtp.exmail.qq.com:465 -u boomer_kefu@truxing.com -p  NyxQGrXTpjiad18H -t $emails -s true"

function findExit(){
	supervisorctl status | grep -v "RUNNING"
}

function mails(){
	echo "$2" | $sendmail -j "$1"
}

msg=`findExit | awk '{print "lp3 service " $1 " has been " $2 }'`
mails "lp3 service error" "$msg"


