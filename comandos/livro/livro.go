package livro

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetLivro() string {

	resp, err := http.Get("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var texto string
	livro := doc.Find(".dotd-title").Text()
	tempo := doc.Find(".packt-js-countdown").Text()

	texto = texto + "Confira o livro gratuito de hoje: \n"
	texto = texto + "Nome: " + strings.TrimSpace(livro) + " \n"
	texto = texto + "Proximo livro em " + strings.TrimSpace(tempo) + " \n \n"
	texto = texto + "Resgate seu livro em: https://www.packtpub.com/packt/offers/free-learning"

	return texto
}
