#!/bin/bash
pkill -f little-server
gp url 8081
bash -c "exec -a little-server node server.js"
