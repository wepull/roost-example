# Base image
FROM redis:5.0.5

# To use your own redis config, uncomment the below line
# COPY redis.conf /usr/local/etc/redis/redis.conf

# Expose redis-server port 6379
EXPOSE 6379

# Start redis-server on container run
CMD ["redis-server"]
#IF using custom redis config
# CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
# End of Dockerfile

