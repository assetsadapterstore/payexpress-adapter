package payexpress

import (
	"encoding/hex"
	"github.com/assetsadapterstore/payexpress-adapter/payexpress_addrdec"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	pub, _ := hex.DecodeString("2a3569c72e7a1d77ef8cc4e2e1a68a2d5ea6221f6486b1704d19295715e2ea74")
	accountID, _ := payexpress_addrdec.Default.AddressEncode(pub)
	t.Logf("accountID: %s", accountID)

}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	accountID := "GAVDK2OHFZ5B257PRTCOFYNGRIWV5JRCD5SINMLQJUMSSVYV4LVHI4CN"
	pub, _ := payexpress_addrdec.Default.AddressDecode(accountID)
	t.Logf("pub: %s", hex.EncodeToString(pub))

}