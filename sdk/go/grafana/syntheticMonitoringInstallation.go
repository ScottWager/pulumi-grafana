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

// Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token.
// Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring.
// This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues.
//
// **Note that this resource must be used on a provider configured with Grafana Cloud credentials.**
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/installation/)
// * [API documentation](https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall)
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
//			smStackCloudStack, err := grafana.NewCloudStack(ctx, "smStackCloudStack", &grafana.CloudStackArgs{
//				Slug:       pulumi.String("<stack-slug>"),
//				RegionSlug: pulumi.String("us"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewCloudApiKey(ctx, "metricsPublish", &grafana.CloudApiKeyArgs{
//				Role:         pulumi.String("MetricsPublisher"),
//				CloudOrgSlug: pulumi.String("<org-slug>"),
//			})
//			if err != nil {
//				return err
//			}
//			smStackSyntheticMonitoringInstallation, err := grafana.NewSyntheticMonitoringInstallation(ctx, "smStackSyntheticMonitoringInstallation", &grafana.SyntheticMonitoringInstallationArgs{
//				StackId: smStackCloudStack.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewProvider(ctx, "sm", &grafana.ProviderArgs{
//				SmAccessToken: smStackSyntheticMonitoringInstallation.SmAccessToken,
//				SmUrl:         smStackSyntheticMonitoringInstallation.StackSmApiUrl,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SyntheticMonitoringInstallation struct {
	pulumi.CustomResourceState

	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	MetricsPublisherKey pulumi.StringOutput `pulumi:"metricsPublisherKey"`
	// Generated token to access the SM API.
	SmAccessToken pulumi.StringOutput `pulumi:"smAccessToken"`
	// The ID or slug of the stack to install SM on.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmApiUrl pulumi.StringOutput `pulumi:"stackSmApiUrl"`
}

// NewSyntheticMonitoringInstallation registers a new resource with the given unique name, arguments, and options.
func NewSyntheticMonitoringInstallation(ctx *pulumi.Context,
	name string, args *SyntheticMonitoringInstallationArgs, opts ...pulumi.ResourceOption) (*SyntheticMonitoringInstallation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricsPublisherKey == nil {
		return nil, errors.New("invalid value for required argument 'MetricsPublisherKey'")
	}
	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	if args.MetricsPublisherKey != nil {
		args.MetricsPublisherKey = pulumi.ToSecret(args.MetricsPublisherKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"metricsPublisherKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyntheticMonitoringInstallation
	err := ctx.RegisterResource("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyntheticMonitoringInstallation gets an existing SyntheticMonitoringInstallation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyntheticMonitoringInstallation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyntheticMonitoringInstallationState, opts ...pulumi.ResourceOption) (*SyntheticMonitoringInstallation, error) {
	var resource SyntheticMonitoringInstallation
	err := ctx.ReadResource("grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyntheticMonitoringInstallation resources.
type syntheticMonitoringInstallationState struct {
	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	MetricsPublisherKey *string `pulumi:"metricsPublisherKey"`
	// Generated token to access the SM API.
	SmAccessToken *string `pulumi:"smAccessToken"`
	// The ID or slug of the stack to install SM on.
	StackId *string `pulumi:"stackId"`
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmApiUrl *string `pulumi:"stackSmApiUrl"`
}

type SyntheticMonitoringInstallationState struct {
	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	MetricsPublisherKey pulumi.StringPtrInput
	// Generated token to access the SM API.
	SmAccessToken pulumi.StringPtrInput
	// The ID or slug of the stack to install SM on.
	StackId pulumi.StringPtrInput
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmApiUrl pulumi.StringPtrInput
}

func (SyntheticMonitoringInstallationState) ElementType() reflect.Type {
	return reflect.TypeOf((*syntheticMonitoringInstallationState)(nil)).Elem()
}

type syntheticMonitoringInstallationArgs struct {
	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	MetricsPublisherKey string `pulumi:"metricsPublisherKey"`
	// The ID or slug of the stack to install SM on.
	StackId string `pulumi:"stackId"`
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmApiUrl *string `pulumi:"stackSmApiUrl"`
}

// The set of arguments for constructing a SyntheticMonitoringInstallation resource.
type SyntheticMonitoringInstallationArgs struct {
	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	MetricsPublisherKey pulumi.StringInput
	// The ID or slug of the stack to install SM on.
	StackId pulumi.StringInput
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmApiUrl pulumi.StringPtrInput
}

func (SyntheticMonitoringInstallationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syntheticMonitoringInstallationArgs)(nil)).Elem()
}

type SyntheticMonitoringInstallationInput interface {
	pulumi.Input

	ToSyntheticMonitoringInstallationOutput() SyntheticMonitoringInstallationOutput
	ToSyntheticMonitoringInstallationOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationOutput
}

func (*SyntheticMonitoringInstallation) ElementType() reflect.Type {
	return reflect.TypeOf((**SyntheticMonitoringInstallation)(nil)).Elem()
}

func (i *SyntheticMonitoringInstallation) ToSyntheticMonitoringInstallationOutput() SyntheticMonitoringInstallationOutput {
	return i.ToSyntheticMonitoringInstallationOutputWithContext(context.Background())
}

func (i *SyntheticMonitoringInstallation) ToSyntheticMonitoringInstallationOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringInstallationOutput)
}

func (i *SyntheticMonitoringInstallation) ToOutput(ctx context.Context) pulumix.Output[*SyntheticMonitoringInstallation] {
	return pulumix.Output[*SyntheticMonitoringInstallation]{
		OutputState: i.ToSyntheticMonitoringInstallationOutputWithContext(ctx).OutputState,
	}
}

// SyntheticMonitoringInstallationArrayInput is an input type that accepts SyntheticMonitoringInstallationArray and SyntheticMonitoringInstallationArrayOutput values.
// You can construct a concrete instance of `SyntheticMonitoringInstallationArrayInput` via:
//
//	SyntheticMonitoringInstallationArray{ SyntheticMonitoringInstallationArgs{...} }
type SyntheticMonitoringInstallationArrayInput interface {
	pulumi.Input

	ToSyntheticMonitoringInstallationArrayOutput() SyntheticMonitoringInstallationArrayOutput
	ToSyntheticMonitoringInstallationArrayOutputWithContext(context.Context) SyntheticMonitoringInstallationArrayOutput
}

type SyntheticMonitoringInstallationArray []SyntheticMonitoringInstallationInput

func (SyntheticMonitoringInstallationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyntheticMonitoringInstallation)(nil)).Elem()
}

