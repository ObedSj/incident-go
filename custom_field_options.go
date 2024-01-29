package incident

import (
	"context"
	"fmt"
)

// CustomFieldOptionsService handles communication with the custom field options related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/tag/Custom-Field-Options-V1
type CustomFieldOptionsService service

// List list all custom field options for a custom field.
//
// customFieldID represents the unique identifier for the custom field
// pageSize represents the number of items to return
// after represents the cursor for pagination
//
// API docs: https://api-docs.incident.io/tag/Custom-Field-Options-V1#operation/Custom%20Field%20Options%20V1_List
func (s *CustomFieldOptionsService) List(ctx context.Context, customFieldID string, pageSize int, after string) (*CustomFieldsList, *Response, error) {
	u := fmt.Sprintf("v1/custom_field_options?custom_field_id=%s&page_size=%d&after=%s", customFieldID, pageSize, after)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &CustomFieldsList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Get returns a single custom field.
//
// id represents the unique identifier for the custom field
//
// API docs: https://api-docs.incident.io/tag/Custom-Field-Options-V1#operation/Custom%20Field%20Options%20V1_Show
func (s *CustomFieldOptionsService) Get(ctx context.Context, id string) (*CustomFieldResponse, *Response, error) {
	u := fmt.Sprintf("v2/custom_field_options/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &CustomFieldResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
