package model

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Decks struct {
	Decks []Deck `json:"decks"`
}

type Deck struct {
	DeckName string `json:"deck_name"`
	Cards    []Card `json:"cards"`
}

type Card struct {
	CardID     int    `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	MeaningUp  string `json:"meaning_up"`
	MeaningRev string `json:"meaning_rev"`
	Desc       string `json:"desc"`
	Img        string `json:"img"`
}

func (c *Card) GetImage() []byte {
	path := filepath.Join("./images", c.Img)
	path += ".png"
	img, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return img
}

func GetDeck(deckName string) []Card {
	Decks := Decks{}
	fileData, err := os.ReadFile("./model/decks.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileData, &Decks)
	if err != nil {
		log.Fatal(err)
	}

	for _, deck := range Decks.Decks {
		if deckName == deck.DeckName {
			return deck.Cards
		}
	}
	log.Print("Deck not found")
	return nil
}

func GetCard(deckName string, ID int) Card {
	Cards := GetDeck(deckName)
	if len(Cards) < ID {
		log.Print("Card not found")
		return Card{}
	}

	return Cards[ID]
}
