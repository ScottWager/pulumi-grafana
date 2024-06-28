// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// **Note:** This resource is available only with Grafana 9.1+.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewServiceAccount(ctx, "admin", &oss.ServiceAccountArgs{
//				IsDisabled: pulumi.Bool(false),
//				Role:       pulumi.String("Admin"),
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
// $ pulumi import grafana:index/serviceAccountToken:ServiceAccountToken name "{{ id }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/serviceAccountToken:ServiceAccountToken name "{{ orgID }}:{{ id }}"
// ```
//
// Deprecated: grafana.index/serviceaccounttoken.ServiceAccountToken has been deprecated in favor of grafana.oss/serviceaccounttoken.ServiceAccountToken
type ServiceAccountToken struct {
	pulumi.CustomResourceState

	// The expiration date of the service account token.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The status of the service account token.
	HasExpired pulumi.BoolOutput `pulumi:"hasExpired"`
	// The key of the service account token.
	Key pulumi.StringOutput `pulumi:"key"`
	// The name of the service account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
	// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
	// expire.
	SecondsToLive pulumi.IntPtrOutput `pulumi:"secondsToLive"`
	// The ID of the service account to which the token belongs.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewServiceAccountToken registers a new resource with the given unique name, arguments, and options.
func NewServiceAccountToken(ctx *pulumi.Context,
	name string, args *ServiceAccountTokenArgs, opts ...pulumi.ResourceOption) (*ServiceAccountToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/serviceAccountToken:ServiceAccountToken"),
		},
	})
	opts = append(opts, aliases)
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAccountToken
	err := ctx.RegisterResource("grafana:index/serviceAccountToken:ServiceAccountToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccountToken gets an existing ServiceAccountToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccountToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountTokenState, opts ...pulumi.ResourceOption) (*ServiceAccountToken, error) {
	var resource ServiceAccountToken
	err := ctx.ReadResource("grafana:index/serviceAccountToken:ServiceAccountToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccountToken resources.
type serviceAccountTokenState struct {
	// The expiration date of the service account token.
	Expiration *string `pulumi:"expiration"`
	// The status of the service account token.
	HasExpired *bool `pulumi:"hasExpired"`
	// The key of the service account token.
	Key *string `pulumi:"key"`
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
	// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
	// expire.
	SecondsToLive *int `pulumi:"secondsToLive"`
	// The ID of the service account to which the token belongs.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type ServiceAccountTokenState struct {
	// The expiration date of the service account token.
	Expiration pulumi.StringPtrInput
	// The status of the service account token.
	HasExpired pulumi.BoolPtrInput
	// The key of the service account token.
	Key pulumi.StringPtrInput
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
	// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
	// expire.
	SecondsToLive pulumi.IntPtrInput
	// The ID of the service account to which the token belongs.
	ServiceAccountId pulumi.StringPtrInput
}

func (ServiceAccountTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountTokenState)(nil)).Elem()
}

type serviceAccountTokenArgs struct {
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
	// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
	// expire.
	SecondsToLive *int `pulumi:"secondsToLive"`
	// The ID of the service account to which the token belongs.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a ServiceAccountToken resource.
type ServiceAccountTokenArgs struct {
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
	// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
	// expire.
	SecondsToLive pulumi.IntPtrInput
	// The ID of the service account to which the token belongs.
	ServiceAccountId pulumi.StringInput
}

func (ServiceAccountTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountTokenArgs)(nil)).Elem()
}

type ServiceAccountTokenInput interface {
	pulumi.Input

	ToServiceAccountTokenOutput() ServiceAccountTokenOutput
	ToServiceAccountTokenOutputWithContext(ctx context.Context) ServiceAccountTokenOutput
}

func (*ServiceAccountToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountToken)(nil)).Elem()
}

func (i *ServiceAccountToken) ToServiceAccountTokenOutput() ServiceAccountTokenOutput {
	return i.ToServiceAccountTokenOutputWithContext(context.Background())
}

func (i *ServiceAccountToken) ToServiceAccountTokenOutputWithContext(ctx context.Context) ServiceAccountTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountTokenOutput)
}

// ServiceAccountTokenArrayInput is an input type that accepts ServiceAccountTokenArray and ServiceAccountTokenArrayOutput values.
// You can construct a concrete instance of `ServiceAccountTokenArrayInput` via:
//
//	ServiceAccountTokenArray{ ServiceAccountTokenArgs{...} }
type ServiceAccountTokenArrayInput interface {
	pulumi.Input

	ToServiceAccountTokenArrayOutput() ServiceAccountTokenArrayOutput
	ToServiceAccountTokenArrayOutputWithContext(context.Context) ServiceAccountTokenArrayOutput
}

