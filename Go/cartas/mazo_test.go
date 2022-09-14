package main

import "testing"
import "os"

func TestNuevoMazo(t *testing.T) {
	m := nuevoMazo()

	if len(m) != 40 {
		t.Errorf("Duracion esperada de mazo 40, pero se obtuvo %v", len(m))
	}

	if m[0] != "Ancho de Espada" {
		t.Errorf("Primer carta esperada era Ancho de Espada, pero se obtuvo %v", m[0])
	}

	if m[len(m)-1] != "Doce de Oro" {
		t.Errorf("Ultima carta esperada era Doce de Oro, pero se obtuvo %v", m[len(m)-1])
	}

}

func TestGuardarAArchivoYNuevoMazoDesdeArchivo(t *testing.T) {
	os.Remove("_mazotesting")

	mazo := nuevoMazo()
	mazo.guardarAArchivo("_mazotesting")

	mazoCargado := nuevoMazoDesdeArchivo("_mazotesting")

	if len(mazoCargado) != 40 {
		t.Errorf("Duracion esperada de mazo 40, pero se obtuvo %v", len(mazoCargado))
	}

	os.Remove("_mazotesting")
}
