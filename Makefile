push:
	git push
	docker build . -t joeyliu086/hushtell
	docker push joeyliu086/hushtell

update_docker:
	docker build . -t joeyliu086/hushtell
	docker push joeyliu086/hushtell

check_swagger:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve:
	go run .