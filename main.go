package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type Response struct {
// 	Name    string    `json:"name"`
// 	Pokemon []Pokemon `json:"pokemon_entries"`
// }

// type Pokemon struct {
// 	EntryNo int            `json:"entry_number"`
// 	Species PokemonSpecies `json:"pokemon_species"`
// }
// type PokemonSpecies struct {
// 	Name string
// }

// type Person struct {
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Location string `json:"location"`
// }

type Person struct {
	Name     string
	Age      int
	Location string
}

/* func (c *Person) abc() {

	c.Name = "abcd"
	c.Age = 20
	c.Location = "HIyd"
	fmt.Println(c)

} */
func GetApi(w http.ResponseWriter, r *http.Request) {

	var person []Person

	// myName := "manikanta"
	// w.Write([]byte(myName))
	//a:= []byte(myName)

	array := []string{"mani", "saad"}

	array2 := []int{2, 5, 9, 2, 4}

	// newarray, err := json.Marshal(array)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// newarray2, err := json.Marshal(array2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(w.Write(newarray))
	// fmt.Println("\n",newarray)
	json.NewEncoder(w).Encode(array2)
	json.NewEncoder(w).Encode(array)

	json.NewEncoder(w).Encode(person)
	// abc:= json.NewDecoder(r).Decode(array)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Write(newarray2)
	fmt.Println(array2, array)
	// w.Write(newarray)
	// fmt.Println(myName)

}

func main() {

	//  := Person{"mani", 32, "Hyd"}
	var manikanta Person
	fmt.Println(manikanta)

	// manikanta.abc()
	var mm Person
	mm.Name="xyz"
	fmt.Println()
	// fmt.Println(manikanta)

	Route()

}

func Route() {
	newRoute := mux.NewRouter()
	newRoute.HandleFunc("/get", GetApi).Methods("GET")
	http.ListenAndServe(":2525", newRoute)
}

