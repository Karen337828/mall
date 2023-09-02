package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
)

const sm4key = "1234567890abcdef"
const sm2PriKey = ""
const sm2PubKey = ""

// SM2EncodeBase64 sm2公钥加密，密文base64编码
func SM2EncodeBase64(pubKey string, plaintext string) (string, error) {
	pub, err := x509.ReadPublicKeyFromHex(pubKey)
	if err != nil {
		return "", err
	}
	cipher, err := sm2.Encrypt(pub, []byte(plaintext), rand.Reader, 0)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipher), nil
}

// SM2DecodeBase64 sm2私钥解密，密文base64编码
func SM2DecodeBase64(privKey string, data string) (string, error) {
	priv, err := x509.ReadPrivateKeyFromHex(privKey)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintext, err := sm2.Decrypt(priv, []byte(ciphertext), 0)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// SM2Encode sm2公钥加密
func SM2Encode(pubKey string, plaintext string) (string, error) {
	pubMen, err := x509.ReadPublicKeyFromHex(pubKey)
	if err != nil {
		return "", err
	}
	msg := []byte(plaintext)
	ciphertxt, err := sm2.Encrypt(pubMen, msg, rand.Reader, 0)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertxt), nil
}

// SM2Decode sm2私钥解密
func SM2Decode(privKey string, data string) (string, error) {
	priv, err := x509.ReadPrivateKeyFromHex(privKey)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintext, err := sm2.Decrypt(priv, []byte(ciphertext), 0)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func SM4Encrypt(originalStr string) string {
	iv := make([]byte, sm4.BlockSize)
	key := []byte(sm4key)
	plainText := []byte(originalStr)
	block, err := sm4.NewCipher(key)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)

}

func SM4Decrypt(originalStr string) ([]byte, error) {
	iv := make([]byte, sm4.BlockSize)
	key := []byte(sm4key)
	cipherText := []byte(originalStr)
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
