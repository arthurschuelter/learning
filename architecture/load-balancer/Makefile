start:
	go run server-1/main.go &
	go run server-2/main.go &
	go run server-3/main.go &
	nginx -c "$PWD/nginx/nginx.conf"

stop:
	nginx -s quit -c "$PWD/nginx/nginx.conf"