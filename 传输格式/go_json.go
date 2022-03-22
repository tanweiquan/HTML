package main

import (
	"encoding/json"
	"fmt"
)

//反射
//反射是指在程序运行期对程序本身进行访问和修改的能力。
//支持反射的语言可以在程序编译期将变量的反射信息，如字段名、类型信息、结构体信息等整合到可执行文件中，
//并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改他们。
//json
/* json是一种“轻量级”数据格式
做个对比吧，比如xml数据，同为文本格式的内容，xml数据每个数据都有标签套在外面，而json数据只需要外面加个 []或者{}，里面的数据用,隔开，键和值中间只要一个：隔开就行。
要是数据庞大的话，无论怎么说，几个简单的符号都比一堆标签更简单，占内存更小。 */

/*1、reflect.Typeof()和reflect.Valueof()
接口类型的变量底层是分为两个部分：动态类型和动态值
reflect.Typeof(参数)拿到动态类型(如mian.xx类型/其他包.xx类型)。
而reflect.Typeof(参数).Kind()拿到的是具体类型(struct)。

reflect.Valueof(参数)拿到动态值
*/

/* 2、NumField()和Field()
NumField()方法获取一个struct所有的fields。
reflect.Valueof(参数).NumField()拿到该结构体里所有的字段的数量

Field(i int)获取指定第i个field的reflect.Value。
reflect.Valueof(参数).Field(i)拿到第i个结构体元素的值
*/

/* 3、Int() 和String()
reflect.Valueof(参数).Int()
从reflect.Value提取对应值转换成int64类型

reflect.Valueof(参数).String()
从reflect.Value提取对应值转换成string类型
*/
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	/*
		Json是一种异常简单易懂的数据格式，关于json的规定，仅仅如下而已：
		1、并列的数据之间用逗号（", "）分隔。
		2、映射用冒号（": "）表示。
		3、并列数据的集合（数组）用方括号("[]")表示。
		4、映射的集合（对象）用大括号（"{}"）表示。
		5、键值对的键要用双引号包裹起来（"key"）
	*/
	/*
		前端取json格式的值的方式：
		假如返回的是json，如{"adb":{"xy":23},"cxz":2}或{"bys":[10,16,3,4,5],"jtc":123}
		json对象：如拿到xy对应的值23，方法是 adb.xy
		json数组：如拿到bys对应数组的值16： bys[1]
	*/
	/*
	    json数组：str := `[111,"aaa",123,"xxy","ccc",26]`
	   	json对象：str := `{"name"："周琳","age"：26}`
	   	json对象里的元素是数组：
	   	str := `{"name"："周琳","hello"：[123,"sdx","sdxc",1234]}`
	*/
	/*
		json对象首先是字符串，它遵循字符串的操作规则，并且里面的key是有双引号的。
		而go对象首先它是go语言的对象，它则遵循的go本身语言的操作规则
	*/

	p1 := Person{
		Name: "周玲",
		Age:  18,
	}
	//序列化用json.Marshal()
	b, err := json.Marshal(p1) //序列化---这里要把p1这个变量取到json包里,所以字段名要大写
	if err != nil {
		fmt.Printf("marshal failed，err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b)) //"{\"Name\":\"周玲\",\"Age\":18}"
	fmt.Println(string(b))         //{"Name":"周玲","Age":18}//这种就是json格式的字符串了

	// 反序列化：拿到json格式字符串，并抓换为go语言格式的对象并存起来
	str := `{"name":"周琳","age":26}` //json对象
	var p Person                    //这里的person的type name:person,kind name:struct
	_ = json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

	stx := `[124,1234,6789,"xsg",1246]` //json数组
	var x interface{}
	_ = json.Unmarshal([]byte(stx), &x)
	fmt.Println(x)
}
