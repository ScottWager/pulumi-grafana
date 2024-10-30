// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.LookupOrganizationPreferences(ctx, &oss.LookupOrganizationPreferencesArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: grafana.index/getorganizationpreferences.getOrganizationPreferences has been deprecated in favor of grafana.oss/getorganizationpreferences.getOrganizationPreferences
func LookupOrganizationPreferences(ctx *pulumi.Context, args *LookupOrganizationPreferencesArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationPreferencesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationPreferencesResult
	err := ctx.Invoke("grafana:index/getOrganizationPreferences:getOrganizationPreferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationPreferences.
type LookupOrganizationPreferencesArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getOrganizationPreferences.
type LookupOrganizationPreferencesResult struct {
	// The Organization home dashboard UID. This is only available in Grafana 9.0+.
	HomeDashboardUid string `pulumi:"homeDashboardUid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The Organization theme. Available values are `light`, `dark`, `system`, or an empty string for the default.
	Theme string `pulumi:"theme"`
	// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
	Timezone string `pulumi:"timezone"`
	// The Organization week start day. Available values are `sunday`, `monday`, `saturday`, or an empty string for the default.
	WeekStart string `pulumi:"weekStart"`
}

func LookupOrganizationPreferencesOutput(ctx *pulumi.Context, args LookupOrganizationPreferencesOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationPreferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationPreferencesResultOutput, error) {
			args := v.(LookupOrganizationPreferencesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupOrganizationPreferencesResult
			secret, err := ctx.InvokePackageRaw("grafana:index/getOrganizationPreferences:getOrganizationPreferences", args, &rv, "", opts...)
			if err != nil {
				return LookupOrganizationPreferencesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupOrganizationPreferencesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupOrganizationPreferencesResultOutput), nil
			}
			return output, nil
		}).(LookupOrganizationPreferencesResultOutput)
}

// A collection of arguments for invoking getOrganizationPreferences.
type LookupOrganizationPreferencesOutputArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupOrganizationPreferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationPreferencesArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationPreferences.
type LookupOrganizationPreferencesResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationPreferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationPreferencesResult)(nil)).Elem()
}

func (o LookupOrganizationPreferencesResultOutput) ToLookupOrganizationPreferencesResultOutput() LookupOrganizationPreferencesResultOutput {
	return o
}

func (o LookupOrganizationPreferencesResultOutput) ToLookupOrganizationPreferencesResultOutputWithContext(ctx context.Context) LookupOrganizationPreferencesResultOutput {
	return o
}

// The Organization home dashboard UID. This is only available in Grafana 9.0+.
func (o LookupOrganizationPreferencesResultOutput) HomeDashboardUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) string { return v.HomeDashboardUid }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrganizationPreferencesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o LookupOrganizationPreferencesResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The Organization theme. Available values are `light`, `dark`, `system`, or an empty string for the default.
func (o LookupOrganizationPreferencesResultOutput) Theme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) string { return v.Theme }).(pulumi.StringOutput)
}

// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
func (o LookupOrganizationPreferencesResultOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) string { return v.Timezone }).(pulumi.StringOutput)
}

// The Organization week start day. Available values are `sunday`, `monday`, `saturday`, or an empty string for the default.
func (o LookupOrganizationPreferencesResultOutput) WeekStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPreferencesResult) string { return v.WeekStart }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationPreferencesResultOutput{})
}
