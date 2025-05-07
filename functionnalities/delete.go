package functionnalities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Delete(Idstr string) {

	jsonContent, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	var listeTaches []tache

	json.Unmarshal(jsonContent, &listeTaches)

	Id, err := strconv.ParseUint(Idstr, 0, 8)

	if err != nil {
		log.Fatal(err)
	}

	var index int
	//IdFound := false

	for i := range listeTaches {

		if listeTaches[i].Id == uint8(Id) {
			index = i
			//IdFound = true

			fmt.Println("Est-ce la bonne tache à supprimer ? Y/N : ")
			fmt.Println(listeTaches[i])
			var confirm string
			fmt.Scanf("%s", &confirm)

			if confirm == "Y" || confirm == "y" {
				continue
			} else {
				log.Fatal("Pas la bonne tâche, essayez de les afficher avec la commande list")
			}
		}
	}

	listeTaches = append(listeTaches[:index], listeTaches[index+1:]...)

	for i := index; i < len(listeTaches); i++ {
		listeTaches[i].Id--
	}

	taskjson, err := json.MarshalIndent(listeTaches, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath, taskjson, 0644)

	if err != nil {
		fmt.Println("Erreur d'écriture fichier :", err)
		return
	}

}
