package security

import (
  "fmt"
  "github.com/lucky-finger/core/env"
  "github.com/lucky-finger/core/security/rsa"
  "os"
  "path/filepath"
  "strings"
  "testing"
)

const (
  localExchangePrivateKeyFile = "exchange-private.pem"
  localExchangePublicKeyFile  = "key.ts"
)

func TestGeneratorExchangeKeyPair(t *testing.T) {

  generateExchangeKey := env.GetWithDefault("GENERATE_EXCHANGE_KEY", false)
  if !generateExchangeKey {
    t.Skip()
    return
  }

  keyPair, err := rsa.GenerateKeyPair()
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

  publicKeyOutTsDir := env.Get[string]("PUBLIC_KEY_OUT_TS_DIR")
  tsOutFilePath, err := filepath.Abs(filepath.Join(publicKeyOutTsDir, localExchangePublicKeyFile))
  if err != nil {
    t.Errorf("Error while getting absolute path: %v", err)
    return
  }

  if err = os.WriteFile(tsOutFilePath, []byte(fmt.Sprintf("export const RSA_PUBLIC_KEY_PEM = `%s`;\n", strings.TrimSpace(keyPair.PublicKey().ToPemStringMust()))), 0644); err != nil {
    t.Errorf("Error while writing key to file: %v", err)
    return
  }

  t.Logf("rsa public key written to file: %s", tsOutFilePath)
  _ = os.Remove(localExchangePublicKeyFile)

}
