package service

import (
	"example/com/app/model"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const lnk = "https://api2.binance.com/api/v3/ticker/24hr"

func FetchApi() []byte {

	response, err := http.Get(lnk)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)
	return data
}

func GetModel() model.CoinModel {

	var cModel model.CoinModel

	cModel = model.UnMarshallCoins(FetchApi())

	return cModel

}

func CreateCoinIndexMap(model model.CoinModel) map[string]int {

	res := model

	coinList := make(map[string]int)

	for i, s := range res {

		coinList[s.Symbol] = i
	}

	return coinList

}

func CoinNameIdentifier() string {

	var coinName string

	fmt.Scanf("%s", &coinName)

	coinName = strings.ToUpper(coinName)

	return coinName
}

func FindCoinFromIndexMap(str string) (model.CoinModel, int) {

	//cName := CoinNameIdentifier()
	cName := strings.ToUpper(str)
	nFetch := GetModel()
	cMap := CreateCoinIndexMap(nFetch)
	index := cMap[cName]

	if index == 0 && cName != "ETHBTC" {
		fmt.Println("Coin is doesn't exist \n Try Again.")
		log.Fatal()

	}

	//fmt.Println(cMap)

	return nFetch, index

}

func CreateData(cModel model.CoinModel, index int) [][]string {

	res := cModel[index]

	data := [][]string{
		{res.Symbol, res.AskPrice, res.PriceChange, res.HighPrice},
	}

	return data

}

func RefreshData(data [][]string) [][]string {

	res := data[0][0]
	nFetch := GetModel()
	index := CreateCoinIndexMap(nFetch)[res]
	nData := CreateData(nFetch, index)
	return nData
}

func CreateTable(data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Coin Name", "Current Price", "Price Change", "Price Change Percent"})
	table.SetFooter([]string{"Options", "For Refresh Press: 1 ", "For Another Coin Press: 2", "For Exit Press: Any"}) // Add Footer
	table.SetBorder(false)                                                                                            // Set Border to false

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.BgHiBlueColor, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor})

	table.SetFooterColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor}, tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgHiMagentaColor})

	table.AppendBulk(data)
	table.Render()

}
