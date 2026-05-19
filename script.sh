# Building the Docker image
docker image build -f Dockerfile -t dockerize-img .

# Running the Docker container based on the built image
docker container run -p 8080 --detach --name dockerize-con-v1 dockerize-img

# Checking the running containers
docker ps -a

# Accessing the container's shell
docker exec -it dockerize /bin/bash

# Stopping the container
docker container stop dockerize-con-v1