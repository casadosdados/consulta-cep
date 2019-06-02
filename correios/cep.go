package correios

import (
	"encoding/base64"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const URL = "http://www.buscacep.correios.com.br/sistemas/buscacep/resultadoBuscaCepEndereco.cfm"
const PAGINATION = 90

const MUNICIPIO_FIELD = "municipio"

type ResponseService struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro string `json:"bairro"`
	Municipio string `json:"municipio"`
	UF string `json:"uf"`
}

type CollectionCEP []ResponseService

func SearchALL(q string) *CollectionCEP {
	c := CollectionCEP{}
	for i:=1; i<=11; i++ {
		s, err := Search(q, i)
		if err == nil {
			for _, n := range  *s {
				c = append(c, n)
			}
		}
		if len(*s) < 90 {
			break
		}
	}
	return &c
}

func Search(q string, page int) (*CollectionCEP, error) {
	pageIni := PAGINATION * (page - 1)
	pageFim := page * PAGINATION
	if pageIni == 0 {
		pageIni = 1
	}
	auth := "cdd:cdd"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	res, body, errs := Request().Post(URL).
		Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36").
		Set("Origin","http://www.buscacep.correios.com.br").
		Set("Referer","http://www.buscacep.correios.com.br/sistemas/buscacep/buscaCepEndereco.cfm").
		Set("Proxy-Authorization", basicAuth).
		Proxy("http://191.252.186.230:8090").
		Type("form-data").
		Send("relaxation=" + q).
		Send("tipoCEP=ALL").
		Send("semelhante=Y").
		Send("qtdrow=50").
		Send("pagini=" + strconv.Itoa(pageIni)).
		Send("pagfim=" + strconv.Itoa(pageFim)).
		End()
	c := new(CollectionCEP)
	if len(errs) > 0 {
		return c, errs[0]
	}
	if res.StatusCode != 200 {
		return c, errors.New("status code diferente de 200 retornado: " + strconv.Itoa(res.StatusCode))
	}
	return ParseHtml(strings.NewReader(string(ToUtf8([]byte(body)))))
}

var space = regexp.MustCompile("\u00a0")

func ParseHtml(body io.Reader) (*CollectionCEP, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		//
	}
	data := CollectionCEP{}
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			if indextr == 0 {
				return
			}
			d := ResponseService{}
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				content := space.ReplaceAllString(tablecell.Text(), "")
				switch indexth {
				case 0:
					d.Logradouro = content
					break
				case 1:
					d.Bairro = content
					break
				case 2:
					s := strings.Split(content, "/")
					if len(s) == 2 {
						d.Municipio = s[0]
						d.UF = s[1]
					}
					break
				case 3:
					d.Cep = strings.Replace(content, "-", "", 1)
					break
				}
			})
			if len(d.Cep) == 8 {
				data = append(data, d)
			}
		})
	})
	return &data, nil
}

func ToUtf8(str []byte) []byte {
	buf := make([]rune, len(str))
	for i, b := range str {
		buf[i] = rune(b)
	}
	return []byte(string(buf))
}

func Request() *gorequest.SuperAgent {
	g := gorequest.New()
	g.Timeout(20 * time.Second)
	g.Set("user-agent", userAgentRandom())
	return g
}

var userAgents =[]string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0.1; SM-G920V Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A5370a Safari/604.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
	"Dalvik/2.1.0"}

func userAgentRandom() string {
	return userAgents[randInt(0, len(userAgents))]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

