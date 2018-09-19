# go-esport-common

基础公共目录

### 编码风格：
 1. 命名大体和google go语言基本一致，
    比如：大小写开头分别代表该函数或者变量可以不可以被其他包调用
	比如：命名用高低位（EMyDef or eMyDef）
    不同的如下：
    - 1.1 go语言没有原生的枚举定义， 但是可以用const+iota去类似实现。
	    其他语言标准枚举定义为全大写加下划线，这种存在着中国人看大写不习惯的问题，
        所以我改良为每个词首字大写，下划线连接各词，或者E开头首字大字，
        ex： EMyDef  or  My_Def
    - 1.2 类的函数，必须以this代表该类，
	    ex：
		type CXX struct {
		}
		func (this *CXX) GetThis() *CXX {
			return this
        }
    - 1.3 类名需用C开头，结构体以T开头，
	    类与普通结构的区别在于，普通结果没有逻辑的操作方法，最多只有类似初始化的函数
 2. 代码在不同情况采用不同的编写标准，在保证功能的前提下，
    在效率要求很高的情况下，以执行效率为先，编码优美简洁优雅易懂为其次。
    在业务逻辑复杂的情况下，以代码优雅通俗易懂为先，执行效率其次
 3. 注释采用`doxygen`的风格，这是大部分语言比较通用的写法，且可以方便生成文档
 4. 代码目录尽量控制在两层，且不要建太多的文件夹，比如一个文件一个文件夹
    ex： 代码根目录为 C:\src, 那尽量在C:\src\a\b
 5. 文件行数尽量不用超过`1200`行,如果超过,可以平行分拆文件
 6. 可以使用golint检测代码

### 功能划分
#### 1. 日志功能
* 用法: 当前目录
```
 import "github.com/yanue/go-esport-common"

 ...
 common.Logs.Info
 common.Logs.Error
 ...
```
* With用法见zap库

#### 2. 工具类
- 目录: `util`
- 用法:
```
 import "github.com/yanue/go-esport-common/util"

 ...
 util.ToString()
 util.ToInt()
 ...

```

#### 3. 错误码
- 目录: `errcode`
- 用法:
```
 import "github.com/yanue/go-esport-common/errcode"

 ...
 errcode.No_Error
 errcode.ErrAccountNotExist
 ...

```

#### 4. 验证处理
- 目录: `validator`
- 用法:
```
 import "github.com/yanue/go-esport-common/validator"

 ...
 validator.Verify.IsUrl()
 validator.Verify.IsNickname()
 ...

```

#### 5. proto定义及生产
- 目录: `proto`
- 用法:
```
 import "github.com/yanue/go-esport-common/proto"

 ...
 // 根据定义情况
 proto.PMessage{}
 ...

```

