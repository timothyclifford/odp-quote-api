package main

type Quote struct {
	Id            string
	SalesPersonId string
	Created       int
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	Street        string
	Suburb        string
	Postcode      string
}

var repository = []Quote{
	{"abc", "abc", 0, "first", "last", "email", "phone", "address", "suburb", "postcode"},
}

func getAllQuotes() []Quote {
	return repository
}

func saveQuote(quote Quote) Quote {
	repository = append(repository, quote)
	return quote
}
