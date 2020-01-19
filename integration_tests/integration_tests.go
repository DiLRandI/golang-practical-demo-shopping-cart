package integration_tests

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
)

//PackingDetail ..
type PackingDetailResponse struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type ItemResponse struct {
	ID             int           `json:"id"`
	ItemCode       string        `json:"item_code"`
	Description    string        `json:"description"`
	UnitPrice      float32       `json:"unit_price"`
	PackingDetails PackingDetailResponse `json:"packing_details"`
}

type apiFeature struct {
	serverURI string
	
	respCode int
	err string

	itemResponse ItemResponse 
	cartResponse []int

}

func (a *apiFeature) clear(interface{}) {
	a.serverURI = ""
	}

	func init() {
		godog.BindFlags("godog.", flag.CommandLine, &opt)
	}
	
	var opt = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress", // can define default values
	}

	func FeatureContext(s *godog.Suite) {
		a := &apiFeature{}
		s.BeforeScenario(a.clear)
		itemFeatureContext(s,a)
	
	}
	
	// Use "func TestMain(m *testing.M) {" to debug godog unit tests
	func testMain(m *testing.M) {
		flag.Parse()
		path, _ := filepath.Abs(".")
		path = path + "/features"
		_, err := ioutil.ReadDir(path)
	
		if err != nil {
			log.Fatal(err)
		}
		opt.Paths = make([]string, 0)
		
		status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
			FeatureContext(s)
		}, opt)
	
		if st := m.Run(); st > status {
			status = st
		}
		os.Exit(status)
	}