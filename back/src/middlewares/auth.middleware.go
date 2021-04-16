package middlewares

import (
	"back/lib"
	"back/lib/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var user lib.IUser
		var dbUser lib.IUser
		json.NewDecoder(r.Body).Decode(&user)
		collection:= lib.MyDB.Collection("user")
		err:=collection.FindOne(context.TODO(), bson.M{"email":user.Email}).Decode(&dbUser)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"`+err.Error()+`"}`))
			return
		}
		userPass:= []byte(user.Password)
		dbPass:= []byte(dbUser.Password)
		passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)
		if passErr != nil{
			log.Println(passErr)
			w.Write([]byte(`{"response":"Wrong Password!"}`))
			return
		}
		jwtToken, err := utils.GenerateJWT(lib.Environment.JwtSecret)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"`+err.Error()+`"}`))
			return
		}
		w.Write([]byte(`{"token":"`+jwtToken+`"}`))
		next.ServeHTTP(w, r)
	})
}
