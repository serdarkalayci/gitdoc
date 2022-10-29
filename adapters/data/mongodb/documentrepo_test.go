package mongodb

import (
	"context"
	"errors"
	"testing"

	"github.com/serdarkalayci/gitdoc/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/gitdoc/domain"
	"github.com/stretchr/testify/assert"
)

// // MockHTTPClient is the client that mocks original http.Client
type MockMongoHelper struct {
}

var (
	// GetDeleteFunc will be used to get different Delete functions for testing purposes
	GetDeleteFunc func(ctx context.Context, id string) (int, error)
	// GetDeleteFunc will be used to get different Update functions for testing purposes
	GetUpdateFunc func(ctx context.Context, id string, update interface{}) (int, error)
	// GetFindOneFunc will be used to get different FindOne functions for testing purposes
	GetFindOneFunc func(ctx context.Context, id string) (dao.DocumentDAO, error)
	// GetInsertOneFunc will be used to get different InsertOne functions for testing purposes
	GetInsertOneFunc func(ctx context.Context, document interface{}) (string, error)
	// GetListFunc will be used to get different List functions for testing purposes
	GetListFunc func(ctx context.Context) ([]dao.DocumentDAO, error)
)

// func (client MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
// 	return GetDoFunc(req)
// }

func (mh MockMongoHelper) Find(ctx context.Context) ([]dao.DocumentDAO, error) {
	return GetListFunc(ctx)
}
func (mh MockMongoHelper) InsertOne(ctx context.Context, document interface{}) (string, error) {
	return GetInsertOneFunc(ctx, document)
}
func (mh MockMongoHelper) FindOne(ctx context.Context, id string) (dao.DocumentDAO, error) {
	return GetFindOneFunc(ctx, id)
}
func (mh MockMongoHelper) UpdateOne(ctx context.Context, id string, update interface{}) (int, error) {
	return GetUpdateFunc(ctx, id, update)
}
func (mh MockMongoHelper) DeleteOne(ctx context.Context, id string) (int, error) {
	return GetDeleteFunc(ctx, id)
}

func TestDocumentRepository_Delete_Error(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetDeleteFunc = func(ctx context.Context, id string) (int, error) {
		return 0, errors.New("Whatever error")
	}
	err := pr.Delete("id")
	assert.EqualError(t, err, "Error deleting the document")
}

func TestDocumentRepository_Delete_ResultNotOne(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetDeleteFunc = func(ctx context.Context, id string) (int, error) {
		return 0, nil
	}
	err := pr.Delete("this_id")
	assert.EqualError(t, err, "Cannot find the document with the ID this_id")
}

func TestDocumentRepository_Delete_ResultSuccess(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetDeleteFunc = func(ctx context.Context, id string) (int, error) {
		return 1, nil
	}
	err := pr.Delete("this_id")
	assert.Nil(t, err)
}

func TestDocumentRepository_Update_Error(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetUpdateFunc = func(ctx context.Context, id string, update interface{}) (int, error) {
		return 0, errors.New("Whatever error")
	}
	err := pr.Update("id", domain.Document{})
	assert.EqualError(t, err, "Error updating the document")
}

func TestDocumentRepository_Update_ResultNotOne(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetUpdateFunc = func(ctx context.Context, id string, update interface{}) (int, error) {
		return 0, nil
	}
	err := pr.Update("this_id", domain.Document{})
	assert.EqualError(t, err, "Cannot find the document with the ID this_id")
}

func TestDocumentRepository_Update_ResultSuccess(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetUpdateFunc = func(ctx context.Context, id string, update interface{}) (int, error) {
		return 1, nil
	}
	err := pr.Update("id", domain.Document{})
	assert.Nil(t, err)
}

