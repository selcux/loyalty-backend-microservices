package product

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
	Create(model *CreateDto) (*Product, error)
	Read(id string) (*Product, error)
	ReadAll() ([]*Product, error)
	Update(id string, model *UpdateDto) (*Product, error)
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
		collection: client.Database(mongoProps.DbName).Collection(mongoProps.Collections["product"]),
	}, nil
}

func (db *Db) Create(model *CreateDto) (*Product, error) {
	companyID, err := primitive.ObjectIDFromHex(model.Company)
	if err != nil {
		return nil, err
	}

	newProduct := &CreateProduct{
		Company: companyID,
		Name:    model.Name,
		Point:   model.Point,
		Code:    model.Code,
	}

	res, err := db.collection.InsertOne(db.ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:      res.InsertedID.(primitive.ObjectID),
		Name:    newProduct.Name,
		Company: newProduct.Company,
		Point:   newProduct.Point,
		Code:    newProduct.Code,
	}, nil
}

func (db *Db) Read(id string) (*Product, error) {
	filterId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	filter := bson.M{"_id": filterId}

	productData := new(Product)
	err = db.collection.FindOne(db.ctx, filter).Decode(&productData)
	if err != nil {
		return nil, errors.New("not found")
	}

	return productData, nil
}

func (db *Db) ReadAll() ([]Product, error) {
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

	products := make([]Product, 0)

	for cur.Next(db.ctx) {
		var result Product
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		products = append(products, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return products, rerr
}

func (db *Db) Update(id string, model *UpdateDto) (*Product, error) {
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
	if model.Code != "" {
		updateContent = append(updateContent, bson.E{Key: "code", Value: model.Code})
	}

	update := bson.M{"$set": updateContent}
	_, err = db.collection.UpdateOne(db.ctx, filter, update)
	if err != nil {
		return nil, err
	}

	productData, err := db.Read(id)
	if err != nil {
		return nil, err
	}

	return productData, nil
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
