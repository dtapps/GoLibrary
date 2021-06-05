#使用vendor文件夹构建
export GO15VENDOREXPERIMENT=1
export GOPROXY=https://goproxy.cn,direct
##在工作目录创建源文件夹
mkdir -p $GOPATH/src/dtapps/golibrary/
##拷贝代码到创建好的目录
cp -rf . $GOPATH/src/dtapps/golibrary/
##未使用vendor时，下载依赖包
go get
##进入项目
cd $GOPATH/src/dtapps/golibrary/
##构建，在GOPATH下生成构建包
go mod tidy
##列出文件
ls -lh