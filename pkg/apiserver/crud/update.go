package crud

import (
	"context"
	"errors"
	"github.com/applike/gosoline/pkg/apiserver"
	"github.com/gin-gonic/gin"
)

type updateHandler struct {
	transformer Handler
}

func NewUpdateHandler(transformer Handler) gin.HandlerFunc {
	uh := updateHandler{
		transformer: transformer,
	}

	return apiserver.CreateJsonHandler(uh)
}

func (uh updateHandler) GetInput() interface{} {
	return uh.transformer.GetUpdateInput()
}

func (uh updateHandler) Handle(ctx context.Context, request *apiserver.Request) (*apiserver.Response, error) {
	id, valid := apiserver.GetUintFromRequest(request, "id")

	if !valid {
		return nil, errors.New("no valid id provided")
	}

	repo := uh.transformer.GetRepository()
	model := uh.transformer.GetModel()
	err := repo.Read(ctx, id, model)

	if err != nil {
		return nil, err
	}

	err = uh.transformer.TransformUpdate(request.Body, model)

	if err != nil {
		return nil, err
	}

	err = repo.Update(ctx, model)

	if err != nil {
		return nil, err
	}

	reload := uh.transformer.GetModel()
	err = repo.Read(ctx, model.GetId(), reload)

	if err != nil {
		return nil, err
	}

	apiView := getApiViewFromHeader(request.Header)
	out, err := uh.transformer.TransformOutput(reload, apiView)

	if err != nil {
		return nil, err
	}

	return apiserver.NewJsonResponse(out), nil
}
