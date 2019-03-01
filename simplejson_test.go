package simplejson

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var jsonStr = map[string]string{
	"simpleObject":      `{"name": "test1"}`,
	"simpleArray":       `["test2.01", "test2.02"]`,
	"simpleArrayObject": `[{"name": "test3.小明"},{"name": "test3.小花"}]`,
	"userProfileAddressMobile": `
{
	"id": 10,
	"name": "test4",
	"age": 18,
	"address": {
		"country": "中国",
		"city": "成都"
	},
	"mobile": [
		"15400012301",
		"15400012302",
		15400012303
	]
}
`,
	"simpleArrayArray": `[["test5.01"], ["test5.02", "test5.03"]]`,
	"userProfile": `
{
	"id": 10,
	"name": "test4",
	"age": 18,
	"address": {
		"country": "中国",
		"city": "成都"
	},
	"mobile": [
		"15400012301",
		"15400012302",
		15400012303
	],
	"books": [
		{
			"bookNo": "book1",
			"bookName": "c程序设计",
			"authors": [
				"小明",
				"哈哈"
			]
		},
		{
			"bookNo": "book2",
			"bookName": "go程序设计",
			"authors": [
				"小红",
				"小黑"
			]
		}
	]
}
`,
}

// TestNewJson 测试创建json
func TestNewJson(t *testing.T) {
	body := jsonStr["simpleObject"]
	m := map[string]interface{}{
		"name": "test1",
	}
	// 创建json对象
	j, err := NewJson([]byte(body))
	assert.Equal(t, err, nil)
	assert.Equal(t, j, &Json{m})
}

// TestNew 测试New函数
func TestNew(t *testing.T) {
	j := New()
	assert.Equal(t, j, &Json{map[string]interface{}{}})
}

// TestJson_Get 测试获取值json对象值
func TestJson_Get(t *testing.T) {
	body := jsonStr["userProfileAddressMobile"]
	j, err := NewJson([]byte(body))
	v, err := j.Get("name").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, "test4")

	v1, err := j.Get("address").Get("city").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, v1, "成都")
}

// 测试获取多级键的值
func TestJson_Get2(t *testing.T) {
	body := jsonStr["userProfileAddressMobile"]
	j, err := NewJson([]byte(body))
	v, err := j.Get("address").Get("city").String()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, "成都")
}

// TestJson_GetArrayIndex 测试获取json数组值
func TestJson_GetArrayIndex(t *testing.T) {
	body := jsonStr["simpleArray"]
	j, err := NewJson([]byte(body))
	v, err := j.GetArrayIndex(0).String()
	assert.Equal(t, err, nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, v, "test2.01")

	v1, err := j.GetArrayIndex(1).String()
	assert.Equal(t, err, nil)
	assert.Equal(t, v1, "test2.02")
}
// TestJson_GetArrayIndex2 测试获取多层数组值
func TestJson_GetArrayIndex2(t *testing.T) {
	body := jsonStr["simpleArrayArray"]
	j, err := NewJson([]byte(body))
	v, err := j.GetArrayIndex(1).GetArrayIndex(0).String()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, "test5.02")
}

// 测试简单的json对象操作
func TestJson_MustSetSimpleObject(t *testing.T) {
	body := jsonStr["simpleObject"]
	j, err := NewJson([]byte(body))
	j.MustSet("test1.change value", "name")
	v, err := j.Get("name").String()
	assert.Equal(t, v, "test1.change value")
	assert.Equal(t, err, nil)
}

// 测试简单MustSet数组操作
func TestJson_MustSetSimpleArray(t *testing.T) {
	body := jsonStr["simpleArray"]
	j, err := NewJson([]byte(body))
	j.MustSet("test2.02.change value", 1)
	v, err := j.GetArrayIndex(1).String()
	assert.Equal(t, v, "test2.02.change value")
	assert.Equal(t, err, nil)
}



// 测试设置多级json对象数组值
func TestJson_MustSetKeys(t *testing.T) {
	body := jsonStr["userProfile"]
	j, err := NewJson([]byte(body))
	log.Println("aaaaaa", j.data)
	j.MustSet("小明set.object.array.object.array", "books", 0, "authors", 0)
	j.MustSet("小明setNew.object.array.object.array", "books", 0, "authors", -1)
	j.MustSet("c程序设计.set.object.array.object", "books", 0, "bookName")

	v, err := j.Get("books").GetArrayIndex(0).Get("authors").GetArrayIndex(0).String()
	assert.Equal(t, v, "小明set.object.array.object.array")
	assert.Equal(t, err, nil)

	v1, err := j.Get("books").GetArrayIndex(0).Get("authors").GetArrayIndex(2).String()
	assert.Equal(t, v1, "小明setNew.object.array.object.array")
	assert.Equal(t, err, nil)

	v2, err := j.Get("books").GetArrayIndex(0).Get("bookName").String()
	assert.Equal(t, v2, "c程序设计.set.object.array.object")
	assert.Equal(t, err, nil)
}