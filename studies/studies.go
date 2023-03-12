package studies

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Car struct {
	Name  string
	Model string
	Price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100, Model: "Sei nao"})
	cars = append(cars, Car{Name: "Mercedes", Price: 200, Model: "Sei nao"})
	cars = append(cars, Car{Name: "Porsche", Price: 300, Model: "Sei nao"})
}

// tem que colocar o retorno no ()type senão não pode ter return na função
// funcao exportada com S maiusculo
func Some(a int, b int) int {
	if a+b > 10 {
		return 0
	}
	return a + b
}

// tem que colocar o retorno no ()type senão não pode ter return na função e pode ter dois retornos
// result, err := soma(1, 20)
// funcao privada soma em minusculo
func soma(a int, b int) (int, error) {

	//maior que 10 retorna 0 e lança erro
	if a+b > 10 {
		return 0, fmt.Errorf("Soma maior que 10")
	}
	//retorna a soma e o erro em branco pq declarou que ia retornar erro
	return a + b, nil
}

// nil é vazio, como undefined
func somar(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Erro")

	}
	return res, nil
}

// retorno NOMEADO
// dessa forma eu nao preciso retornar o valor "return a+b" ou return result, pq ja falei oq retorna
func retornoNomeado(a int, b int) /* retorna */ (result int) {
	result = a + b
	return
}

// retorno NOMEADO
// dessa forma eu nao preciso retornar o valor "return a+b" ou return result, pq ja falei oq retorna
func somoutro(a int, b int) /* retorna */ (result int) {

	//aqui não precisa := pq ja declarou no retorno
	result = a + b
	return
}

// retorno NOMEADO
// dessa forma eu nao preciso retornar o valor "return a+b" ou return result, pq ja falei oq retorna
func somaTudo(a ...int) int {
	resultado := 0

	for _, v := range a {
		resultado += v
	}

	return resultado

}

/* somaTudo(10, 20, 30) retorna = 60*/

type Carror struct {
	Nome string
}

// composições // heranca de outras structs
type CarroComHeranca struct {
	Carror
	//quando for json vai mostrar assim, ou outro nome que quiser
	Chassi string `json:"chassi"`
}

// ///relaciona com carro //nome da função
func (c Carror) andar() {
	/* fmt.Println(c.Nome, "Andow") */
}

func main() {

	carro := Carror{
		Nome: "Fusca",
		//syntax error: unexpected newline in composite literal; possibly missing comma or } =>>> A porra da virgula depois do nome do corno
	}

	carro.andar()

	carroherdado := CarroComHeranca{
		Carror: Carror{
			Nome: "seilá",
		},
		Chassi: "123456",
	}

	fmt.Println(carroherdado.Nome)
	fmt.Println(carroherdado.Chassi)

	carroJson, err := json.Marshal(carroherdado)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(carroJson))

	carroemjaison := `{"Nome":"seilá","Chassi":"123456"}`
	fmt.Println(carroemjaison)
	json.Unmarshal([]byte(carroemjaison), &carroemjaison)
	fmt.Println(carroemjaison)

}

func ponteiro() {

	a := 10

	fmt.Println(&a)

	var ponteiro *int = &a

	//retorna o hash da memoria
	fmt.Println(ponteiro)
	//retorna o valor dentro do hash ===>> valor:10
	fmt.Println(*ponteiro)

	//redeclarar através do poneitor ===>> valor:50
	*ponteiro = 50
	fmt.Println(*ponteiro)

}

// & antes da variavel mostra o endereço dela(HASH)

func switche() {

	/* a := 10 */
	a := 5
	switch a < 5 {
	case true:
		fmt.Println("A é maior que a condição do o switch")
		break

	case false:
		fmt.Println("A é maior que o switch")
		break

	default:
		fmt.Println("Deu erro parça")
		break
	}

	b := 10
	switch b {
	case 5:
		fmt.Println("É menor que 5")
		break
	case 15:
		fmt.Println("case b é 15")
		break
	default:
		fmt.Println("caiu no default")
		break
	}

}
