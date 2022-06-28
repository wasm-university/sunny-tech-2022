#!/bin/bash
pkill -f little-server
gp url 8082
bash -c "exec -a little-server-8082 node server.js"
