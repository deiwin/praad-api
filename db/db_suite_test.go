package db_test

import (
	"github.com/Lunchr/luncher-api/db"
	"github.com/Lunchr/luncher-api/db/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	dbClient                           *db.Client
	offersCollection                   db.Offers
	offerGroupPostsCollection          db.OfferGroupPosts
	tagsCollection                     db.Tags
	regionsCollection                  db.Regions
	restaurantsCollection              db.Restaurants
	usersCollection                    db.Users
	registrationAccessTokensCollection db.RegistrationAccessTokens
	mocks                              *Mocks
)

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}

var _ = BeforeSuite(func() {
	mocks = createMocks()
	createClient()
	wipeDb()
	initCollections()
})

// RebuildDBAfterEach can be used on tests or test blocks that mess up the data
// in the DB. Most commonly inserts and updates
var RebuildDBAfterEach = func() {
	AfterEach(func() {
		wipeDb()
		initCollections()
	})
}
var _ = AfterSuite(func() {
	wipeDb()
	dbClient.Disconnect()
})

var _ = It("should work", func() {})

func createClient() {
	dbConfig := createTestDbConf()
	dbClient = db.NewClient(dbConfig)
	err := dbClient.Connect()
	Expect(err).NotTo(HaveOccurred())
}

func initCollections() {
	initOffersCollection()
	initOfferGroupPostsCollection()
	initTagsCollection()
	initRegionsCollection()
	initRestaurantsCollection()
	initUsersCollection()
	initRegistrationAccessTokensCollection()
}

func initOffersCollection() {
	var err error
	offersCollection, err = db.NewOffers(dbClient)
	Expect(err).NotTo(HaveOccurred())
	_, err = insertOffers()
	Expect(err).NotTo(HaveOccurred())
}

func initOfferGroupPostsCollection() {
	offerGroupPostsCollection = db.NewOfferGroupPosts(dbClient)
}

func initTagsCollection() {
	tagsCollection = db.NewTags(dbClient)
	err := insertTags()
	Expect(err).NotTo(HaveOccurred())
}

func initRegionsCollection() {
	regionsCollection = db.NewRegions(dbClient)
	err := insertRegions()
	Expect(err).NotTo(HaveOccurred())
}

func initRestaurantsCollection() {
	restaurantsCollection = db.NewRestaurants(dbClient)
	_, err := insertRestaurants()
	Expect(err).NotTo(HaveOccurred())
}

func initUsersCollection() {
	usersCollection = db.NewUsers(dbClient)
	err := insertUsers()
	Expect(err).NotTo(HaveOccurred())
}

func initRegistrationAccessTokensCollection() {
	var err error
	registrationAccessTokensCollection, err = db.NewRegistrationAccessTokens(dbClient)
	Expect(err).NotTo(HaveOccurred())
}

func createTestDbConf() (dbConfig *db.Config) {
	dbConfig = &db.Config{
		DbURL:  "127.0.0.1",
		DbName: "test",
	}
	return
}

func insertTags() (err error) {
	return tagsCollection.Insert(mocks.tags...)
}

func insertRegions() (err error) {
	return regionsCollection.Insert(mocks.regions...)
}

func insertRestaurants() ([]*model.Restaurant, error) {
	return restaurantsCollection.Insert(mocks.restaurants...)
}

func insertOffers() ([]*model.Offer, error) {
	return offersCollection.Insert(mocks.offers...)
}

func insertUsers() (err error) {
	return usersCollection.Insert(mocks.users...)
}

func wipeDb() {
	err := dbClient.WipeDb()
	Expect(err).NotTo(HaveOccurred())
}
