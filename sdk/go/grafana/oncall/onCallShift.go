// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oncall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/on_call_shifts/)
//
// ## Import
//
// ```sh
// $ pulumi import grafana:onCall/onCallShift:OnCallShift name "{{ id }}"
// ```
type OnCallShift struct {
	pulumi.CustomResourceState

	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	ByDays pulumi.StringArrayOutput `pulumi:"byDays"`
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	ByMonthdays pulumi.IntArrayOutput `pulumi:"byMonthdays"`
	// This parameter takes a list of months. Valid values are 1 to 12
	ByMonths pulumi.IntArrayOutput `pulumi:"byMonths"`
	// The duration of the event.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The priority level. The higher the value, the higher the priority.
	Level pulumi.IntPtrOutput `pulumi:"level"`
	// The shift's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of lists with on-call users (for rollingUsers event type)
	RollingUsers pulumi.StringArrayArrayOutput `pulumi:"rollingUsers"`
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start pulumi.StringOutput `pulumi:"start"`
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex pulumi.IntPtrOutput `pulumi:"startRotationFromUserIndex"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
	// The shift's type. Can be rolling*users, recurrent*event, single_event
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of on-call users (for single*event and recurrent*event event type).
	Users pulumi.StringArrayOutput `pulumi:"users"`
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart pulumi.StringPtrOutput `pulumi:"weekStart"`
}

// NewOnCallShift registers a new resource with the given unique name, arguments, and options.
func NewOnCallShift(ctx *pulumi.Context,
	name string, args *OnCallShiftArgs, opts ...pulumi.ResourceOption) (*OnCallShift, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/oncallOnCallShift:OncallOnCallShift"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OnCallShift
	err := ctx.RegisterResource("grafana:onCall/onCallShift:OnCallShift", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOnCallShift gets an existing OnCallShift resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOnCallShift(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnCallShiftState, opts ...pulumi.ResourceOption) (*OnCallShift, error) {
	var resource OnCallShift
	err := ctx.ReadResource("grafana:onCall/onCallShift:OnCallShift", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OnCallShift resources.
type onCallShiftState struct {
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	ByDays []string `pulumi:"byDays"`
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	ByMonthdays []int `pulumi:"byMonthdays"`
	// This parameter takes a list of months. Valid values are 1 to 12
	ByMonths []int `pulumi:"byMonths"`
	// The duration of the event.
	Duration *int `pulumi:"duration"`
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency *string `pulumi:"frequency"`
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval *int `pulumi:"interval"`
	// The priority level. The higher the value, the higher the priority.
	Level *int `pulumi:"level"`
	// The shift's name.
	Name *string `pulumi:"name"`
	// The list of lists with on-call users (for rollingUsers event type)
	RollingUsers [][]string `pulumi:"rollingUsers"`
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start *string `pulumi:"start"`
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex *int `pulumi:"startRotationFromUserIndex"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
	TeamId *string `pulumi:"teamId"`
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone *string `pulumi:"timeZone"`
	// The shift's type. Can be rolling*users, recurrent*event, single_event
	Type *string `pulumi:"type"`
	// The list of on-call users (for single*event and recurrent*event event type).
	Users []string `pulumi:"users"`
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart *string `pulumi:"weekStart"`
}

type OnCallShiftState struct {
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	ByDays pulumi.StringArrayInput
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	ByMonthdays pulumi.IntArrayInput
	// This parameter takes a list of months. Valid values are 1 to 12
	ByMonths pulumi.IntArrayInput
	// The duration of the event.
	Duration pulumi.IntPtrInput
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency pulumi.StringPtrInput
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval pulumi.IntPtrInput
	// The priority level. The higher the value, the higher the priority.
	Level pulumi.IntPtrInput
	// The shift's name.
	Name pulumi.StringPtrInput
	// The list of lists with on-call users (for rollingUsers event type)
	RollingUsers pulumi.StringArrayArrayInput
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start pulumi.StringPtrInput
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex pulumi.IntPtrInput
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
	TeamId pulumi.StringPtrInput
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone pulumi.StringPtrInput
	// The shift's type. Can be rolling*users, recurrent*event, single_event
	Type pulumi.StringPtrInput
	// The list of on-call users (for single*event and recurrent*event event type).
	Users pulumi.StringArrayInput
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart pulumi.StringPtrInput
}

func (OnCallShiftState) ElementType() reflect.Type {
	return reflect.TypeOf((*onCallShiftState)(nil)).Elem()
}

