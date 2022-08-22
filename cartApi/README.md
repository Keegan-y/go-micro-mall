# CartApi api 

当前服务 名称为 CartApi 类型 api 

创建初始化模版请使用

```
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) -e ICODE=xxxxxx cap1573/cap-micro new git.imooc.com/cap1573/cartApi
```

## 快速开始

- [配置信息](#配置信息)
- [使用](#使用)

## 配置信息

- 服务名称: go.micro.api.cartApi
- 类型: api
- 简称: cartApi

 

## 使用
根据 proto 自动生成
```
make proto
```

编译
```
make proto
```

构建镜像
```
make docker
```