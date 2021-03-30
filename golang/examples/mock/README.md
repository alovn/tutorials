# mockgen

## install

    # Go version < 1.16
    GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0

    # Go 1.16+
    go install github.com/golang/mock/mockgen@v1.5.0

## Useage

    mockgen -source demo.go -destination demo_mock.go -package demo

    -source： 指定接口文件
    -destination: 生成的文件名
    -package:生成文件的包名
    -imports: 依赖的需要import的包
    -aux_files:接口文件不止一个文件时附加文件
    -build_flags: 传递给build工具的参数
