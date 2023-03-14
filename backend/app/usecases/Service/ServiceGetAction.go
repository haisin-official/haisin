package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Service"
	responses "github.com/haisin-official/haisin/app/responses/Service"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/service"
	"github.com/haisin-official/haisin/ent/user"
)

func (ServiceUseCase) ServiceGetAction(req requests.ServiceGetRequest) (responses.ServiceGetResponse, int, error) {
	userId := req.UserID

	services, httpCode, err := getService(userId)
	if err != nil {
		return responses.ServiceGetResponse{}, httpCode, err
	}

	rs := new([]responses.ServiceGetResponseServices)
	for _, service := range services {
		*rs = append(*rs, responses.ServiceGetResponseServices{
			Service: service.Service.String(),
			Url:     service.URL,
		})
	}

	user, httpCode, err := getSlug(userId)
	if err != nil {
		return responses.ServiceGetResponse{}, httpCode, err
	}

	res := responses.ServiceGetResponse{
		Slug:     user.Slug,
		Services: rs,
	}
	return res, http.StatusOK, nil
}

func getService(userId uuid.UUID) ([]*ent.Service, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	service, err := client.Service.
		Query().
		Select(service.FieldService, service.FieldURL).
		Where(
			service.HasUserIDWith(user.IDEQ(userId)),
		).All(ctx)

	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	} else if ent.IsValidationError(err) {
		return nil, http.StatusUnprocessableEntity, err
	} else if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return service, http.StatusOK, nil
}

func getSlug(userId uuid.UUID) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select(user.FieldSlug).
		Where(
			user.IDEQ(userId),
		).Only(ctx)

	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	} else if ent.IsValidationError(err) {
		return nil, http.StatusUnprocessableEntity, err
	} else if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
