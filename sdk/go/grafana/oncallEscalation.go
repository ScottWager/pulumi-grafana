// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/oncall/latest/escalation-chains-and-routes/)
// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_policies/)
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
//			_, err := grafana.NewOncallEscalationChain(ctx, "default", nil, pulumi.Provider(grafana.Oncall))
//			if err != nil {
//				return err
//			}
//			alex, err := grafana.GetOncallUser(ctx, &grafana.GetOncallUserArgs{
//				Username: "alex",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewOncallEscalation(ctx, "exampleNotifyStepOncallEscalation", &grafana.OncallEscalationArgs{
//				EscalationChainId: _default.ID(),
//				Type:              pulumi.String("notify_persons"),
//				PersonsToNotifies: pulumi.StringArray{
//					*pulumi.String(alex.Id),
//				},
//				Position: pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewOncallEscalation(ctx, "exampleNotifyStepIndex/oncallEscalationOncallEscalation", &grafana.OncallEscalationArgs{
//				EscalationChainId: _default.ID(),
//				Type:              pulumi.String("wait"),
//				Duration:          pulumi.Int(300),
//				Position:          pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewOncallEscalation(ctx, "exampleNotifyStepGrafanaIndex/oncallEscalationOncallEscalation", &grafana.OncallEscalationArgs{
//				EscalationChainId: _default.ID(),
//				Type:              pulumi.String("notify_persons"),
//				Important:         pulumi.Bool(true),
//				PersonsToNotifies: pulumi.StringArray{
//					*pulumi.String(alex.Id),
//				},
//				Position: pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import grafana:index/oncallEscalation:OncallEscalation escalation_name {{escalation_id}}
//
// ```
type OncallEscalation struct {
	pulumi.CustomResourceState

	// The ID of an Action for triggerAction type step.
	ActionToTrigger pulumi.StringPtrOutput `pulumi:"actionToTrigger"`
	// The duration of delay for wait type step.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The ID of the escalation chain.
	EscalationChainId pulumi.StringOutput `pulumi:"escalationChainId"`
	// The ID of a User Group for notify*user*group type step.
	GroupToNotify pulumi.StringPtrOutput `pulumi:"groupToNotify"`
	// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
	Important pulumi.BoolPtrOutput `pulumi:"important"`
	// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom pulumi.StringPtrOutput `pulumi:"notifyIfTimeFrom"`
	// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo pulumi.StringPtrOutput `pulumi:"notifyIfTimeTo"`
	// ID of a Schedule for notify*on*call*from*schedule type step.
	NotifyOnCallFromSchedule pulumi.StringPtrOutput `pulumi:"notifyOnCallFromSchedule"`
	// The list of ID's of users for notifyPersons type step.
	PersonsToNotifies pulumi.StringArrayOutput `pulumi:"personsToNotifies"`
	// The list of ID's of users for notify*person*next*each*time type step.
	PersonsToNotifyNextEachTimes pulumi.StringArrayOutput `pulumi:"personsToNotifyNextEachTimes"`
	// The position of the escalation step (starts from 0).
	Position pulumi.IntOutput `pulumi:"position"`
	// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewOncallEscalation registers a new resource with the given unique name, arguments, and options.
func NewOncallEscalation(ctx *pulumi.Context,
	name string, args *OncallEscalationArgs, opts ...pulumi.ResourceOption) (*OncallEscalation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EscalationChainId == nil {
		return nil, errors.New("invalid value for required argument 'EscalationChainId'")
	}
	if args.Position == nil {
		return nil, errors.New("invalid value for required argument 'Position'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OncallEscalation
	err := ctx.RegisterResource("grafana:index/oncallEscalation:OncallEscalation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOncallEscalation gets an existing OncallEscalation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOncallEscalation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OncallEscalationState, opts ...pulumi.ResourceOption) (*OncallEscalation, error) {
	var resource OncallEscalation
	err := ctx.ReadResource("grafana:index/oncallEscalation:OncallEscalation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OncallEscalation resources.
type oncallEscalationState struct {
	// The ID of an Action for triggerAction type step.
	ActionToTrigger *string `pulumi:"actionToTrigger"`
	// The duration of delay for wait type step.
	Duration *int `pulumi:"duration"`
	// The ID of the escalation chain.
	EscalationChainId *string `pulumi:"escalationChainId"`
	// The ID of a User Group for notify*user*group type step.
	GroupToNotify *string `pulumi:"groupToNotify"`
	// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
	Important *bool `pulumi:"important"`
	// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom *string `pulumi:"notifyIfTimeFrom"`
	// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo *string `pulumi:"notifyIfTimeTo"`
	// ID of a Schedule for notify*on*call*from*schedule type step.
	NotifyOnCallFromSchedule *string `pulumi:"notifyOnCallFromSchedule"`
	// The list of ID's of users for notifyPersons type step.
	PersonsToNotifies []string `pulumi:"personsToNotifies"`
	// The list of ID's of users for notify*person*next*each*time type step.
	PersonsToNotifyNextEachTimes []string `pulumi:"personsToNotifyNextEachTimes"`
	// The position of the escalation step (starts from 0).
	Position *int `pulumi:"position"`
	// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
	Type *string `pulumi:"type"`
}

type OncallEscalationState struct {
	// The ID of an Action for triggerAction type step.
	ActionToTrigger pulumi.StringPtrInput
	// The duration of delay for wait type step.
	Duration pulumi.IntPtrInput
	// The ID of the escalation chain.
	EscalationChainId pulumi.StringPtrInput
	// The ID of a User Group for notify*user*group type step.
	GroupToNotify pulumi.StringPtrInput
	// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
	Important pulumi.BoolPtrInput
	// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom pulumi.StringPtrInput
	// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo pulumi.StringPtrInput
	// ID of a Schedule for notify*on*call*from*schedule type step.
	NotifyOnCallFromSchedule pulumi.StringPtrInput
	// The list of ID's of users for notifyPersons type step.
	PersonsToNotifies pulumi.StringArrayInput
	// The list of ID's of users for notify*person*next*each*time type step.
	PersonsToNotifyNextEachTimes pulumi.StringArrayInput
	// The position of the escalation step (starts from 0).
	Position pulumi.IntPtrInput
	// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
	Type pulumi.StringPtrInput
}

func (OncallEscalationState) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallEscalationState)(nil)).Elem()
}

type oncallEscalationArgs struct {
	// The ID of an Action for triggerAction type step.
	ActionToTrigger *string `pulumi:"actionToTrigger"`
	// The duration of delay for wait type step.
	Duration *int `pulumi:"duration"`
	// The ID of the escalation chain.
	EscalationChainId string `pulumi:"escalationChainId"`
	// The ID of a User Group for notify*user*group type step.
	GroupToNotify *string `pulumi:"groupToNotify"`
	// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
	Important *bool `pulumi:"important"`
	// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom *string `pulumi:"notifyIfTimeFrom"`
	// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo *string `pulumi:"notifyIfTimeTo"`
	// ID of a Schedule for notify*on*call*from*schedule type step.
	NotifyOnCallFromSchedule *string `pulumi:"notifyOnCallFromSchedule"`
	// The list of ID's of users for notifyPersons type step.
	PersonsToNotifies []string `pulumi:"personsToNotifies"`
	// The list of ID's of users for notify*person*next*each*time type step.
	PersonsToNotifyNextEachTimes []string `pulumi:"personsToNotifyNextEachTimes"`
	// The position of the escalation step (starts from 0).
	Position int `pulumi:"position"`
	// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OncallEscalation resource.
type OncallEscalationArgs struct {
	// The ID of an Action for triggerAction type step.
	ActionToTrigger pulumi.StringPtrInput
	// The duration of delay for wait type step.
	Duration pulumi.IntPtrInput
	// The ID of the escalation chain.
	EscalationChainId pulumi.StringInput
	// The ID of a User Group for notify*user*group type step.
	GroupToNotify pulumi.StringPtrInput
	// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
	Important pulumi.BoolPtrInput
	// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom pulumi.StringPtrInput
	// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo pulumi.StringPtrInput
	// ID of a Schedule for notify*on*call*from*schedule type step.
	NotifyOnCallFromSchedule pulumi.StringPtrInput
	// The list of ID's of users for notifyPersons type step.
	PersonsToNotifies pulumi.StringArrayInput
	// The list of ID's of users for notify*person*next*each*time type step.
	PersonsToNotifyNextEachTimes pulumi.StringArrayInput
	// The position of the escalation step (starts from 0).
	Position pulumi.IntInput
	// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
	Type pulumi.StringPtrInput
}

func (OncallEscalationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oncallEscalationArgs)(nil)).Elem()
}

type OncallEscalationInput interface {
	pulumi.Input

	ToOncallEscalationOutput() OncallEscalationOutput
	ToOncallEscalationOutputWithContext(ctx context.Context) OncallEscalationOutput
}

func (*OncallEscalation) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallEscalation)(nil)).Elem()
}

