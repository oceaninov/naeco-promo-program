deploy-sit:
	podman build -t 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-sit-$(shell git rev-parse --short HEAD) .
	podman push 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-sit-$(shell git rev-parse --short HEAD)
	podman rm nobita-promo-program-sit -f
	podman run --pod promo-engine --rm --name nobita-promo-program-sit -dt 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-sit-$(shell git rev-parse --short HEAD)
	podman image prune -a

deploy-uat:
	podman build -t 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-uat-$(shell git rev-parse --short HEAD) .
	podman push 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-uat-$(shell git rev-parse --short HEAD)
	podman rm nobita-promo-program-uat -f
	podman run --pod promo-engine --rm --name nobita-promo-program-uat -dt 35.219.77.34:8082/nbdg-promo/nobita-promo-program:1.0.0-uat-$(shell git rev-parse --short HEAD)
	podman image prune -a