func (i SyntheticMonitoringInstallationArray) ToSyntheticMonitoringInstallationArrayOutput() SyntheticMonitoringInstallationArrayOutput {
	return i.ToSyntheticMonitoringInstallationArrayOutputWithContext(context.Background())
}

func (i SyntheticMonitoringInstallationArray) ToSyntheticMonitoringInstallationArrayOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringInstallationArrayOutput)
}

func (i SyntheticMonitoringInstallationArray) ToOutput(ctx context.Context) pulumix.Output[[]*SyntheticMonitoringInstallation] {
	return pulumix.Output[[]*SyntheticMonitoringInstallation]{
		OutputState: i.ToSyntheticMonitoringInstallationArrayOutputWithContext(ctx).OutputState,
	}
}

// SyntheticMonitoringInstallationMapInput is an input type that accepts SyntheticMonitoringInstallationMap and SyntheticMonitoringInstallationMapOutput values.
// You can construct a concrete instance of `SyntheticMonitoringInstallationMapInput` via:
//
//	SyntheticMonitoringInstallationMap{ "key": SyntheticMonitoringInstallationArgs{...} }
type SyntheticMonitoringInstallationMapInput interface {
	pulumi.Input

	ToSyntheticMonitoringInstallationMapOutput() SyntheticMonitoringInstallationMapOutput
	ToSyntheticMonitoringInstallationMapOutputWithContext(context.Context) SyntheticMonitoringInstallationMapOutput
}

type SyntheticMonitoringInstallationMap map[string]SyntheticMonitoringInstallationInput

func (SyntheticMonitoringInstallationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyntheticMonitoringInstallation)(nil)).Elem()
}

func (i SyntheticMonitoringInstallationMap) ToSyntheticMonitoringInstallationMapOutput() SyntheticMonitoringInstallationMapOutput {
	return i.ToSyntheticMonitoringInstallationMapOutputWithContext(context.Background())
}

func (i SyntheticMonitoringInstallationMap) ToSyntheticMonitoringInstallationMapOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringInstallationMapOutput)
}

func (i SyntheticMonitoringInstallationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SyntheticMonitoringInstallation] {
	return pulumix.Output[map[string]*SyntheticMonitoringInstallation]{
		OutputState: i.ToSyntheticMonitoringInstallationMapOutputWithContext(ctx).OutputState,
	}
}

type SyntheticMonitoringInstallationOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringInstallationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyntheticMonitoringInstallation)(nil)).Elem()
}

func (o SyntheticMonitoringInstallationOutput) ToSyntheticMonitoringInstallationOutput() SyntheticMonitoringInstallationOutput {
	return o
}

func (o SyntheticMonitoringInstallationOutput) ToSyntheticMonitoringInstallationOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationOutput {
	return o
}

func (o SyntheticMonitoringInstallationOutput) ToOutput(ctx context.Context) pulumix.Output[*SyntheticMonitoringInstallation] {
	return pulumix.Output[*SyntheticMonitoringInstallation]{
		OutputState: o.OutputState,
	}
}

// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
func (o SyntheticMonitoringInstallationOutput) MetricsPublisherKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringInstallation) pulumi.StringOutput { return v.MetricsPublisherKey }).(pulumi.StringOutput)
}

// Generated token to access the SM API.
func (o SyntheticMonitoringInstallationOutput) SmAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringInstallation) pulumi.StringOutput { return v.SmAccessToken }).(pulumi.StringOutput)
}

// The ID or slug of the stack to install SM on.
func (o SyntheticMonitoringInstallationOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringInstallation) pulumi.StringOutput { return v.StackId }).(pulumi.StringOutput)
}

// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
func (o SyntheticMonitoringInstallationOutput) StackSmApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringInstallation) pulumi.StringOutput { return v.StackSmApiUrl }).(pulumi.StringOutput)
}

type SyntheticMonitoringInstallationArrayOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringInstallationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyntheticMonitoringInstallation)(nil)).Elem()
}

func (o SyntheticMonitoringInstallationArrayOutput) ToSyntheticMonitoringInstallationArrayOutput() SyntheticMonitoringInstallationArrayOutput {
	return o
}

func (o SyntheticMonitoringInstallationArrayOutput) ToSyntheticMonitoringInstallationArrayOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationArrayOutput {
	return o
}

func (o SyntheticMonitoringInstallationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SyntheticMonitoringInstallation] {
	return pulumix.Output[[]*SyntheticMonitoringInstallation]{
		OutputState: o.OutputState,
	}
}

func (o SyntheticMonitoringInstallationArrayOutput) Index(i pulumi.IntInput) SyntheticMonitoringInstallationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyntheticMonitoringInstallation {
		return vs[0].([]*SyntheticMonitoringInstallation)[vs[1].(int)]
	}).(SyntheticMonitoringInstallationOutput)
}

type SyntheticMonitoringInstallationMapOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringInstallationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyntheticMonitoringInstallation)(nil)).Elem()
}

func (o SyntheticMonitoringInstallationMapOutput) ToSyntheticMonitoringInstallationMapOutput() SyntheticMonitoringInstallationMapOutput {
	return o
}

func (o SyntheticMonitoringInstallationMapOutput) ToSyntheticMonitoringInstallationMapOutputWithContext(ctx context.Context) SyntheticMonitoringInstallationMapOutput {
	return o
}

func (o SyntheticMonitoringInstallationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SyntheticMonitoringInstallation] {
	return pulumix.Output[map[string]*SyntheticMonitoringInstallation]{
		OutputState: o.OutputState,
	}
}

func (o SyntheticMonitoringInstallationMapOutput) MapIndex(k pulumi.StringInput) SyntheticMonitoringInstallationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyntheticMonitoringInstallation {
		return vs[0].(map[string]*SyntheticMonitoringInstallation)[vs[1].(string)]
	}).(SyntheticMonitoringInstallationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringInstallationInput)(nil)).Elem(), &SyntheticMonitoringInstallation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringInstallationArrayInput)(nil)).Elem(), SyntheticMonitoringInstallationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringInstallationMapInput)(nil)).Elem(), SyntheticMonitoringInstallationMap{})
	pulumi.RegisterOutputType(SyntheticMonitoringInstallationOutput{})
	pulumi.RegisterOutputType(SyntheticMonitoringInstallationArrayOutput{})
	pulumi.RegisterOutputType(SyntheticMonitoringInstallationMapOutput{})
}
