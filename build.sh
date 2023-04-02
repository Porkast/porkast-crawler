gf pack resource,manifest internal/packed/packed.go -n packed -y
gf build main.go -n guoshaofm-crawler -trimpath -a amd64 -s linux,darwin -p ./bin