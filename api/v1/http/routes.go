package http

import (
	"net/http"

	infra "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
)

var (
	GroupV1                       = infra.NewGroup("v1")
	GroupV1P                      = GroupV1.Group("p")
	Alive                         = GroupV1.NewEndpoint("/alive", http.MethodGet)
	GroupV1Uploader               = GroupV1.Group("uploader")
	GroupV1PUploader              = GroupV1P.Group("uploader")
	GroupV1PDocumentation         = GroupV1P.Group("documentation")
	GroupV1PDocumentationUploader = GroupV1PDocumentation.Group("uploader")

	PublicAliveUploader = GroupV1P.NewEndpoint("/alive/uploader", http.MethodGet)

	SwaggerUploaderDocs    = GroupV1PDocumentationUploader.NewEndpoint("/docs", http.MethodGet)
	SwaggerUploaderSwagger = GroupV1PDocumentationUploader.NewEndpoint("/swagger", http.MethodGet)

	// Image Routes
	ImageUpload  = GroupV1Uploader.NewEndpoint("/image/upload", http.MethodPost)
	ImagePUpload = GroupV1PUploader.NewEndpoint("/image/upload", http.MethodPost)
)
