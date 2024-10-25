package foreman

import (
	"context"
	"fmt"

	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/helper"
	"github.com/HanseMerkur/terraform-provider-utils/log"
	"github.com/terraform-coop/terraform-provider-foreman/foreman/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceForemanDiscoveryRule() *schema.Resource {
	// copy attributes from resource definition
	r := resourceForemanDiscoveryRule()
	ds := helper.DataSourceSchemaFromResourceSchema(r.Schema)

	// define searchable attributes for the data source

	ds["title"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		Description: fmt.Sprintf(
			"The title is the fullname of the discoveryrule.  A "+
				"discoveryrule's title is a path-like string from the head "+
				"of the discoveryrule tree down to this discoveryrule.  The title will be "+
				"in the form of: \"<parent 1>/<parent 2>/.../<name>\". "+
				"%s \"BO1/VM/DEVP4\"",
			autodoc.MetaExample,
		),
	}

	return &schema.Resource{

		ReadContext: dataSourceForemanDiscoveryRuleRead,

		// NOTE(ALL): See comments in the corresponding resource file
		Schema: ds,
	}
}

func dataSourceForemanDiscoveryRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("data_source_foreman_discoveryrule.go#Read")

	client := meta.(*api.Client)
	h := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	queryResponse, queryErr := client.QueryDiscoveryRule(ctx, h)
	if queryErr != nil {
		return diag.FromErr(queryErr)
	}

	if queryResponse.Subtotal == 0 {
		return diag.Errorf("Data source discoveryrule returned no results")
	} else if queryResponse.Subtotal > 1 {
		return diag.Errorf("Data source discoveryrule returned more than 1 result")
	}

	var queryDiscoveryRule api.ForemanDiscoveryRule
	var ok bool
	if queryDiscoveryRule, ok = queryResponse.Results[0].(api.ForemanDiscoveryRule); !ok {
		return diag.Errorf(
			"Data source results contain unexpected type. Expected "+
				"[api.ForemanDiscoveryRule], got [%T]",
			queryResponse.Results[0],
		)
	}
	h = &queryDiscoveryRule

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	setResourceDataFromForemanDiscoveryRule(d, h)

	return nil
}
