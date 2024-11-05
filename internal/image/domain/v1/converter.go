package v1

import (
	uploader "bitbucket.org/asadventure/be-uploader-service/api/v1/grpc/uploader_service_uploader"
	"bitbucket.org/asadventure/be-uploader-service/api/v1/http/envelope/response"
)

func (a *Upload) FromDomainToApi() *response.ImageResponse {
	if a == nil {
		return nil
	}

	return &response.ImageResponse{
		ImagePath: a.ImagePath,
	}
}

func (a *Upload) FromDomainToGrpc() *uploader.ImageUploadGrpcResponse {
	if a == nil {
		return nil
	}

	return &uploader.ImageUploadGrpcResponse{
		ImagePath: a.ImagePath,
	}
}
