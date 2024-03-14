run: clean dock
	go build .
	./jrport.almosheriff

clean: 
	rm -f jrport.almosheriff
	sudo docker stop postgres_container

dock:
	sudo docker start postgres_container
