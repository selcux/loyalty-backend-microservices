package item

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
	Create(model *CreateDto) (*Entity, error)
	Read(id string) (*Entity, error)
	ReadAll() ([]*Entity, error)
	Update(id string, model *UpdateDto) (*Entity, error)
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
		//collection: client.Database(mongoProps.DbName).Collection(mongoProps.Collections["item"]),
		collection: client.Database("loyalty-dlt").Collection("items"),
	}, nil
}

func (db *Db) Create(model *CreateDto) (*Entity, error) {
	res, err := db.collection.InsertOne(db.ctx, model)
	if err != nil {
		return nil, err
	}

	return &Entity{
		ID:      res.InsertedID.(primitive.ObjectID),
		Name:    model.Name,
		Company: model.Company,
		Product: model.Product,
		Point:   model.Point,
		Code:    model.Code,
	}, nil
}

func (db *Db) Read(id string) (*Entity, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	filter := bson.M{"_id": filterId}

	var itemData *Entity
	err = db.collection.FindOne(db.ctx, filter).Decode(&itemData)
	if err != nil {
		return nil, errors.New("not found")
	}

	return itemData, nil
}

func (db *Db) ReadAll() ([]*Entity, error) {
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

	items := make([]*Entity, 0)

	for cur.Next(db.ctx) {
		var result *Entity
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		items = append(items, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return items, rerr
}

func (db *Db) Update(id string, model *UpdateDto) (*Entity, error) {
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

	// TODO: Needs improvement
	if model.Point != 0 {
		updateContent = append(updateContent, bson.E{Key: "point", Value: model.Point})
	}

	update := bson.M{"$set": updateContent}
	_, err = db.collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	itemData, err := db.Read(id)

	return itemData, err
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
