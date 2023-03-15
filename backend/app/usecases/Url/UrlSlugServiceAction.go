package usecases

import (
	"context"
	"net/http"

	requests "github.com/haisin-official/haisin/app/requests/Url"
	responses "github.com/haisin-official/haisin/app/responses/Url"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/service"
	"github.com/haisin-official/haisin/ent/user"
)

func (UrlUseCase) UrlSlugServiceGetAction(req requests.UrlSlugServiceGetRequest) (responses.UrlSlugServiceGetResponse, int, error) {
	slug := req.Slug
	name := req.Service

	service, httpCode, err := getURLBySlugAndName(slug, name)
	if err != nil {
		return responses.UrlSlugServiceGetResponse{}, httpCode, err
	}

	res := responses.UrlSlugServiceGetResponse{
		Url: service.URL,
	}

	return res, http.StatusOK, nil
}

func getURLBySlugAndName(slug string, name string) (*ent.Service, int, error) {

	client := database.GetClient()
	ctx := context.Background()

	// slugからuserIdを取得
	u, err := client.User.
		Query().
		Select(user.FieldID).
		Where(user.SlugEQ(slug)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	}

	userId := u.ID

	service, err := client.Service.
		Query().
		Select(service.FieldURL).
		Where(
			service.And(
				service.ServiceEQ(service.Service(name)),
				service.HasUserIDWith(user.IDEQ(userId)),
			),
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

	return service, http.StatusOK, nil
}
