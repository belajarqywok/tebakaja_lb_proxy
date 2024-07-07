DEV_ENDPOINT=http://192.168.137.1:7860

PROD_ENDPOINT_0=https://qywok-tebakaja-proxy-space-0.hf.space
PROD_ENDPOINT_1=https://qywok-tebakaja-proxy-space-1.hf.space
PROD_ENDPOINT_2=https://qywok-tebakaja-proxy-space-2.hf.space

start:
	go run main.go

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

endpoints_check:
	bash -c 'chmod +x endpoints_test.sh && ./endpoints_test.sh'

# 
#    --- Development Testing ---
# 

stock-list-test:
	curl -X GET $(DEV_ENDPOINT)/stock/lists

stock-prediction-test:
	curl -X POST $(DEV_ENDPOINT)/stock/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"

crypto-list-test:
	curl -X GET $(DEV_ENDPOINT)/crypto/lists

crypto-prediction-test:
	curl -X POST $(DEV_ENDPOINT)/crypto/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"

natcurr-list-test:
	curl -X GET $(DEV_ENDPOINT)/national-currency/lists

natcurr-prediction-test:
	curl -X POST $(DEV_ENDPOINT)/national-currency/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 2, \"currency\": \"BTC-USD\"}"

