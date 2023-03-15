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

func (UrlUseCase) UrlSlugGetAction(req requests.UrlSlugGetRequest) (responses.UrlSlugGetResponse, int, error) {
	slug := req.Slug

	services, httpCode, err := getAllServiceBySlug(slug)
	if err != nil {
		return responses.UrlSlugGetResponse{}, httpCode, err
	}

	rs := new([]responses.UrlSlugGetResponseServices)
	for _, service := range services {
		*rs = append(*rs, responses.UrlSlugGetResponseServices{
			Service: service.Service.String(),
			Url:     service.URL,
		})
	}

	res := responses.UrlSlugGetResponse{
		Slug:     slug,
		Services: rs,
	}

	return res, http.StatusOK, nil
}

func getAllServiceBySlug(slug string) ([]*ent.Service, int, error) {
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

	services, err := client.Service.
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

	return services, http.StatusOK, nil
}
