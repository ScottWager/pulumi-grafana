// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/cloud"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloud.GetOrganization(ctx, &cloud.GetOrganizationArgs{
//				Slug: pulumi.StringRef("my-org"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationResult
	err := ctx.Invoke("grafana:cloud/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	// The ID of this resource.
	Id   *string `pulumi:"id"`
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The ID of this resource.
	Id        string `pulumi:"id"`
	Name      string `pulumi:"name"`
	Slug      string `pulumi:"slug"`
	UpdatedAt string `pulumi:"updatedAt"`
	Url       string `pulumi:"url"`
}

func GetOrganizationOutput(ctx *pulumi.Context, args GetOrganizationOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationResult, error) {
			args := v.(GetOrganizationArgs)
			r, err := GetOrganization(ctx, &args, opts...)
			var s GetOrganizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationOutputArgs struct {
	// The ID of this resource.
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (GetOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Slug }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}