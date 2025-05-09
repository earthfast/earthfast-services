#!/bin/bash

if [ "$DISABLE_NGINX" = "true" ]; then
    echo "Nginx is disabled. Starting content node directly."

    export HTTP_PORT=80
    /armada-content-node
else
    echo "Starting with Nginx enabled."
    nginx

    export HTTP_PORT=5000
    /armada-content-node
fi
