#!/usr/bin/env sh
set -ex

consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=agent-one

i="0"
while [ $i -lt 10 ] #waiting for cluster 10 seconds
do
    xterm &
    i=$[$i+1]
    sleep 1

    consul join service.api; consul join service.project
    RESULT=$?
    if [ $RESULT -eq 0 ]; then
        echo "Consul cluster is ready"
        consul members
        exit 1
    else
        echo "cluster doesn't ready. waiting..\r\n"
    fi

    consul members
done

echo "Error: Consul cluster doesn't start"
consul members
exit 1
