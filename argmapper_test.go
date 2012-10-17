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
