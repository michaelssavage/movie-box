package repository

import (
	"context"
	"fmt"
	"time"

	"boxd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveFavourites(ctx context.Context, client *mongo.Client, movies []models.Movie) error {
	collection := client.Database("letterboxd").Collection("favorites")
	
	favoriteMovies := models.FavoriteMovies{
		ID:          "latest",
		Movies:      movies,
		LastUpdated: time.Now(),
	}

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: favoriteMovies.ID}}
	
	_, err := collection.ReplaceOne(ctx, filter, favoriteMovies, opts)
	return err
}

func GetFavourites(ctx context.Context, client *mongo.Client) ([]models.Movie, error) {
	collection := client.Database("letterboxd").Collection("favorites")

	var favoriteMovies models.FavoriteMovies

	err := collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: "latest"}}).Decode(&favoriteMovies)
	if err != nil {
			return nil, fmt.Errorf("failed to find favorites: %v", err)
	}


	return favoriteMovies.Movies, nil
}