/*
	 BeUploaderService Service

	 # BeUploaderService Service API

	 Schemes: http, https
	 BasePath: /v1
	 Version: 1.0

	 Consumes:
	 - application/json

	 Produces:
	 - application/json

	 SecurityDefinitions:
		Bearer:
		  type: apiKey
		  name: Authorization
		  in: header

	 swagger:meta
*/
package swagger

import (
	_ "bitbucket.org/asadventure/be-uploader-service/internal/alive/controller/v1"   // alive controller
	_ "bitbucket.org/asadventure/be-uploader-service/internal/image/controller/v1"   // uploader controller
	_ "bitbucket.org/asadventure/be-uploader-service/internal/swagger/controller/v1" // swagger controller
)