func (i *OncallEscalation) ToOncallEscalationOutput() OncallEscalationOutput {
	return i.ToOncallEscalationOutputWithContext(context.Background())
}

func (i *OncallEscalation) ToOncallEscalationOutputWithContext(ctx context.Context) OncallEscalationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationOutput)
}

func (i *OncallEscalation) ToOutput(ctx context.Context) pulumix.Output[*OncallEscalation] {
	return pulumix.Output[*OncallEscalation]{
		OutputState: i.ToOncallEscalationOutputWithContext(ctx).OutputState,
	}
}

// OncallEscalationArrayInput is an input type that accepts OncallEscalationArray and OncallEscalationArrayOutput values.
// You can construct a concrete instance of `OncallEscalationArrayInput` via:
//
//	OncallEscalationArray{ OncallEscalationArgs{...} }
type OncallEscalationArrayInput interface {
	pulumi.Input

	ToOncallEscalationArrayOutput() OncallEscalationArrayOutput
	ToOncallEscalationArrayOutputWithContext(context.Context) OncallEscalationArrayOutput
}

type OncallEscalationArray []OncallEscalationInput

func (OncallEscalationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallEscalation)(nil)).Elem()
}

func (i OncallEscalationArray) ToOncallEscalationArrayOutput() OncallEscalationArrayOutput {
	return i.ToOncallEscalationArrayOutputWithContext(context.Background())
}

