package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Crear un nuevo tipo de "mazo"
// el cual es un segmento de strings

type mazo []string

func nuevoMazo() mazo {
	cartas := mazo{}
	paloCartas := []string{"Espada", "Basto", "Copa", "Oro"}
	valoresCartas := []string{"Ancho", "Dos", "Tres", "Cuatro", "Cinco", "Seis", "Siete", "Diez", "Once", "Doce"}

	for _, paloCarta := range paloCartas {
		for _, valorCarta := range valoresCartas {
			cartas = append(cartas, valorCarta+" de "+paloCarta)
		}
	}

	return cartas
}

func repartir(m mazo, tamanoMano int) (mazo, mazo) {
	return m[:tamanoMano], m[tamanoMano:]
}

func (m mazo) print() {

	for i, carta := range m {
		fmt.Println(i, carta)
	}
}

func (m mazo) toString() string {
	return strings.Join([]string(m), ",")
}

func (m mazo) guardarAArchivo(nombreArchivo string) error {
	return ioutil.WriteFile(nombreArchivo, []byte(m.toString()), 0666)
}

func nuevoMazoDesdeArchivo(nombreArchivo string) mazo {
	bs, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return mazo(s)

}

func (m mazo) mezclar() {

	fuente := rand.NewSource(time.Now().UnixNano())
	r := rand.New(fuente)
	
	for i := range m {
		nuevaPosicion := r.Intn(len(m) - 1)
		m[i], m[nuevaPosicion] = m[nuevaPosicion], m[i]
	}
}
