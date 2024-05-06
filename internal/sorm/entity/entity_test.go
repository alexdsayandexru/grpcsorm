package entity

import (
	"testing"
)

func testEntity(e IEntity, correlationId string, name string, t *testing.T) {
	if e.GetCorrelationId() != correlationId {
		t.Error()
	}
	if e.GetName() != name {
		t.Error()
	}
	rules := e.GetRules()
	for _, rule := range rules {
		_, _ = rule()
	}
	if len(ToBytes(e)) == 0 {
		t.Error()
	}
}

func Test_ValidateAdditional(t *testing.T) {
	if _, err := ValidateAdditional([]Additional{
		{Title: "Aвфыпаауя Ё ё  - , __ fgefagaefr", Content: ""},
	}); err != nil {
		t.Error(err)
	}
	if _, err := ValidateAdditional([]Additional{
		{Title: "Aвф5ыпаdsgg6546456ауя Ё ё  - , __ fgefagaefr", Content: ""},
	}); err == nil {
		t.Error()
	}
}

func Test_ValidateUserAgent(t *testing.T) {
	text := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.83 Safari/537.36"
	if ok, err := ValidateUserAgent(text); !ok {
		t.Error(err)
	}
}

func Test_ValidateMsisdnLogin(t *testing.T) {
	phone := "9778889900"
	email := ""
	if ok, err := ValidateMsisdnLogin(phone, email); !ok {
		t.Error(err)
	}
}

func Test_DeleteUserAccount(t *testing.T) {
	entity := &DeleteUserAccount{CorrelationId: "1"}
	testEntity(entity, "1", "DeleteUserAccount", t)
}

func Test_DeleteAccount(t *testing.T) {
	entity := &DeleteAccount{CorrelationId: "1"}
	testEntity(entity, "1", "DeleteAccount", t)
}

func Test_DeleteUserAccountAdmin(t *testing.T) {
	entity := &DeleteUserAccountAdmin{CorrelationId: "1"}
	testEntity(entity, "1", "DeleteUserAccountAdmin", t)
}

func Test_DirectoryData(t *testing.T) {
	entity := &DirectoryData{CorrelationId: "1"}
	testEntity(entity, "1", "DirectoryData", t)
}

func Test_LoginUser(t *testing.T) {
	entity := &LoginUser{CorrelationId: "1"}
	testEntity(entity, "1", "LoginUser", t)
}

func Test_LogoutUser(t *testing.T) {
	entity := &LogoutUser{CorrelationId: "1"}
	testEntity(entity, "1", "LogoutUser", t)
}

func Test_MigrateUser(t *testing.T) {
	entity := &MigrateUser{CorrelationId: "1"}
	testEntity(entity, "1", "MigrateUser", t)
}

func Test_RegisterUser(t *testing.T) {
	entity := &RegisterUser{CorrelationId: "1"}
	testEntity(entity, "1", "RegisterUser", t)
}

func Test_UpdateUserData(t *testing.T) {
	entity := &UpdateUserData{CorrelationId: "1"}
	testEntity(entity, "1", "UpdateUserData", t)
}

func Test_UpdateUserDataAdmin(t *testing.T) {
	entity := &UpdateUserDataAdmin{CorrelationId: "1"}
	testEntity(entity, "1", "UpdateUserDataAdmin", t)
}

func Test_UserAccountRecovery(t *testing.T) {
	entity := &UserAccountRecovery{CorrelationId: "1"}
	testEntity(entity, "1", "UserAccountRecovery", t)
}

func Test_EmptyValidator(t *testing.T) {
	v := NewEmptyValidator()
	v.Validate(nil)
}

func Test_Validator(t *testing.T) {
	v := NewValidator()
	if ok, info := v.Validate(
		&LogoutUser{}); !ok {
		_ = info.ToJson()
		_ = info.ToBytes()
	}
}

func Test_Arrays(t *testing.T) {
	_, _ = ValidateAdditional(nil)
	_, _ = ValidateAdditional([]Additional{{Title: "", Content: ""}})
	if ok, err := ValidateMsisdns([]string{"79998887766"}, []string{""}); !ok {
		t.Error(err)
	}
	_, _ = ValidateEmails([]string{"aaa@ru"}, []string{""})
	_, _ = ValidateImId([]ImId{{ServiceId: "1", ServiceName: "2"}})
	_, _ = ValidateOriServices([]Service{{ServiceId: 1, BeginTime: "05-03-2024:23.23.23.003"}})
	_, _ = ValidateOriServices([]Service{{ServiceId: 1, EndTime: "05-03-2024:23.23.23.003"}})
}
