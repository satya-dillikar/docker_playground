docker build -t test1 -f Test1.Dockerfile .
docker run test1 hello
docker run test1 printf hello

docker build -t test2 -f Test2.Dockerfile .
docker run test2 Satya
docker run test2 Satya AnotherSatya
