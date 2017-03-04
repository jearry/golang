#!/usr/bin/env bash

program=fm.sh
sn=`ps -ef | grep $program | grep -v grep |awk '{print $2}'`  
if [ "${sn}" = "" ]
then
NOW="$(date)"
echo $NOW |mail -s 'fm service has stopped, restarting...' jearry@163.com
cd /fm/
nohup /fm/fm.sh >/dev/null 2>&1 &
fi