default:
	just --list

clean:
	-docker rm searchrefiner.sr-accelerator.com-a
	-docker rm searchrefiner.sr-accelerator.com-b
	-docker image rm ielab-searchrefiner

stop:
	-pm2 stop searchrefiner.sr-accelerator.com-a searchrefiner.sr-accelerator.com-b
	-docker stop searchrefiner.sr-accelerator.com-a
	-docker stop searchrefiner.sr-accelerator.com-b

build:
	cd /sites/searchrefiner.sr-accelerator.com
	docker build -t ielab-searchrefiner .

restart:
	pm2 restart searchrefiner.sr-accelerator.com-a searchrefiner.sr-accelerator.com-b

rebuild: stop clean build restart
