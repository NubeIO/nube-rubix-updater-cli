#!/bin/bash

for ARGUMENT in "$@"
do
    KEY=$(echo "$ARGUMENT" | cut -f1 -d=)
    VALUE=$(echo "$ARGUMENT" | cut -f2 -d=)

    case "$KEY" in
            ip)                ip=${VALUE} ;;
            port)              port=${VALUE} ;;
            user)              user=${VALUE} ;;
            password)          password=${VALUE} ;;
            *)
    esac
done
if [ -z "$ip" ]
then
      ip="192.168.15.100"
fi
if [ -z "$port" ]
then
      port=22
fi
if [ -z "$user" ]
then
      user="pi"
fi
if [ -z "$password" ]
then
      password="N00B"
fi


echo "user = $user"
echo "pass = $password"
echo "port = $port"

sshpass -p "$password" ssh -o StrictHostKeyChecking=no $user@$ip -p $port

