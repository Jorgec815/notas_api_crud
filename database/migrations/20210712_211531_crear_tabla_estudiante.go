package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEstudiante_20210712_211531 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEstudiante_20210712_211531{}
	m.Created = "20210712_211531"

	migration.Register("CrearTablaEstudiante_20210712_211531", m)
}

// Run the migrations
func (m *CrearTablaEstudiante_20210712_211531) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210712_211531_crear_tabla_estudiante_up.sql")

	if err != nil {
	// handle error
	fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
	fmt.Println(request)
	m.SQL(request)
	// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearTablaEstudiante_20210712_211531) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210712_211531_crear_tabla_estudiante_down.sql")

	if err != nil {
	// handle error
	fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
	fmt.Println(request)
	m.SQL(request)
	// do whatever you need with result and error
	}
}