type ServiceAccountTokenArray []ServiceAccountTokenInput

func (ServiceAccountTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountToken)(nil)).Elem()
}

func (i ServiceAccountTokenArray) ToServiceAccountTokenArrayOutput() ServiceAccountTokenArrayOutput {
	return i.ToServiceAccountTokenArrayOutputWithContext(context.Background())
}

func (i ServiceAccountTokenArray) ToServiceAccountTokenArrayOutputWithContext(ctx context.Context) ServiceAccountTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountTokenArrayOutput)
}

// ServiceAccountTokenMapInput is an input type that accepts ServiceAccountTokenMap and ServiceAccountTokenMapOutput values.
// You can construct a concrete instance of `ServiceAccountTokenMapInput` via:
//
//	ServiceAccountTokenMap{ "key": ServiceAccountTokenArgs{...} }
type ServiceAccountTokenMapInput interface {
	pulumi.Input

	ToServiceAccountTokenMapOutput() ServiceAccountTokenMapOutput
	ToServiceAccountTokenMapOutputWithContext(context.Context) ServiceAccountTokenMapOutput
}

type ServiceAccountTokenMap map[string]ServiceAccountTokenInput

func (ServiceAccountTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountToken)(nil)).Elem()
}

func (i ServiceAccountTokenMap) ToServiceAccountTokenMapOutput() ServiceAccountTokenMapOutput {
	return i.ToServiceAccountTokenMapOutputWithContext(context.Background())
}

func (i ServiceAccountTokenMap) ToServiceAccountTokenMapOutputWithContext(ctx context.Context) ServiceAccountTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountTokenMapOutput)
}

type ServiceAccountTokenOutput struct{ *pulumi.OutputState }

func (ServiceAccountTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountToken)(nil)).Elem()
}

func (o ServiceAccountTokenOutput) ToServiceAccountTokenOutput() ServiceAccountTokenOutput {
	return o
}

func (o ServiceAccountTokenOutput) ToServiceAccountTokenOutputWithContext(ctx context.Context) ServiceAccountTokenOutput {
	return o
}

// The expiration date of the service account token.
func (o ServiceAccountTokenOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// The status of the service account token.
func (o ServiceAccountTokenOutput) HasExpired() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.BoolOutput { return v.HasExpired }).(pulumi.BoolOutput)
}

// The key of the service account token.
func (o ServiceAccountTokenOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The name of the service account.
func (o ServiceAccountTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The key expiration in seconds. It is optional. If it is a positive number an expiration date for the key is set. If it
// is null, zero or is omitted completely (unless `apiKeyMaxSecondsToLive` configuration option is set) the key will never
// expire.
func (o ServiceAccountTokenOutput) SecondsToLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.IntPtrOutput { return v.SecondsToLive }).(pulumi.IntPtrOutput)
}

// The ID of the service account to which the token belongs.
func (o ServiceAccountTokenOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccountToken) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

type ServiceAccountTokenArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountToken)(nil)).Elem()
}

func (o ServiceAccountTokenArrayOutput) ToServiceAccountTokenArrayOutput() ServiceAccountTokenArrayOutput {
	return o
}

func (o ServiceAccountTokenArrayOutput) ToServiceAccountTokenArrayOutputWithContext(ctx context.Context) ServiceAccountTokenArrayOutput {
	return o
}

func (o ServiceAccountTokenArrayOutput) Index(i pulumi.IntInput) ServiceAccountTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccountToken {
		return vs[0].([]*ServiceAccountToken)[vs[1].(int)]
	}).(ServiceAccountTokenOutput)
}

type ServiceAccountTokenMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountToken)(nil)).Elem()
}

func (o ServiceAccountTokenMapOutput) ToServiceAccountTokenMapOutput() ServiceAccountTokenMapOutput {
	return o
}

func (o ServiceAccountTokenMapOutput) ToServiceAccountTokenMapOutputWithContext(ctx context.Context) ServiceAccountTokenMapOutput {
	return o
}

func (o ServiceAccountTokenMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccountToken {
		return vs[0].(map[string]*ServiceAccountToken)[vs[1].(string)]
	}).(ServiceAccountTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountTokenInput)(nil)).Elem(), &ServiceAccountToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountTokenArrayInput)(nil)).Elem(), ServiceAccountTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountTokenMapInput)(nil)).Elem(), ServiceAccountTokenMap{})
	pulumi.RegisterOutputType(ServiceAccountTokenOutput{})
	pulumi.RegisterOutputType(ServiceAccountTokenArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountTokenMapOutput{})
}
