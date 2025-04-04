package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main(){
	
	exibeIntroducao()
	registraLogs("site-falso", false)

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido!")	
			os.Exit(-1)
		}

	}

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2{
	// 	fmt.Println("Exibindo Logs..")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Comando desconhecido!")
	// } 
}

func exibeIntroducao(){
	nome := "Vinicius"
	versao := 1.1
	fmt.Println("hello world, olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMenu(){
	fmt.Println("1-Iniciar o monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do programa")
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	
	sites := leSitesDoArquivo()

	for i:=0; i < monitoramentos; i++ {
		for i, site:= range sites {
			fmt.Println("Posição:", i, site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string){
	resp, err := http.Get(site)

if err != nil {
	fmt.Println("Ocorreu um erro:", err)
}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!", resp.StatusCode)
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode )
		registraLogs(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sites)
	arquivo.Close()
	return sites
}

func registraLogs(site string, status bool){

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: "+ strconv.FormatBool(status)+ "\n")

	arquivo.Close()
}

func imprimeLogs(){
	fmt.Println("Exibindo logs...")
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}