LOGS_DIR=/home/guoshaofm-crawler/logs
if [[ ! -e $LOGS_DIR ]]; then
    mkdir -p $LOGS_DIR
elif [[ ! -d $LOGS_DIR ]]; then
    echo "$LOGS_DIR already exists but is not a directory" 1>&2
fi
docker run --name guoshaofm-crawler --network host --log-opt max-size=500m --log-opt max-file=3 -v $LOGS_DIR:/app/logs -d beegedelow/guoshaofm-crawler