test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build Structs

stop:
	sudo docker-compose stop
	#скачать git bash
#makefile в main.go