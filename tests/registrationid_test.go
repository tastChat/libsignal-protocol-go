package tests

import (
	"fmt"
	"github.com/tastChat/libsignal-protocol-go/util/keyhelper"
	"testing"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
