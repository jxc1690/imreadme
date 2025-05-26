## 1. 发起电话请求

1客户端发送 1号请求2号
```json
{
  "typecode": 1,
  "typecode2": 9,
  "formid": 1,
  "toid": 2,
  "msg": null
}
```

## 2. 回应
2收到请求 做出回应 如果apply不是0则直接发送给1号该消息 如果为0则继续下面步骤
```json
{
  "typecode": 1,
  "typecode2": 10,
  "formid": 2,
  "toid": 1,
  "msg": {
    "apply":0
  }
}
```

## 3. 服务器发送双方房间号
```json
{
  "typecode": 1,
  "typecode2": 11,
  "formid": 0,
  "toid": 1,
  "msg": {
    "room":"xxxxx"
  }
}
{
  "typecode": 1,
  "typecode2": 11,
  "formid": 0,
  "toid": 2,
  "msg": {
    "room":"xxxxx"
  }
}
```
## 4. 客户端开始走例子的流程
将东西封装在结构中 这个为快速的接口 省略了很多的流程
```json
{
  "typecode": 1,
  "typecode2": 11,
  "formid": 1,
  "msg": {
    "room":"xxxxx",
    "roomMsg": {
      ...
    }
  }
}
```