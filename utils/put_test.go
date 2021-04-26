package utils

import (
	"sp/test"
	"testing"
)

func TestPut(t *testing.T) {
	data := CreateEmptyMap()
	Put(data, "name", "steven")
	test.AssertEqual(t, (*data).(map[string]interface{}),
		map[string]interface{}{"name": "steven"})

	data = CreateEmptyMap()
	Put(data, "year", 2021)
	test.AssertEqual(t, (*data).(map[string]interface{}),
		map[string]interface{}{"year": 2021})

	data = CreateEmptyMap()
	Put(data, "nested.year", 2021)
	Put(data, "nested.hello", "world")
	test.AssertEqual(t, (*data).(map[string]interface{}),
		map[string]interface{}{
			"nested": map[string]interface{}{
				"year":  2021,
				"hello": "world",
			},
		})

	data = CreateEmptyMap()
	Put(data, "nested.year", 2020)
	Put(data, "nested.year", 2021)
	test.AssertEqual(t, (*data).(map[string]interface{}),
		map[string]interface{}{
			"nested": map[string]interface{}{
				"year": 2021,
			},
		})

	data = CreateEmptyMap()
	Put(data, "nested.alphabet.[0]", "a")
	Put(data, "nested.alphabet.[+]", "b")
	Put(data, "nested.alphabet.[+]", "e")
	Put(data, "nested.alphabet.[2:]", "c")
	Put(data, "nested.alphabet.[:2]", "d")
	test.AssertEqual(t, (*data).(map[string]interface{}),
		map[string]interface{}{
			"nested": map[string]interface{}{
				"alphabet": []string{"a", "b", "c", "d", "e"},
			},
		})

	data = CreateEmptyArray()
	Put(data, "[+]", "a")
	Put(data, "[+]", "b")
	Put(data, "[+]", "e")
	Put(data, "[2:]", "c")
	Put(data, "[:2]", "d")
	test.AssertEqual(t, (*data).([]interface{}),
		[]interface{}{"a", "b", "c", "d", "e"})

	data = CreateEmptyArray()
	Put(data, "[+]", map[string]string{"english": "hello"})
	Put(data, "[+]", map[string]string{"german": "hallo"})
	Put(data, "[+]", map[string]string{"polish": "cześć"})
	Put(data, "[2:]", map[string]string{"japanese": "kon'nichiwa"})
	Put(data, "[:2]", map[string]string{"spanish": "hola"})
	test.AssertEqual(t, (*data).([]interface{}),
		[]interface{}{
			map[string]string{"english": "hello"},
			map[string]string{"german": "hallo"},
			map[string]string{"japanese": "kon'nichiwa"},
			map[string]string{"spanish": "hola"},
			map[string]string{"polish": "cześć"},
		})
}
