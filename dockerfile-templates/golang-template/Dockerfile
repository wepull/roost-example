#Base image
FROM golang:latest 

# The author
LABEL AUTHOR="Author Name"
LABEL app="app-name"

#Copy your app files into src folder
RUN  mkdir /src
COPY src/main.go /src
WORKDIR /src

#expose port 
EXPOSE 8080

#Execution
CMD ["go","run","./main.go"]