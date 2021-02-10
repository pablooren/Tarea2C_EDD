package main

import (
	"fmt"

	"./lista"
)

func main() {
	ls := lista.NuevaLista()
	var nm, ape string

	ls.Insertar("juan", "orellana", 25, "las ilusiones", "sistemas", "edd", "Activo")
	ls.Insertar("jorge", "salazar", 25, "bosque", "sistemas", "io2", "Activo")
	ls.Insertar("jose", "pinto", 20, "casamia", "civil", "mb2", "Activo")
	ls.Insertar("cristian", "salazar", 26, "las ilusiones", "farmacia", "mate4", "Activo")
	ls.Insertar("jonathan", "lopez", 23, "las ilusiones", "farmacia", "analisis", "Activo")
	ls.Insertar("kristhian", "valenzuela", 20, "zona 17", "farmacia", "bioquimica", "Activo")
	ls.Insertar("arturo", "cañas", 28, "zona 10", "auditoria", "control", "Desactivo")
	ls.Insertar("antonio", "cañas", 30, "antigua", "marino", "vientos", "Desactivo")
	ls.Insertar("luisa", "diaz", 22, "zona 2", "civil", "estructural", "Desactivo")
	ls.Insertar("kathy", "monterrosa", 29, "san francisco", "auditoria", "planillas", "Desactivo")

	//	ls.Imprimir()
	fmt.Println("Ingrese nombre y apellido de estudiante a modificar en minusculas: ")
	fmt.Println("Nombre: ")
	fmt.Scanln(&nm)
	fmt.Println("Apellido: ")
	fmt.Scanln(&ape)

	ls.Modificar(nm, ape)

	ls.Imprimir()

}
