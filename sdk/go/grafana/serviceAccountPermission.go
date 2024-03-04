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

// **Note:** This resource is available from Grafana 9.2.4 onwards.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/#manage-users-and-teams-permissions-for-a-service-account-in-grafana)
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
//			test, err := grafana.NewServiceAccount(ctx, "test", &grafana.ServiceAccountArgs{
//				Role:       pulumi.String("Editor"),
//				IsDisabled: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			testTeam, err := grafana.NewTeam(ctx, "testTeam", nil)
//			if err != nil {
//				return err
//			}
//			testUser, err := grafana.NewUser(ctx, "testUser", &grafana.UserArgs{
//				Email:    pulumi.String("tf_user@test.com"),
//				Login:    pulumi.String("tf_user@test.com"),
//				Password: pulumi.String("password"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewServiceAccountPermission(ctx, "testPermissions", &grafana.ServiceAccountPermissionArgs{
//				ServiceAccountId: test.ID(),
//				Permissions: grafana.ServiceAccountPermissionPermissionArray{
//					&grafana.ServiceAccountPermissionPermissionArgs{
//						UserId:     testUser.ID(),
//						Permission: pulumi.String("Edit"),
//					},
//					&grafana.ServiceAccountPermissionPermissionArgs{
//						TeamId:     testTeam.ID(),
//						Permission: pulumi.String("Admin"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ServiceAccountPermission struct {
	pulumi.CustomResourceState

	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions ServiceAccountPermissionPermissionArrayOutput `pulumi:"permissions"`
	// The id of the service account.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewServiceAccountPermission registers a new resource with the given unique name, arguments, and options.
func NewServiceAccountPermission(ctx *pulumi.Context,
	name string, args *ServiceAccountPermissionArgs, opts ...pulumi.ResourceOption) (*ServiceAccountPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAccountPermission
	err := ctx.RegisterResource("grafana:index/serviceAccountPermission:ServiceAccountPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccountPermission gets an existing ServiceAccountPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccountPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountPermissionState, opts ...pulumi.ResourceOption) (*ServiceAccountPermission, error) {
	var resource ServiceAccountPermission
	err := ctx.ReadResource("grafana:index/serviceAccountPermission:ServiceAccountPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccountPermission resources.
type serviceAccountPermissionState struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []ServiceAccountPermissionPermission `pulumi:"permissions"`
	// The id of the service account.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type ServiceAccountPermissionState struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions ServiceAccountPermissionPermissionArrayInput
	// The id of the service account.
	ServiceAccountId pulumi.StringPtrInput
}

func (ServiceAccountPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountPermissionState)(nil)).Elem()
}

type serviceAccountPermissionArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []ServiceAccountPermissionPermission `pulumi:"permissions"`
	// The id of the service account.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a ServiceAccountPermission resource.
type ServiceAccountPermissionArgs struct {
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions ServiceAccountPermissionPermissionArrayInput
	// The id of the service account.
	ServiceAccountId pulumi.StringInput
}

func (ServiceAccountPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountPermissionArgs)(nil)).Elem()
}

type ServiceAccountPermissionInput interface {
	pulumi.Input

	ToServiceAccountPermissionOutput() ServiceAccountPermissionOutput
	ToServiceAccountPermissionOutputWithContext(ctx context.Context) ServiceAccountPermissionOutput
}

func (*ServiceAccountPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountPermission)(nil)).Elem()
}

func (i *ServiceAccountPermission) ToServiceAccountPermissionOutput() ServiceAccountPermissionOutput {
	return i.ToServiceAccountPermissionOutputWithContext(context.Background())
}

func (i *ServiceAccountPermission) ToServiceAccountPermissionOutputWithContext(ctx context.Context) ServiceAccountPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPermissionOutput)
}

func (i *ServiceAccountPermission) ToOutput(ctx context.Context) pulumix.Output[*ServiceAccountPermission] {
	return pulumix.Output[*ServiceAccountPermission]{
		OutputState: i.ToServiceAccountPermissionOutputWithContext(ctx).OutputState,
	}
}

// ServiceAccountPermissionArrayInput is an input type that accepts ServiceAccountPermissionArray and ServiceAccountPermissionArrayOutput values.
// You can construct a concrete instance of `ServiceAccountPermissionArrayInput` via:
//
//	ServiceAccountPermissionArray{ ServiceAccountPermissionArgs{...} }
type ServiceAccountPermissionArrayInput interface {
	pulumi.Input

	ToServiceAccountPermissionArrayOutput() ServiceAccountPermissionArrayOutput
	ToServiceAccountPermissionArrayOutputWithContext(context.Context) ServiceAccountPermissionArrayOutput
}

type ServiceAccountPermissionArray []ServiceAccountPermissionInput

func (ServiceAccountPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountPermission)(nil)).Elem()
}

func (i ServiceAccountPermissionArray) ToServiceAccountPermissionArrayOutput() ServiceAccountPermissionArrayOutput {
	return i.ToServiceAccountPermissionArrayOutputWithContext(context.Background())
}

