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
// $ pulumi import grafana:index/serviceAccount:ServiceAccount name "{{ id }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/serviceAccount:ServiceAccount name "{{ orgID }}:{{ id }}"
// ```
//
// Deprecated: grafana.index/serviceaccount.ServiceAccount has been deprecated in favor of grafana.oss/serviceaccount.ServiceAccount
type ServiceAccount struct {
	pulumi.CustomResourceState

	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// The name of the service account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The basic role of the service account in the organization.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewServiceAccount(ctx *pulumi.Context,
	name string, args *ServiceAccountArgs, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/serviceAccount:ServiceAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAccount
	err := ctx.RegisterResource("grafana:index/serviceAccount:ServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccount gets an existing ServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountState, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	var resource ServiceAccount
	err := ctx.ReadResource("grafana:index/serviceAccount:ServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccount resources.
type serviceAccountState struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The basic role of the service account in the organization.
	Role *string `pulumi:"role"`
}

type ServiceAccountState struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrInput
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The basic role of the service account in the organization.
	Role pulumi.StringPtrInput
}

func (ServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountState)(nil)).Elem()
}

type serviceAccountArgs struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of the service account.
	Name *string `pulumi:"name"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The basic role of the service account in the organization.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ServiceAccount resource.
type ServiceAccountArgs struct {
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled pulumi.BoolPtrInput
	// The name of the service account.
	Name pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The basic role of the service account in the organization.
	Role pulumi.StringInput
}

func (ServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountArgs)(nil)).Elem()
}

type ServiceAccountInput interface {
	pulumi.Input

	ToServiceAccountOutput() ServiceAccountOutput
	ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput
}

func (*ServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (i *ServiceAccount) ToServiceAccountOutput() ServiceAccountOutput {
	return i.ToServiceAccountOutputWithContext(context.Background())
}

func (i *ServiceAccount) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountOutput)
}

// ServiceAccountArrayInput is an input type that accepts ServiceAccountArray and ServiceAccountArrayOutput values.
// You can construct a concrete instance of `ServiceAccountArrayInput` via:
//
//	ServiceAccountArray{ ServiceAccountArgs{...} }
type ServiceAccountArrayInput interface {
	pulumi.Input

	ToServiceAccountArrayOutput() ServiceAccountArrayOutput
	ToServiceAccountArrayOutputWithContext(context.Context) ServiceAccountArrayOutput
}

type ServiceAccountArray []ServiceAccountInput

func (ServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountArray) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return i.ToServiceAccountArrayOutputWithContext(context.Background())
}

func (i ServiceAccountArray) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountArrayOutput)
}

// ServiceAccountMapInput is an input type that accepts ServiceAccountMap and ServiceAccountMapOutput values.
// You can construct a concrete instance of `ServiceAccountMapInput` via:
//
//	ServiceAccountMap{ "key": ServiceAccountArgs{...} }
type ServiceAccountMapInput interface {
	pulumi.Input

	ToServiceAccountMapOutput() ServiceAccountMapOutput
	ToServiceAccountMapOutputWithContext(context.Context) ServiceAccountMapOutput
}

type ServiceAccountMap map[string]ServiceAccountInput

func (ServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountMap) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return i.ToServiceAccountMapOutputWithContext(context.Background())
}

func (i ServiceAccountMap) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountMapOutput)
}

type ServiceAccountOutput struct{ *pulumi.OutputState }

func (ServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountOutput) ToServiceAccountOutput() ServiceAccountOutput {
	return o
}

func (o ServiceAccountOutput) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return o
}

// The disabled status for the service account. Defaults to `false`.
func (o ServiceAccountOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// The name of the service account.
func (o ServiceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o ServiceAccountOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The basic role of the service account in the organization.
func (o ServiceAccountOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type ServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) Index(i pulumi.IntInput) ServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].([]*ServiceAccount)[vs[1].(int)]
	}).(ServiceAccountOutput)
}

type ServiceAccountMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].(map[string]*ServiceAccount)[vs[1].(string)]
	}).(ServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountInput)(nil)).Elem(), &ServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountArrayInput)(nil)).Elem(), ServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountMapInput)(nil)).Elem(), ServiceAccountMap{})
	pulumi.RegisterOutputType(ServiceAccountOutput{})
	pulumi.RegisterOutputType(ServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountMapOutput{})
}
