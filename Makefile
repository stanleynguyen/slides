build:
	docker build . -t registry.heroku.com/stanslides/web
deploy:
	sudo heroku container:push web --app stanslides && sudo heroku container:release web --app stanslides
