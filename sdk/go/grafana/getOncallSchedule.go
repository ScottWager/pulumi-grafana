// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// * [Official documentation](https://grafana.com/docs/grafana-cloud/oncall/calendar-schedules/)
// * [HTTP API](https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/schedules/)
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
//			_, err = grafana.LookupOncallSchedule(ctx, &GetOncallScheduleArgs{
//				Name: "example_schedule",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOncallSchedule(ctx *pulumi.Context, args *LookupOncallScheduleArgs, opts ...pulumi.InvokeOption) (*LookupOncallScheduleResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOncallScheduleResult
	err := ctx.Invoke("grafana:index/getOncallSchedule:getOncallSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOncallSchedule.
type LookupOncallScheduleArgs struct {
	// The schedule name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOncallSchedule.
type LookupOncallScheduleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The schedule name.
	Name string `pulumi:"name"`
	// The schedule type.
	Type string `pulumi:"type"`
}

func LookupOncallScheduleOutput(ctx *pulumi.Context, args LookupOncallScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupOncallScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOncallScheduleResult, error) {
			args := v.(LookupOncallScheduleArgs)
			r, err := LookupOncallSchedule(ctx, &args, opts...)
			var s LookupOncallScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOncallScheduleResultOutput)
}

// A collection of arguments for invoking getOncallSchedule.
type LookupOncallScheduleOutputArgs struct {
	// The schedule name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOncallScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOncallScheduleArgs)(nil)).Elem()
}

// A collection of values returned by getOncallSchedule.
type LookupOncallScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupOncallScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOncallScheduleResult)(nil)).Elem()
}

func (o LookupOncallScheduleResultOutput) ToLookupOncallScheduleResultOutput() LookupOncallScheduleResultOutput {
	return o
}

func (o LookupOncallScheduleResultOutput) ToLookupOncallScheduleResultOutputWithContext(ctx context.Context) LookupOncallScheduleResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOncallScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOncallScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The schedule name.
func (o LookupOncallScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOncallScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The schedule type.
func (o LookupOncallScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOncallScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOncallScheduleResultOutput{})
}
