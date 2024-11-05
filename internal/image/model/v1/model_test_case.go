package v1

import (
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/request"
	"github.com/guilhermealegre/go-clean-arch-core-lib/test"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"

	v1Domain "bitbucket.org/asadventure/be-uploader-service/internal/image/domain/v1"
	"github.com/stretchr/testify/mock"
)

type TestCase struct {
	test.BaseTestCase
	Repository test.MapCall
}

func testCaseImageModel() *TestCase {
	// expected data
	expected := &v1Domain.Upload{
		ImagePath: "/user-upload/file-mock.png",
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Upload File",
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}, &request.ImageUploadRequest{
					Section:     "test-upload",
					Image:       []byte{},
					FileName:    "file-mock.png",
					ContentType: "image/png",
				}},
				Expected: []interface{}{expected, nil},
			},
		},
		Repository: test.MapCall{
			"Upload": test.CallList{
				test.Call{
					Arguments: []interface{}{mock.IsType(&context.Context{}), &request.ImageUploadRequest{
						Section:     "test-upload",
						Image:       []byte{},
						FileName:    "file-mock.png",
						ContentType: "image/png",
					},
					},
					Expected: []interface{}{expected, nil},
				},
			},
		},
	}
}
