package v1

import (
	"testing"

	domainImage "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	"github.com/stretchr/testify/assert"

	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"

	v1Repo "bitbucket.org/asadventure/be-uploader-service/internal/image/repository/v1"
)

type TModel struct {
	model      domainImage.IModel
	repository *v1Repo.RepositoryMock
}

func newModel() *TModel {
	repository := v1Repo.NewRepositoryMock()

	return &TModel{
		model:      NewModel(app.NewAppMock(), repository),
		repository: repository,
	}
}

func TestImageModel(t *testing.T) {
	testCases := []*TestCase{
		testCaseImageModel(),
	}

	for _, test := range testCases {
		test.Log(t)
		m := newModel()

		// setup
		test.Repository.Setup(m.repository)

		// model
		result, err := m.model.Upload(
			test.Arguments[0].(*context.Context),
			test.Arguments[1].(*request.ImageUploadRequest),
		)

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}
