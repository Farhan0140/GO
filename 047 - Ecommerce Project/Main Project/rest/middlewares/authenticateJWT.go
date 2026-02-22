package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 || headerArr[0] != "JWT" {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		messageByteArr := []byte(message)

		cnf := config.GetConfig()

		hash := hmac.New(sha256.New, []byte(cnf.SecretKey))
		hash.Write(messageByteArr)

		newSignature := hash.Sum(nil)
		newSignatureBase64 := base64_URL_Encode(newSignature)

		if newSignatureBase64 != signature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64_URL_Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
