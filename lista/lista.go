package lista

import (
	"fmt"
	// "strings" se usa para manejo de cadenas
)

type nodo struct {
	siguiente, anterior *nodo
	nombre              string
	apellido            string
	edad                int
	direccion           string
	carrera             string
	curso               string
	estado              string
}
type lista struct {
	inicio, ultimo *nodo
	tamaño         int
}

func NuevaLista() *lista {
	return &lista{nil, nil, 0}

}

func (list *lista) Insertar(nombre_ string, apellido_ string, edad_ int, direccion_ string, carrera_ string, curso_ string, estado_ string) {
	nuevo := &nodo{nil, nil, nombre_, apellido_, edad_, direccion_, carrera_, curso_, estado_}

	if list.inicio == nil {
		list.inicio = nuevo
		list.ultimo = nuevo

	} else {
		list.ultimo.siguiente = nuevo
		nuevo.anterior = list.ultimo
		list.ultimo = nuevo
	}
	list.tamaño++
}

func (list *lista) Buscar(nombre_ string, apellido_ string) *nodo {
	aux := list.inicio
	for aux != nil {
		if (aux.nombre == nombre_) && (aux.apellido == apellido_) {
			fmt.Println("lo encontramos ")
			return aux

		} //fin de if
		aux = aux.siguiente
	} // fin de for
	//fmt.Println("no se encontro el nodo")
	return nil

}

func (list *lista) Imprimir() {
	aux := list.inicio
	for aux != nil {
		fmt.Print("[Nombre: ", aux.nombre, ", ")
		fmt.Print("Apellido: ", aux.apellido, ", ")
		fmt.Print("Edad: ", aux.edad, ", ")
		fmt.Print("Direccion: ", aux.direccion, ", ")
		fmt.Print("Carrera: ", aux.carrera, ", ")
		fmt.Print("Curso: ", aux.curso, ", ")
		fmt.Print("Estado: ", aux.estado, " ]")
		fmt.Println()
		aux = aux.siguiente
	}
	fmt.Println()
	fmt.Println("tamaño de la lista: ", list.tamaño)
}

func (list *lista) Modificar(nombre_ string, apellido_ string) {
	aux := list.Buscar(nombre_, apellido_)
	if aux != nil {
		fmt.Println("Ingrese datos a modificar:")
		var ed int
		var est, direc string
		fmt.Println("Edad: ")
		fmt.Scanln(&ed)
		fmt.Println("Estado: ")
		fmt.Scanln(&est)
		fmt.Println("Direccion: ")
		fmt.Scanln(&direc)
		aux.edad = ed
		aux.estado = est
		aux.direccion = direc
		fmt.Println("cambio realizado en nodo")
	} else {
		fmt.Println("el nodo no pudo ser encontrado")

	}
}
