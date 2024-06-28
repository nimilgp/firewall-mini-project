package main

import (
	"crypto/rand"
	"ebpf-firewall/dbLayer"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (app *application) generateBearerToken(w http.ResponseWriter, acc dbLayer.Account) {
	b := make([]byte, 128)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("<ERROR>\t\t[(gen bearer-token)failed to get random bytes]\n%s\n\n", err)
		return
	}
	tokenString := base64.URLEncoding.EncodeToString(b)[:128]
	curTime := time.Now()
	validTill := curTime.Add(time.Minute * 30)

	timeStamp := pgtype.Timestamp{
		Time:  validTill,
		Valid: true,
	}
	arg := dbLayer.CreateBearerTokenParams{
		Tokenstring: tokenString,
		Validtill:   timeStamp,
		Username:    acc.Username,
	}
	if err := app.queries.CreateBearerToken(app.ctx, arg); err != nil {
		log.Printf("<ERROR>\t\t[(gen bearer-token)failed to create bearer token]\n%s\n\n", err)
		return
	} else {
		if err := json.NewEncoder(w).Encode(arg); err != nil {
			log.Printf("<ERROR>\t\t[(gen bearer-token)failed to send bearer token]\n%s\n\n", err)
			return
		}
		log.Printf("<INFO>\t\t[(gen bearer-token)succesfully generate bearer token]\ntoken sting :%s\n\n", tokenString)
		return
	}
}
