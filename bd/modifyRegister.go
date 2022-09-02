package bd

import (
	"context"
	"time"

	"github.com/jonathanludena/tgotter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Function update profile user in DB */
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tgotter")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Lastname) > 0 {
		register["lastname"] = u.Lastname
	}
	register["birthday"] = u.Birthday
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.Siteweb) > 0 {
		register["siteweb"] = u.Siteweb
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
