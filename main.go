package main


import (
    "github.com/DorryIT/dbapi"
    "github.com/SpeculationFund/eccrawler"
//    "fmt"
)

/*
    Database configuration
*/
const (
    DBADDR string = "127.0.0.1"
    DBNAME string = "SpeculationFund"
    TICKER_COLLECTION string = "Ticker"
)

func main() {

    // Init db session
    sm := dbapi.Init(DBADDR)
    // Init crawler processor
    crawlersfactory := &crawlers.C_CrawlerFactory{}
    cpu := crawlersfactory.CreateCPU()


    // Retrieve data as save
    // TODO: It will hang since I did nothing to control chan
    // TODO: Work on time control
    result := make(chan interface{})
    cpu.GetTickers(result)
    for {
        sm.Insert(DBNAME, TICKER_COLLECTION, <-result)
    }
}
