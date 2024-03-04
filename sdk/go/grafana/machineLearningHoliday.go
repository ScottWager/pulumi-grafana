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

// A holiday describes time periods where a time series is expected to behave differently to normal.
//
// To use a holiday in a job, use its id in the `holidays` attribute of a `MachineLearningJob`:
type MachineLearningHoliday struct {
	pulumi.CustomResourceState

	// A list of custom periods for the holiday.
	CustomPeriods MachineLearningHolidayCustomPeriodArrayOutput `pulumi:"customPeriods"`
	// A description of the holiday.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrOutput `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrOutput `pulumi:"icalUrl"`
	// The name of the custom period.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewMachineLearningHoliday registers a new resource with the given unique name, arguments, and options.
func NewMachineLearningHoliday(ctx *pulumi.Context,
	name string, args *MachineLearningHolidayArgs, opts ...pulumi.ResourceOption) (*MachineLearningHoliday, error) {
	if args == nil {
		args = &MachineLearningHolidayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineLearningHoliday
	err := ctx.RegisterResource("grafana:index/machineLearningHoliday:MachineLearningHoliday", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineLearningHoliday gets an existing MachineLearningHoliday resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineLearningHoliday(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningHolidayState, opts ...pulumi.ResourceOption) (*MachineLearningHoliday, error) {
	var resource MachineLearningHoliday
	err := ctx.ReadResource("grafana:index/machineLearningHoliday:MachineLearningHoliday", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineLearningHoliday resources.
type machineLearningHolidayState struct {
	// A list of custom periods for the holiday.
	CustomPeriods []MachineLearningHolidayCustomPeriod `pulumi:"customPeriods"`
	// A description of the holiday.
	Description *string `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl *string `pulumi:"icalUrl"`
	// The name of the custom period.
	Name *string `pulumi:"name"`
}

type MachineLearningHolidayState struct {
	// A list of custom periods for the holiday.
	CustomPeriods MachineLearningHolidayCustomPeriodArrayInput
	// A description of the holiday.
	Description pulumi.StringPtrInput
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrInput
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrInput
	// The name of the custom period.
	Name pulumi.StringPtrInput
}

func (MachineLearningHolidayState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningHolidayState)(nil)).Elem()
}

type machineLearningHolidayArgs struct {
	// A list of custom periods for the holiday.
	CustomPeriods []MachineLearningHolidayCustomPeriod `pulumi:"customPeriods"`
	// A description of the holiday.
	Description *string `pulumi:"description"`
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `pulumi:"icalTimezone"`
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl *string `pulumi:"icalUrl"`
	// The name of the custom period.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a MachineLearningHoliday resource.
type MachineLearningHolidayArgs struct {
	// A list of custom periods for the holiday.
	CustomPeriods MachineLearningHolidayCustomPeriodArrayInput
	// A description of the holiday.
	Description pulumi.StringPtrInput
	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone pulumi.StringPtrInput
	// A URL to an iCal file containing all occurrences of the holiday.
	IcalUrl pulumi.StringPtrInput
	// The name of the custom period.
	Name pulumi.StringPtrInput
}

func (MachineLearningHolidayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningHolidayArgs)(nil)).Elem()
}

type MachineLearningHolidayInput interface {
	pulumi.Input

	ToMachineLearningHolidayOutput() MachineLearningHolidayOutput
	ToMachineLearningHolidayOutputWithContext(ctx context.Context) MachineLearningHolidayOutput
}

func (*MachineLearningHoliday) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningHoliday)(nil)).Elem()
}

func (i *MachineLearningHoliday) ToMachineLearningHolidayOutput() MachineLearningHolidayOutput {
	return i.ToMachineLearningHolidayOutputWithContext(context.Background())
}

func (i *MachineLearningHoliday) ToMachineLearningHolidayOutputWithContext(ctx context.Context) MachineLearningHolidayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningHolidayOutput)
}

func (i *MachineLearningHoliday) ToOutput(ctx context.Context) pulumix.Output[*MachineLearningHoliday] {
	return pulumix.Output[*MachineLearningHoliday]{
		OutputState: i.ToMachineLearningHolidayOutputWithContext(ctx).OutputState,
	}
}

// MachineLearningHolidayArrayInput is an input type that accepts MachineLearningHolidayArray and MachineLearningHolidayArrayOutput values.
// You can construct a concrete instance of `MachineLearningHolidayArrayInput` via:
//
//	MachineLearningHolidayArray{ MachineLearningHolidayArgs{...} }
type MachineLearningHolidayArrayInput interface {
	pulumi.Input

	ToMachineLearningHolidayArrayOutput() MachineLearningHolidayArrayOutput
	ToMachineLearningHolidayArrayOutputWithContext(context.Context) MachineLearningHolidayArrayOutput
}

type MachineLearningHolidayArray []MachineLearningHolidayInput

func (MachineLearningHolidayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineLearningHoliday)(nil)).Elem()
}

func (i MachineLearningHolidayArray) ToMachineLearningHolidayArrayOutput() MachineLearningHolidayArrayOutput {
	return i.ToMachineLearningHolidayArrayOutputWithContext(context.Background())
}

func (i MachineLearningHolidayArray) ToMachineLearningHolidayArrayOutputWithContext(ctx context.Context) MachineLearningHolidayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningHolidayArrayOutput)
}

func (i MachineLearningHolidayArray) ToOutput(ctx context.Context) pulumix.Output[[]*MachineLearningHoliday] {
	return pulumix.Output[[]*MachineLearningHoliday]{
		OutputState: i.ToMachineLearningHolidayArrayOutputWithContext(ctx).OutputState,
	}
}

// MachineLearningHolidayMapInput is an input type that accepts MachineLearningHolidayMap and MachineLearningHolidayMapOutput values.
// You can construct a concrete instance of `MachineLearningHolidayMapInput` via:
//
//	MachineLearningHolidayMap{ "key": MachineLearningHolidayArgs{...} }
type MachineLearningHolidayMapInput interface {
	pulumi.Input

	ToMachineLearningHolidayMapOutput() MachineLearningHolidayMapOutput
	ToMachineLearningHolidayMapOutputWithContext(context.Context) MachineLearningHolidayMapOutput
}

type MachineLearningHolidayMap map[string]MachineLearningHolidayInput

func (MachineLearningHolidayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineLearningHoliday)(nil)).Elem()
}

func (i MachineLearningHolidayMap) ToMachineLearningHolidayMapOutput() MachineLearningHolidayMapOutput {
	return i.ToMachineLearningHolidayMapOutputWithContext(context.Background())
}

func (i MachineLearningHolidayMap) ToMachineLearningHolidayMapOutputWithContext(ctx context.Context) MachineLearningHolidayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningHolidayMapOutput)
}

func (i MachineLearningHolidayMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineLearningHoliday] {
	return pulumix.Output[map[string]*MachineLearningHoliday]{
		OutputState: i.ToMachineLearningHolidayMapOutputWithContext(ctx).OutputState,
	}
}

type MachineLearningHolidayOutput struct{ *pulumi.OutputState }

func (MachineLearningHolidayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningHoliday)(nil)).Elem()
}

func (o MachineLearningHolidayOutput) ToMachineLearningHolidayOutput() MachineLearningHolidayOutput {
	return o
}

func (o MachineLearningHolidayOutput) ToMachineLearningHolidayOutputWithContext(ctx context.Context) MachineLearningHolidayOutput {
	return o
}

func (o MachineLearningHolidayOutput) ToOutput(ctx context.Context) pulumix.Output[*MachineLearningHoliday] {
	return pulumix.Output[*MachineLearningHoliday]{
		OutputState: o.OutputState,
	}
}

// A list of custom periods for the holiday.
func (o MachineLearningHolidayOutput) CustomPeriods() MachineLearningHolidayCustomPeriodArrayOutput {
	return o.ApplyT(func(v *MachineLearningHoliday) MachineLearningHolidayCustomPeriodArrayOutput { return v.CustomPeriods }).(MachineLearningHolidayCustomPeriodArrayOutput)
}

// A description of the holiday.
func (o MachineLearningHolidayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningHoliday) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The timezone to use for events in the iCal file pointed to by ical_url.
func (o MachineLearningHolidayOutput) IcalTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningHoliday) pulumi.StringPtrOutput { return v.IcalTimezone }).(pulumi.StringPtrOutput)
}

// A URL to an iCal file containing all occurrences of the holiday.
func (o MachineLearningHolidayOutput) IcalUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningHoliday) pulumi.StringPtrOutput { return v.IcalUrl }).(pulumi.StringPtrOutput)
}

// The name of the custom period.
func (o MachineLearningHolidayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningHoliday) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type MachineLearningHolidayArrayOutput struct{ *pulumi.OutputState }

func (MachineLearningHolidayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineLearningHoliday)(nil)).Elem()
}

func (o MachineLearningHolidayArrayOutput) ToMachineLearningHolidayArrayOutput() MachineLearningHolidayArrayOutput {
	return o
}

func (o MachineLearningHolidayArrayOutput) ToMachineLearningHolidayArrayOutputWithContext(ctx context.Context) MachineLearningHolidayArrayOutput {
	return o
}

func (o MachineLearningHolidayArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MachineLearningHoliday] {
	return pulumix.Output[[]*MachineLearningHoliday]{
		OutputState: o.OutputState,
	}
}

func (o MachineLearningHolidayArrayOutput) Index(i pulumi.IntInput) MachineLearningHolidayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineLearningHoliday {
		return vs[0].([]*MachineLearningHoliday)[vs[1].(int)]
	}).(MachineLearningHolidayOutput)
}

type MachineLearningHolidayMapOutput struct{ *pulumi.OutputState }

func (MachineLearningHolidayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineLearningHoliday)(nil)).Elem()
}

func (o MachineLearningHolidayMapOutput) ToMachineLearningHolidayMapOutput() MachineLearningHolidayMapOutput {
	return o
}

func (o MachineLearningHolidayMapOutput) ToMachineLearningHolidayMapOutputWithContext(ctx context.Context) MachineLearningHolidayMapOutput {
	return o
}

func (o MachineLearningHolidayMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MachineLearningHoliday] {
	return pulumix.Output[map[string]*MachineLearningHoliday]{
		OutputState: o.OutputState,
	}
}

func (o MachineLearningHolidayMapOutput) MapIndex(k pulumi.StringInput) MachineLearningHolidayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineLearningHoliday {
		return vs[0].(map[string]*MachineLearningHoliday)[vs[1].(string)]
	}).(MachineLearningHolidayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningHolidayInput)(nil)).Elem(), &MachineLearningHoliday{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningHolidayArrayInput)(nil)).Elem(), MachineLearningHolidayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningHolidayMapInput)(nil)).Elem(), MachineLearningHolidayMap{})
	pulumi.RegisterOutputType(MachineLearningHolidayOutput{})
	pulumi.RegisterOutputType(MachineLearningHolidayArrayOutput{})
	pulumi.RegisterOutputType(MachineLearningHolidayMapOutput{})
}
