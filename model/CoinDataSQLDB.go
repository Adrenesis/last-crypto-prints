package model

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"strings"
	//"time"
)

func ReadCryptoSQLDB(id int64) CoinDatum {
	fmt.Println("reading database...")
	db, err := sql.Open("sqlite", "./cryptoDB.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, name, slug, symbol, logo, price, vol24, date_added, explorer, bscscan, ethscan, xrpscan, bsccontract, ethcontract, xrpcontract, twitter, website, facebook, chat, message_board, technical, source_code, announcement, tag, max_supply , circulating_supply, vol24, volchange24, percentchange24, percentchange7d, percentchange30d, percentchange60d, percentchange90d, market_cap, market_cap_dominance, fully_diluted_market_cap from cryptos where id = ? order by date_added desc;", fmt.Sprintf("%d", id))
	if err != nil {
		log.Fatal(err)
	}

	var coinDatum CoinDatum
	for rows.Next() {
		var slug interface{}
		var logo interface{}
		var tags interface{}
		var explorer interface{}
		var twitter interface{}
		var website interface{}
		var facebook interface{}
		var chat interface{}
		var messageBoard interface{}
		var technical interface{}
		var sourceCode interface{}
		var announcement interface{}
		var bscScan interface{}
		var ethScan interface{}
		var xrpScan interface{}
		var bscContract interface{}
		var ethContract interface{}
		var xrpContract interface{}
		if err = rows.Scan(
			&coinDatum.Id,
			&coinDatum.Name,
			&slug,
			&coinDatum.Symbol,
			&logo,
			&coinDatum.Properties.Dollar.Price,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.DateAdded,
			&explorer,
			&bscScan,
			&ethScan,
			&xrpScan,
			&bscContract,
			&ethContract,
			&xrpContract,
			&twitter,
			&website,
			&facebook,
			&chat,
			&messageBoard,
			&technical,
			&sourceCode,
			&announcement,
			&tags,
			&coinDatum.MaxSupply,
			&coinDatum.CirculatingSupply,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.Properties.Dollar.VolumeChange24,
			&coinDatum.Properties.Dollar.PercentChange24,
			&coinDatum.Properties.Dollar.PercentChange7d,
			&coinDatum.Properties.Dollar.PercentChange30d,
			&coinDatum.Properties.Dollar.PercentChange60d,
			&coinDatum.Properties.Dollar.PercentChange90d,
			&coinDatum.Properties.Dollar.MarketCap,
			&coinDatum.Properties.Dollar.MarketCapDominance,
			&coinDatum.Properties.Dollar.FullyDilutedMarketPrice); err != nil {
			log.Fatal(err)
		}
		coinDatum.Tags = strings.Split(fmt.Sprintf("%v", tags), ",")
		coinDatum.Explorers = strings.Split(fmt.Sprintf("%v", explorer), ",")
		coinDatum.Twitters = strings.Split(fmt.Sprintf("%v", twitter), ",")
		coinDatum.Facebooks = strings.Split(fmt.Sprintf("%v", facebook), ",")
		coinDatum.Websites = strings.Split(fmt.Sprintf("%v", website), ",")
		coinDatum.MessageBoards = strings.Split(fmt.Sprintf("%v", messageBoard), ",")
		coinDatum.Chats = strings.Split(fmt.Sprintf("%v", chat), ",")
		coinDatum.Technicals = strings.Split(fmt.Sprintf("%v", technical), ",")
		coinDatum.SourceCodes = strings.Split(fmt.Sprintf("%v", sourceCode), ",")
		coinDatum.Announcements = strings.Split(fmt.Sprintf("%v", announcement), ",")
		coinDatum.Slug = fmt.Sprintf("%v", slug)
		if coinDatum.Slug == "<nil>" {
			coinDatum.Slug = strings.ToLower(coinDatum.Name)
			coinDatum.Slug = strings.Replace(coinDatum.Slug, " ", "-", -1)
		}
		coinDatum.Logo = fmt.Sprintf("%v", logo)
		//fmt.Println("logo", coinDatum.Logo)
		coinDatum.BscScan = fmt.Sprintf("%v", bscScan)
		coinDatum.EthScan = fmt.Sprintf("%v", ethScan)
		coinDatum.XrpScan = fmt.Sprintf("%v", xrpScan)
		coinDatum.BscContract = fmt.Sprintf("%v", bscContract)
		coinDatum.EthContract = fmt.Sprintf("%v", ethContract)
		coinDatum.XrpContract = fmt.Sprintf("%v", xrpContract)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(fmt.Sprintf("%v", coinData))

	CloseDB(db)
	return coinDatum
}

func ReadCryptoByBSCContractSQLDB(contract string) CoinDatum {
	fmt.Println("reading database...")
	db, err := sql.Open("sqlite", "./cryptoDB.db")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select id, name, slug, symbol, logo, price, vol24, date_added, explorer, bscscan, ethscan, xrpscan, bsccontract, ethcontract, xrpcontract, twitter, website, facebook, chat, message_board, technical, source_code, announcement, tag, max_supply , circulating_supply, vol24, volchange24, percentchange24, percentchange7d, percentchange30d, percentchange60d, percentchange90d, market_cap, market_cap_dominance, fully_diluted_market_cap from cryptos where UPPER(bsccontract) LIKE UPPER('" + contract + "') order by date_added desc;")
	if err != nil {
		log.Fatal(err)
	}
	var coinDatum CoinDatum
	for rows.Next() {
		var slug interface{}
		var logo interface{}
		var tags interface{}
		var explorer interface{}
		var twitter interface{}
		var website interface{}
		var facebook interface{}
		var chat interface{}
		var messageBoard interface{}
		var technical interface{}
		var sourceCode interface{}
		var announcement interface{}
		var bscScan interface{}
		var ethScan interface{}
		var xrpScan interface{}
		var bscContract interface{}
		var ethContract interface{}
		var xrpContract interface{}
		if err = rows.Scan(
			&coinDatum.Id,
			&coinDatum.Name,
			&slug,
			&coinDatum.Symbol,
			&logo,
			&coinDatum.Properties.Dollar.Price,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.DateAdded,
			&explorer,
			&bscScan,
			&ethScan,
			&xrpScan,
			&bscContract,
			&ethContract,
			&xrpContract,
			&twitter,
			&website,
			&facebook,
			&chat,
			&messageBoard,
			&technical,
			&sourceCode,
			&announcement,
			&tags,
			&coinDatum.MaxSupply,
			&coinDatum.CirculatingSupply,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.Properties.Dollar.VolumeChange24,
			&coinDatum.Properties.Dollar.PercentChange24,
			&coinDatum.Properties.Dollar.PercentChange7d,
			&coinDatum.Properties.Dollar.PercentChange30d,
			&coinDatum.Properties.Dollar.PercentChange60d,
			&coinDatum.Properties.Dollar.PercentChange90d,
			&coinDatum.Properties.Dollar.MarketCap,
			&coinDatum.Properties.Dollar.MarketCapDominance,
			&coinDatum.Properties.Dollar.FullyDilutedMarketPrice); err != nil {
			log.Fatal(err)
		}
		coinDatum.Tags = strings.Split(fmt.Sprintf("%v", tags), ",")
		coinDatum.Explorers = strings.Split(fmt.Sprintf("%v", explorer), ",")
		coinDatum.Twitters = strings.Split(fmt.Sprintf("%v", twitter), ",")
		coinDatum.Facebooks = strings.Split(fmt.Sprintf("%v", facebook), ",")
		coinDatum.Websites = strings.Split(fmt.Sprintf("%v", website), ",")
		coinDatum.MessageBoards = strings.Split(fmt.Sprintf("%v", messageBoard), ",")
		coinDatum.Chats = strings.Split(fmt.Sprintf("%v", chat), ",")
		coinDatum.Technicals = strings.Split(fmt.Sprintf("%v", technical), ",")
		coinDatum.SourceCodes = strings.Split(fmt.Sprintf("%v", sourceCode), ",")
		coinDatum.Announcements = strings.Split(fmt.Sprintf("%v", announcement), ",")
		coinDatum.Slug = fmt.Sprintf("%v", slug)
		if coinDatum.Slug == "<nil>" {
			coinDatum.Slug = strings.ToLower(coinDatum.Name)
			coinDatum.Slug = strings.Replace(coinDatum.Slug, " ", "-", -1)
		}
		coinDatum.Logo = fmt.Sprintf("%v", logo)
		//fmt.Println("logo", coinDatum.Logo)
		coinDatum.BscScan = fmt.Sprintf("%v", bscScan)
		coinDatum.EthScan = fmt.Sprintf("%v", ethScan)
		coinDatum.XrpScan = fmt.Sprintf("%v", xrpScan)
		coinDatum.BscContract = fmt.Sprintf("%v", bscContract)
		coinDatum.EthContract = fmt.Sprintf("%v", ethContract)
		coinDatum.XrpContract = fmt.Sprintf("%v", xrpContract)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(fmt.Sprintf("%v", coinData))

	CloseDB(db)
	return coinDatum
}

func ReadCryptosSQLDB() CoinData {
	fmt.Println("reading database...")
	db, err := sql.Open("sqlite", "./cryptoDB.db")
	if err != nil {
		log.Fatal(err)
	}
	CreateCryptoTable()
	rows, err := db.Query("select id, name, slug, symbol, logo, price, vol24, date_added, explorer, bscscan, ethscan, xrpscan, bsccontract, ethcontract, xrpcontract, twitter, website, facebook, chat, message_board, technical, source_code, announcement, tag, max_supply , circulating_supply, vol24, volchange24, percentchange24, percentchange7d, percentchange30d, percentchange60d, percentchange90d, market_cap, market_cap_dominance, fully_diluted_market_cap from cryptos order by date_added desc;")
	if err != nil {
		log.Fatal(err)
	}
	var coinData CoinData
	for rows.Next() {
		var coinDatum CoinDatum
		var slug interface{}
		var logo interface{}
		var explorer interface{}
		var twitter interface{}
		var website interface{}
		var facebook interface{}
		var chat interface{}
		var messageBoard interface{}
		var technical interface{}
		var sourceCode interface{}
		var announcement interface{}
		var tags interface{}
		var bscScan interface{}
		var ethScan interface{}
		var xrpScan interface{}
		var bscContract interface{}
		var ethContract interface{}
		var xrpContract interface{}
		if err = rows.Scan(
			&coinDatum.Id,
			&coinDatum.Name,
			&slug,
			&coinDatum.Symbol,
			&logo,
			&coinDatum.Properties.Dollar.Price,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.DateAdded,
			&explorer,
			&bscScan,
			&ethScan,
			&xrpScan,
			&bscContract,
			&ethContract,
			&xrpContract,
			&twitter,
			&website,
			&facebook,
			&chat,
			&messageBoard,
			&technical,
			&sourceCode,
			&announcement,
			&tags,
			&coinDatum.MaxSupply,
			&coinDatum.CirculatingSupply,
			&coinDatum.Properties.Dollar.Volume24,
			&coinDatum.Properties.Dollar.VolumeChange24,
			&coinDatum.Properties.Dollar.PercentChange24,
			&coinDatum.Properties.Dollar.PercentChange7d,
			&coinDatum.Properties.Dollar.PercentChange30d,
			&coinDatum.Properties.Dollar.PercentChange60d,
			&coinDatum.Properties.Dollar.PercentChange90d,
			&coinDatum.Properties.Dollar.MarketCap,
			&coinDatum.Properties.Dollar.MarketCapDominance,
			&coinDatum.Properties.Dollar.FullyDilutedMarketPrice); err != nil {
			log.Fatal(err)
		}
		coinDatum.Tags = strings.Split(fmt.Sprintf("%v", tags), ",")
		coinDatum.Explorers = strings.Split(fmt.Sprintf("%v", explorer), ",")
		coinDatum.Twitters = strings.Split(fmt.Sprintf("%v", twitter), ",")
		coinDatum.Facebooks = strings.Split(fmt.Sprintf("%v", facebook), ",")
		coinDatum.Websites = strings.Split(fmt.Sprintf("%v", website), ",")
		coinDatum.MessageBoards = strings.Split(fmt.Sprintf("%v", messageBoard), ",")
		coinDatum.Chats = strings.Split(fmt.Sprintf("%v", chat), ",")
		coinDatum.Technicals = strings.Split(fmt.Sprintf("%v", technical), ",")
		coinDatum.SourceCodes = strings.Split(fmt.Sprintf("%v", sourceCode), ",")
		coinDatum.Announcements = strings.Split(fmt.Sprintf("%v", announcement), ",")
		coinDatum.Slug = fmt.Sprintf("%v", slug)
		if coinDatum.Slug == "<nil>" {
			coinDatum.Slug = strings.ToLower(coinDatum.Name)
			coinDatum.Slug = strings.Replace(coinDatum.Slug, " ", "-", -1)
		}
		coinDatum.Logo = fmt.Sprintf("%v", logo)
		//fmt.Println("logo", coinDatum.Logo)
		coinDatum.BscScan = fmt.Sprintf("%v", bscScan)
		coinDatum.EthScan = fmt.Sprintf("%v", ethScan)
		coinDatum.XrpScan = fmt.Sprintf("%v", xrpScan)
		coinDatum.BscContract = fmt.Sprintf("%v", bscContract)
		coinDatum.EthContract = fmt.Sprintf("%v", ethContract)
		coinDatum.XrpContract = fmt.Sprintf("%v", xrpContract)
		coinData.CoinData = append(coinData.CoinData, coinDatum)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(fmt.Sprintf("%v", coinData))

	CloseDB(db)
	return coinData
}

func CreateCryptoTable() {
	db, err := sql.Open("sqlite", "./cryptoDB.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = db.Exec(`
-- drop table if exists cryptos;
create table if not exists cryptos(id INTEGER, name VARCHAR, slug VARCHAR, symbol VARCHAR, logo TEXT, price REAL, vol24 REAL, date_added TEXT, explorer TEXT, bscscan TEXT, ethscan TEXT, xrpscan TEXT, bsccontract TEXT, ethcontract TEXT, xrpcontract TEXT, twitter TEXT, website TEXT, facebook TEXT, chat TEXT, message_board TEXT, technical TEXT, source_code TEXT, announcement TEXT, tag TEXT, max_supply REAL, circulating_supply REAL, volchange24 REAL, percentchange24 REAL, percentchange7d REAL, percentchange30d REAL, percentchange60d REAL, percentchange90d REAL, market_cap REAL, market_cap_dominance REAL, fully_diluted_market_cap REAL, PRIMARY KEY(id));
	`); err != nil {
		log.Fatal(err)
	}
}
func writeCrypto(id int64, name string, symbol string, dateAdded string, properties Property, tags []string, maxSupply float64, circulatingSupply float64) {
	db, tx, stmt := Prepare("INSERT INTO cryptos (id, name, symbol, date_added, tag, max_supply, circulating_supply, price, vol24, volchange24, percentchange24, percentchange7d, percentchange30d, percentchange60d, percentchange90d, market_cap, market_cap_dominance, fully_diluted_market_cap) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	tagString := SerializeStringList(tags)
	ExecIgnoreDuplicate(tx, stmt, id, name, symbol, dateAdded, tagString, maxSupply, circulatingSupply, properties.Dollar.Price, properties.Dollar.Volume24, properties.Dollar.VolumeChange24, properties.Dollar.PercentChange24, properties.Dollar.PercentChange7d, properties.Dollar.PercentChange30d, properties.Dollar.PercentChange60d, properties.Dollar.PercentChange90d, properties.Dollar.MarketCap, properties.Dollar.MarketCapDominance, properties.Dollar.FullyDilutedMarketPrice)
	CloseDB(db)
}
func WriteCryptosSQLDB(coinData CoinData) {
	db, err := sql.Open("sqlite", "./cryptoDB.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	CreateCryptoTable()

	for i := 0; i < len(coinData.CoinData); i++ {
		//fmt.Println("INSERT INTO cryptos (name, symbol, price, vol24, date_added) VALUES ('"+
		//	strings.Replace(coinData.CoinData[i].Name, "'", "''",-1) +"', '" +
		//	strings.Replace(coinData.CoinData[i].Symbol, "'", "''",-1) +"', '" +
		//	fmt.Sprintf("%.7+f", coinData.CoinData[i].Properties.Dollar.Price) +"', '" +
		//	fmt.Sprintf("%.2f", coinData.CoinData[i].Properties.Dollar.Volume24) +"', '" +
		//	coinData.CoinData[i].DateAdded +"');")
		//
		writeCrypto(coinData.CoinData[i].Id,
			coinData.CoinData[i].Name,
			coinData.CoinData[i].Symbol,
			coinData.CoinData[i].DateAdded,
			coinData.CoinData[i].Properties,
			coinData.CoinData[i].Tags,
			coinData.CoinData[i].MaxSupply,
			coinData.CoinData[i].CirculatingSupply,
		)
	}
	CloseDB(db)

}

func writeUrls(explorer string, twitter string, website string, facebook string, chat string, messageBoard string, technical string, sourceCode string, announcement string, id int64) {
	db, tx, updateExplorers := Prepare("UPDATE cryptos SET explorer =?, twitter = ?, website = ?, facebook = ?, chat = ?, message_board = ?, technical = ?, source_code = ?, announcement = ? WHERE id = ?;")
	Exec(tx, updateExplorers, explorer, twitter, website, facebook, chat, messageBoard, technical, sourceCode, announcement, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func writeExplorer(explorer string, id int64) {
	db, tx, updateExplorers := Prepare("UPDATE cryptos SET explorer = ? WHERE id = ?;")
	Exec(tx, updateExplorers, explorer, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func writeBscScan(bscScan string, bscContract string, id int64) {
	db, tx, stmt := Prepare("UPDATE cryptos SET bscscan = ?, bsccontract = ? WHERE id = ?;")
	Exec(tx, stmt, bscScan, bscContract, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func writeEthScan(ethScan string, ethContract string, id int64) {
	db, tx, stmt := Prepare("UPDATE cryptos SET ethscan = ?, ethcontract = ? WHERE id = ?;")
	Exec(tx, stmt, ethScan, ethContract, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func writeXrpScan(xrpScan string, xrpContract string, id int64) {
	db, tx, stmt := Prepare("UPDATE cryptos SET xrpscan = ?, xrpcontract = ? WHERE id = ?;")
	Exec(tx, stmt, xrpScan, xrpContract, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func writeMap(slug string, logo string, id int64) {
	db, tx, stmt := Prepare("UPDATE cryptos SET slug = ?, logo = ? WHERE id = ?;")
	Exec(tx, stmt, slug, logo, fmt.Sprintf("%d", id))
	CloseDB(db)
}

func WriteCryptosMapSQLDB(coinDataMap CoinDataMap) {

	CreateCryptoTable()

	for i := 0; i < len(coinDataMap.CoinDataMap); i++ {
		//fmt.Println("INSERT INTO cryptos (name, symbol, price, vol24, date_added) VALUES ('"+
		//	strings.Replace(coinData.CoinData[i].Name, "'", "''",-1) +"', '" +
		//	strings.Replace(coinData.CoinData[i].Symbol, "'", "''",-1) +"', '" +
		//	fmt.Sprintf("%.7+f", coinData.CoinData[i].Properties.Dollar.Price) +"', '" +
		//	fmt.Sprintf("%.2f", coinData.CoinData[i].Properties.Dollar.Volume24) +"', '" +
		//	coinData.CoinData[i].DateAdded +"');")
		explorer := ""
		bscScan := ""
		ethScan := ""
		xrpScan := ""
		bscContract := ""
		ethContract := ""
		xrpContract := ""
		for j := 0; j < len(coinDataMap.CoinDataMap[i].URLs.Explorer); j++ {
			if explorer == "" {
				explorer = coinDataMap.CoinDataMap[i].URLs.Explorer[j]

			} else {
				explorer += "," + coinDataMap.CoinDataMap[i].URLs.Explorer[j]
			}
			//fmt.Println(coinDataMap.CoinDataMap[i].URLs.Explorer[j])
			//fmt.Println(strings.Contains(coinDataMap.CoinDataMap[i].URLs.Explorer[j], "bscscan"))
			if strings.Contains(coinDataMap.CoinDataMap[i].URLs.Explorer[j], "bscscan") {
				bscScan = coinDataMap.CoinDataMap[i].URLs.Explorer[j]

				bscContract = strings.TrimPrefix(bscScan, "https://www.bscscan.com/token/")
				bscContract = strings.TrimPrefix(bscContract, "https://bscscan.com/token/")
				bscContract = strings.TrimPrefix(bscContract, "https://bscscan.com/address/")
				bscContract = strings.TrimPrefix(bscContract, "https://www.bscscan.com/address/")
				if len(bscContract) > 42 {
					bscContract = bscContract[:42]
				}

				if strings.HasPrefix(bscContract, "0x") {
					writeBscScan(bscScan, bscContract, coinDataMap.CoinDataMap[i].Id)
				}

			}
			if strings.Contains(coinDataMap.CoinDataMap[i].URLs.Explorer[j], "etherscan") {
				ethScan = coinDataMap.CoinDataMap[i].URLs.Explorer[j]
				ethContract = strings.TrimPrefix(ethScan, "https://etherscan.io/token/")
				ethContract = strings.TrimPrefix(ethContract, "https://www.etherscan.io/token/")
				writeEthScan(ethScan, ethContract, coinDataMap.CoinDataMap[i].Id)
			}
			if strings.Contains(coinDataMap.CoinDataMap[i].URLs.Explorer[j], "xrpscan") {
				xrpScan = coinDataMap.CoinDataMap[i].URLs.Explorer[j]
				xrpContract = strings.TrimPrefix(xrpScan, "https://xrpscan.com/account/")
				xrpContract = strings.TrimPrefix(xrpContract, "https://www.xrpscan.com/account/")
				writeXrpScan(xrpScan, xrpContract, coinDataMap.CoinDataMap[i].Id)
			}
			fmt.Println(coinDataMap.CoinDataMap[i].DateAdded)
			//fmt.Println(coinDataMap.CoinDataMap[i].Logo)

			writeMap(coinDataMap.CoinDataMap[i].Slug, coinDataMap.CoinDataMap[i].Logo, coinDataMap.CoinDataMap[i].Id)
		}
		twitter := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Twitter)
		website := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Website)
		facebook := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Facebook)
		chat := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Chat)
		messageBoard := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.MessageBoard)
		technical := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Technical)
		sourceCode := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.SourceCode)
		announcement := SerializeStringList(coinDataMap.CoinDataMap[i].URLs.Announcement)
		//fmt.Println(explorer)

		//args[1] = coinDataMap.CoinDataMap[i].Id
		//args.Explorer = explorer
		//args.Id = coinDataMap.CoinDataMap[i].Id
		//writeExplorer(explorer, coinDataMap.CoinDataMap[i].Id)
		writeUrls(explorer, twitter, website, facebook, chat, messageBoard, technical, sourceCode, announcement, coinDataMap.CoinDataMap[i].Id)
		//writeUrls(website, "website", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(facebook, "facebook", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(chat, "chat", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(messageBoard, "message_board", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(technical, "technical", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(sourceCode, "source_code", coinDataMap.CoinDataMap[i].Id)
		//writeUrls(announcement, "announcement", coinDataMap.CoinDataMap[i].Id)
		//fmt.Println(fmt.Sprintf("%v", res))
	}

}
