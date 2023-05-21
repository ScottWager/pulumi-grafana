// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/user_groups/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.GetOncallUserGroup(ctx, &grafana.GetOncallUserGroupArgs{
//				SlackHandle: "example_slack_handle",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOncallUserGroup(ctx *pulumi.Context, args *GetOncallUserGroupArgs, opts ...pulumi.InvokeOption) (*GetOncallUserGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOncallUserGroupResult
	err := ctx.Invoke("grafana:index/getOncallUserGroup:getOncallUserGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallUserGroup.
type GetOncallUserGroupArgs struct {
	SlackHandle string `pulumi:"slackHandle"`
}

// A collection of values returned by getOncallUserGroup.
type GetOncallUserGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	SlackHandle string `pulumi:"slackHandle"`
	SlackId     string `pulumi:"slackId"`
}

func GetOncallUserGroupOutput(ctx *pulumi.Context, args GetOncallUserGroupOutputArgs, opts ...pulumi.InvokeOption) GetOncallUserGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOncallUserGroupResult, error) {
			args := v.(GetOncallUserGroupArgs)
			r, err := GetOncallUserGroup(ctx, &args, opts...)
			var s GetOncallUserGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOncallUserGroupResultOutput)
}

// A collection of arguments for invoking getOncallUserGroup.
type GetOncallUserGroupOutputArgs struct {
	SlackHandle pulumi.StringInput `pulumi:"slackHandle"`
}

func (GetOncallUserGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallUserGroupArgs)(nil)).Elem()
}

// A collection of values returned by getOncallUserGroup.
type GetOncallUserGroupResultOutput struct{ *pulumi.OutputState }

func (GetOncallUserGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOncallUserGroupResult)(nil)).Elem()
}

func (o GetOncallUserGroupResultOutput) ToGetOncallUserGroupResultOutput() GetOncallUserGroupResultOutput {
	return o
}

func (o GetOncallUserGroupResultOutput) ToGetOncallUserGroupResultOutputWithContext(ctx context.Context) GetOncallUserGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOncallUserGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOncallUserGroupResultOutput) SlackHandle() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserGroupResult) string { return v.SlackHandle }).(pulumi.StringOutput)
}

func (o GetOncallUserGroupResultOutput) SlackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOncallUserGroupResult) string { return v.SlackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOncallUserGroupResultOutput{})
}
