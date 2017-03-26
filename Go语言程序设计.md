# Go语言程序设计

## 一、概述

* 大道至简的设计哲学

    * 没有继承、构造、析构、虚构、函数重载、默认参数等。
    * 少即是多，只有通过简洁的设计，才能让一个系统保持稳定、安全和持续的进化。
    * Go项目是在Google公司维护超级复杂的几个软件系统遇到的一些问题的反思。

* 为并发而生
    
    * 语言层次支持并发模型：goroutine
    ```
    go func(){
        ...
    }()
    ```
    * goroutine比线程更轻量，可以轻松跑上万个goroutine

* 支持垃圾回收

    消除了并发编程中的对象生命周期管理的负担

* 非侵入式接口

    * 鸭子类型
    * 支持接口查询
    ```
    if v, ok := v.(IFile); ok { 
        ...
    }
    ```
* 极度简化但完备的面向对象方法

    废弃了大量OOP特性：等

* 标准化的错误处理规范

    * 内置error
    * defer语句编写异常安全代码

* 功能内聚

    匿名组合完成继承
    ```
    type Foo struct {
        *Base
        ...
    }
    ```
* 适合云计算
    * 性能大幅领先python、ruby、php等脚本语言，接近C、C++
    * 腾讯、阿里、京东、360、网易、新浪、金山、豆瓣等都有团队对go做服务端开发进行实践
    * 目前用Go实现比较火的应用：Docker、TiDB
    * 2016再次获得年度编程语言


## 二、布尔与数字类型

### 开始

* 命令行运行 go version，如果出错则把如下脚本加入 ～/.profile
```
export GOROOT=/HOME/opt/go
export PATH=$PATH:$GOROOT/bin
```
* 编译 go build，编译速度秒杀C++几条街
* go程序做脚本用：gonow gorun
* IDE: VS Code、LiteIDE、Gogland

### 基础
* 关键字

    break default func interface select case defer go map 
    struct chan else goto package switch const fallthrough if 
    range type contiue for import return const var

* 预定义标识符

    * 内建常量: true false iota nil

    * 内建类型: int int8 int16 int32 int64
            uint uint8 uint16 uint32 uint64 uintptr
            float32 float64 complex128 complex64
            bool byte rune string error

    * 内建函数: make len cap new append copy close delete
            complex real imag
            panic recover

* 常量、变量
```
const(
    Cyan = iota
    Magenta
    Yelow
)
```

* 不支持隐式类型转换，不同类型必须显式类型转换

    type(value)

* 大数值类型

    * big.Int 
    * big.Rat

* 不支持操作符重载

* ++、--只支持后缀方式

* 除非特殊说明，math包所有函数都用float64

## 三、字符串

* unicode码点用rune表示（4字节）
* 字符串用双引号 “ 或 反引号` 创建
* []rune(s) 将字符串转换成Unicode码点
* += 拼凑低效，建议用 strings.Join 或 bytes.Buffer
* 字符串保存为utf8，用for...range遍历，非ASCII索引更新的步长将超过1个字节
（建议先转[]rune），
* utf8.DecodeRuneInString()获取第一个字符的位置和大小
* strings.Map()可用来替换或去掉字符串中的字符（返回负数则原来字符删除）
* 相关包：fmt、strings、strconv、utf8、unicode、regexp

## 四、集合类型

* 对于chan、func、map、slice变量，持有的为引用，其他都持有值
* 传递数组按值传递，代价非常大，通常不用数组，用slice
* 创建变量同时获取指针：new(Type)、&Type{}
* 数组创建方式：
```
[len]Type
[len]Type{v1, v2, v3..., vn}
[...]Type{v1, v2, v3..., vn}
```
* s[:cap(s)] :增加切片长度到其容量
* 切片创建方式：
```
make([]Type, len, cap)
make([]Type, len)
[]Type{}
[]Type{v1, v2, v3..., vn}
```
* 使用...加在切片后用于把切片当成多个元素(同不定长参数正好相反)
```
s = append(s, u[2:5]...)

```
* 相关函数：append、copy、len、cap
* 相关包：sort
* map的操作
    ```
    m[k] = v1
    delete(m, k)
    v := m[k]
    v, found := m[k]
    len(m)
    ```
* map比切片的字节索引慢2个数量级（100倍），不过也足够快
* map的创建方式：
```
make(map[KeyType]ValueType, cap)
make(map[KeyType]ValueType)
map[KeyType]ValueType{}
map[KeyType]ValueType{k1: v1, k2: v2..., kn: vn}
```
* struct 可以作为map的key，只要它的成员都支持==和!=运算即可

## 五、过程式编程

## 六、面向对象编程

## 七、并发编程

## 八、文件处理

## 九、包