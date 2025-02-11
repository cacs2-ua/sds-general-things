/*
Ejemplo 1 (base para el cifrado del césar y la escítala espartana)

Este programa copia de la entrada a la salida carácter a carácter,
restringiéndose a un alfabeto limitado y pasando a mayúsculas.
Permite leer de la entrada y salida estándar o usar ficheros.


ejemplos de uso:

go run ejemplo1.go

go run ejemplo1.go fichentrada.txt fichsalida.txt


-lectura y escritura
-entrada y salida estándar
-ficheros
-parámetros en línea de comandos (os.Args)
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	var fin *os.File  // fichero de entrada
	var fout *os.File // fichero de salida
	var err error     // receptor de error
	var lineas int
	var matriz [][]rune
	var numeroColumna int
	var numeroFila int

	// alfabeto con el que vamos a trabajar
	alfabeto := map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
		'K': 10, 'L': 11, 'M': 12, 'N': 13, 'Ñ': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19,
		'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}

	alfabetoInverso := make(map[int]rune, len(alfabeto))

	for letra, valor := range alfabeto {
		alfabetoInverso[valor] = letra
	}

	if len(os.Args) == 1 { // no hay parámetros, usamos entrada (teclado) y salida estándar (pantalla)
		fin = os.Stdin
		fout = os.Stdout
		lineas = 0

	} else if len(os.Args) == 2 {
		fin = os.Stdin
		fout = os.Stdout
		lineas, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: el argumento no es un número válido")
			os.Exit(1)
		}

	} else if len(os.Args) == 3 { // tenemos los nombres de los ficheros de entrada y salida en los parámetros
		fin, err = os.Open(os.Args[1]) // abrimos el primer fichero (entrada)
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		fout, err = os.Create(os.Args[2]) // abrimos el segundo fichero (salida)
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	} else if len(os.Args) == 4 { // tenemos los nombres de los ficheros de entrada y salida en los parámetros
		fin, err = os.Open(os.Args[1]) // abrimos el primer fichero (entrada)
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		fout, err = os.Create(os.Args[2]) // abrimos el segundo fichero (salida)
		if err != nil {
			panic(err)
		}
		defer fout.Close()
		lineas, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Error: el argumento no es un número válido")
			os.Exit(1)
		}
	} else { // error de parámetros
		fmt.Println("Número de parámetros incorrecto: se espera los ficheros de origen y destino (1 y 2, opcionales)")
		os.Exit(1)
	}

	for { // bucle infinito
		var c rune // carácter a leer

		_, err = fmt.Fscanf(fin, "%c", &c) // lectura de la entrada

		fmt.Printf("Leído: '%c' (ASCII: %d)\n", c, c)

		if c == '\n' || c == 0 {
			for j := 0; j < lineas; j++ {
				for i := 0; i <= numeroFila; i++ {
					if matriz[i][j] != 0 {
						fmt.Fprintf(fout, "%c", matriz[i][j])
					}

				}
			}
			numeroColumna = 0
			numeroFila = 0
			matriz = [][]rune{}

			if c == '\n' {
				continue
			} else {
				break
			}

		}

		if numeroColumna >= lineas {
			numeroColumna = 0
			numeroFila++
		}

		C := unicode.ToUpper(c) // pasamos a mayúsculas

		if numeroFila >= len(matriz) {
			matriz = append(matriz, make([]rune, lineas))
		}

		matriz[numeroFila][numeroColumna] = C
		numeroColumna++

	}

}
