#!/bin/sh
#./check_port.sh -w "ls -al " -h 127.0.0.1 -p 13000
#while :
#do
#
#sleep 3
#  flage=$(nc -v -z -w 2  127.0.0.1 13000 > /dev/null 2>&1)
#  if [ $? -eq 0 ]
#    then
#    echo "success"
#    sleep 1
#    else
#    echo "Port UnOnline"
#    fi
#done
# while ! timeout 1 bash -c "echo > /dev/tcp/localhost/13000"; do sleep 10; done
V_COMMAND=${V_COMMAND:-""}
V_HOST=${V_HOST:-""}
V_PORT=${V_PORT:-""}
while getopts "h:p:w:" OPT; do
    case $OPT in
        w)
            V_COMMAND=$OPTARG;;
        h)
            V_HOST=$OPTARG;;
        p)
            V_PORT=$OPTARG;;
    esac
done

while ! echo exit | nc ${V_HOST} ${V_PORT}; do sleep 3; echo "Port UnOnline"; done
${V_COMMAND}

