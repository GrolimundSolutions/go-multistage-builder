init:
	@echo make bad
	@echo make good

bad:
	docker build -f .Dockerfile_Bad -t bad_go_image .

good:
	docker build -f .Dockerfile_Good -t good_go_image .


.PHONY: all