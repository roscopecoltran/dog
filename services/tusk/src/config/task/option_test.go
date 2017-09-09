package task

import (
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/rliebz/tusk/appyaml"
	yaml "gopkg.in/yaml.v2"
)

func TestOption_Dependencies(t *testing.T) {
	option := &Option{DefaultValues: valueList{
		{When: falseWhen, Value: "foo"},
		{When: appyaml.When{
			Equal: map[string]appyaml.StringList{
				"foo": {"foovalue"},
				"bar": {"barvalue"},
			},
		}, Value: "bar"},
		{When: appyaml.When{
			NotEqual: map[string]appyaml.StringList{
				"baz": {"bazvalue"},
			},
		}, Value: "bar"},
	}}

	expected := []string{"foo", "bar", "baz"}
	actual := option.Dependencies()
	if !equalUnordered(expected, actual) {
		t.Errorf(
			"Option.Dependencies(): expected %s, actual %s",
			expected, actual,
		)
	}
}

// nolint: dupl
func equalUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[string]interface{})
	for _, val := range a {
		aMap[val] = struct{}{}
	}

	bMap := make(map[string]interface{})
	for _, val := range b {
		bMap[val] = struct{}{}
	}

	return reflect.DeepEqual(aMap, bMap)
}

// TODO: Make these more accessible to other tests
var trueWhen = appyaml.When{OS: appyaml.StringList{runtime.GOOS}}
var falseWhen = appyaml.When{OS: appyaml.StringList{"FAKE"}}

// Env var `OPTION_VAR` will be set to `option_val`
var valuetests = []struct {
	desc     string
	input    *Option
	expected string
}{
	{"nil", nil, ""},
	{"empty option", &Option{}, ""},
	{
		"default only",
		&Option{DefaultValues: valueList{
			{Value: "default"},
		}},
		"default",
	},
	{
		"command only",
		&Option{DefaultValues: valueList{
			{Command: "echo command"},
		}},
		"command",
	},
	{
		"environment variable only",
		&Option{Environment: "OPTION_VAR"},
		"option_val",
	},
	{
		"passed variable only",
		&Option{Passed: "passed"},
		"passed",
	},
	{
		"conditional value",
		&Option{DefaultValues: valueList{
			{When: falseWhen, Value: "foo"},
			{When: trueWhen, Value: "bar"},
			{When: falseWhen, Value: "baz"},
		}},
		"bar",
	},
	{
		"passed when all settings are defined",
		&Option{
			Environment: "OPTION_VAR",
			DefaultValues: valueList{
				{When: trueWhen, Value: "when"},
			},
			Passed: "passed",
		},
		"passed",
	},
}

func TestOption_Value(t *testing.T) {
	if err := os.Setenv("OPTION_VAR", "option_val"); err != nil {
		t.Fatalf("unexpected err setting environment variable: %s", err)
	}

	for _, tt := range valuetests {
		actual, err := tt.input.Value()
		if err != nil {
			t.Errorf(
				`Option.Value() for %s: unexpected err: %s`,
				tt.desc, err,
			)
			continue
		}

		if tt.expected != actual {
			t.Errorf(
				`Option.Value() for %s: expected "%s", actual "%s"`,
				tt.desc, tt.expected, actual,
			)
		}
	}
}
func TestOption_Value_default_and_command(t *testing.T) {
	option := Option{DefaultValues: valueList{
		{Value: "foo", Command: "echo bar"},
	}}
	_, err := option.Value()
	if err == nil {
		t.Fatalf(
			"option.Value() for %s: expected err, actual nil",
			"both Default and Command defined",
		)
	}
}

func TestOption_Value_private_and_environment(t *testing.T) {
	option := Option{Private: true, Environment: "OPTION_VAR"}
	_, err := option.Value()
	if err == nil {
		t.Fatalf(
			"option.Value() for %s: expected err, actual nil",
			"both Private and Environment variable defined",
		)
	}
}

func TestValue_UnmarshalYAML(t *testing.T) {
	s1 := []byte(`value: example`)
	s2 := []byte(`example`)
	v1 := value{}
	v2 := value{}

	if err := yaml.Unmarshal(s1, &v1); err != nil {
		t.Fatalf("yaml.Unmarshal(%s, ...): unexpcted error: %s", s1, err)
	}

	if err := yaml.Unmarshal(s2, &v2); err != nil {
		t.Fatalf("yaml.Unmarshal(%s, ...): unexpcted error: %s", s2, err)
	}

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf(
			"Unmarshalling of values `%s` and `%s` not equal:\n%#v != %#v",
			s1, s2, v1, v2,
		)
	}

	if v1.Value != "example" {
		t.Errorf(
			"yaml.Unmarshal(%s, ...): expected member `%s`, actual `%s`",
			s1, "example", v1.Command,
		)
	}
}

func TestValueList_UnmarshalYAML(t *testing.T) {
	s1 := []byte(`example`)
	s2 := []byte(`[example]`)
	v1 := valueList{}
	v2 := valueList{}

	if err := yaml.Unmarshal(s1, &v1); err != nil {
		t.Fatalf("yaml.Unmarshal(%s, ...): unexpcted error: %s", s1, err)
	}

	if err := yaml.Unmarshal(s2, &v2); err != nil {
		t.Fatalf("yaml.Unmarshal(%s, ...): unexpcted error: %s", s2, err)
	}

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf(
			"Unmarshalling of valueLists `%s` and `%s` not equal:\n%#v != %#v",
			s1, s2, v1, v2,
		)
	}

	if len(v1) != 1 {
		t.Errorf(
			"yaml.Unmarshal(%s, ...): expected 1 item, actual %d",
			s1, len(v1),
		)
	}

	if v1[0].Value != "example" {
		t.Errorf(
			"yaml.Unmarshal(%s, ...): expected member `%s`, actual `%s`",
			s1, "example", v1[0].Value,
		)
	}
}
