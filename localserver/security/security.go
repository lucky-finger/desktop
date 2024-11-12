package security

import "github.com/lucky-finger/core/security/rsa"

const (
	// rsaBits rsa密钥长度
	rsaBits = 2048
)

// GenerateRsaKeyPair 生成rsa密钥对
func GenerateRsaKeyPair() (*rsa.KeyPair, error) {
	return rsa.GenerateKeyPair(rsaBits)
}