type onCallShiftArgs struct {
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	ByDays []string `pulumi:"byDays"`
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	ByMonthdays []int `pulumi:"byMonthdays"`
	// This parameter takes a list of months. Valid values are 1 to 12
	ByMonths []int `pulumi:"byMonths"`
	// The duration of the event.
	Duration int `pulumi:"duration"`
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency *string `pulumi:"frequency"`
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval *int `pulumi:"interval"`
	// The priority level. The higher the value, the higher the priority.
	Level *int `pulumi:"level"`
	// The shift's name.
	Name *string `pulumi:"name"`
	// The list of lists with on-call users (for rollingUsers event type)
	RollingUsers [][]string `pulumi:"rollingUsers"`
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start string `pulumi:"start"`
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex *int `pulumi:"startRotationFromUserIndex"`
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
	TeamId *string `pulumi:"teamId"`
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone *string `pulumi:"timeZone"`
	// The shift's type. Can be rolling*users, recurrent*event, single_event
	Type string `pulumi:"type"`
	// The list of on-call users (for single*event and recurrent*event event type).
	Users []string `pulumi:"users"`
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart *string `pulumi:"weekStart"`
}

// The set of arguments for constructing a OnCallShift resource.
type OnCallShiftArgs struct {
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	ByDays pulumi.StringArrayInput
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	ByMonthdays pulumi.IntArrayInput
	// This parameter takes a list of months. Valid values are 1 to 12
	ByMonths pulumi.IntArrayInput
	// The duration of the event.
	Duration pulumi.IntInput
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency pulumi.StringPtrInput
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval pulumi.IntPtrInput
	// The priority level. The higher the value, the higher the priority.
	Level pulumi.IntPtrInput
	// The shift's name.
	Name pulumi.StringPtrInput
	// The list of lists with on-call users (for rollingUsers event type)
	RollingUsers pulumi.StringArrayArrayInput
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start pulumi.StringInput
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex pulumi.IntPtrInput
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
	TeamId pulumi.StringPtrInput
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone pulumi.StringPtrInput
	// The shift's type. Can be rolling*users, recurrent*event, single_event
	Type pulumi.StringInput
	// The list of on-call users (for single*event and recurrent*event event type).
	Users pulumi.StringArrayInput
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart pulumi.StringPtrInput
}

func (OnCallShiftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onCallShiftArgs)(nil)).Elem()
}

type OnCallShiftInput interface {
	pulumi.Input

	ToOnCallShiftOutput() OnCallShiftOutput
	ToOnCallShiftOutputWithContext(ctx context.Context) OnCallShiftOutput
}

func (*OnCallShift) ElementType() reflect.Type {
	return reflect.TypeOf((**OnCallShift)(nil)).Elem()
}

func (i *OnCallShift) ToOnCallShiftOutput() OnCallShiftOutput {
	return i.ToOnCallShiftOutputWithContext(context.Background())
}

func (i *OnCallShift) ToOnCallShiftOutputWithContext(ctx context.Context) OnCallShiftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnCallShiftOutput)
}

// OnCallShiftArrayInput is an input type that accepts OnCallShiftArray and OnCallShiftArrayOutput values.
// You can construct a concrete instance of `OnCallShiftArrayInput` via:
//
//	OnCallShiftArray{ OnCallShiftArgs{...} }
type OnCallShiftArrayInput interface {
	pulumi.Input

	ToOnCallShiftArrayOutput() OnCallShiftArrayOutput
	ToOnCallShiftArrayOutputWithContext(context.Context) OnCallShiftArrayOutput
}

type OnCallShiftArray []OnCallShiftInput

func (OnCallShiftArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OnCallShift)(nil)).Elem()
}

func (i OnCallShiftArray) ToOnCallShiftArrayOutput() OnCallShiftArrayOutput {
	return i.ToOnCallShiftArrayOutputWithContext(context.Background())
}

func (i OnCallShiftArray) ToOnCallShiftArrayOutputWithContext(ctx context.Context) OnCallShiftArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnCallShiftArrayOutput)
}

// OnCallShiftMapInput is an input type that accepts OnCallShiftMap and OnCallShiftMapOutput values.
// You can construct a concrete instance of `OnCallShiftMapInput` via:
//
//	OnCallShiftMap{ "key": OnCallShiftArgs{...} }
type OnCallShiftMapInput interface {
	pulumi.Input

	ToOnCallShiftMapOutput() OnCallShiftMapOutput
	ToOnCallShiftMapOutputWithContext(context.Context) OnCallShiftMapOutput
}

type OnCallShiftMap map[string]OnCallShiftInput

func (OnCallShiftMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OnCallShift)(nil)).Elem()
}

