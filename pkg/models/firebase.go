package models

import (
	"context"
	"errors"
	"funhackathon2022-backend/pkg/config"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	ErrFirebaseInit = errors.New("failed to initializing of Firebase")
	ErrFirestore    = errors.New("failed to establish connection to Firestore")
)

// defer (*firestore.Client).Close()を忘れずに
func NewFirestore() (*firestore.Client, *context.Context, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	conf := &firebase.Config{ProjectID: "funhackathon22"}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, nil, ErrFirebaseInit
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
		return nil, nil, ErrFirestore
	}
	return client, &ctx, nil
}
