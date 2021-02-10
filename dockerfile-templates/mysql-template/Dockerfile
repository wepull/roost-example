#From official mysql image (our base image)
FROM mysql

# Add a database
ENV MYSQL_DATABASE company

# set the password , alternatively you can pass the password in the docker run command
ENV MYSQL_ROOT_PASSWORD="supersecret"

# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
#So , lets create a table and add some data to it using .sql scripts
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

#Expose port 
EXPOSE 3306

# A sample docker run command to bind mount volume
# docker run -d -p 3306:3306 --name my-mysql \
# -v /Users/hardik/templates/mysql-1-template/sql-scripts:/docker-entrypoint-initdb.d/ \
# -e MYSQL_ROOT_PASSWORD=supersecret \
# -e MYSQL_DATABASE=company \
# mysql:latest