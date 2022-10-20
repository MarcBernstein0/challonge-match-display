help: # Show this help.
	@fgrep -h "#" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/#//'

docker_build: # build docker image
	docker build -t match-display-frontend -f ./container/Dockerfile .

docker_run: # run docker image
	docker run --rm -p 3000:80 match-display-frontend

run_container: docker_build docker_run # build docker image and run container
