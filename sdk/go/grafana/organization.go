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

// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/organization-management/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/org/)
//
// This resource represents an instance-scoped resource and uses Grafana's admin APIs.
// It does not work with API tokens or service accounts which are org-scoped.
// You must use basic auth.
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
//			_, err := grafana.NewOrganization(ctx, "test", &grafana.OrganizationArgs{
//				AdminUser: pulumi.String("admin"),
//				Admins: pulumi.StringArray{
//					pulumi.String("admin@example.com"),
//				},
//				CreateUsers: pulumi.Bool(true),
//				Editors: pulumi.StringArray{
//					pulumi.String("editor-01@example.com"),
//					pulumi.String("editor-02@example.com"),
//				},
//				Viewers: pulumi.StringArray{
//					pulumi.String("viewer-01@example.com"),
//					pulumi.String("viewer-02@example.com"),
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
//
// ## Import
//
// ```sh
//
//	$ pulumi import grafana:index/organization:Organization org_name {{org_id}}
//
// ```
type Organization struct {
	pulumi.CustomResourceState

	// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
	// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
	// this parameter keeps Terraform from removing it from organizations.
	AdminUser pulumi.StringPtrOutput `pulumi:"adminUser"`
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Admins pulumi.StringArrayOutput `pulumi:"admins"`
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers pulumi.BoolPtrOutput `pulumi:"createUsers"`
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Editors pulumi.StringArrayOutput `pulumi:"editors"`
	// The display name for the Grafana organization created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization id assigned to this organization by Grafana.
	OrgId pulumi.IntOutput `pulumi:"orgId"`
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Viewers pulumi.StringArrayOutput `pulumi:"viewers"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		args = &OrganizationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Organization
	err := ctx.RegisterResource("grafana:index/organization:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("grafana:index/organization:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
	// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
	// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
	// this parameter keeps Terraform from removing it from organizations.
	AdminUser *string `pulumi:"adminUser"`
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Admins []string `pulumi:"admins"`
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers *bool `pulumi:"createUsers"`
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Editors []string `pulumi:"editors"`
	// The display name for the Grafana organization created.
	Name *string `pulumi:"name"`
	// The organization id assigned to this organization by Grafana.
	OrgId *int `pulumi:"orgId"`
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Viewers []string `pulumi:"viewers"`
}

type OrganizationState struct {
	// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
	// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
	// this parameter keeps Terraform from removing it from organizations.
	AdminUser pulumi.StringPtrInput
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Admins pulumi.StringArrayInput
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers pulumi.BoolPtrInput
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Editors pulumi.StringArrayInput
	// The display name for the Grafana organization created.
	Name pulumi.StringPtrInput
	// The organization id assigned to this organization by Grafana.
	OrgId pulumi.IntPtrInput
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Viewers pulumi.StringArrayInput
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
	// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
	// this parameter keeps Terraform from removing it from organizations.
	AdminUser *string `pulumi:"adminUser"`
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Admins []string `pulumi:"admins"`
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers *bool `pulumi:"createUsers"`
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Editors []string `pulumi:"editors"`
	// The display name for the Grafana organization created.
	Name *string `pulumi:"name"`
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Viewers []string `pulumi:"viewers"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
	// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
	// this parameter keeps Terraform from removing it from organizations.
	AdminUser pulumi.StringPtrInput
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Admins pulumi.StringArrayInput
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers pulumi.BoolPtrInput
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Editors pulumi.StringArrayInput
	// The display name for the Grafana organization created.
	Name pulumi.StringPtrInput
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	Viewers pulumi.StringArrayInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

func (i *Organization) ToOutput(ctx context.Context) pulumix.Output[*Organization] {
	return pulumix.Output[*Organization]{
		OutputState: i.ToOrganizationOutputWithContext(ctx).OutputState,
	}
}

// OrganizationArrayInput is an input type that accepts OrganizationArray and OrganizationArrayOutput values.
// You can construct a concrete instance of `OrganizationArrayInput` via:
//
//	OrganizationArray{ OrganizationArgs{...} }
type OrganizationArrayInput interface {
	pulumi.Input

	ToOrganizationArrayOutput() OrganizationArrayOutput
	ToOrganizationArrayOutputWithContext(context.Context) OrganizationArrayOutput
}

type OrganizationArray []OrganizationInput

func (OrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Organization)(nil)).Elem()
}

func (i OrganizationArray) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return i.ToOrganizationArrayOutputWithContext(context.Background())
}

func (i OrganizationArray) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationArrayOutput)
}

