package main

import (
	v1AliveController "bitbucket.org/asadventure/be-uploader-service/internal/alive/controller/v1"
	v1AliveModel "bitbucket.org/asadventure/be-uploader-service/internal/alive/model/v1"
	v1ImageController "bitbucket.org/asadventure/be-uploader-service/internal/image/controller/v1"
	v1UImageModel "bitbucket.org/asadventure/be-uploader-service/internal/image/model/v1"
	v1ImageRepository "bitbucket.org/asadventure/be-uploader-service/internal/image/repository/v1"
	v1Middleware "bitbucket.org/asadventure/be-uploader-service/internal/middleware/v1"
	v1SwaggerController "bitbucket.org/asadventure/be-uploader-service/internal/swagger/controller/v1"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/aws"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/grpc"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/logger"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/meter"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/s3"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/sqs"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/tracer"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/validator"

	"fmt"

	"os"
)

func main() {
	newApp := app.New(nil)
	newLogger := logger.New(newApp, nil)
	newTracer := tracer.New(newApp, nil)
	newMeter := meter.New(newApp, nil)
	newHttp := http.New(newApp, nil)
	newAws := aws.New(newApp, nil)
	newS3 := s3.New(newApp, nil)
	newGrpc := grpc.New(newApp, nil)
	newValidator := validator.New(newApp).
		AddFieldValidators().
		AddStructValidators()
	newSQS := sqs.New(newApp, nil)

	imageRepository := v1ImageRepository.NewRepository(newApp)

	//model
	aliveModel := v1AliveModel.NewModel(newApp)
	imageModel := v1UImageModel.NewModel(newApp, imageRepository)

	newHttp.
		// middlewares
		WithMiddleware(v1Middleware.NewAuthenticateMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewPrintRequestMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewPrepareContextMiddleware(newApp)).
		// controllers
		WithController(v1SwaggerController.NewController(newApp)).
		WithController(v1AliveController.NewController(newApp, aliveModel)).
		WithController(v1ImageController.NewController(newApp, imageModel))

	// grpc
	newGrpc.
		WithController(v1ImageController.NewGrpcController(newApp, imageModel))

	newApp.
		WithLogger(newLogger).
		WithSQS(newSQS).
		WithTracer(newTracer).
		WithMeter(newMeter).
		WithValidator(newValidator).
		WithAws(newAws).
		WithS3(newS3).
		WithGrpc(newGrpc).
		WithHttp(newHttp)

	// start app
	if err := newApp.Start(); err != nil {
		fmt.Println("Error while starting: ", err)
		os.Exit(1)
	}
}
