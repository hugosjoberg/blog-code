# Build Remote
FROM node:18-alpine as remote-builder

RUN apk update \
    && apk add openssl

WORKDIR /app
COPY remote .
RUN npm install
RUN npm run build

# Build Host
FROM node:18-alpine as host-builder

RUN apk update \
    && apk add openssl

WORKDIR /app
COPY host .
RUN npm install
RUN npm run build

# Move the build react components into nginx
FROM nginx:1.25.3-alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=host-builder /app/dist /usr/share/nginx/html/host
COPY --from=remote-builder /app/dist /usr/share/nginx/html/remote

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
