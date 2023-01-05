package parser

import "testing"

func TestTransformNameGo(t *testing.T) {
	type testcase struct {
		name   string
		expect string
	}

	cases := []testcase{
		{"post_id", "PostID"},
		{"nsfw", "NSFW"},
		{"test_url", "TestURL"},
		{"some_complex_name_with_id_and_nsfw_and_url", "SomeComplexNameWithIDAndNSFWAndURL"},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := TransformNameGo(testcase.name)
			if got != testcase.expect {
				t.Errorf("Expected %s, got %s", testcase.expect, got)
			}
		})
	}
}

func TestTransformTypeGo(t *testing.T) {
	type testcase struct {
		typeName string
		expect   string
	}

	cases := []testcase{
		{"i16", "int16"},
		{"Option<Vec<i64>>", "Optional[[]int64]"},
		{"Url", "string"},
		{"Sensitive<String>", "string"},
	}

	for _, testcase := range cases {
		t.Run(testcase.typeName, func(t *testing.T) {
			got := TransformTypeGo(testcase.typeName)
			if got != testcase.expect {
				t.Errorf("Expected %s, got %s", testcase.expect, got)
			}
		})
	}
}
