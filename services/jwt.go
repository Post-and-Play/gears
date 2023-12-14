package services

import (
	"crypto/sha256"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secret string
	issuer string
}

type Claim struct {
	Sub    int64 `json:"sub"`
	Claims jwt.StandardClaims
}

func NewJWTService() *jwtService {
	return &jwtService{
		secret: os.Getenv("SECRET_KEY"),
		issuer: "website-login",
	}
}

func (s *jwtService) GenerateToken(userid int64) (string, error) {
	claim := &Claim{
		Sub: userid,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim.Claims)

	t, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secret), nil
	})

	return err == nil
}

func SHA256Encoder(pass string) string {
	str := sha256.Sum256([]byte(pass))
	return fmt.Sprintf("%x", str)
}

func Encrypt(text string) string {
	key := []byte("15f03dedbb5b8a09210c95249244fb67")
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("blogPostGeek")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	//fmt.Printf("Ciphertext: %x\n", ciphertext)
	//return string(ciphertext)
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(text string) string {
	key := []byte("15f03dedbb5b8a09210c95249244fb67")
	ciphertext, _ := hex.DecodeString(text)
	nonce := []byte("blogPostGeek")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("Plaintext: %s\n", string(plaintext))
	//return string(plaintext)
	return fmt.Sprintf("%s", plaintext)
}