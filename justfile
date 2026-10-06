default:
	just --list

clean:
	-docker rm searchrefiner.sr-accelerator.com-a
	-docker rm searchrefiner.sr-accelerator.com-b
	-docker image rm ielab-searchrefiner

stop:
	-runwisp stop searchrefiner.sr-accelerator.com-*
	-docker stop searchrefiner.sr-accelerator.com-a
	-docker stop searchrefiner.sr-accelerator.com-b

build:
	cd /sites/searchrefiner.sr-accelerator.com
	docker build -t ielab-searchrefiner .

restart:
	runwisp restart searchrefiner.sr-accelerator.com-*

rebuild: stop clean build restart
