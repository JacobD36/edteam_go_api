package model

//Community es la estructura de una comunidad
type Community struct {
	//Name nombre de una comunidad. Ej: EDTeam
	Name string `json:"name"`
}

//Communities es un slice de comunidades
type Communities []Community

//Person es la estructura de una persona
type Person struct {
	//name es el nombre de una persona Ej.:Alexys
	Name string `json:"name"`
	//Age es la edad de la persona Ej.: 40
	Age uint8 `json:"age"`
	//Communities son las comunidades a las que pertenece la persona
	Communities Communities `json:"communities"`
}

//Persons es un slice de personas
type Persons []Person