func TestDocumentRepository_FindOne_Error(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetFindOneFunc = func(ctx context.Context, id string) (dao.DocumentDAO, error) {
		return dao.DocumentDAO{}, errors.New("Cannot find the document with the ID this_id")
	}
	pDAO, err := pr.Get("this_id")
	assert.Equal(t, pDAO, domain.Document{})
	assert.EqualError(t, err, "Cannot find the document with the ID this_id")
}

func TestDocumentRepository_FindOne_Success(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetFindOneFunc = func(ctx context.Context, id string) (dao.DocumentDAO, error) {
		return dao.DocumentDAO{
			ID:                      "id",
			Name:                    "name",
			documentClass:           1,
			Sex:                     "male",
			Age:                     34,
			Survived:                false,
			SiblingsOrSpousesAboard: 3,
			ParentsOrChildrenAboard: 2,
			Fare:                    4.97,
		}, nil
	}
	pDAO, err := pr.Get("this_id")
	assert.Equal(t, pDAO, domain.Document{
		ID:                      "id",
		Name:                    "name",
		documentClass:           1,
		Sex:                     "male",
		Age:                     34,
		Survived:                false,
		SiblingsOrSpousesAboard: 3,
		ParentsOrChildrenAboard: 2,
		Fare:                    4.97,
	})
	assert.Nil(t, err)
}

func TestDocumentRepository_InsertOne_Error(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetInsertOneFunc = func(ctx context.Context, document interface{}) (string, error) {
		return "", errors.New("Whatever error")
	}
	document, err := pr.Add(domain.Document{
		ID: "this_id",
	})
	assert.Equal(t, document, domain.Document{})
	assert.EqualError(t, err, "Cannot insert the document")
}

func TestDocumentRepository_InsertOne_Success(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetInsertOneFunc = func(ctx context.Context, document interface{}) (string, error) {
		return "new_id", nil
	}
	document, err := pr.Add(domain.Document{
		ID: "this_id",
	})
	assert.Equal(t, document, domain.Document{
		ID: "this_id",
	})
	assert.Nil(t, err)
}

func TestDocumentRepository_List_Error(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	GetListFunc = func(ctx context.Context) ([]dao.DocumentDAO, error) {
		return nil, errors.New("Whatever error")
	}
	result, err := pr.List()
	assert.Nil(t, result)
	assert.EqualError(t, err, "Error getting documents")
}

func TestDocumentRepository_List_Success(t *testing.T) {
	pr := DocumentRepository{MockMongoHelper{}}
	pDAOs := []dao.DocumentDAO{
		dao.DocumentDAO{
			ID:                      "id1",
			Name:                    "name1",
			documentClass:           1,
			Sex:                     "male",
			Age:                     14,
			Survived:                false,
			SiblingsOrSpousesAboard: 1,
			ParentsOrChildrenAboard: 2,
			Fare:                    1.97,
		},
		dao.DocumentDAO{
			ID:                      "id2",
			Name:                    "name2",
			documentClass:           2,
			Sex:                     "male",
			Age:                     24,
			Survived:                false,
			SiblingsOrSpousesAboard: 2,
			ParentsOrChildrenAboard: 3,
			Fare:                    2.97,
		},
	}

	GetListFunc = func(ctx context.Context) ([]dao.DocumentDAO, error) {
		return pDAOs, nil
	}
	result, err := pr.List()
	assert.Nil(t, err)
	assert.Equal(t, result, []domain.Document{
		domain.Document{
			ID:                      "id1",
			Name:                    "name1",
			documentClass:           1,
			Sex:                     "male",
			Age:                     14,
			Survived:                false,
			SiblingsOrSpousesAboard: 1,
			ParentsOrChildrenAboard: 2,
			Fare:                    1.97,
		},
		domain.Document{
			ID:                      "id2",
			Name:                    "name2",
			documentClass:           2,
			Sex:                     "male",
			Age:                     24,
			Survived:                false,
			SiblingsOrSpousesAboard: 2,
			ParentsOrChildrenAboard: 3,
			Fare:                    2.97,
		},
	})
}
