#!/usr/bin/env bash

i="0"
while [ $i -lt 25 ] #waiting for cluster 25 seconds
do
    i=$[$i+1]
    sleep 1
    echo -e ".\c"
done

echo ""
