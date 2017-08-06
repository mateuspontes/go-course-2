package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type People struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

type SpacePeople struct {
	Message string   `json:"message"`
	People  []People `json:"people"`
	Number  int      `json:"number"`
}

func GetPeopleFromSpace() (*SpacePeople, error) {
	url := "http://api.open-notify.org/astros.json"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	spacePeople := &SpacePeople{}
	err = json.NewDecoder(res.Body).Decode(spacePeople)

	return spacePeople, nil
}

func main() {
	spacePeople, err := GetPeopleFromSpace()

	if err == nil {
		fmt.Printf("Astronautas no espaço: %d\n", spacePeople.Number)
		fmt.Println("---------------------------")
		for i := 0; i < len(spacePeople.People); i++ {
			fmt.Printf("%s | %s\n", spacePeople.People[i].Name, spacePeople.People[i].Craft)
		}
	} else {
		fmt.Printf("Erro ao buscar astronautas.")
	}
}
