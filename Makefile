.PHONY: restart
restart:
	sudo docker-compose down
	sudo rm -rf mysql
	sudo rm -rf phpmyadmin
	sudo docker compose up -d --build