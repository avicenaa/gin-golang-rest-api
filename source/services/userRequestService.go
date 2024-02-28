package services

import (
	"api-request/config"
	"api-request/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	client, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	dbName := "avicenDB"
	collectionName := "RequestData"

	// Connect to the collection
	collection = client.Database(dbName).Collection(collectionName)
}

func GetAllRequestData() ([]model.Form, error) {
	var forms []model.Form

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &forms)
	if err != nil {
		return nil, err
	}

	return forms, nil
}

func GetRequestByID(id string) (model.Form, error) {
	var form model.Form

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return form, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&form)
	if err != nil {
		return form, err
	}

	return form, nil
}

func AddRequestData(form model.Form) error {

	if form.Status == "" {
		form.Status = "Waiting"
	}
	_, err := collection.InsertOne(context.Background(), form)
	return err
}

func UpdateRequestData(id string, updatForm model.Form) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updatForm.ID = objID
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": updatForm},
	)
	return err
}

func DeleteRequestData(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
