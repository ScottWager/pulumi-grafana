// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
//   - [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
func LookupServiceAccount(ctx *pulumi.Context, args *LookupServiceAccountArgs, opts ...pulumi.InvokeOption) (*LookupServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceAccountResult
	err := ctx.Invoke("grafana:index/getServiceAccount:getServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type LookupServiceAccountArgs struct {
	// The name of the Service Account.
	Name string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getServiceAccount.
type LookupServiceAccountResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The disabled status for the service account.
	IsDisabled bool `pulumi:"isDisabled"`
	// The name of the Service Account.
	Name string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The basic role of the service account in the organization.
	Role string `pulumi:"role"`
}

func LookupServiceAccountOutput(ctx *pulumi.Context, args LookupServiceAccountOutputArgs, opts ...pulumi.InvokeOption) LookupServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceAccountResult, error) {
			args := v.(LookupServiceAccountArgs)
			r, err := LookupServiceAccount(ctx, &args, opts...)
			var s LookupServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceAccountResultOutput)
}

// A collection of arguments for invoking getServiceAccount.
type LookupServiceAccountOutputArgs struct {
	// The name of the Service Account.
	Name pulumi.StringInput `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getServiceAccount.
type LookupServiceAccountResultOutput struct{ *pulumi.OutputState }

func (LookupServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountResult)(nil)).Elem()
}

func (o LookupServiceAccountResultOutput) ToLookupServiceAccountResultOutput() LookupServiceAccountResultOutput {
	return o
}

func (o LookupServiceAccountResultOutput) ToLookupServiceAccountResultOutputWithContext(ctx context.Context) LookupServiceAccountResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The disabled status for the service account.
func (o LookupServiceAccountResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

// The name of the Service Account.
func (o LookupServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o LookupServiceAccountResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The basic role of the service account in the organization.
func (o LookupServiceAccountResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceAccountResultOutput{})
}
