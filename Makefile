dep-up:
	docker-compose up -d --force-recreate  mongo-primary mongo-secondary mongo-express

dep-stop:
	docker-compose down && docker system prune -f && docker volume prune -f && docker container prune -f

dep-logs: 
	docker-compose logs -f

run:
	go run main.go

app-run:
	go mod vendor
	docker-compose up --force-recreate  mongo-app