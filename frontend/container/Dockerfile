FROM node:latest AS builder

WORKDIR /app

ENV PATH /app/node_modules/.bin:${PATH} 

COPY package.json ./
COPY package-lock.json ./
COPY tsconfig.json ./

RUN npm install


COPY . ./

RUN npm run build 

# Pull from nginx to run frontend
FROM nginx:latest

WORKDIR /usr/share/nginx/html

# Remove default nginx static resources
RUN rm -rf ./*

# Copies static resources from builder stage
COPY --from=builder /app/build .

# Container run nginx with global directives and daemon off
ENTRYPOINT [ "nginx", "-g", "daemon off;" ]
