
sudo docker run -d   --name roach   --hostname db   --network mynet   -p 26257:26257   -p 8080:8080   -v roach:/cockroach/cockroach-data   cockroachdb/cockroach:latest-v20.1 start-single-node   --insecure

sudo docker exec -it roach ./cockroach sql --insecure

CREATE USER postgres;
GRANT ALL ON DATABASE postgres TO postgres;

sudo docker run -it --rm -d --network mynet --name lana-sre-challenge -p 3000:3000 lana-sre-challenge