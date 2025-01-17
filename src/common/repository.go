package common

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommonRepository struct {
	Collection *mongo.Collection
}

func NewCommonRepository(collection *mongo.Collection) *CommonRepository {
	return &CommonRepository{Collection: collection}
}

// FindAll retrieves all documents in the collection with optional filter and sorting.
func (r *CommonRepository) FindAll(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOptions) error {
	cursor, err := r.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, result)
}

// FindOne retrieves a single document by a filter.
func (r *CommonRepository) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
	return r.Collection.FindOne(ctx, filter).Decode(result)
}

// InsertOne inserts a document into the collection.
func (r *CommonRepository) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return r.Collection.InsertOne(ctx, document)
}

// InsertMany inserts multiple documents into the collection.
func (r *CommonRepository) InsertMany(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error) {
	return r.Collection.InsertMany(ctx, documents)
}
// BatchInsert inserts multiple documents into the collection.
func (r *CommonRepository) BatchInsert(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error) {
    return r.Collection.InsertMany(ctx, documents)
}

// UpdateOne updates a document by a filter.
func (r *CommonRepository) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
}

// UpdateMany updates multiple documents by a filter.
func (r *CommonRepository) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateMany(ctx, filter, bson.M{"$set": update})
}
// Upsert updates or inserts a document.
func (r *CommonRepository) Upsert(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
    opts := options.Update().SetUpsert(true)
    return r.Collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts)
}
// PushToArray pushes an element to an array field.
func (r *CommonRepository) PushToArray(ctx context.Context, filter interface{}, field string, value interface{}) (*mongo.UpdateResult, error) {
    return r.Collection.UpdateOne(ctx, filter, bson.M{"$push": bson.M{field: value}})
}
// AddToSet adds a value to an array only if it doesn't already exist.
func (r *CommonRepository) AddToSet(ctx context.Context, filter interface{}, field string, value interface{}) (*mongo.UpdateResult, error) {
    return r.Collection.UpdateOne(ctx, filter, bson.M{"$addToSet": bson.M{field: value}})
}

// DeleteOne deletes a document by a filter.
func (r *CommonRepository) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(ctx, filter)
}

// DeleteMany deletes multiple documents by a filter.
func (r *CommonRepository) DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteMany(ctx, filter)
}

// Count counts documents matching a filter.
func (r *CommonRepository) Count(ctx context.Context, filter interface{}) (int64, error) {
	return r.Collection.CountDocuments(ctx, filter)
}

// Aggregate performs an aggregation pipeline and stores the results in the provided interface.
func (r *CommonRepository) Aggregate(ctx context.Context, pipeline mongo.Pipeline, result interface{}) error {
	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, result)
}

// Paginate retrieves paginated results from the collection.
func (r *CommonRepository) Paginate(ctx context.Context, filter interface{}, result interface{}, page int64, pageSize int64, sort interface{}) (int64, error) {
	// Calculate skip and limit
	skip := (page - 1) * pageSize
	findOptions := options.Find().SetSkip(skip).SetLimit(pageSize)

	// Add sorting if provided
	if sort != nil {
		findOptions.SetSort(sort)
	}

	// Query the collection
	cursor, err := r.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	// Decode results
	if err := cursor.All(ctx, result); err != nil {
		return 0, err
	}

	// Get the total count
	totalCount, err := r.Count(ctx, filter)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

// Exists checks if a document exists matching the filter.
func (r *CommonRepository) Exists(ctx context.Context, filter interface{}) (bool, error) {
	count, err := r.Count(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ConvertID converts a string ID to a MongoDB ObjectID.
func (r *CommonRepository) ConvertID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
// Distinct retrieves distinct values for a specified field.
func (r *CommonRepository) Distinct(ctx context.Context, field string, filter interface{}) ([]interface{}, error) {
    return r.Collection.Distinct(ctx, field, filter)
}
// BulkWrite performs multiple write operations in a single batch.
func (r *CommonRepository) BulkWrite(ctx context.Context, operations []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
    return r.Collection.BulkWrite(ctx, operations, opts...)
}
// FindAndModify atomically finds and modifies a document.
func (r *CommonRepository) FindAndModify(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
    return r.Collection.FindOneAndUpdate(ctx, filter, update, opts...), nil
}

// Watch listens to changes on the collection and streams them.
func (r *CommonRepository) Watch(ctx context.Context, pipeline mongo.Pipeline, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
    return r.Collection.Watch(ctx, pipeline, opts...)
}
// IncrementField increments a numeric field atomically.
func (r *CommonRepository) IncrementField(ctx context.Context, filter interface{}, field string, incrementValue int) (*mongo.UpdateResult, error) {
    return r.Collection.UpdateOne(ctx, filter, bson.M{"$inc": bson.M{field: incrementValue}})
}
// CountDocuments counts the number of documents matching the filter.
func (r *CommonRepository) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
    return r.Collection.CountDocuments(ctx, filter)
}
// EstimatedDocumentCount provides an estimated count of documents in the collection.
func (r *CommonRepository) EstimatedDocumentCount(ctx context.Context) (int64, error) {
    return r.Collection.EstimatedDocumentCount(ctx)
}
// FindAndDelete finds and deletes a document, returning the deleted document.
func (r *CommonRepository) FindAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) (*mongo.SingleResult, error) {
    return r.Collection.FindOneAndDelete(ctx, filter, opts...), nil
}
// ReplaceDocument replaces a document with a new one.
func (r *CommonRepository) ReplaceDocument(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
    return r.Collection.ReplaceOne(ctx, filter, replacement, opts...)
}
// TextSearch performs a text search query on indexed fields.
func (r *CommonRepository) TextSearch(ctx context.Context, query string, opts ...*options.FindOptions) (*mongo.Cursor, error) {
    return r.Collection.Find(ctx, bson.M{"$text": bson.M{"$search": query}}, opts...)
}