func (i OnCallShiftMap) ToOnCallShiftMapOutput() OnCallShiftMapOutput {
	return i.ToOnCallShiftMapOutputWithContext(context.Background())
}

func (i OnCallShiftMap) ToOnCallShiftMapOutputWithContext(ctx context.Context) OnCallShiftMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnCallShiftMapOutput)
}

type OnCallShiftOutput struct{ *pulumi.OutputState }

func (OnCallShiftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnCallShift)(nil)).Elem()
}

func (o OnCallShiftOutput) ToOnCallShiftOutput() OnCallShiftOutput {
	return o
}

func (o OnCallShiftOutput) ToOnCallShiftOutputWithContext(ctx context.Context) OnCallShiftOutput {
	return o
}

// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
func (o OnCallShiftOutput) ByDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringArrayOutput { return v.ByDays }).(pulumi.StringArrayOutput)
}

// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
func (o OnCallShiftOutput) ByMonthdays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntArrayOutput { return v.ByMonthdays }).(pulumi.IntArrayOutput)
}

// This parameter takes a list of months. Valid values are 1 to 12
func (o OnCallShiftOutput) ByMonths() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntArrayOutput { return v.ByMonths }).(pulumi.IntArrayOutput)
}

// The duration of the event.
func (o OnCallShiftOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

// The frequency of the event. Can be hourly, daily, weekly, monthly
func (o OnCallShiftOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringPtrOutput { return v.Frequency }).(pulumi.StringPtrOutput)
}

// The positive integer representing at which intervals the recurrence rule repeats.
func (o OnCallShiftOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// The priority level. The higher the value, the higher the priority.
func (o OnCallShiftOutput) Level() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntPtrOutput { return v.Level }).(pulumi.IntPtrOutput)
}

// The shift's name.
func (o OnCallShiftOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of lists with on-call users (for rollingUsers event type)
func (o OnCallShiftOutput) RollingUsers() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringArrayArrayOutput { return v.RollingUsers }).(pulumi.StringArrayArrayOutput)
}

// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
func (o OnCallShiftOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// The index of the list of users in rolling_users, from which on-call rotation starts.
func (o OnCallShiftOutput) StartRotationFromUserIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.IntPtrOutput { return v.StartRotationFromUserIndex }).(pulumi.IntPtrOutput)
}

// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `onCall.getTeam` datasource.
func (o OnCallShiftOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

// The shift's timezone.  Overrides schedule's timezone.
func (o OnCallShiftOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

// The shift's type. Can be rolling*users, recurrent*event, single_event
func (o OnCallShiftOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of on-call users (for single*event and recurrent*event event type).
func (o OnCallShiftOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
func (o OnCallShiftOutput) WeekStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnCallShift) pulumi.StringPtrOutput { return v.WeekStart }).(pulumi.StringPtrOutput)
}

type OnCallShiftArrayOutput struct{ *pulumi.OutputState }

func (OnCallShiftArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OnCallShift)(nil)).Elem()
}

func (o OnCallShiftArrayOutput) ToOnCallShiftArrayOutput() OnCallShiftArrayOutput {
	return o
}

func (o OnCallShiftArrayOutput) ToOnCallShiftArrayOutputWithContext(ctx context.Context) OnCallShiftArrayOutput {
	return o
}

func (o OnCallShiftArrayOutput) Index(i pulumi.IntInput) OnCallShiftOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OnCallShift {
		return vs[0].([]*OnCallShift)[vs[1].(int)]
	}).(OnCallShiftOutput)
}

type OnCallShiftMapOutput struct{ *pulumi.OutputState }

func (OnCallShiftMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OnCallShift)(nil)).Elem()
}

func (o OnCallShiftMapOutput) ToOnCallShiftMapOutput() OnCallShiftMapOutput {
	return o
}

func (o OnCallShiftMapOutput) ToOnCallShiftMapOutputWithContext(ctx context.Context) OnCallShiftMapOutput {
	return o
}

func (o OnCallShiftMapOutput) MapIndex(k pulumi.StringInput) OnCallShiftOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OnCallShift {
		return vs[0].(map[string]*OnCallShift)[vs[1].(string)]
	}).(OnCallShiftOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OnCallShiftInput)(nil)).Elem(), &OnCallShift{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnCallShiftArrayInput)(nil)).Elem(), OnCallShiftArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnCallShiftMapInput)(nil)).Elem(), OnCallShiftMap{})
	pulumi.RegisterOutputType(OnCallShiftOutput{})
	pulumi.RegisterOutputType(OnCallShiftArrayOutput{})
	pulumi.RegisterOutputType(OnCallShiftMapOutput{})
}