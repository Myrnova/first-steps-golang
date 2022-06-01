package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // define a chave que o campo vai representar quando converter para JSON
	Raca  string `json:"raca"` // caso voce queira que ignore um campo é só fazer `json:"-"`
	Idade uint   `json:"idade"`
}

func main() {

	cachorroEmJSON := `{
		"nome":"Rex",
		"raca":"Dálmata",
		"idade":3
	}`

	var c cachorro

	erro := json.Unmarshal([]byte(cachorroEmJSON), &c) // segundo parametro é o endereço de memória onde o valor vai ser armazenado, então precisa do &
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorroEmJSON2 := `{"nome":"Bola", "raca":"Vira-lata"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
