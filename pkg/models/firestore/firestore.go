package firestore

import (
	"context"
	"errors"
	"funhackathon2022-backend/pkg/config"
	"funhackathon2022-backend/pkg/models"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
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

func Set(userid models.UserId, c map[string]interface{}) error {
	if fs == nil {
		return ErrFirestore
	}
	_, err := fs.Client.Collection("users").Doc(userid.UserId).Set(fs.Context, c)
	return err
}

func Get(userid models.UserId) (map[string]interface{}, error) {
	dsnap, err := fs.Client.Collection("users").Doc(userid.UserId).Get(fs.Context)
	if err != nil {
		return nil, err
	}
	c := dsnap.Data()

	return c, nil
}

func CheckAll() error {

	iter := fs.Client.Collection("users").Documents(fs.Context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		log.Printf(doc.Ref.ID)
		log.Println(doc.Data())
	}

	//iter := ref.Limit(batchSize).Documents(ctx)

	/*
		for {
		    // Get a batch of documents
		    iter := ref.Limit(batchSize).Documents(ctx)
		    numDeleted := 0

		    // Iterate through the documents, adding
		    // a delete operation for each one to a
		    // WriteBatch.
		    batch := client.Batch()
		    for {
		        doc, err := iter.Next()
		        if err == iterator.Done {
		            break
		        }
		        if err != nil {
		            return err
		        }

		        batch.Delete(doc.Ref)
		        numDeleted++
		    }

		    // If there are no documents to delete,
		    // the process is over.
		    if numDeleted == 0 {
		        return nil
		    }

		    _, err := batch.Commit(ctx)
		    if err != nil {
		        return err
		    }
		}
	*/

	return nil
}

func Close() {
	fs.Client.Close()
}
