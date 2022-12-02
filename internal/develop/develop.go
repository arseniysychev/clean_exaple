package develop

type Element interface {
	GetValue() interface{}
	GetEntity() Entity
	GetOrder() uint8
	GetName() string
}

type Entity struct {
	title    string
	elements []Element
}

func Develop() {
	m := make(map[string]interface{})

	m["name"] = "Donald"
	m["surname"] = "Trump"
	m["age"] = 72

}
