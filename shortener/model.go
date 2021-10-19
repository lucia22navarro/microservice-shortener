package shortener

type Redirect struct {
	Code      string `json:"code" bson:"code"`                                   //código pequeño que devuelve el acortador
	URL       string `json:"url" bson:"url" validate:"empty=false & format=url"` //URL que le pasamos
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}
