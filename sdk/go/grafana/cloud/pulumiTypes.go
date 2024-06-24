// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

var _ = internal.GetEnvOrDefault

type AccessPolicyRealm struct {
	// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	Identifier    string                         `pulumi:"identifier"`
	LabelPolicies []AccessPolicyRealmLabelPolicy `pulumi:"labelPolicies"`
	// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
	Type string `pulumi:"type"`
}

// AccessPolicyRealmInput is an input type that accepts AccessPolicyRealmArgs and AccessPolicyRealmOutput values.
// You can construct a concrete instance of `AccessPolicyRealmInput` via:
//
//	AccessPolicyRealmArgs{...}
type AccessPolicyRealmInput interface {
	pulumi.Input

	ToAccessPolicyRealmOutput() AccessPolicyRealmOutput
	ToAccessPolicyRealmOutputWithContext(context.Context) AccessPolicyRealmOutput
}

type AccessPolicyRealmArgs struct {
	// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	Identifier    pulumi.StringInput                     `pulumi:"identifier"`
	LabelPolicies AccessPolicyRealmLabelPolicyArrayInput `pulumi:"labelPolicies"`
	// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (AccessPolicyRealmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRealm)(nil)).Elem()
}

func (i AccessPolicyRealmArgs) ToAccessPolicyRealmOutput() AccessPolicyRealmOutput {
	return i.ToAccessPolicyRealmOutputWithContext(context.Background())
}

func (i AccessPolicyRealmArgs) ToAccessPolicyRealmOutputWithContext(ctx context.Context) AccessPolicyRealmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyRealmOutput)
}

// AccessPolicyRealmArrayInput is an input type that accepts AccessPolicyRealmArray and AccessPolicyRealmArrayOutput values.
// You can construct a concrete instance of `AccessPolicyRealmArrayInput` via:
//
//	AccessPolicyRealmArray{ AccessPolicyRealmArgs{...} }
type AccessPolicyRealmArrayInput interface {
	pulumi.Input

	ToAccessPolicyRealmArrayOutput() AccessPolicyRealmArrayOutput
	ToAccessPolicyRealmArrayOutputWithContext(context.Context) AccessPolicyRealmArrayOutput
}

type AccessPolicyRealmArray []AccessPolicyRealmInput

func (AccessPolicyRealmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyRealm)(nil)).Elem()
}

func (i AccessPolicyRealmArray) ToAccessPolicyRealmArrayOutput() AccessPolicyRealmArrayOutput {
	return i.ToAccessPolicyRealmArrayOutputWithContext(context.Background())
}

func (i AccessPolicyRealmArray) ToAccessPolicyRealmArrayOutputWithContext(ctx context.Context) AccessPolicyRealmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyRealmArrayOutput)
}

type AccessPolicyRealmOutput struct{ *pulumi.OutputState }

func (AccessPolicyRealmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRealm)(nil)).Elem()
}

func (o AccessPolicyRealmOutput) ToAccessPolicyRealmOutput() AccessPolicyRealmOutput {
	return o
}

func (o AccessPolicyRealmOutput) ToAccessPolicyRealmOutputWithContext(ctx context.Context) AccessPolicyRealmOutput {
	return o
}

// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
func (o AccessPolicyRealmOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyRealm) string { return v.Identifier }).(pulumi.StringOutput)
}

func (o AccessPolicyRealmOutput) LabelPolicies() AccessPolicyRealmLabelPolicyArrayOutput {
	return o.ApplyT(func(v AccessPolicyRealm) []AccessPolicyRealmLabelPolicy { return v.LabelPolicies }).(AccessPolicyRealmLabelPolicyArrayOutput)
}

// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
func (o AccessPolicyRealmOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyRealm) string { return v.Type }).(pulumi.StringOutput)
}

type AccessPolicyRealmArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyRealmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyRealm)(nil)).Elem()
}

func (o AccessPolicyRealmArrayOutput) ToAccessPolicyRealmArrayOutput() AccessPolicyRealmArrayOutput {
	return o
}

func (o AccessPolicyRealmArrayOutput) ToAccessPolicyRealmArrayOutputWithContext(ctx context.Context) AccessPolicyRealmArrayOutput {
	return o
}

func (o AccessPolicyRealmArrayOutput) Index(i pulumi.IntInput) AccessPolicyRealmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicyRealm {
		return vs[0].([]AccessPolicyRealm)[vs[1].(int)]
	}).(AccessPolicyRealmOutput)
}

type AccessPolicyRealmLabelPolicy struct {
	// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	Selector string `pulumi:"selector"`
}

