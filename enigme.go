package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("enigme.txt") //lecture du fichier
	if err != nil {                            //si err différent de la valeur nil
		log.Fatal(err) //affiche l'erreur
	}

	fmt.Println(Read(string(file))[0])
	fmt.Println(string(file[23])) //affiche le caractère à l'index 23 car l'indice dit que le fragment se trouve derriere le mot before  le caractère est la première lettre du mot random donc c'est un fragment de l'enigme

	//fmt.Println(string(file[237]))  l'opposé du commencement de tout
}

func Read(file string) []string {
	var arr []string
	var tmp string

	for _, num := range file {

		if num == '\n' { //si le caractère est une fin de ligne  13 en byte
			arr = append(arr, tmp)
			tmp = ""
		} else { //sinon on ajoute le caractère au mot jusqu'a la  prochaine fin de ligne
			tmp += string(num) //concatenation des lettres

		}
		if tmp == "generate" { // le commencement de tout
			fmt.Println(tmp)
		}
		if tmp == "a" { // l'opposé du commencement de tout
			fmt.Println(tmp)
		}
		if tmp == "number" { //
			fmt.Println(tmp)
		}

	}

	return arr
}

//fmt.Println(string(file)) // print the content as a 'string'
