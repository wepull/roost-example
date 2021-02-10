# Base image
FROM postgres

#Setting postgres env variables
ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD postgres 

# Add a database
ENV POSTGRES_DB testdb 

# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
COPY init.sql /docker-entrypoint-initdb.d/