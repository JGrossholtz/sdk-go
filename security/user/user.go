package user

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityUser struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves an User using its provided unique id.
*/
func (su SecurityUser) Fetch(id string, options types.QueryOptions) (types.User, error) {
	if id == "" {
		return types.User{}, errors.New("Security.User.Fetch: user id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getUser",
		Id:         id,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Gets the rights of an User using its provided unique id.
 */
func (su SecurityUser) GetRights(kuid string, options *types.Options) ([]types.UserRights, error) {
	if kuid == "" {
		return []types.UserRights{}, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return []types.UserRights{}, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}

/*
  Indicates whether an action is allowed, denied or conditional based on user rights provided as the first argument.
  An action is defined as a couple of action and controller (mandatory), plus an index and a collection(optional).
*/
func (su SecurityUser) IsActionAllowed(rights []types.UserRights, controller string, action string, index string, collection string) (string, error) {
	if rights == nil {
		return "", errors.New("Security.User.IsActionAllowed: Rights parameter is mandatory")
	}
	if controller == "" {
		return "", errors.New("Security.User.IsActionAllowed: Controller parameter is mandatory")
	}
	if action == "" {
		return "", errors.New("Security.User.IsActionAllowed: Action parameter is mandatory")
	}

	filteredUserRights := []types.UserRights{}

	for _, ur := range rights {
		if (ur.Controller == controller || ur.Controller == "*") && (ur.Action == action || ur.Action == "*") && (ur.Index == index || ur.Index == "*") && (ur.Collection == collection || ur.Collection == "*") {
			filteredUserRights = append(filteredUserRights, ur)
		}
	}

	for _, ur := range filteredUserRights {
		if ur.Value == "allowed" {
			return "allowed", nil
		}
		if ur.Value == "conditional" {
			return "conditional", nil
		}
	}

	return "denied", nil
}

