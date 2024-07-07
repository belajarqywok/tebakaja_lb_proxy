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


# 
#    --- Production Testing (Proxy 0) ---
# 

stock-list-prod-0:
	curl -X GET $(PROD_ENDPOINT_0)/stock/lists

stock-prediction-prod-0:
	curl -X POST $(PROD_ENDPOINT_0)/stock/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

crypto-list-prod-0:
	curl -X GET $(PROD_ENDPOINT_0)/crypto/lists

crypto-prediction-prod-0:
	curl -X POST $(PROD_ENDPOINT_0)/crypto/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

natcurr-list-prod-0:
	curl -X GET $(PROD_ENDPOINT_0)/national-currency/lists

natcurr-prediction-prod-0:
	curl -X POST $(PROD_ENDPOINT_0)/national-currency/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

# 
#    --- Production Testing (Proxy 1) ---
# 

stock-list-prod-1:
	curl -X GET $(PROD_ENDPOINT_1)/stock/lists

stock-prediction-prod-1:
	curl -X POST $(PROD_ENDPOINT_1)/stock/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

crypto-list-prod-1:
	curl -X GET $(PROD_ENDPOINT_1)/crypto/lists

crypto-prediction-prod-1:
	curl -X POST $(PROD_ENDPOINT_1)/crypto/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

natcurr-list-prod-1:
	curl -X GET $(PROD_ENDPOINT_1)/national-currency/lists

natcurr-prediction-prod-1:
	curl -X POST $(PROD_ENDPOINT_1)/national-currency/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

# 
#    --- Production Testing (Proxy 2) ---
# 

stock-list-prod-2:
	curl -X GET $(PROD_ENDPOINT_2)/stock/lists

stock-prediction-prod-2:
	curl -X POST $(PROD_ENDPOINT_2)/stock/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

crypto-list-prod-2:
	curl -X GET $(PROD_ENDPOINT_2)/crypto/lists

crypto-prediction-prod-2:
	curl -X POST $(PROD_ENDPOINT_2)/crypto/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"

natcurr-list-prod-2:
	curl -X GET $(PROD_ENDPOINT_2)/national-currency/lists

natcurr-prediction-prod-2:
	curl -X POST $(PROD_ENDPOINT_2)/national-currency/prediction \
		-H "Content-Type: application/json" \
		-d "{\"days\": 7, \"currency\": \"BTC-USD\"}"
		