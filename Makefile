help: # Show this help.
	@fgrep -h "#" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/#//'

backend_docker_build: # build backend docker image
	docker build --no-cache -t match-display -f ./backend/container/Dockerfile .

backend_docker_run: # run backend docker image
	docker run -d --env-file ./backend/.env -p 8080:8080 match-display

backend_build_and_run: backend_docker_build backend_docker_run # build and run backend container

frontend_docker_build: # build frontend docker image
	docker build --no-cache -t match-display-frontend -f ./frontend/container/Dockerfile .

frontend_docker_run: # run frontend docker image
	docker run --rm --env-file ./frontend/.env -e PORT=80 -p 3000:80 match-display-frontend 

frontend_build_and_run: frontend_docker_build frontend_docker_run # build and run frontend container

start: # run docker-compose up headless
	docker compose up -d 

start_no_cache: # run docker-compose without cache
	docker compose build --no-cache && docker compose up -d 
	

stop: # stock docker-compose
	docker compose stop

clean_compose: # remove non-running docker composed containers
	docker compose rm

stop_n_clean: # stop docker compose and remove containers
	docker compose rm -s 