func (i OncallEscalationArray) ToOncallEscalationArrayOutputWithContext(ctx context.Context) OncallEscalationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationArrayOutput)
}

func (i OncallEscalationArray) ToOutput(ctx context.Context) pulumix.Output[[]*OncallEscalation] {
	return pulumix.Output[[]*OncallEscalation]{
		OutputState: i.ToOncallEscalationArrayOutputWithContext(ctx).OutputState,
	}
}

// OncallEscalationMapInput is an input type that accepts OncallEscalationMap and OncallEscalationMapOutput values.
// You can construct a concrete instance of `OncallEscalationMapInput` via:
//
//	OncallEscalationMap{ "key": OncallEscalationArgs{...} }
type OncallEscalationMapInput interface {
	pulumi.Input

	ToOncallEscalationMapOutput() OncallEscalationMapOutput
	ToOncallEscalationMapOutputWithContext(context.Context) OncallEscalationMapOutput
}

type OncallEscalationMap map[string]OncallEscalationInput

func (OncallEscalationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallEscalation)(nil)).Elem()
}

func (i OncallEscalationMap) ToOncallEscalationMapOutput() OncallEscalationMapOutput {
	return i.ToOncallEscalationMapOutputWithContext(context.Background())
}

func (i OncallEscalationMap) ToOncallEscalationMapOutputWithContext(ctx context.Context) OncallEscalationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OncallEscalationMapOutput)
}

func (i OncallEscalationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OncallEscalation] {
	return pulumix.Output[map[string]*OncallEscalation]{
		OutputState: i.ToOncallEscalationMapOutputWithContext(ctx).OutputState,
	}
}

type OncallEscalationOutput struct{ *pulumi.OutputState }

func (OncallEscalationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OncallEscalation)(nil)).Elem()
}

func (o OncallEscalationOutput) ToOncallEscalationOutput() OncallEscalationOutput {
	return o
}

func (o OncallEscalationOutput) ToOncallEscalationOutputWithContext(ctx context.Context) OncallEscalationOutput {
	return o
}

