FROM nginx:alpine

COPY ../config/nginx.conf /etc/nginx/nginx.conf
COPY ./deployment/server.crt /etc/nginx/server.crt
COPY ./deployment/server.key /etc/nginx/server.key