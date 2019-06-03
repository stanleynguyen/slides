build:
	docker build . -t registry.heroku.com/stanslides/web
deploy:
	sudo heroku container:push web --app stanslides && sudo heroku container:release web --app stanslides
dev:
	docker run -it -p 8080:8080 -v $(shell pwd):/go/src/golang.org/x/tools/cmd/present/content/ registry.heroku.com/stanslides/web bash
