#!/bin/bash

# Create env file
set -e

echo Setting Env File

if [ -n "$REACT_APP_MATCH_DISPLAY_URL" ]; then
echo $REACT_APP_MATCH_DISPLAY_URL
sed -i -e "s|REPLACE_MATCH_DISPLAY_URL|$REACT_APP_MATCH_DISPLAY_URL|g" /app/container/templates/sample.env.template 
cat /app/container/templates/sample.env.template  > /app/.env 
fi

# if [ -n "$PORT" ]; then
# echo $PORT
# sed -i -e "s|PORT|$PORT|g" /app/container/templates/default.conf.template
# fi
