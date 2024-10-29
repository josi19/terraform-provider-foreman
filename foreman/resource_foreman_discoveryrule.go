package foreman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/HanseMerkur/terraform-provider-utils/autodoc"
	"github.com/HanseMerkur/terraform-provider-utils/conv"
	"github.com/HanseMerkur/terraform-provider-utils/log"
	"github.com/terraform-coop/terraform-provider-foreman/foreman/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceForemanDiscoveryRule() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceForemanDiscoveryRuleCreate,
		ReadContext:   resourceForemanDiscoveryRuleRead,
		UpdateContext: resourceForemanDiscoveryRuleUpdate,
		DeleteContext: resourceForemanDiscoveryRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			autodoc.MetaAttribute: {
				Type:     schema.TypeBool,
				Computed: true,
				Description: fmt.Sprintf(
					"Discovery Rules are configurations within the Foreman tool that automate " +
						"the provisioning of newly discovered hosts on your network." +
						"They specify criteria—like hardware characteristics or network details." +
						"When matched by a discovered host, trigger automatic actions such as assigning " +
						"it to a host group or initiating a specific installation process." +
						"This streamlines adding new servers by reducing manual setup." +
						autodoc.MetaSummary,
				),
			},

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(8, 256),
				Description: fmt.Sprintf(
					"DiscoveryRule name. "+
						"%s \"compute\"",
					autodoc.MetaExample,
				),
			},

			"search": {
				Type:     schema.TypeString,
				Required: true,
				//Sensitive:    false,
				ValidateFunc: validation.StringLenBetween(8, 256),
				Description:  "Search query that matches specific hosts",
			},

			"hostgroup_ids": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(0),
				//Computed:    true,
				Description: "Assing target hostgroup by ID ",
			},

			"hostname": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(8, 256),
				Description:  "Specifies the name of the new host. Can be a string or extracted via facts.",
			},

			"max_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "Sets the Host Limit, which defines, how many host can be provisioned wiht this rule. (0 = unlimited)",
			},

			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "Rule priority (lower integer means higher priority).",
			},

			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Enables or Disables the Discovery rule.",
			},

			"location_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "List of all locations the Discovery rule can be used.",
			},

			"organization_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "List of all organizations the Discovery rule can be used.",
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Conversion Helpers
// -----------------------------------------------------------------------------

// buildForemanDiscoveryRule constructs a ForemanDiscoveryRule struct from a resource
// data reference. The struct's members are populated from the data populated
// in the resource data. Missing members will be left to the zero value for
// that member's type.
func buildForemanDiscoveryRule(d *schema.ResourceData) *api.ForemanDiscoveryRule {
	log.Tracef("resource_foreman_discovery_rule.go#buildForemanDiscoveryRule")

	discovery_rule := api.ForemanDiscoveryRule{}

	obj := buildForemanObject(d)
	discovery_rule.ForemanObject = *obj

	var attr interface{}
	var ok bool

	if attr, ok = d.GetOk("name"); ok {
		discovery_rule.Name = attr.(string)
	}

	if attr, ok = d.GetOk("search"); ok {
		discovery_rule.Search = attr.(string)
	}

	if attr, ok = d.GetOk("hostgroup_ids"); ok {
		discovery_rule.HostGroupId = attr.(int)
	}

	if attr, ok = d.GetOk("hostname"); ok {
		discovery_rule.Hostname = attr.(string)
	}

	if attr, ok = d.GetOk("max_count"); ok {
		discovery_rule.HostLimitMaxCount = attr.(int)
	}

	if attr, ok = d.GetOk("priority"); ok {
		discovery_rule.Priority = attr.(int)
	}

	if attr, ok = d.GetOk("enabled"); ok {
		discovery_rule.Enabled = attr.(bool)
	}

	if attr, ok = d.GetOk("location_ids"); ok {
		attrSet := attr.(*schema.Set)
		discovery_rule.LocationIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	if attr, ok = d.GetOk("organization_ids"); ok {
		attrSet := attr.(*schema.Set)
		discovery_rule.OrganizationIds = conv.InterfaceSliceToIntSlice(attrSet.List())
	}

	return &discovery_rule
}

// setResourceDataFromForemanDiscoveryRule sets a ResourceData's attributes from
// the attributes of the supplied ForemanDiscoveryRule struct
func setResourceDataFromForemanDiscoveryRule(d *schema.ResourceData, fh *api.ForemanDiscoveryRule) {
	log.Tracef("resource_foreman_discovery_rule.go#setResourceDataFromForemanDiscoveryRule")

	d.SetId(strconv.Itoa(fh.Id))
	d.Set("name", fh.Name)
	d.Set("search", fh.Search)
	d.Set("hostgroup_ids", fh.HostGroupId)
	d.Set("hostname", fh.Hostname)
	d.Set("max_count", fh.HostLimitMaxCount)
	d.Set("priority", fh.Priority)
	d.Set("enabled", fh.Enabled)
	d.Set("location_ids", fh.LocationIds)
	d.Set("organization_ids", fh.OrganizationIds)
}

// -----------------------------------------------------------------------------
// Resource CRUD Operations
// -----------------------------------------------------------------------------

func resourceForemanDiscoveryRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("resource_foreman_discovery_rule.go#Create")

	client := meta.(*api.Client)
	h := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	createdDiscoveryRule, createErr := client.CreateDiscoveryRule(ctx, h)
	if createErr != nil {
		return diag.FromErr(createErr)
	}

	log.Debugf("Created ForemanDiscoveryRule: [%+v]", createdDiscoveryRule)

	setResourceDataFromForemanDiscoveryRule(d, createdDiscoveryRule)

	return nil
}

func resourceForemanDiscoveryRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("resource_foreman_discovery_rule.go#Read")

	client := meta.(*api.Client)
	h := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	readDiscoveryRule, readErr := client.ReadDiscoveryRule(ctx, h.Id)
	if readErr != nil {
		return diag.FromErr(api.CheckDeleted(d, readErr))
	}

	log.Debugf("Read ForemanDiscoveryRule: [%+v]", readDiscoveryRule)

	setResourceDataFromForemanDiscoveryRule(d, readDiscoveryRule)

	return nil
}

func resourceForemanDiscoveryRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("resource_foreman_discovery_rule.go#Update")

	client := meta.(*api.Client)
	h := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	updatedDiscoveryRule, updateErr := client.UpdateDiscoveryRule(ctx, h)
	if updateErr != nil {
		return diag.FromErr(updateErr)
	}

	log.Debugf("Updated ForemanDiscoveryRule: [%+v]", updatedDiscoveryRule)

	setResourceDataFromForemanDiscoveryRule(d, updatedDiscoveryRule)

	return nil
}

func resourceForemanDiscoveryRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Tracef("resource_foreman_discovery_rule.go#Delete")

	client := meta.(*api.Client)
	h := buildForemanDiscoveryRule(d)

	log.Debugf("ForemanDiscoveryRule: [%+v]", h)

	// NOTE(ALL): d.SetId("") is automatically called by terraform assuming delete
	//   returns no errors
	return diag.FromErr(api.CheckDeleted(d, client.DeleteDiscoveryRule(ctx, h.Id)))
}
