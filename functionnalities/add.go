package functionnalities

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const filepath string = "./tasks.json"

type tache struct {
	Id            uint8
	Nom           string
	En_cours      bool
	Terminee      bool
	Date_creation string
	Date_update   string
}

func Add(nomTache string) {

	var id uint8 = 0

	jsonContent, err := os.ReadFile(filepath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {

			fmt.Println("Fichier introuvable, création...")

			task := creerTache(id, nomTache)

			f, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()

			f, err = os.OpenFile(filepath, os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}

			f.WriteString("[")
			f.WriteString("]")

			f.Close()

			ecrireTache(task)
		}
	} else {

		var listeTaches []tache

		json.Unmarshal(jsonContent, &listeTaches)

		id = listeTaches[len(listeTaches)-1].Id + 1

		ecrireTache(creerTache(id, nomTache))
	}

}

func creerTache(id uint8, nomTache string) *tache {
	currentTime := time.Now()

	maTache := tache{Id: id, Nom: nomTache, En_cours: false, Terminee: false,
		Date_creation: currentTime.Format("2006-01-02 15:04:05"),
		Date_update:   currentTime.Format("2006-01-02 15:04:05")}

	return &maTache
}

func ecrireTache(task *tache) {

	var listeTaches []tache

	taches, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(taches, &listeTaches)

	listeTaches = append(listeTaches, *task)

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
