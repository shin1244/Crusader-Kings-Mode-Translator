package main

import (
	"context"
	"encoding/json"
	"fmt"
	"modules/app"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	m := mux.NewRouter() // 라우터 생성

	client := app.MongoConnect()
	dataBase := client.Database("CK3_mod")
	collections, _ := dataBase.ListCollectionNames(context.Background(), bson.M{})

	defer client.Disconnect(context.Background())

	// 콜렉션 목록을 반환하는 API
	m.HandleFunc("/getCollections", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(collections)
	})

	// 콜렉션의 내용을 반환하는 API
	m.HandleFunc("/{collection}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		collection := vars["collection"]

		isCollectionExist := app.CheckCollectionExist(collections, collection)

		// 콜렉션이 존재하면 콜렉션 보여주고, 아니면 404
		if isCollectionExist {
			result := bson.M{}
			err := dataBase.Collection(collection).FindOne(context.Background(), bson.M{}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// 개같은 코드

			for key := range result {
				for value := range result[key].(bson.M) {
					inMap, ok := app.GetMapValue(result, key)
					if ok {
						event, ok := app.GetMapValue(inMap.(bson.M), value)
						if ok {
							fmt.Println(event)
						}
					}
				}
			}

			data := struct {
				Collection string
			}{
				Collection: collection,
			}

			tmpl, _ := template.ParseFiles("public/collection.html")

			tmpl.Execute(w, data)
		} else {
			http.NotFound(w, r)
		}
	})

	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":8080", n)
}
