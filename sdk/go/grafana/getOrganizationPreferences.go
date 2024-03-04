// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.GetOrganizationPreferences(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganizationPreferences(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationPreferencesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationPreferencesResult
	err := ctx.Invoke("grafana:index/getOrganizationPreferences:getOrganizationPreferences", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganizationPreferences.
type GetOrganizationPreferencesResult struct {
	// The Organization home dashboard ID.
	HomeDashboardId int `pulumi:"homeDashboardId"`
	// The Organization home dashboard UID. This is only available in Grafana 9.0+.
	HomeDashboardUid string `pulumi:"homeDashboardUid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId string `pulumi:"orgId"`
	// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
	Theme string `pulumi:"theme"`
	// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone string `pulumi:"timezone"`
	// The Organization week start.
	WeekStart string `pulumi:"weekStart"`
}

func GetOrganizationPreferencesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationPreferencesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationPreferencesResult, error) {
		r, err := GetOrganizationPreferences(ctx, opts...)
		var s GetOrganizationPreferencesResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetOrganizationPreferencesResultOutput)
}

// A collection of values returned by getOrganizationPreferences.
type GetOrganizationPreferencesResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationPreferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationPreferencesResult)(nil)).Elem()
}

func (o GetOrganizationPreferencesResultOutput) ToGetOrganizationPreferencesResultOutput() GetOrganizationPreferencesResultOutput {
	return o
}

func (o GetOrganizationPreferencesResultOutput) ToGetOrganizationPreferencesResultOutputWithContext(ctx context.Context) GetOrganizationPreferencesResultOutput {
	return o
}

func (o GetOrganizationPreferencesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetOrganizationPreferencesResult] {
	return pulumix.Output[GetOrganizationPreferencesResult]{
		OutputState: o.OutputState,
	}
}

// The Organization home dashboard ID.
func (o GetOrganizationPreferencesResultOutput) HomeDashboardId() pulumi.IntOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) int { return v.HomeDashboardId }).(pulumi.IntOutput)
}

// The Organization home dashboard UID. This is only available in Grafana 9.0+.
func (o GetOrganizationPreferencesResultOutput) HomeDashboardUid() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.HomeDashboardUid }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationPreferencesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o GetOrganizationPreferencesResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
func (o GetOrganizationPreferencesResultOutput) Theme() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.Theme }).(pulumi.StringOutput)
}

// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
func (o GetOrganizationPreferencesResultOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.Timezone }).(pulumi.StringOutput)
}

// The Organization week start.
func (o GetOrganizationPreferencesResultOutput) WeekStart() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationPreferencesResult) string { return v.WeekStart }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationPreferencesResultOutput{})
}
