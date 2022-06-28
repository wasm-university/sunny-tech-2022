#!/bin/bash
pkill -f little-server
gp url 8081
bash -c "exec -a little-server-8081 node server.js"
