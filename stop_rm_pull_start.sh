echo "stop porkast-crawler"
docker stop porkast-crawler
echo "remove porkast-crawler container"
docker container rm porkast-crawler
echo "remove porkast-crawler image"
docker rmi porkast-crawler
echo "pull porkast-crawler image"
docker pull beegedelow/porkast-crawler

LOGS_DIR=/home/porkast-crawler/logs
if [[ ! -e $LOGS_DIR ]]; then
    mkdir -p $LOGS_DIR
elif [[ ! -d $LOGS_DIR ]]; then
    echo "$LOGS_DIR already exists but is not a directory" 1>&2
fi
echo "run porkast-crawler container"
docker run --name porkast-crawler --network host --log-opt max-size=500m --log-opt max-file=3 -v $LOGS_DIR:/app/logs -d beegedelow/porkast-crawler
