package foreman

import (
	"context"

	"github.com/HanseMerkur/terraform-provider-utils/helper"
	"github.com/HanseMerkur/terraform-provider-utils/log"
	"github.com/terraform-coop/terraform-provider-foreman/foreman/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceForemanDiscoveryRule() *schema.Resource {
	r := resourceForemanDiscoveryRule()
	ds := helper.DataSourceSchemaFromResourceSchema(r.Schema)

	ds["name"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the Discovery Rule.",
	}

	return &schema.Resource{
		ReadContext: dataSourceForemanDiscoveryRuleRead,
		Schema:      ds,
	}
}

func dataSourceForemanDiscoveryRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("data_source_foreman_discovery_rule.go#Read")

	client := meta.(*api.Client)
	discovery_rule := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", discovery_rule)

	queryResponse, queryErr := client.QueryDiscoveryRule(ctx, discovery_rule)
	if queryErr != nil {
		return diag.FromErr(queryErr)
	}

	if queryResponse.Subtotal == 0 {
		return diag.Errorf("data source discovery_rule returned no results")
	}
	if queryResponse.Subtotal > 1 {
		return diag.Errorf("data source discovery_rule returned more than 1 result")
	}

	var queryDiscoveryRule api.ForemanDiscoveryRule
	var ok bool
	if queryDiscoveryRule, ok = queryResponse.Results[0].(api.ForemanDiscoveryRule); !ok {
		return diag.Errorf(
			"data source results contain unexpected type. Expected "+
				"[api.ForemanDiscoveryRule], got [%T]",
			queryResponse.Results[0],
		)
	}
	discovery_rule = &queryDiscoveryRule

	log.Debugf("ForemanDiscoveryRule: [%+v]", discovery_rule)

	setResourceDataFromForemanDiscoveryRule(d, discovery_rule)

	return nil
}
