package merchant

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbInterface interface {
	Create(model *CreateDto) (*Merchant, error)
	Read(id string) (*Merchant, error)
	ReadAll() ([]*Merchant, error)
	Update(id string, model *UpdateDto) (*Merchant, error)
	Delete(id string) error
	Close() error
}

type Db struct {
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDb() (*Db, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoProps := conf.MongoProps()

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoProps.ConnectionString))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Db{
		ctx:        ctx,
		client:     client,
		collection: client.Database(mongoProps.DbName).Collection(mongoProps.Collections["merchant"]),
	}, nil
}

func (db *Db) Create(model *CreateDto) (*Merchant, error) {
	res, err := db.collection.InsertOne(db.ctx, model)
	if err != nil {
		return nil, err
	}

	return &Merchant{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     model.Name,
		Location: model.Location,
	}, nil
}

func (db *Db) Read(id string) (*Merchant, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	filter := bson.M{"_id": filterId}

	var merchantData *Merchant

	err = db.collection.FindOne(db.ctx, filter).Decode(&merchantData)
	if err != nil {
		return nil, errors.New("not found")
	}

	return merchantData, nil
}

func (db *Db) ReadAll() ([]*Merchant, error) {
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

	merchants := make([]*Merchant, 0)

	for cur.Next(db.ctx) {
		var result *Merchant
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return merchants, rerr
}

func (db *Db) Update(id string, model *UpdateDto) (*Merchant, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": filterId}
	updateContent := bson.D{}

	// TODO: Needs improvement
	if model.Name != "" {
		updateContent = append(updateContent, bson.E{Key: "name", Value: model.Name})
	}

	update := bson.M{"$set": updateContent}
	_, err = db.collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	merchantData, err := db.Read(id)

	return merchantData, err
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