func (i ServiceAccountPermissionArray) ToServiceAccountPermissionArrayOutputWithContext(ctx context.Context) ServiceAccountPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPermissionArrayOutput)
}

func (i ServiceAccountPermissionArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceAccountPermission] {
	return pulumix.Output[[]*ServiceAccountPermission]{
		OutputState: i.ToServiceAccountPermissionArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceAccountPermissionMapInput is an input type that accepts ServiceAccountPermissionMap and ServiceAccountPermissionMapOutput values.
// You can construct a concrete instance of `ServiceAccountPermissionMapInput` via:
//
//	ServiceAccountPermissionMap{ "key": ServiceAccountPermissionArgs{...} }
type ServiceAccountPermissionMapInput interface {
	pulumi.Input

	ToServiceAccountPermissionMapOutput() ServiceAccountPermissionMapOutput
	ToServiceAccountPermissionMapOutputWithContext(context.Context) ServiceAccountPermissionMapOutput
}

type ServiceAccountPermissionMap map[string]ServiceAccountPermissionInput

func (ServiceAccountPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountPermission)(nil)).Elem()
}

func (i ServiceAccountPermissionMap) ToServiceAccountPermissionMapOutput() ServiceAccountPermissionMapOutput {
	return i.ToServiceAccountPermissionMapOutputWithContext(context.Background())
}

func (i ServiceAccountPermissionMap) ToServiceAccountPermissionMapOutputWithContext(ctx context.Context) ServiceAccountPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountPermissionMapOutput)
}

func (i ServiceAccountPermissionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceAccountPermission] {
	return pulumix.Output[map[string]*ServiceAccountPermission]{
		OutputState: i.ToServiceAccountPermissionMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceAccountPermissionOutput struct{ *pulumi.OutputState }

func (ServiceAccountPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccountPermission)(nil)).Elem()
}

func (o ServiceAccountPermissionOutput) ToServiceAccountPermissionOutput() ServiceAccountPermissionOutput {
	return o
}

func (o ServiceAccountPermissionOutput) ToServiceAccountPermissionOutputWithContext(ctx context.Context) ServiceAccountPermissionOutput {
	return o
}

func (o ServiceAccountPermissionOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceAccountPermission] {
	return pulumix.Output[*ServiceAccountPermission]{
		OutputState: o.OutputState,
	}
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o ServiceAccountPermissionOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAccountPermission) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The permission items to add/update. Items that are omitted from the list will be removed.
func (o ServiceAccountPermissionOutput) Permissions() ServiceAccountPermissionPermissionArrayOutput {
	return o.ApplyT(func(v *ServiceAccountPermission) ServiceAccountPermissionPermissionArrayOutput { return v.Permissions }).(ServiceAccountPermissionPermissionArrayOutput)
}

// The id of the service account.
func (o ServiceAccountPermissionOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccountPermission) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

type ServiceAccountPermissionArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccountPermission)(nil)).Elem()
}

func (o ServiceAccountPermissionArrayOutput) ToServiceAccountPermissionArrayOutput() ServiceAccountPermissionArrayOutput {
	return o
}

func (o ServiceAccountPermissionArrayOutput) ToServiceAccountPermissionArrayOutputWithContext(ctx context.Context) ServiceAccountPermissionArrayOutput {
	return o
}

func (o ServiceAccountPermissionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServiceAccountPermission] {
	return pulumix.Output[[]*ServiceAccountPermission]{
		OutputState: o.OutputState,
	}
}

func (o ServiceAccountPermissionArrayOutput) Index(i pulumi.IntInput) ServiceAccountPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccountPermission {
		return vs[0].([]*ServiceAccountPermission)[vs[1].(int)]
	}).(ServiceAccountPermissionOutput)
}

type ServiceAccountPermissionMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccountPermission)(nil)).Elem()
}

func (o ServiceAccountPermissionMapOutput) ToServiceAccountPermissionMapOutput() ServiceAccountPermissionMapOutput {
	return o
}

func (o ServiceAccountPermissionMapOutput) ToServiceAccountPermissionMapOutputWithContext(ctx context.Context) ServiceAccountPermissionMapOutput {
	return o
}

func (o ServiceAccountPermissionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServiceAccountPermission] {
	return pulumix.Output[map[string]*ServiceAccountPermission]{
		OutputState: o.OutputState,
	}
}

func (o ServiceAccountPermissionMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccountPermission {
		return vs[0].(map[string]*ServiceAccountPermission)[vs[1].(string)]
	}).(ServiceAccountPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPermissionInput)(nil)).Elem(), &ServiceAccountPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPermissionArrayInput)(nil)).Elem(), ServiceAccountPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountPermissionMapInput)(nil)).Elem(), ServiceAccountPermissionMap{})
	pulumi.RegisterOutputType(ServiceAccountPermissionOutput{})
	pulumi.RegisterOutputType(ServiceAccountPermissionArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountPermissionMapOutput{})
}
