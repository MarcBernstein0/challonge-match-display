FROM node:16.16.0-alpine3.16 AS builder

WORKDIR /app

# Set env variable
ENV PATH /app/node_modules/.bin:${PATH} 

COPY ./frontend ./

RUN chmod +x ./container/scripts/set-env.sh \
    ./container/scripts/set-env.sh && \
    npm install 


RUN npm run build

# Pull from nginx to run frontend
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

# Remove default nginx static resources
RUN rm -rf ./*

# Copies static resources from builder stage
COPY --from=builder /app/build .

COPY --from=builder /app/container/templates/default.conf.template /etc/nginx/conf.d/default.conf.template

# Container run nginx with global directives and daemon off
CMD sh -c "envsubst \"`env | awk -F = '{printf \" \\\\$%s\", $1}'`\" < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"
