package main

import (
	cuenta "objetos-structs/Entidades"
)

func main() {
	cuenta1 := cuenta.CrearCuenta()
	cuenta.RetirarDinero(cuenta1)
}
