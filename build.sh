gf pack resource,manifest internal/packed/data.go -n packed -y
gf build main.go -n porkast-crawler -trimpath -a amd64 -s linux,darwin -p ./bin
rm -f internal/packed/data.go 
