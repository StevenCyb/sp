package utils

import (
	"sp/test"
	"testing"
)

func TestByteIsValidJson(t *testing.T) {
	test.AssertEqual(t, ByteIsValidJson([]byte("{}")), true)
	test.AssertEqual(t, ByteIsValidJson([]byte("{]")), false)
}

func TestStringIsValidJson(t *testing.T) {
	test.AssertEqual(t, StringIsValidJson("{}"), true)
	test.AssertEqual(t, StringIsValidJson("{]"), false)
}

func TestByteIsValidToml(t *testing.T) {
	test.AssertEqual(t, ByteIsValidToml([]byte(`a = "A"`)), true)
	test.AssertEqual(t, ByteIsValidToml([]byte(`a  "A"`)), false)
}

func TestStringIsValidToml(t *testing.T) {
	test.AssertEqual(t, StringIsValidToml(`a = "A"`), true)
	test.AssertEqual(t, StringIsValidToml(`a  "A"`), false)
}

func TestByteIsValidYaml(t *testing.T) {
	test.AssertEqual(t, ByteIsValidYaml([]byte("a: A")), true)
	test.AssertEqual(t, ByteIsValidYaml([]byte(`a: [] A`)), false)
}

func TestStringIsValidYaml(t *testing.T) {
	test.AssertEqual(t, StringIsValidYaml("a: A"), true)
	test.AssertEqual(t, StringIsValidYaml(`a: [] A`), false)
}
