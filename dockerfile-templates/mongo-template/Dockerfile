# Set the base image
FROM mongo:4.0.4

# The author
LABEL AUTHOR="Author"

# To use custon mongodb config
# ADD mongodb.conf /etc/mongodb.conf

VOLUME /data/db

# Expose ports.
#   - 27017: process
#   - 28017: http
EXPOSE 27017
EXPOSE 28017

CMD ["mongod"]
## If using custom config
# CMD ["/usr/bin/mongod", "--config", "/etc/mongodb.conf"]