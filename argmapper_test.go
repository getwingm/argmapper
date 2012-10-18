package argmapper

import "testing"

func Test_mapArgs(t *testing.T) {
	e := Map{"foo": "bar"}
	o := []Map{
		New([]string{"--foo=bar"}),
		New([]string{"-foo=bar"}),
		New([]string{"--foo", "bar"}),
		New([]string{"-foo", "bar"})}

	for i, v := range o {
		if len(v) != len(e) {
			t.Fatalf("%v: %v != %v", i, len(v), len(e))
		}

		for k, _ := range v {
			if k != "foo" {
				t.Fatalf("o[%v]: got %#v", i, k)
			}
		}

		for _, ov := range v {
			if ov != "bar" {
				t.Fatalf("o[%v]: got %#v", i, ov)
			}
		}
	}
}

func Test_mapArgs_nuanced(t *testing.T) {
	m := New([]string{
		"subcommand",         // not an option
		"--foo=bar",          // first option
		"--follow",           // second option
		"-silent",            // third option
		"--path", "./foobar", // fourth option
		"iam-something-else", // not an option
	})

	if len(m) != 4 {
		t.Fatalf("got %#v", len(m))
	}

	if v, ok := m["foo"]; !ok || v != "bar" {
		t.Fatalf("got %#v", v)
	}

	if v, ok := m["follow"]; !ok || v != "" {
		t.Fatalf("got %#v", v)
	}

	if v, ok := m["silent"]; !ok || v != "" {
		t.Fatalf("got %#v", v)
	}

	if v, ok := m["path"]; !ok || v != "./foobar" {
		t.Fatalf("got %#v", v)
	}
}
