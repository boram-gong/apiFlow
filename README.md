**通用接口服务**

### JSON转换接口


**json转换地址  GET /lgi/adapter/json**

##### 请求参数

| 参数 | 含义                                       |
| ---- | ------------------------------------------ |
| name | 要使用的规则名称                           |
| data | 要处理的json（原始json格式，非字符串格式） |

##### 返回参数

| 参数 | 含义                                           |
| ---- | ---------------------------------------------- |
| code | 响应好（200表示成功，40x表示用户请求存在问题） |
| msg  | 响应描述                                       |
| data | json变化结果数据                               |

请求示例

```json
{
    "name": "demo2",
    "data": {
        "head":"head",
        "demo": {
            "key":"value"
        },
        "list": [
                {"d1":"d1"},
                {"d2":"d2"}
        ]
    }
}
```

返回示例

```json
{
    "code": 200,
    "msg": "成功",
    "data": {
        "demo": {
            "key": "value",
            "head": "value"
        },
        "list1": [
            {
                "d1": "d1"
            },
            {
                "d2": "d2"
            }
        ]
    }
}
```

**获取规则  GET /lgi/adapter/rule**

请求示例

```
获取所有规则    GET 127.0.0.1:29989/lgi/responseAdapter/rule 
获取指定id规则  GET 127.0.0.1:29989/lgi/responseAdapter/rule?id=1 
```

响应示例

```json
{
    "code": 200,
    "msg": "成功",
    "data": [
        {
            "id": 11,
            "rule_name": "demo2",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        },
        {
            "id": 12,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey2"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
    ]
}
```

```json
{
    "code": 200,
    "msg": "成功",
    "data": {
        "id": 11,
        "rule_name": "demo2",
        "rules": [
            {
                "key": "demo.key",
                "operation": "rename",
                "content": "demo.newKey"
            }
        ],
        "stat": 1,
        "start_time": "",
        "end_time": ""
    }
}
```

**新增规则  POST /lgi/adapter/rule**

请求体（json）参数

| 参数       | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| id         | post新增规则，id为0，必有参数                                |
| rule_name  | 规则名称，必有参数，唯一                                     |
| rules      | 详细规则数组                                                 |
| stat       | 状态，有效为1，无效为其他数字，必有字段                      |
| start_time | 规则有效起始时间，格式有两种: "2006-01-02 15:04:05" 和 "15:04:05"，，举例："2006-01-02 15:04:05"~"2006-01-03 15:04:05" 表示这段时间有效，"15:00:00"~"20:00:00" 表示每天这个时间段有效 |
| end_time   | 规则有效结束时间                                             |

详细规则参数（可参考规则转换说明）

| 参数      | 含义              |
| --------- | ----------------- |
| key       | 表示要操作的键    |
| operation | 对key（键）的操作 |
| content   | 操作的内容        |

请求示例

```json
        {
            "id":0,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
```

响应示例（响应的data中是当前所有的规则）

```json
{
    "code": 200,
    "msg": "成功",
    "data": [
        {
            "id": 11,
            "rule_name": "demo2",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        },
        {
            "id": 13,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
    ]
}
```

**修改规则  PUT  /lgi/adapter/rule**

修改规则和新增规则几乎一样，就是id需要明确

**删除规则  DELETE  /lgi/adapter/rule**

请求示例

```
需指定规则id  DELETE 127.0.0.1:29989/lgi/responseAdapter/rule?id=1 
```

**重新加载规则  GET /lgi/adapter/re**

