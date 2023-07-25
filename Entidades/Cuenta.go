package cuenta

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Cuenta struct {
	saldo   float64
	titular string
	id      int
}

func CrearCuenta() Cuenta {
	var saldoInput float64
	var titularInput string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese el titular de la cuenta")
	titularInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada", err)
	}

	fmt.Println("Ingrese el saldo del titular de la cuenta")
	fmt.Scan(&saldoInput)

	cuenta1 := Cuenta{
		titular: titularInput,
		saldo:   saldoInput,
		id:      rand.Int(),
	}

	return cuenta1
}

func RetirarDinero(cuenta1 Cuenta) {
	var retiro float64
	var confirm string

	fmt.Printf("Bienvenido %v\n", cuenta1.titular)
	for {
		fmt.Println("Ingrese el monto a retirar")
		fmt.Scan(&retiro)
		if retiro > cuenta1.saldo || retiro <= 0 {
			fmt.Println("Error: monto invalido.")

		} else {
			fmt.Printf("Se han retirado %v de su cuenta %v\n", retiro, cuenta1.id)
			cuenta1.saldo -= retiro
			fmt.Printf("Actualmente le quedan $%v en su cuenta", cuenta1.saldo)
		}
		fmt.Println("Si desea ingresar otro monto presione 's' sino, presione otra tecla")
		fmt.Scan(&confirm)
		if confirm != "s" {
			break
		}
	}
}
