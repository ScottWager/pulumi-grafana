// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/onCall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := onCall.GetUserGroup(ctx, &oncall.GetUserGroupArgs{
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
func GetUserGroup(ctx *pulumi.Context, args *GetUserGroupArgs, opts ...pulumi.InvokeOption) (*GetUserGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserGroupResult
	err := ctx.Invoke("grafana:onCall/getUserGroup:getUserGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserGroup.
type GetUserGroupArgs struct {
	SlackHandle string `pulumi:"slackHandle"`
}

// A collection of values returned by getUserGroup.
type GetUserGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	SlackHandle string `pulumi:"slackHandle"`
	SlackId     string `pulumi:"slackId"`
}

func GetUserGroupOutput(ctx *pulumi.Context, args GetUserGroupOutputArgs, opts ...pulumi.InvokeOption) GetUserGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserGroupResult, error) {
			args := v.(GetUserGroupArgs)
			r, err := GetUserGroup(ctx, &args, opts...)
			var s GetUserGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserGroupResultOutput)
}

// A collection of arguments for invoking getUserGroup.
type GetUserGroupOutputArgs struct {
	SlackHandle pulumi.StringInput `pulumi:"slackHandle"`
}

func (GetUserGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserGroupArgs)(nil)).Elem()
}

// A collection of values returned by getUserGroup.
type GetUserGroupResultOutput struct{ *pulumi.OutputState }

func (GetUserGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserGroupResult)(nil)).Elem()
}

func (o GetUserGroupResultOutput) ToGetUserGroupResultOutput() GetUserGroupResultOutput {
	return o
}

func (o GetUserGroupResultOutput) ToGetUserGroupResultOutputWithContext(ctx context.Context) GetUserGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserGroupResultOutput) SlackHandle() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserGroupResult) string { return v.SlackHandle }).(pulumi.StringOutput)
}

func (o GetUserGroupResultOutput) SlackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserGroupResult) string { return v.SlackId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserGroupResultOutput{})
}