sleep 1
docker-compose up -d vault

sleep 1
export VAULT_ADDR='http://127.0.0.1:8200'

sleep 1
export VAULT_TOKEN=myroot



