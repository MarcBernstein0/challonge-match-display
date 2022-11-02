#!/bin/bash

# Create env file
set -e

if [ -n "$REACT_APP_MATCH_DISPLAY_URL" ]; then
sed -i -e "s|REPLACE_MATCH_DISPLAY_URL|$REACT_APP_MATCH_DISPLAY_URL|g" /app/container/scripts/sample.env 
fi

cat /app/container/scripts/sample.env > /app/.env 
