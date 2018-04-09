package main

func main() {
	Host := []string{
		"127.0.0.1:27017",
	}
	const (
		Username   = "root"
		Database   = "isprime-cdn-database"
		Collection = "resource-record-collection"
	)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//collection
	c := session.DB(Database).C(Collection)

	// Remove
}
