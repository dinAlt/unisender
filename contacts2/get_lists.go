package contacts2

import "github.com/alexeyco/unisender/api"

// GetListsResponse response of createList request.
type GetListsResponse []List

// GetListsRequest request to get the list of all available campaign lists.
type GetListsRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetListsRequest) Execute() (lists []List, err error) {
	err = r.request.Execute("getLists", &lists)
	return
}

// GetLists returns request to get the list of all available campaign lists.
//
// See https://www.unisender.com/en/support/api/partners/getlists/
func GetLists(request *api.Request) *GetListsRequest {
	return &GetListsRequest{
		request: request,
	}
}