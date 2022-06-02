package connector

import (
	"testing"
)

func TestNewBrokerPilot(t *testing.T) {
	_, err := NewBrokerPilot(&Settings{HostURL: "https://yandex.ru"})
	if err != nil {
		t.Errorf(`error does not have be returned`)
	}

	if _, err = NewBrokerPilot(&Settings{HostURL: ""}); err == nil {
		t.Errorf(`have to be returned error`)
	}
}
