ps:
	sudo docker-compose ps -a

up:
	sudo docker-compose up -d

start:
	sudo docker-compose start

restart:
	sudo docker-compose restart

stop:
	sudo docker-compose stop

down:
	sudo docker-compose down

deploy:
	podman build -t 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-$(shell git rev-parse --short HEAD) .
	podman push 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-$(shell git rev-parse --short HEAD)
	podman rm nobita-promo-program-dev -f
	podman run --pod promo-engine --rm --name nobita-promo-program-dev -dt 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-$(shell git rev-parse --short HEAD)
	podman image prune -a