package consumer

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
	AddToWallet(id string, itemId string) error
	RemoveFromWallet(id string, itemId string) error
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
		//collection: client.Database(mongoProps.DbName).Collection(mongoProps.Collections["consumer"]),
		collection: client.Database("loyalty-dlt").Collection("consumers"),
	}, nil
}

func (db *Db) Create(model *CreateDto) (*Entity, error) {
	res, err := db.collection.InsertOne(db.ctx, model)
	if err != nil {
		return nil, err
	}

	return &Entity{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     model.Name,
		Lastname: model.Lastname,
		Email:    model.Email,
	}, nil
}

func (db *Db) Read(id string) (*Entity, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	filter := bson.M{"_id": filterId}

	var consumerData *Entity
	err = db.collection.FindOne(db.ctx, filter).Decode(&consumerData)
	if err != nil {
		return nil, errors.New("not found")
	}

	return consumerData, nil
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

	consumers := make([]*Entity, 0)

	for cur.Next(db.ctx) {
		var result *Entity
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		consumers = append(consumers, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return consumers, rerr
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
	if model.Lastname != "" {
		updateContent = append(updateContent, bson.E{Key: "lastname", Value: model.Lastname})
	}
	if model.Email != "" {
		updateContent = append(updateContent, bson.E{Key: "email", Value: model.Email})
	}

	update := bson.M{"$set": updateContent}
	_, err = db.collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	consumerData, err := db.Read(id)
	if err != nil {
		return nil, err
	}

	return consumerData, nil
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

func (db *Db) AddToWallet(id string, itemId string) error {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	filter := bson.M{"_id": filterId}
	incrementKey := fmt.Sprintf("wallet.%s", itemId)
	addToWallet := bson.M{"$inc": bson.M{incrementKey: 1}}

	_, err = db.collection.UpdateOne(db.ctx, filter, addToWallet)
	return err
}

func (db *Db) RemoveFromWallet(id string, itemId string, val int) error {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	decrementKey := fmt.Sprintf("wallet.%s", itemId)
	filter := bson.M{"_id": filterId, decrementKey: bson.M{"$gt": 0}}
	// TODO: Should return error(not exist) if 0 - security vulnerability
	removeFromWallet := bson.M{"$inc": bson.M{decrementKey: -val}}

	_, err = db.collection.UpdateOne(db.ctx, filter, removeFromWallet)
	return err
}
