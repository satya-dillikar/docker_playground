cd ../python-docker
docker build --tag python-docker .
docker images
docker run python-docker
docker run -d -p 5000:5000 python-docker
docker ps
curl http://0.0.0.0:5000/
