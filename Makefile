ENDPOINT=http://192.168.137.1:7860

start: go run main.go

haproxy-test:
	haproxy -f ./haproxy/haproxy.cfg

nginx-test:
	nginx -c nginx.conf

traefik-test:
	traefik \
		--configFile=./traefik/traefik.yaml \
		--entryPoints.web.address=":7860" \
		--entryPoints.websecure.address=":443" \
		--entryPoints.web.http.redirections.entryPoint.to=websecure \
		--entryPoints.web.http.redirections.entryPoint.scheme=https \
		--api.dashboard=true \
		--api.insecure=false

stock-list-test:
	curl -X GET $(ENDPOINT)/stock/lists

stock-prediction-test:
	curl -X POST $(ENDPOINT)/stock/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"

crypto-list-test:
	curl -X GET $(ENDPOINT)/crypto/lists

crypto-prediction-test:
	curl -X POST $(ENDPOINT)/crypto/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"

natcurr-list-test:
	curl -X GET $(ENDPOINT)/national-currency/lists

natcurr-prediction-test:
	curl -X POST $(ENDPOINT)/national-currency/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"
	