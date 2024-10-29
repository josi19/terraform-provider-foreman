package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HanseMerkur/terraform-provider-utils/log"
)

const (
	DiscoveryRuleEndpointPrefix = "v2/discovery_rules"
)

type ForemanDiscoveryRule struct {
	// Inherits the base object's attributes
	ForemanObject
	Name                  string `json:"name"`
	Search                string `json:"search,omitempty"`
	HostGroupId           int    `json:"hostgroup_id,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	HostLimitMaxCount     int    `json:"max_count,omitempty"`
	Priority              int    `json:"priority"`
	Enabled               bool   `json:"enabled"`
	LocationIds           []int  `json:"location_ids,omitempty"`
	OrganizationIds       []int  `json:"organization_ids,omitempty"`
	DefaultLocationId     int    `json:"location_id,omitempty"`
	DefaultOrganizationId int    `json:"organization_id,omitempty"`
}

// -----------------------------------------------------------------------------
// CRUD Implementation
// -----------------------------------------------------------------------------

func (c *Client) CreateDiscoveryRule(ctx context.Context, d *ForemanDiscoveryRule) (*ForemanDiscoveryRule, error) {
	log.Tracef("foreman/api/discovery_rule.go#Create")

	reqEndpoint := fmt.Sprintf("/%s", DiscoveryRuleEndpointPrefix)

	if d.DefaultLocationId == 0 {
		d.DefaultLocationId = c.clientConfig.LocationID
	}

	if d.DefaultOrganizationId == 0 {
		d.DefaultOrganizationId = c.clientConfig.OrganizationID
	}

	dJSONBytes, jsonEncErr := c.WrapJSON("discovery_rule", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("discoveryruleJSONBytes: [%s]", dJSONBytes)

	req, reqErr := c.NewRequestWithContext(
		ctx,
		http.MethodPost,
		reqEndpoint,
		bytes.NewBuffer(dJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var createdDiscoveryRule ForemanDiscoveryRule
	sendErr := c.SendAndParse(req, &createdDiscoveryRule)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("createdDiscoveryRule: [%+v]", createdDiscoveryRule)

	return &createdDiscoveryRule, nil
}

func (c *Client) ReadDiscoveryRule(ctx context.Context, id int) (*ForemanDiscoveryRule, error) {
	log.Tracef("foreman/api/discovery_rule.go#Read")

	reqEndpoint := fmt.Sprintf("/%s/%d", DiscoveryRuleEndpointPrefix, id)

	req, reqErr := c.NewRequestWithContext(
		ctx,
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var readDiscoveryRule ForemanDiscoveryRule
	sendErr := c.SendAndParse(req, &readDiscoveryRule)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("readDiscoveryRule: [%+v]", readDiscoveryRule)

	return &readDiscoveryRule, nil
}

func (c *Client) UpdateDiscoveryRule(ctx context.Context, d *ForemanDiscoveryRule) (*ForemanDiscoveryRule, error) {
	log.Tracef("foreman/api/discovery_rule.go#Update")

	reqEndpoint := fmt.Sprintf("/%s/%d", DiscoveryRuleEndpointPrefix, d.Id)

	discoveryruleJSONBytes, jsonEncErr := c.WrapJSONWithTaxonomy("discovery_rule", d)
	if jsonEncErr != nil {
		return nil, jsonEncErr
	}

	log.Debugf("discoveryruleJSONBytes: [%s]", discoveryruleJSONBytes)

	req, reqErr := c.NewRequestWithContext(
		ctx,
		http.MethodPut,
		reqEndpoint,
		bytes.NewBuffer(discoveryruleJSONBytes),
	)
	if reqErr != nil {
		return nil, reqErr
	}

	var updatedDiscoveryRule ForemanDiscoveryRule
	sendErr := c.SendAndParse(req, &updatedDiscoveryRule)
	if sendErr != nil {
		return nil, sendErr
	}

	log.Debugf("updatedDiscoveryRule: [%+v]", updatedDiscoveryRule)

	return &updatedDiscoveryRule, nil
}

// DeleteDiscoveryRule deletes the ForemanDiscoveryRule identified by the supplied ID
func (c *Client) DeleteDiscoveryRule(ctx context.Context, id int) error {
	log.Tracef("foreman/api/discovery_rule.go#Delete")

	reqEndpoint := fmt.Sprintf("/%s/%d", DiscoveryRuleEndpointPrefix, id)

	req, reqErr := c.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return reqErr
	}

	return c.SendAndParse(req, nil)
}

// -----------------------------------------------------------------------------
// Query Implementation
// -----------------------------------------------------------------------------

func (c *Client) QueryDiscoveryRule(ctx context.Context, d *ForemanDiscoveryRule) (QueryResponse, error) {
	log.Tracef("foreman/api/discovery_rule.go#Search")

	queryResponse := QueryResponse{}

	reqEndpoint := fmt.Sprintf("/%s", DiscoveryRuleEndpointPrefix)
	req, reqErr := c.NewRequestWithContext(
		ctx,
		http.MethodGet,
		reqEndpoint,
		nil,
	)
	if reqErr != nil {
		return queryResponse, reqErr
	}

	// dynamically build the query based on the attributes
	reqQuery := req.URL.Query()
	name := `"` + d.Name + `"`
	reqQuery.Set("search", "name="+name)

	req.URL.RawQuery = reqQuery.Encode()
	sendErr := c.SendAndParse(req, &queryResponse)
	if sendErr != nil {
		return queryResponse, sendErr
	}

	log.Debugf("queryResponse: [%+v]", queryResponse)

	// Results will be Unmarshaled into a []map[string]interface{}
	//
	// Encode back to JSON, then Unmarshal into []ForemanDiscoveryRule for
	// the results
	results := []ForemanDiscoveryRule{}
	resultsBytes, jsonEncErr := json.Marshal(queryResponse.Results)
	if jsonEncErr != nil {
		return queryResponse, jsonEncErr
	}
	jsonDecErr := json.Unmarshal(resultsBytes, &results)
	if jsonDecErr != nil {
		return queryResponse, jsonDecErr
	}
	// convert the search results from []ForemanDiscoveryRule to []interface
	// and set the search results on the query
	iArr := make([]interface{}, len(results))
	for idx, val := range results {
		iArr[idx] = val
	}
	queryResponse.Results = iArr

	return queryResponse, nil
}
