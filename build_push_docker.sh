./build.sh

echo "remove guoshaofm-cralwer image"
docker rmi beegedelow/guoshaofm-crawler

./build_docker.sh

docker push beegedelow/guoshaofm-crawler
