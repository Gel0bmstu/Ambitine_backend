#!/bin/bash

# remove exited containers:
sudo docker ps -aq --filter status=dead --filter status=exited | xargs -r docker rm -v

# remove unused images:
sudo docker images -qf dangling=true --no-trunc | grep '<none>' | awk '{ print $3 }' | xargs -r docker rmi

# remove unused volumes:
sudo docker volume ls -qf dangling=true | xargs -r docker volume rm