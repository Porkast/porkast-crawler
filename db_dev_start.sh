docker run -d \
  --name porkast-postgres \
  -e POSTGRES_USER=porkast \
  -e POSTGRES_PASSWORD=1qaz!QAZ \
  -e POSTGRES_DB=porkastdb \
  -p 5432:5432 \
  --net=host \
  -v /root/develop/porkast/db-data:/var/lib/postgresql/data \
  beegedelow/porkcast-postgresql:latest