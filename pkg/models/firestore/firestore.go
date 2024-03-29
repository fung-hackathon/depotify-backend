package firestore

import (
	"context"
	"errors"
	"funhackathon2022-backend/pkg/config"
	"funhackathon2022-backend/pkg/models/dto"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	ErrFirebaseInit = errors.New("failed to initializing of Firebase")
	ErrFirestore    = errors.New("failed to establish connection to Firestore")
)

type Firestore struct {
	Client  *firestore.Client
	Context context.Context
}

var fs *Firestore

func Initialize() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	conf := &firebase.Config{ProjectID: "funhackathon22"}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return ErrFirebaseInit
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
		return ErrFirestore
	}

	fs = &Firestore{Client: client, Context: ctx}

	return nil
}

func Set(userid dto.UserId, c map[string]interface{}) error {
	if fs == nil {
		return ErrFirestore
	}
	_, err := fs.Client.Collection("users").Doc(userid.UserId).Set(fs.Context, c)
	return err
}

func Update(userid dto.UserId, path string, value interface{}) error {
	if fs == nil {
		return ErrFirestore
	}
	_, err := fs.Client.Collection("users").Doc(userid.UserId).Update(fs.Context, []firestore.Update{{
		Path:  path,
		Value: value,
	}})
	return err
}

func Get(userid dto.UserId) (map[string]interface{}, error) {
	dsnap, err := fs.Client.Collection("users").Doc(userid.UserId).Get(fs.Context)
	if err != nil {
		return nil, err
	}
	c := dsnap.Data()

	return c, nil
}

func Close() {
	fs.Client.Close()
}
