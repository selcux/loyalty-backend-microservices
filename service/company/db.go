package company

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbInterface interface {
	Create(model *CreateDto) (*Company, error)
	Read(id string) (*Company, error)
	ReadAll() ([]*Company, error)
	Update(id string, model *UpdateDto) error
	Delete(id string) error
	Close() error
}

//Db provides required parameters for db operations
type Db struct {
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDb() (*Db, error) {
	//conf := di.InitializeConfig()
	//mongoProps := conf.MongoProps()

	ctx := context.TODO()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoProps.ConnectionString))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:37017"))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Db{
		ctx:    ctx,
		client: client,
		//collection: client.Database(mongoProps.DbName).Collection(mongoProps.Collections["company"]),
		collection: client.Database("loyalty-dlt").Collection("companies"),
	}, nil
}

func (db *Db) Create(model *CreateDto) (*Company, error) {
	res, err := db.collection.InsertOne(db.ctx, model)
	if err != nil {
		return nil, err
	}

	return &Company{
		ID:   res.InsertedID.(primitive.ObjectID),
		Name: model.Name,
	}, nil
}

func (db *Db) Read(id string) (*Company, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	filter := bson.M{"_id": filterId}

	var companyData *Company
	err = db.collection.FindOne(db.ctx, filter).Decode(&companyData)
	if err != nil {
		return nil, err
	}

	return companyData, nil
}

func (db *Db) ReadAll() ([]*Company, error) {
	cur, err := db.collection.Find(db.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var rerr error = nil
	defer func() {
		if ferr := cur.Close(db.ctx); ferr != nil {
			rerr = ferr
		}
	}()

	companies := make([]*Company, 0)

	for cur.Next(db.ctx) {
		var result *Company
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		companies = append(companies, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return companies, rerr
}

func (db *Db) Update(id string, model *UpdateDto) error {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": filterId}
	update := bson.M{"$set": bson.M{"name": model.Name}}

	_, err = db.collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) Delete(id string) error {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": filterId}

	res, err := db.collection.DeleteOne(db.ctx, filter)
	if err != nil {
		return err
	}

	log.Debug(fmt.Sprintf("Deleted result count: %d", res.DeletedCount))

	return nil
}

func (db *Db) Close() error {
	return db.client.Disconnect(db.ctx)
}
