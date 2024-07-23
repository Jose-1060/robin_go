package repository

import (
	"context"
	"errors"

	"github.com/Jose-1060/robin_go/internal/domain/entities"
	"github.com/Jose-1060/robin_go/internal/domain/mappers"
	"github.com/Jose-1060/robin_go/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepositoryMongo struct {
	collection *mongo.Collection
}

func NewDeviceRepositoryMongo(db *mongo.Database, collectionName string) *DeviceRepositoryMongo {
	return &DeviceRepositoryMongo{
		collection: db.Collection(collectionName),
	}
}

func (r *DeviceRepositoryMongo) GetByID(id string) (*entities.Device, error){
	var deviceModel models.DeviceModel
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&deviceModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return mappers.ToDeviceEntity(&deviceModel), err
}

func (r *DeviceRepositoryMongo) GetAll() ([]*entities.Device, error){
	var devices []*entities.Device
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var deviceModel models.DeviceModel
		if err := cursor.Decode(&deviceModel); err != nil {
			return nil, err
		}
		devices = append(devices, mappers.ToDeviceEntity(&deviceModel))
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	return devices, nil
}

func (r *DeviceRepositoryMongo) Save(device *entities.Device) error{
	deviceModel := mappers.ToDeviceModel(device)
	_, err := r.collection.InsertOne(context.Background(), deviceModel)
	return err
}

func (r *DeviceRepositoryMongo) Update(device *entities.Device) error{
	filter := bson.M{"_id": device.Imei}
	update := bson.M{"$set": mappers.ToDeviceModel(device)}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *DeviceRepositoryMongo) Delete(id string) error{
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}