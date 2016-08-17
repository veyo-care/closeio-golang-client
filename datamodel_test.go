package closeio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetContact_BadPhone(t *testing.T) {

	contacts := GetContact("asdasd", "test@gmai.fr")
	assert.Nil(t, contacts[0].Phones)
}

func Test_GetContact_GoodPhone(t *testing.T) {

	phone := "+123123"
	contacts := GetContact(phone, "test@gmai.fr")
	assert.Equal(t, phone, contacts[0].Phones[0].Phone)
}

func Test_GetContact_ReplaceGoodPhone(t *testing.T) {

	phone := "+123 / 123"
	contacts := GetContact(phone, "test@gmai.fr")
	assert.Equal(t, "+123123", contacts[0].Phones[0].Phone)
}

func Test_ParseCloseIo(t *testing.T) {

	phone := "+123123"
	contacts := GetContact(phone, "test@gmai.fr")
	assert.Equal(t, phone, contacts[0].Phones[0].Phone)
}
