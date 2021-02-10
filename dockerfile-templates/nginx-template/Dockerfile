# Base Image
FROM nginx:alpine

# The author
LABEL author="Author Name"

# Copy html and nginx config 
COPY ./dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# For standard non-secured "http" is port 80, Netscape chose 443 to be the default port used by secure http.
EXPOSE 80 443

# Start NGINX Server
ENTRYPOINT ["nginx", "-g", "daemon off"]

# End of Dockerfile