package repository

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/projects/mongodb/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client *mongo.Client
}

//NewMongoRepository ...
func NewMongoRepository(mongoURL string) (product.ProductRepository, error) {

	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//conect client

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {

		return nil, errors.Wrap(err, "repository.NewMongoRepo")

	}

	err = c.Ping(ctx, readpref.Primary())

	if err != nil {

		log.Fatal(err)
	}

	repo := &mongoRepository{}
	repo.client = c
	return repo, nil

}

func (r *mongoRepository) Find(code string) (*product.Product, error) {

	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	prod := &product.Product{}
	collection := r.client.Database("product").Collection("items")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&prod)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("repository.Product.Find")
		}
		return nil, errors.Wrap(err, "repository.Product.Find")
	}
	return prod, nil

}
func (r *mongoRepository) Store(product *product.Product) error {
	timeout := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	collection := r.client.Database("product").Collection("items")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":  product.Code,
			"name":  product.Name,
			"price": product.Price,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Product.Store")
	}
	return nil

}

func (r *mongoRepository) FindAll() ([]*product.Product, error) {

	var items []*product.Product

	collection := r.client.Database("product").Collection("items")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item product.Product
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil

}
func (r *mongoRepository) Delete(code string) error {

	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	filter := bson.M{"code": code}

	coll := r.client.Database("product").Collection("items")
	_, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
