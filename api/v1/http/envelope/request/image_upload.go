package request

import (
	uploader "bitbucket.org/asadventure/be-uploader-service/api/v1/grpc/uploader_service_uploader"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

const (
	FormRequestSection    = "section"
	FormRequestImage      = "image"
	HeaderFileContentType = "Content-Type"
)

//swagger:parameters ImageUploadRequest
type ImageUploadRequest struct {
	// Section
	// in: form
	// required: false
	Section string `form:"section"`
	// Image
	Image []byte `form:"image"`
	// FileName
	FileName string `form:"-" json:"fileName"`
	// FileName
	ContentType string `form:"-" json:"contentType"`
}

//swagger:parameters PublicImageUploadRequest
type PublicImageUploadRequest ImageUploadRequest

func (i *ImageUploadRequest) LoadFromGrpc(grpcRequest *uploader.ImageUploadGrpcRequest) {
	i.Image = grpcRequest.Image
	i.Section = grpcRequest.Section
	i.FileName = grpcRequest.Filename
	i.ContentType = grpcRequest.ContentType
}

func (i *ImageUploadRequest) LoadFromApi(gCtx dCtx.IContext) error {
	// Bind Parameters
	i.Section = gCtx.PostForm(FormRequestSection)
	// get image multipart header
	imgFileHeader, err := gCtx.FormFile(FormRequestImage)
	if err != nil {
		return err
	}
	// set file name and Content Type
	i.FileName = imgFileHeader.Filename
	i.ContentType = imgFileHeader.Header.Get(HeaderFileContentType)
	// open Image
	file, err := imgFileHeader.Open()
	if err != nil {
		return err
	}
	i.Image = make([]byte, imgFileHeader.Size)
	// read bytes to the request object
	_, err = file.Read(i.Image)
	if err != nil {
		return err
	}
	return nil
}
