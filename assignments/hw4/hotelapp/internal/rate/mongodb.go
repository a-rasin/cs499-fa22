//go:build mongodb

package rate

import (
	log "github.com/sirupsen/logrus"

	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
}

func NewDatabaseSession(db_addr string) *DatabaseSession {
	// DONE?: Implement me
	session, err := mgo.Dial(db_addr)
	if err != nil{
		log.Fatal(err)
	}
	log.Info("New session successfull.")

	return &DatabaseSession{
		MongoSession: session,
	}
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "rate-db", "inventory", rateJsonPath)
}

// GetRates gets rates for hotels for specific date range.
func (db *DatabaseSession) GetRates(hotelIds []string) (RatePlans, error) {
	// DONE: Implement me
	session := db.MongoSession.Copy()
	defer session.Close()
	c := session.DB("rate-db").C("inventory")

	rates := make([]*RatePlans, 0)

	for _, := range hotelIds {
		temprate := new(RatePlans)
		err := c.Find(bson.M{"id" : id}).One(&temorate)
		if err != nil{
			log.Ftalf("Failed to get rate data: ", err)
		}
		rates = append(rates, temprate)
	}
	return rates, ni
}
