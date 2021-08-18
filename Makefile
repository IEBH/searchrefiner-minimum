run:
	go run .

build:
	go build -o server .

docker-build:
	docker build -t ielab-searchrefiner .
	-docker network rm search-refiner-net
	docker network create search-refiner-net --driver=bridge

docker-build-force:
	docker build --no-cache -t ielab-searchrefiner .
	-docker network rm search-refiner-net
	docker network create search-refiner-net --driver=bridge

# Deploy a SearchRefiner docker setup on the SRA server
docker-deploy:
	docker run --rm --net=search-refiner-net --publish=8001:4853/tcp --publish=8080:80/tcp  --name=sr-1 ielab-searchrefiner
	# docker run --rm --net=search-refiner-net --publish=8002:4853/tcp --publish=8080:80/tcp --name=sr-2 ielab-searchrefiner
	# docker run --rm --net=search-refiner-net --publish=8003:4853/tcp --publish=8080:80/tcp --name=sr-3 ielab-searchrefiner
	# docker run --rm --net=search-refiner-net --publish=8004:4853/tcp --publish=8080:80/tcp --name=sr-4 ielab-searchrefiner

# Run the SearchRefiner in the foreground on the local terminal
docker-run:
	-docker run --rm --net=search-refiner-net --publish=8001:4853/tcp --publish=8080:80/tcp --name=sr-1 ielab-searchrefiner

# Dial into a paused SearchRefiner instance (use ./server to run)
docker-shell:
	-docker run --rm --net=search-refiner-net --publish=8001:4853/tcp --publish=8080:80/tcp --name=sr-1 -it ielab-searchrefiner /bin/sh
