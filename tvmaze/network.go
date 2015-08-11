package tvmaze

type network struct {
	Id      int
	Name    string
	Country country
}

type country struct {
	Name     string
	Code     string
	Timezone string
}
