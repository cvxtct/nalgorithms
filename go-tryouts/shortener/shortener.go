package shortener

import (
	"crypto/sha256"
	"github.com/bwmarrin/go-alone"
	"strings"
	"fmt"
	"os"
	"math/big"
	// b64 "encoding/base64"
	"github.com/itchyny/base58-go"
	uuid "github.com/nu7hatch/gouuid"
)

// storage
// 4 chars: 1679616
// 8 chars: 2821109907456
// 11 chars: 1.316x10^17

type Shortener struct {

}

func (s *Shortener) GenerateShortedURL(shortID string) string {
	// find a short domain like cv.ly something
	return "https://cvxtct.com/" + shortID
}
 
func (s *Shortener) GenerateShortId(intitialURL string) string {
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	// either sign or add uuid
	// and sign earlier then the shor id generation
	signedURL := s.GenerateTokenFromString(intitialURL)
	hashedURL := s.generateSHA256(signedURL + uid.String())
	number := new(big.Int).SetBytes(hashedURL).Uint64()
	shortId := s.base58Encoder([]byte(fmt.Sprintf("%d", number)))

	return shortId	
}

func (s *Shortener) generateSHA256(uuid string) []byte {
	h := sha256.New()
	h.Write([]byte(uuid))
	bs := h.Sum(nil)

	return bs
}

func (s *Shortener) base58Encoder(bytes []byte) string {
	encoding := base58.FlickrEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// store in config
const secret = "DDOqaKzScsp5s1ZqO5RNqDIgDr-d4_3z7ETlW1QtPLI="

var secretKey []byte

// NewURLSigner creates a new signer
func NewURLSigner() {
	secretKey = []byte(secret)
}

// GenerateTokenFromString generates a signed token
func (s *Shortener) GenerateTokenFromString(data string) string {
	var urlToSign string

	sec := goalone.New(secretKey, goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlToSign = fmt.Sprintf("%s&hash=", data)
	} else {
		urlToSign = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := sec.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return token
}

// VerifyToken verifies a signed token
func (s *Shortener) VerifyToken(token string) bool {
	sec := goalone.New(secretKey, goalone.Timestamp)
	_, err := sec.Unsign([]byte(token))

	if err != nil {
		// signature is not valid. Token was tampered with, forged, or maybe it's
		// not even a token at all! Either way, it's not safe to use it.
		return false
	}
	// valid hash
	return true

}