package foreman

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-coop/terraform-provider-foreman/foreman/api"
)

func Test_resourceForemanDiscoveryRule(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceForemanDiscoveryRule(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceForemanDiscoveryRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildForemanDiscoveryRule(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name string
		args args
		want *api.ForemanDiscoveryRule
	}{
		{
			name: "Basic Test",
			args: args{d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{
				"name":             "test-rule",
				"search":           "name ~ test",
				"hostgroup_ids":    1,
				"hostname":         "test-host",
				"max_count":        10,
				"priority":         1,
				"enabled":          true,
				"location_ids":     []interface{}{1, 2},
				"organization_ids": []interface{}{1, 2},
			})},
			want: &api.ForemanDiscoveryRule{
				Name:              "test-rule",
				Search:            "name ~ test",
				HostGroupId:       1,
				Hostname:          "test-host",
				HostLimitMaxCount: 10,
				Priority:          1,
				Enabled:           true,
				LocationIds:       []int{1, 2},
				OrganizationIds:   []int{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildForemanDiscoveryRule(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildForemanDiscoveryRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setResourceDataFromForemanDiscoveryRule(t *testing.T) {
	type args struct {
		d  *schema.ResourceData
		fh *api.ForemanDiscoveryRule
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic Test",
			args: args{
				d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{}),
				fh: &api.ForemanDiscoveryRule{
					Name:              "test-rule",
					Search:            "name ~ test",
					HostGroupId:       1,
					Hostname:          "test-host",
					HostLimitMaxCount: 10,
					Priority:          1,
					Enabled:           true,
					LocationIds:       []int{1, 2},
					OrganizationIds:   []int{1, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setResourceDataFromForemanDiscoveryRule(tt.args.d, tt.args.fh)
			if got := tt.args.d.Get("name"); got != tt.args.fh.Name {
				t.Errorf("name = %v, want %v", got, tt.args.fh.Name)
			}
		})
	}
}

func Test_resourceForemanDiscoveryRuleCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		{
			name: "Basic Test",
			args: args{
				ctx: context.TODO(),
				d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{
					"name":             "test-rule",
					"search":           "name ~ test",
					"hostgroup_ids":    1,
					"hostname":         "test-host",
					"max_count":        10,
					"priority":         1,
					"enabled":          true,
					"location_ids":     []int{1, 2},
					"organization_ids": []int{1, 2},
				}),
				meta: &api.Client{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceForemanDiscoveryRuleCreate(tt.args.ctx, tt.args.d, tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceForemanDiscoveryRuleCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceForemanDiscoveryRuleRead(t *testing.T) {
	type args struct {
		ctx  context.Context
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		{
			name: "Basic Test",
			args: args{
				ctx: context.TODO(),
				d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{
					"name":             "test-rule",
					"search":           "name ~ test",
					"hostgroup_ids":    1,
					"hostname":         "test-host",
					"max_count":        10,
					"priority":         1,
					"enabled":          true,
					"location_ids":     []int{1, 2},
					"organization_ids": []int{1, 2},
				}),
				meta: &api.Client{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceForemanDiscoveryRuleRead(tt.args.ctx, tt.args.d, tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceForemanDiscoveryRuleRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceForemanDiscoveryRuleUpdate(t *testing.T) {
	type args struct {
		ctx  context.Context
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		{
			name: "Basic Test",
			args: args{
				ctx: context.TODO(),
				d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{
					"name":             "test-rule",
					"search":           "name ~ test",
					"hostgroup_ids":    1,
					"hostname":         "test-host",
					"max_count":        10,
					"priority":         1,
					"enabled":          true,
					"location_ids":     []int{1, 2},
					"organization_ids": []int{1, 2},
				}),
				meta: &api.Client{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceForemanDiscoveryRuleUpdate(tt.args.ctx, tt.args.d, tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceForemanDiscoveryRuleUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceForemanDiscoveryRuleDelete(t *testing.T) {
	type args struct {
		ctx  context.Context
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		{
			name: "Basic Test",
			args: args{
				ctx: context.TODO(),
				d: schema.TestResourceDataRaw(t, resourceForemanDiscoveryRule().Schema, map[string]interface{}{
					"name":             "test-rule",
					"search":           "name ~ test",
					"hostgroup_ids":    1,
					"hostname":         "test-host",
					"max_count":        10,
					"priority":         1,
					"enabled":          true,
					"location_ids":     []int{1, 2},
					"organization_ids": []int{1, 2},
				}),
				meta: &api.Client{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceForemanDiscoveryRuleDelete(tt.args.ctx, tt.args.d, tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceForemanDiscoveryRuleDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
