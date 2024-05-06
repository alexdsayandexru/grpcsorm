package validator

import (
	"errors"
	"testing"
)

func Test_Length(t *testing.T) {
	v := Validate("А-Яа-яЁё z")

	if ok, err := v.Length(10).GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v.Length(9).GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.Length(10).GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_MaxLength(t *testing.T) {
	v := Validate("А-Яа-яЁё z")

	if ok, err := v.MaxLength(10).GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v.MaxLength(9).GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.MaxLength(10).GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_Regex(t *testing.T) {
	v := Validate("А-Яа-яЁё z")

	if ok, err := v.Regex("^[A-Za-zА-Яа-яЁё -]+$").GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v.Regex("^[A-Za-zА-Яа-я -]+$").GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.Regex("^[A-Za-zА-Яа-яЁё -]+$").GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_Required(t *testing.T) {
	vs1 := []*Validator{Validate(1), Validate("1"), Validate([]string{"1"})}
	for _, v := range vs1 {
		if ok, err := v.Required().GetResult(); !ok {
			t.Error(err)
		}
		v.error = errors.New("")
		if _, err := v.Required().GetResult(); err.Error() != "" {
			t.Error(err)
		}
	}
	var arr []string
	vs2 := []*Validator{Validate(int32(0)), Validate(""), Validate(arr)}
	for _, v := range vs2 {
		if ok, err := v.Required().GetResult(); ok {
			t.Error(err)
		}
	}
}

func Test_RequiredIf(t *testing.T) {
	v1 := Validate(1)
	if ok, err := v1.RequiredIf(true).GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v1.RequiredIf(false).GetResult(); !ok {
		t.Error(err)
	}
	v1.error = errors.New("")
	if _, err := v1.RequiredIf(true).GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_Maximum(t *testing.T) {
	v := Validate(int32(100))

	if ok, err := v.Maximum(100).GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v.Maximum(99).GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.Maximum(100).GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_Equal(t *testing.T) {
	v := Validate(int32(100))

	if ok, err := v.Equal(100).GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := v.Equal(99).GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.Equal(100).GetResult(); err.Error() != "" {
		t.Error(err)
	}
}

func Test_Uuid(t *testing.T) {
	v := Validate("75e328bf-ddb2-44a8-a76d-67db7bd992be")
	if ok, err := v.Uiid().GetResult(); !ok {
		t.Error(err)
	}
	if ok, err := Validate("75e328bf_ddb2-44a8-a76d-67db7bd992be").Uiid().GetResult(); ok {
		t.Error(err)
	}
	v.error = errors.New("")
	if _, err := v.Uiid().GetResult(); err.Error() != "" {
		t.Error(err)
	}
}