// AccessPolicyRealmLabelPolicyInput is an input type that accepts AccessPolicyRealmLabelPolicyArgs and AccessPolicyRealmLabelPolicyOutput values.
// You can construct a concrete instance of `AccessPolicyRealmLabelPolicyInput` via:
//
//	AccessPolicyRealmLabelPolicyArgs{...}
type AccessPolicyRealmLabelPolicyInput interface {
	pulumi.Input

	ToAccessPolicyRealmLabelPolicyOutput() AccessPolicyRealmLabelPolicyOutput
	ToAccessPolicyRealmLabelPolicyOutputWithContext(context.Context) AccessPolicyRealmLabelPolicyOutput
}

type AccessPolicyRealmLabelPolicyArgs struct {
	// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	Selector pulumi.StringInput `pulumi:"selector"`
}

func (AccessPolicyRealmLabelPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRealmLabelPolicy)(nil)).Elem()
}

func (i AccessPolicyRealmLabelPolicyArgs) ToAccessPolicyRealmLabelPolicyOutput() AccessPolicyRealmLabelPolicyOutput {
	return i.ToAccessPolicyRealmLabelPolicyOutputWithContext(context.Background())
}

func (i AccessPolicyRealmLabelPolicyArgs) ToAccessPolicyRealmLabelPolicyOutputWithContext(ctx context.Context) AccessPolicyRealmLabelPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyRealmLabelPolicyOutput)
}

// AccessPolicyRealmLabelPolicyArrayInput is an input type that accepts AccessPolicyRealmLabelPolicyArray and AccessPolicyRealmLabelPolicyArrayOutput values.
// You can construct a concrete instance of `AccessPolicyRealmLabelPolicyArrayInput` via:
//
//	AccessPolicyRealmLabelPolicyArray{ AccessPolicyRealmLabelPolicyArgs{...} }
type AccessPolicyRealmLabelPolicyArrayInput interface {
	pulumi.Input

	ToAccessPolicyRealmLabelPolicyArrayOutput() AccessPolicyRealmLabelPolicyArrayOutput
	ToAccessPolicyRealmLabelPolicyArrayOutputWithContext(context.Context) AccessPolicyRealmLabelPolicyArrayOutput
}

type AccessPolicyRealmLabelPolicyArray []AccessPolicyRealmLabelPolicyInput

func (AccessPolicyRealmLabelPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyRealmLabelPolicy)(nil)).Elem()
}

func (i AccessPolicyRealmLabelPolicyArray) ToAccessPolicyRealmLabelPolicyArrayOutput() AccessPolicyRealmLabelPolicyArrayOutput {
	return i.ToAccessPolicyRealmLabelPolicyArrayOutputWithContext(context.Background())
}

func (i AccessPolicyRealmLabelPolicyArray) ToAccessPolicyRealmLabelPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyRealmLabelPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyRealmLabelPolicyArrayOutput)
}

type AccessPolicyRealmLabelPolicyOutput struct{ *pulumi.OutputState }

func (AccessPolicyRealmLabelPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRealmLabelPolicy)(nil)).Elem()
}

func (o AccessPolicyRealmLabelPolicyOutput) ToAccessPolicyRealmLabelPolicyOutput() AccessPolicyRealmLabelPolicyOutput {
	return o
}

func (o AccessPolicyRealmLabelPolicyOutput) ToAccessPolicyRealmLabelPolicyOutputWithContext(ctx context.Context) AccessPolicyRealmLabelPolicyOutput {
	return o
}

// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
func (o AccessPolicyRealmLabelPolicyOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyRealmLabelPolicy) string { return v.Selector }).(pulumi.StringOutput)
}

type AccessPolicyRealmLabelPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyRealmLabelPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyRealmLabelPolicy)(nil)).Elem()
}

func (o AccessPolicyRealmLabelPolicyArrayOutput) ToAccessPolicyRealmLabelPolicyArrayOutput() AccessPolicyRealmLabelPolicyArrayOutput {
	return o
}

func (o AccessPolicyRealmLabelPolicyArrayOutput) ToAccessPolicyRealmLabelPolicyArrayOutputWithContext(ctx context.Context) AccessPolicyRealmLabelPolicyArrayOutput {
	return o
}

func (o AccessPolicyRealmLabelPolicyArrayOutput) Index(i pulumi.IntInput) AccessPolicyRealmLabelPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicyRealmLabelPolicy {
		return vs[0].([]AccessPolicyRealmLabelPolicy)[vs[1].(int)]
	}).(AccessPolicyRealmLabelPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRealmInput)(nil)).Elem(), AccessPolicyRealmArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRealmArrayInput)(nil)).Elem(), AccessPolicyRealmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRealmLabelPolicyInput)(nil)).Elem(), AccessPolicyRealmLabelPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyRealmLabelPolicyArrayInput)(nil)).Elem(), AccessPolicyRealmLabelPolicyArray{})
	pulumi.RegisterOutputType(AccessPolicyRealmOutput{})
	pulumi.RegisterOutputType(AccessPolicyRealmArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyRealmLabelPolicyOutput{})
	pulumi.RegisterOutputType(AccessPolicyRealmLabelPolicyArrayOutput{})
}