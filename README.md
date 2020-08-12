## 斗鱼-api

[![Build Status](https://travis-ci.com/Birjemin/douyuapi.svg?branch=master)](https://travis-ci.com/Birjemin/douyuapi) [![Go Report Card](https://goreportcard.com/badge/github.com/birjemin/douyuapi)](https://goreportcard.com/report/github.com/birjemin/douyuapi) [![codecov](https://codecov.io/gh/Birjemin/douyuapi/branch/master/graph/badge.svg)](https://codecov.io/gh/Birjemin/douyuapi)


[开发者中心](https://open.douyu.com/source)

### 引入方式
```
go get github.com/birjemin/douyuapi
```

### 接口列表

- [token获取接口](https://open.douyu.com/source/api/8) ✅
- [直播视频流](https://open.douyu.com/source/api/9) ✅
- 房间信息
   - [指定房间信息](https://open.douyu.com/source/api/15)✅
   - [批量房间获取](https://open.douyu.com/source/api/25)⚠️
- 分类信息
   - [一级分类](https://open.douyu.com/source/api/17)⚠️
   - [二级分类](https://open.douyu.com/source/api/18)⚠️
   - [三级分类](https://open.douyu.com/source/api/19)⚠️
- [批量获取idfa信息](https://open.douyu.com/source/api/26)⚠️
- 分类列表房间信息
   - [直播间列表信息](https://open.douyu.com/source/api/22)✅
   - [分类全量在线房间](https://open.douyu.com/source/api/57)⚠️
- 点播视频
   - [获取点播二级分类列表](https://open.douyu.com/source/api/32)⚠️
   - [点播二级分类视频列表](https://open.douyu.com/source/api/33)⚠️
   - [点播UP主视频列表](https://open.douyu.com/source/api/34)⚠️
   - [点播推荐池视频列表](https://open.douyu.com/source/api/35)⚠️
   - [点播拉流接口](https://open.douyu.com/source/api/36)⚠️
   - [点播视频下载接口](https://open.douyu.com/source/api/38)⚠️
- 智能分类
   - [智能分类列表](https://open.douyu.com/source/api/52)⚠️
   - [智能分类房间详情](https://open.douyu.com/source/api/53)⚠️
- [直播场次信息](https://open.douyu.com/source/api/54)⚠️
- 房间弹幕（一期）
   - [拉取弹幕](https://open.douyu.com/source/api/65)✅
   - [接入弹幕](https://open.douyu.com/source/api/66)✅
- [直播音频流](https://open.douyu.com/source/api/67)⚠️

### 使用方式

- 示例

```golang
    httpClient := &utils.HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	token := &Token{
		BaseClient: BaseClient{
			Client: httpClient,
			Secret: "test-secret",
			AID:    "test-aid",
		},
	}

	timestamp := cast.ToString(utils.GetCurrTime())
	if ret, err := token.Handle(timestamp); err != nil {
		// handle err
	} else {
		if ret.Code != 0 {
			// handle err
		}
		// get ret.Data.Token
	}
	
    live := &Live{
        BaseClient: BaseClient{
            Client: httpClient,
            Secret: "test-secret",
            AID:    "test-aid",
        },
        Token: "test-token",
    }

    msg := `{"cid_type":1,"cid":1,"limit":10,"offset":0}`

    if ret, err := live.Handle(msg, cast.ToString(timestamp)); err != nil {
        // handle err
    } else {
        if ret.Code != 0 {
            // handle err
        }
        // handle
    }
```

### 备注

后续会补足新接口