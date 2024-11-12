package security

import (
	"github.com/lucky-finger/core/env"
	"os"
	"testing"
)

const localExchangePrivateKeyFile = "exchange-private.pem"

func TestGeneratorExchangeKeyPair(t *testing.T) {

	generateExchangeKey := env.GetWithDefault("GENERATE_EXCHANGE_KEY", false)
	if !generateExchangeKey {
		t.Skip()
		return
	}

	keyPair, err := GenerateRsaKeyPair()
	if err != nil {
		t.Errorf("Error while generating key pair: %v", err)
		return
	}

	key := keyPair.PrivateKey()
	if err = os.WriteFile(localExchangePrivateKeyFile, key.ToPemMust(), 0644); err != nil {
		t.Errorf("Error while writing key to file: %v", err)
		return
	}

	t.Logf("Key pair generated successfully and rsa private key written to file: %s", localExchangePrivateKeyFile)
	t.Logf("Public key:\n%s", keyPair.PublicKey().ToPemStringMust())

}
