all: inside-container

## Docker targets
docker-build:
	@echo "\033[1;38;5;155mBuilding Docker image... \033[0m"
	@docker image build -f Dockerfile -t ascii-art-web-docker .
	@echo "\033[1;38;5;155mList All Docker images... \033[0m"
	@docker images

## Running Container
docker-run: docker-build
	@echo "\033[1;38;5;155mRunning Docker Container with Port 8080:8080... \033[0m"
	@docker container run -p 8080:8080 --detach --name ascii-art-web ascii-art-web-docker
	@echo "\033[1;38;5;155mList All Containers: \033[0m"
	@docker ps

## list All && Exit
inside-container: docker-run
	@echo "\033[1;38;5;155mGo inside The Container, list all File System and Exit: \033[0m"
	@docker exec -it ascii-art-web /bin/bash -c "ls -l && exit"

## Clean All
clean:
	@echo "\033[1;38;5;155mStopping and Removing all Docker Containers... \033[0m"
	@docker ps -q | xargs -r docker stop
	@docker ps -a -q | xargs -r docker rm
	@echo "\033[1;38;5;155mCleaning up all Docker images... \033[0m"
	@docker image prune -a -f
	@echo "\033[1;38;5;155mThe Clean-up is Completed. \033[0m"
