# 插件开发模板

```javascript
// [name:Cookie检测（demo 插件开发演示）]

// 第一行为插件名称， 在后台显示的使用会用到

// 返回数据格式
// return {
//      // 代表是否允许通过
//     "bool": true,
//      // 处理后的变量
//     "env": env
// }

// 必须以main为函数名, env为传入变量
function main(env) {
    let result = request({
        "method": "get",
        "url": "https://plogin.m.jd.com/cgi-bin/ml/islogin",
        "headers": {
            "Cookie": env,
            "User-Agent": "jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1",
            "Referer": "https://h5.m.jd.com/"
        },
        "dataType": "json",
        "timeout": 5 * 1000
    })

    if (result) {
        // 判断是否过期
        if (result["islogin"] === "1"){
            // Cookie有效
            return {
                "bool": true,
                "env": env
            }
        } else {
            // Cookie无效
            return {
                "bool": false,
                "env": "Cookie已失效"
            }
        }
    } else {
        return {
            "bool": false,
            "env": "请求失败"
        }
    }
}
```

## 封装可用方法

### *Request*

```
let result = request({
    // 请求方式（默认get）
    "method": "get",
    // 请求地址
    "url": "https://me-api.jd.com/user_new/info/GetJDUserInfoUnion",
    // 数据类型（返回数据如果是JSON那么就需要指定为json，否则默认为location）
    "dataType": "",
    // 请求头
    "headers": {
        "Cookie": env,
        "User-Agent": "jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1",
        "Referer": "https://h5.m.jd.com/"
    },
    // 请求体(body为json请求体， 二选一即可)
    "body": {},
    "formdata": {},
    // 自定义超时(单位：纳秒)(5 * 1000 = 5秒)
    "timeout": 5 * 1000
})
```

## *ReFind*方法

```
// ReFind(正则表达式, 待匹配数据) 返回：匹配结果列表（数组）（string）
let result = ReFind("pt_pin=(.*?);", "pt_pin=xxx")
```

### *consolo方法*

```
// 填写打印信息即可（string）
console.info()
console.debug()
console.warn()
console.error()
console.log()
```

### *Replace方法*

```
// 参数说明：替换文本中的关键词
// s（string）：原始字符串
// old（string）：需要替换的内容
// new（string）：替换后的内容
// count（int）：需要替换的数量，不填写默认为替换第一个
let result = Replace(s, old, new, count)
```

