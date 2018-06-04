package pm

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/openpgp/armor"
)

// ...Armor Type
type ArmorType string

func (at ArmorType) string() string {
	return string(at)
}

const (
	pgpMessageType      ArmorType = "PGP MESSAGE"
	pgpPublicBlockType  ArmorType = "PGP PUBLIC KEY BLOCK"
	pgpPrivateBlockType ArmorType = "PGP PRIVATE KEY BLOCK"
)

// ArmorKey make bytes input key to armor format
func ArmorKey(input []byte) (string, error) {
	return ArmorWithType(input, pgpPublicBlockType.string())
}

// ArmorWithType make bytes input to armor format
func ArmorWithType(input []byte, armorType string) (string, error) {
	var b bytes.Buffer
	w, err := armor.Encode(&b, armorType, nil)
	if err != nil {
		return "", err
	}
	_, err = w.Write(input)
	if err != nil {
		return "", err
	}
	w.Close()
	return b.String(), nil
}

// UnArmor an armored key to bytes key
func UnArmor(input string) ([]byte, error) {
	b, err := unArmor(input)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(b.Body)
}

func unArmor(input string) (*armor.Block, error) {
	io := strings.NewReader(input)
	b, err := armor.Decode(io)
	if err != nil {
		return nil, err
	}
	return b, nil

}