func (o OncallEscalationOutput) ToOutput(ctx context.Context) pulumix.Output[*OncallEscalation] {
	return pulumix.Output[*OncallEscalation]{
		OutputState: o.OutputState,
	}
}

// The ID of an Action for triggerAction type step.
func (o OncallEscalationOutput) ActionToTrigger() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.ActionToTrigger }).(pulumi.StringPtrOutput)
}

// The duration of delay for wait type step.
func (o OncallEscalationOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// The ID of the escalation chain.
func (o OncallEscalationOutput) EscalationChainId() pulumi.StringOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringOutput { return v.EscalationChainId }).(pulumi.StringOutput)
}

// The ID of a User Group for notify*user*group type step.
func (o OncallEscalationOutput) GroupToNotify() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.GroupToNotify }).(pulumi.StringPtrOutput)
}

// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
func (o OncallEscalationOutput) Important() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.BoolPtrOutput { return v.Important }).(pulumi.BoolPtrOutput)
}

// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
func (o OncallEscalationOutput) NotifyIfTimeFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.NotifyIfTimeFrom }).(pulumi.StringPtrOutput)
}

// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
func (o OncallEscalationOutput) NotifyIfTimeTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.NotifyIfTimeTo }).(pulumi.StringPtrOutput)
}

// ID of a Schedule for notify*on*call*from*schedule type step.
func (o OncallEscalationOutput) NotifyOnCallFromSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.NotifyOnCallFromSchedule }).(pulumi.StringPtrOutput)
}

// The list of ID's of users for notifyPersons type step.
func (o OncallEscalationOutput) PersonsToNotifies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringArrayOutput { return v.PersonsToNotifies }).(pulumi.StringArrayOutput)
}

// The list of ID's of users for notify*person*next*each*time type step.
func (o OncallEscalationOutput) PersonsToNotifyNextEachTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringArrayOutput { return v.PersonsToNotifyNextEachTimes }).(pulumi.StringArrayOutput)
}

// The position of the escalation step (starts from 0).
func (o OncallEscalationOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
func (o OncallEscalationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OncallEscalation) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type OncallEscalationArrayOutput struct{ *pulumi.OutputState }

func (OncallEscalationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OncallEscalation)(nil)).Elem()
}

func (o OncallEscalationArrayOutput) ToOncallEscalationArrayOutput() OncallEscalationArrayOutput {
	return o
}

func (o OncallEscalationArrayOutput) ToOncallEscalationArrayOutputWithContext(ctx context.Context) OncallEscalationArrayOutput {
	return o
}

func (o OncallEscalationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OncallEscalation] {
	return pulumix.Output[[]*OncallEscalation]{
		OutputState: o.OutputState,
	}
}

func (o OncallEscalationArrayOutput) Index(i pulumi.IntInput) OncallEscalationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OncallEscalation {
		return vs[0].([]*OncallEscalation)[vs[1].(int)]
	}).(OncallEscalationOutput)
}

type OncallEscalationMapOutput struct{ *pulumi.OutputState }

func (OncallEscalationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OncallEscalation)(nil)).Elem()
}

func (o OncallEscalationMapOutput) ToOncallEscalationMapOutput() OncallEscalationMapOutput {
	return o
}

func (o OncallEscalationMapOutput) ToOncallEscalationMapOutputWithContext(ctx context.Context) OncallEscalationMapOutput {
	return o
}

func (o OncallEscalationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OncallEscalation] {
	return pulumix.Output[map[string]*OncallEscalation]{
		OutputState: o.OutputState,
	}
}

func (o OncallEscalationMapOutput) MapIndex(k pulumi.StringInput) OncallEscalationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OncallEscalation {
		return vs[0].(map[string]*OncallEscalation)[vs[1].(string)]
	}).(OncallEscalationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationInput)(nil)).Elem(), &OncallEscalation{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationArrayInput)(nil)).Elem(), OncallEscalationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OncallEscalationMapInput)(nil)).Elem(), OncallEscalationMap{})
	pulumi.RegisterOutputType(OncallEscalationOutput{})
	pulumi.RegisterOutputType(OncallEscalationArrayOutput{})
	pulumi.RegisterOutputType(OncallEscalationMapOutput{})
}
