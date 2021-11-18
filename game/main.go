package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	//"reflect"
	//"encoding/json"
	"regexp"
	"strings"
)

// Estructura en la que se almacena la informacion de un juego
type Juego struct {
	Gamename    []string //Contendra los juegos con su numero, de la forma: [1;Random 2;Ruleta 5;Last]
	Players     int      //Contiene el numero maximo de jugadores para este juego
	Rungames    int      //Contiene el numero de veces que se juega/ejecuta el juego
	Concurrence int      //Contiene la concurrencia
	Timeout     int      //Este tiempo siempre sera en MINUTOS
}

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	fmt.Println("To exit the command line, press ctrl + c")
	for {
		var entry string
		var ok bool
		// PARA SALIR DEL JUEGO HAY QUE SALIR HACER LA COMBINACION DE TECLAS CTRL+C
		fmt.Printf("Entry: ")
		if entry, ok = readline(fi); ok {

			fmt.Println("Checking entry...")
			if checkParameters(entry) {
				pars := splitString(entry, "--")
				game, err := generateGame(pars)
				if err == nil {
					runGame(game)
				} else {
					fmt.Println(err)
				}
			}
			fmt.Println("\nPress ctrl+c to exit")
		} else {
			break
		}
	}
}

// Ejecuta el juego
func runGame(game Juego) {

	fmt.Println("Running game...")
	fmt.Println(game)
}

// Lee una linea en la consola
func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

// Crea un array de strings con una cadena de entrada:'str' y los separa con el discriminante:'disc'
func splitString(str string, disc string) []string {

	fmt.Println("Getting Parameters...")
	arr_str := strings.Split(str, disc)
	fmt.Println("Result: ", len(arr_str), "parameters were gotten")
	fmt.Println("<---")
	fmt.Println(strings.Join(arr_str, " ,\n"))
	fmt.Println("--->")
	return arr_str
}

// Revisa que todos los parametros del comando se encuentren en la cadena de entrada
func checkParameters(str string) bool {
	cad := strings.ToLower(str)
	if strings.Contains(cad, "rungame") &&
		strings.Contains(cad, "gamename") &&
		strings.Contains(cad, "players") &&
		strings.Contains(cad, "rungames") &&
		strings.Contains(cad, "concurrence") &&
		strings.Contains(cad, "timeout") {
		return true
	}
	fmt.Println("Command is Incorrect, You have missing or incorrect parameters!")
	return false
}

// Genera un struct con la informacion del juego, es la ultima etapa del analisis del comando de entrada, para ejecutarse luego
func generateGame(arr []string) (Juego, error) {
	game := Juego{}
	for _, element := range arr {

		if strings.Contains(element, "gamename") {
			game.Gamename = extractGameNames(element)
		} else if strings.Contains(element, "players ") {
			ply, err := strconv.Atoi(extractSingleParameter(element, "players "))
			if err == nil {
				game.Players = ply
			}
		} else if strings.Contains(element, "rungames ") {
			rngames, err := strconv.Atoi(extractSingleParameter(element, "rungames "))
			if err == nil {
				game.Rungames = rngames
			}
		} else if strings.Contains(element, "concurrence ") {
			concurr, err := strconv.Atoi(extractSingleParameter(element, "concurrence "))
			if err == nil {
				game.Concurrence = concurr
			}
		} else if strings.Contains(element, "timeout ") {
			srr := extractSingleParameter(element, "timeout ")
			re := regexp.MustCompile("[0-9]+")
			result := re.FindStringSubmatch(srr)
			if result != nil {
				tmout, err := strconv.Atoi(result[0])
				if err == nil {
					game.Timeout = tmout
				}
			}
		}
	}

	if len(game.Gamename) != 0 &&
		game.Players != 0 &&
		game.Rungames != 0 &&
		game.Concurrence != 0 &&
		game.Timeout != 0 {

		return game, nil
	}

	return game, errors.New("Something went wrong with the paramaters provided!")

}

// Extrae en un array de string los numero y nombre de los diferetes juegos introducidos en el comando: " 1;UNO | 2;JENGA"
func extractGameNames(str string) []string {
	res := []string{}

	re := regexp.MustCompile("\\\"(.*?)\\\"")
	match := re.FindStringSubmatch(str)
	//fmt.Println(match[1])
	aux := strings.Split(match[1], "|")
	for _, element := range aux {
		new_str := strings.TrimSpace(element)
		res = append(res, new_str)
	}
	//fmt.Println(res)
	return res
}

// Extrae un solo parametro del comando de entrada, es para aquellos parametrso que son una simple cadena: 2m, 20000, etc
func extractSingleParameter(str string, par string) string {
	cad := ""
	cad = str[len(par):]
	cad = strings.TrimSpace(cad)
	//fmt.Println("Extracted:",cad)
	return cad
}

/*
rungame --gamename "1;Random | 2;Ruleta | 5;Last" --players 30 --rungames 30000 --concurrence 10 --timeout 3m

--gamename "1;Random | 2;Ruleta | 5;Last"
*/
