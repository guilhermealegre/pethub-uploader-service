package v1

import (
	"os"

	v1 "bitbucket.org/asadventure/be-uploader-service/internal/alive/domain/v1"
	"github.com/guilhermealegre/go-clean-arch-core-lib/test"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
)

type TestCase struct {
	test.BaseTestCase
}

func testCaseAlive() *TestCase {
	hostName, _ := os.Hostname()

	aliveResponse := &v1.Alive{
		ServerName: "uploader",
		Port:       "80",
		Hostname:   hostName,
		Message:    "I AM ALIVE!!!",
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting alive",
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}},
				Expected:  []interface{}{aliveResponse, nil},
			},
		},
	}
}

func testCasePublicAlive() *TestCase {
	aliveResponse := &v1.PublicAlive{
		Name:    "uploader",
		Message: "I AM ALIVE!!!",
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting alive",
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}},
				Expected:  []interface{}{aliveResponse, nil},
			},
		},
	}
}