func (i OrganizationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Organization] {
	return pulumix.Output[[]*Organization]{
		OutputState: i.ToOrganizationArrayOutputWithContext(ctx).OutputState,
	}
}

// OrganizationMapInput is an input type that accepts OrganizationMap and OrganizationMapOutput values.
// You can construct a concrete instance of `OrganizationMapInput` via:
//
//	OrganizationMap{ "key": OrganizationArgs{...} }
type OrganizationMapInput interface {
	pulumi.Input

	ToOrganizationMapOutput() OrganizationMapOutput
	ToOrganizationMapOutputWithContext(context.Context) OrganizationMapOutput
}

type OrganizationMap map[string]OrganizationInput

func (OrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Organization)(nil)).Elem()
}

func (i OrganizationMap) ToOrganizationMapOutput() OrganizationMapOutput {
	return i.ToOrganizationMapOutputWithContext(context.Background())
}

func (i OrganizationMap) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMapOutput)
}

func (i OrganizationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Organization] {
	return pulumix.Output[map[string]*Organization]{
		OutputState: i.ToOrganizationMapOutputWithContext(ctx).OutputState,
	}
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOutput(ctx context.Context) pulumix.Output[*Organization] {
	return pulumix.Output[*Organization]{
		OutputState: o.OutputState,
	}
}

// The login name of the configured default admin user for the Grafana installation. If unset, this value defaults to
// admin, the Grafana default. Grafana adds the default admin user to all organizations automatically upon creation, and
// this parameter keeps Terraform from removing it from organizations.
func (o OrganizationOutput) AdminUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.AdminUser }).(pulumi.StringPtrOutput)
}

// A list of email addresses corresponding to users who should be given admin
// access to the organization. Note: users specified here must already exist in
// Grafana unless 'create_users' is set to true.
func (o OrganizationOutput) Admins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringArrayOutput { return v.Admins }).(pulumi.StringArrayOutput)
}

// Whether or not to create Grafana users specified in the organization's
// membership if they don't already exist in Grafana. If unspecified, this
// parameter defaults to true, creating placeholder users with the name, login,
// and email set to the email of the user, and a random password. Setting this
// option to false will cause an error to be thrown for any users that do not
// already exist in Grafana.
// Defaults to `true`.
func (o OrganizationOutput) CreateUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.BoolPtrOutput { return v.CreateUsers }).(pulumi.BoolPtrOutput)
}

// A list of email addresses corresponding to users who should be given editor
// access to the organization. Note: users specified here must already exist in
// Grafana unless 'create_users' is set to true.
func (o OrganizationOutput) Editors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringArrayOutput { return v.Editors }).(pulumi.StringArrayOutput)
}

// The display name for the Grafana organization created.
func (o OrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization id assigned to this organization by Grafana.
func (o OrganizationOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v *Organization) pulumi.IntOutput { return v.OrgId }).(pulumi.IntOutput)
}

// A list of email addresses corresponding to users who should be given viewer
// access to the organization. Note: users specified here must already exist in
// Grafana unless 'create_users' is set to true.
func (o OrganizationOutput) Viewers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringArrayOutput { return v.Viewers }).(pulumi.StringArrayOutput)
}

type OrganizationArrayOutput struct{ *pulumi.OutputState }

func (OrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Organization)(nil)).Elem()
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Organization] {
	return pulumix.Output[[]*Organization]{
		OutputState: o.OutputState,
	}
}

func (o OrganizationArrayOutput) Index(i pulumi.IntInput) OrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Organization {
		return vs[0].([]*Organization)[vs[1].(int)]
	}).(OrganizationOutput)
}

type OrganizationMapOutput struct{ *pulumi.OutputState }

func (OrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Organization)(nil)).Elem()
}

func (o OrganizationMapOutput) ToOrganizationMapOutput() OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Organization] {
	return pulumix.Output[map[string]*Organization]{
		OutputState: o.OutputState,
	}
}

func (o OrganizationMapOutput) MapIndex(k pulumi.StringInput) OrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Organization {
		return vs[0].(map[string]*Organization)[vs[1].(string)]
	}).(OrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationInput)(nil)).Elem(), &Organization{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationArrayInput)(nil)).Elem(), OrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMapInput)(nil)).Elem(), OrganizationMap{})
	pulumi.RegisterOutputType(OrganizationOutput{})
	pulumi.RegisterOutputType(OrganizationArrayOutput{})
	pulumi.RegisterOutputType(OrganizationMapOutput{})
}
