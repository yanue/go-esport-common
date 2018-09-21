### 工具集

#### 格式转换相关 convert.go
 - ToInt 转换成int
 - ToString 转换成string
 - StrTo

#### http相关 http.go
 - 用法: util.Http.xxx
```
  RemoteCall
  RemoteCallWithTimeout
  RemoteCallWithTry
  RemoteDeleteWithTimeout
  RemotePost
  RemotePostOctStream
  RemotePutOctStream
  RemotePostProtoStream
  RemotePostURLEncode
  RemotePostJson
  RemoteHead
  HttpRespond
```

#### ip相关 ip.go
 - 用法: util.Ip.xxx
```
 GetActiveIP
 GetPrivateIP
 GetIP
```

#### jwt token相关 jwt.go
 - 用法: util.JwtToken.xxx
```
 Generate
 Verify
```

#### 密码相关 password.go
 - 用法: util.Password.xxx
```
 Generate
 Verify
```

#### 输出响应格式 response.go
 - 用法: util.Response.xxx
```
 WritePbResponse
 WriteErrResponse
 CheckPResult
```

#### rsa 加密解密 rsa.go
 - 用法: util.Rsa.xxx
```
 RsaEncryptPrivate
 RsaDecryptPrivate
 RsaEncryptPublic
 RsaDecryptPublic
 RsaSign
 RsaSignVerify
```

#### slice操作相关 slice.go
 - 用法: util.Slice.xxx
```
ReverseIntSlice
RemoveIntSlice
RemovesIntSlice
RemoveInt32Slice
RemovesInt32Slice
RemoveStringSlice
RemoveStringSliceEx
InsertSortStringSlice
```

#### struct操作相关 struct.go
 - 用法: util.Struct.xxx
```
ToMap
ToJsonByte
ToJsonString
```

#### 其他
 - 用法: util.xxx
 - 可以从util.go扩展或新增文件
```
```
