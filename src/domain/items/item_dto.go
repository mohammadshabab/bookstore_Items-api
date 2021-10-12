package items

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             string      `json:"price"`
	AvailableQuantiry int64       `json:"available_quantity"`
	SoldQuantity      string      `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}
type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
