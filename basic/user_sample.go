package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "Mangunsari",
		City:    "Grobogan",
		Country: "Indonesia",
	}
	u := basic.User{
		Id:       99,
		Username: "Marley",
		IsActive: true,
		Password: []byte("namasayamarley"),
		Address:  &addr,
		//Emails:   []string{"marley@mail.com", "bob@mail.com"},
		//Gender:   basic.Gender_GENDER_MALE,
	}
	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "Sinta",
		IsActive: true,
		Password: []byte("sinta.com"),
		Emails:   []string{"sinta@mail.com", "sinta123@mail.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	jsonString := string(jsonBytes)
	log.Println(jsonString)
}

func JsonToProtoUser() {
	json := `{
		"id": 97,
		"username": "peter",
		"is_active": true,
		"password": "c14ludGEuY149t",
		"emails": [
			"peter@mail.com",
			"peter@email.com"
		],
		"gender": "GENDER_MALE"
	}`

	var p basic.User
	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(&p)
}
