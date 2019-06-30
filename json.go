package main

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Id         int        `json:"id"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
	Options    Options    `json:"options"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	BalloonContent string `json:"balloonContent"`
	ClusterCaption string `json:"clusterCaption"`
	HintContent    string `json:"hintContent"`
	IconContent    string `json:"iconContent"`
}

type Options struct {
	IconLayout    string `json:"iconLayout"`
	IconImageHref string `json:"iconImageHref"`
	IconImageSize []int  `json:"iconImageSize"`
}

type User2 struct {
	FirstName string   `json:"name"`          // свойство FirstName будет преобразовано в ключ "name"
	LastName  string   `json:"lastname"`      // свойство LastName будет преобразовано в ключ "lastname"
	Books     []string `json:"ordered_books"` // свойство Books будет преобразовано в ключ "ordered_books"
}
