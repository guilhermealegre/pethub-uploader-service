package response

import "github.com/guilhermealegre/go-clean-arch-core-lib/response"

// swagger:model SwaggerSkeletonResponse
type swaggerSkeletonResponse struct { //nolint:all
	response.Response
	Data AliveResponse `json:"data"`
}

// swagger:model UploaderResponse
type UploaderResponse struct {
}

// swagger:model ImageResponse
type ImageResponse struct {
	ImagePath string `json:"imagePath"`
}
