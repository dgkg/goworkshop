build: 
	docker build -t go-multi-stage .

run:
	docker run -p 8080:8080 --name go-multi-stage go-multi-stage

clean:
	docker stop go-multi-stage
	docker rm go-multi-stage