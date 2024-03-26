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

// Manages Grafana SSO Settings for OAuth2. SAML support will be added soon.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-authentication/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/sso-settings/)
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
//			_, err := grafana.NewSsoSettings(ctx, "githubSsoSettings", &grafana.SsoSettingsArgs{
//				Oauth2Settings: &grafana.SsoSettingsOauth2SettingsArgs{
//					AllowedOrganizations: pulumi.String("organization1,organization2"),
//					ClientId:             pulumi.String("github_client_id"),
//					ClientSecret:         pulumi.String("github_client_secret"),
//					TeamIds:              pulumi.String("12,50,123"),
//				},
//				ProviderName: pulumi.String("github"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SsoSettings struct {
	pulumi.CustomResourceState

	// The SSO settings set.
	Oauth2Settings SsoSettingsOauth2SettingsOutput `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
}

// NewSsoSettings registers a new resource with the given unique name, arguments, and options.
func NewSsoSettings(ctx *pulumi.Context,
	name string, args *SsoSettingsArgs, opts ...pulumi.ResourceOption) (*SsoSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Oauth2Settings == nil {
		return nil, errors.New("invalid value for required argument 'Oauth2Settings'")
	}
	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SsoSettings
	err := ctx.RegisterResource("grafana:index/ssoSettings:SsoSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSsoSettings gets an existing SsoSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSsoSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SsoSettingsState, opts ...pulumi.ResourceOption) (*SsoSettings, error) {
	var resource SsoSettings
	err := ctx.ReadResource("grafana:index/ssoSettings:SsoSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SsoSettings resources.
type ssoSettingsState struct {
	// The SSO settings set.
	Oauth2Settings *SsoSettingsOauth2Settings `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
	ProviderName *string `pulumi:"providerName"`
}

type SsoSettingsState struct {
	// The SSO settings set.
	Oauth2Settings SsoSettingsOauth2SettingsPtrInput
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
	ProviderName pulumi.StringPtrInput
}

func (SsoSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ssoSettingsState)(nil)).Elem()
}

type ssoSettingsArgs struct {
	// The SSO settings set.
	Oauth2Settings SsoSettingsOauth2Settings `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
	ProviderName string `pulumi:"providerName"`
}

// The set of arguments for constructing a SsoSettings resource.
type SsoSettingsArgs struct {
	// The SSO settings set.
	Oauth2Settings SsoSettingsOauth2SettingsInput
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
	ProviderName pulumi.StringInput
}

func (SsoSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ssoSettingsArgs)(nil)).Elem()
}

type SsoSettingsInput interface {
	pulumi.Input

	ToSsoSettingsOutput() SsoSettingsOutput
	ToSsoSettingsOutputWithContext(ctx context.Context) SsoSettingsOutput
}

func (*SsoSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SsoSettings)(nil)).Elem()
}

func (i *SsoSettings) ToSsoSettingsOutput() SsoSettingsOutput {
	return i.ToSsoSettingsOutputWithContext(context.Background())
}

func (i *SsoSettings) ToSsoSettingsOutputWithContext(ctx context.Context) SsoSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoSettingsOutput)
}

func (i *SsoSettings) ToOutput(ctx context.Context) pulumix.Output[*SsoSettings] {
	return pulumix.Output[*SsoSettings]{
		OutputState: i.ToSsoSettingsOutputWithContext(ctx).OutputState,
	}
}

// SsoSettingsArrayInput is an input type that accepts SsoSettingsArray and SsoSettingsArrayOutput values.
// You can construct a concrete instance of `SsoSettingsArrayInput` via:
//
//	SsoSettingsArray{ SsoSettingsArgs{...} }
type SsoSettingsArrayInput interface {
	pulumi.Input

	ToSsoSettingsArrayOutput() SsoSettingsArrayOutput
	ToSsoSettingsArrayOutputWithContext(context.Context) SsoSettingsArrayOutput
}

type SsoSettingsArray []SsoSettingsInput

func (SsoSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SsoSettings)(nil)).Elem()
}

func (i SsoSettingsArray) ToSsoSettingsArrayOutput() SsoSettingsArrayOutput {
	return i.ToSsoSettingsArrayOutputWithContext(context.Background())
}

func (i SsoSettingsArray) ToSsoSettingsArrayOutputWithContext(ctx context.Context) SsoSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoSettingsArrayOutput)
}

func (i SsoSettingsArray) ToOutput(ctx context.Context) pulumix.Output[[]*SsoSettings] {
	return pulumix.Output[[]*SsoSettings]{
		OutputState: i.ToSsoSettingsArrayOutputWithContext(ctx).OutputState,
	}
}

// SsoSettingsMapInput is an input type that accepts SsoSettingsMap and SsoSettingsMapOutput values.
// You can construct a concrete instance of `SsoSettingsMapInput` via:
//
//	SsoSettingsMap{ "key": SsoSettingsArgs{...} }
type SsoSettingsMapInput interface {
	pulumi.Input

	ToSsoSettingsMapOutput() SsoSettingsMapOutput
	ToSsoSettingsMapOutputWithContext(context.Context) SsoSettingsMapOutput
}

type SsoSettingsMap map[string]SsoSettingsInput

func (SsoSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SsoSettings)(nil)).Elem()
}

func (i SsoSettingsMap) ToSsoSettingsMapOutput() SsoSettingsMapOutput {
	return i.ToSsoSettingsMapOutputWithContext(context.Background())
}

func (i SsoSettingsMap) ToSsoSettingsMapOutputWithContext(ctx context.Context) SsoSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SsoSettingsMapOutput)
}

func (i SsoSettingsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SsoSettings] {
	return pulumix.Output[map[string]*SsoSettings]{
		OutputState: i.ToSsoSettingsMapOutputWithContext(ctx).OutputState,
	}
}

type SsoSettingsOutput struct{ *pulumi.OutputState }

func (SsoSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsoSettings)(nil)).Elem()
}

func (o SsoSettingsOutput) ToSsoSettingsOutput() SsoSettingsOutput {
	return o
}

func (o SsoSettingsOutput) ToSsoSettingsOutputWithContext(ctx context.Context) SsoSettingsOutput {
	return o
}

func (o SsoSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[*SsoSettings] {
	return pulumix.Output[*SsoSettings]{
		OutputState: o.OutputState,
	}
}

// The SSO settings set.
func (o SsoSettingsOutput) Oauth2Settings() SsoSettingsOauth2SettingsOutput {
	return o.ApplyT(func(v *SsoSettings) SsoSettingsOauth2SettingsOutput { return v.Oauth2Settings }).(SsoSettingsOauth2SettingsOutput)
}

// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth.
func (o SsoSettingsOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *SsoSettings) pulumi.StringOutput { return v.ProviderName }).(pulumi.StringOutput)
}

type SsoSettingsArrayOutput struct{ *pulumi.OutputState }

func (SsoSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SsoSettings)(nil)).Elem()
}

func (o SsoSettingsArrayOutput) ToSsoSettingsArrayOutput() SsoSettingsArrayOutput {
	return o
}

func (o SsoSettingsArrayOutput) ToSsoSettingsArrayOutputWithContext(ctx context.Context) SsoSettingsArrayOutput {
	return o
}

func (o SsoSettingsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SsoSettings] {
	return pulumix.Output[[]*SsoSettings]{
		OutputState: o.OutputState,
	}
}

func (o SsoSettingsArrayOutput) Index(i pulumi.IntInput) SsoSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SsoSettings {
		return vs[0].([]*SsoSettings)[vs[1].(int)]
	}).(SsoSettingsOutput)
}

type SsoSettingsMapOutput struct{ *pulumi.OutputState }

func (SsoSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SsoSettings)(nil)).Elem()
}

func (o SsoSettingsMapOutput) ToSsoSettingsMapOutput() SsoSettingsMapOutput {
	return o
}

func (o SsoSettingsMapOutput) ToSsoSettingsMapOutputWithContext(ctx context.Context) SsoSettingsMapOutput {
	return o
}

func (o SsoSettingsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SsoSettings] {
	return pulumix.Output[map[string]*SsoSettings]{
		OutputState: o.OutputState,
	}
}

func (o SsoSettingsMapOutput) MapIndex(k pulumi.StringInput) SsoSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SsoSettings {
		return vs[0].(map[string]*SsoSettings)[vs[1].(string)]
	}).(SsoSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SsoSettingsInput)(nil)).Elem(), &SsoSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SsoSettingsArrayInput)(nil)).Elem(), SsoSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SsoSettingsMapInput)(nil)).Elem(), SsoSettingsMap{})
	pulumi.RegisterOutputType(SsoSettingsOutput{})
	pulumi.RegisterOutputType(SsoSettingsArrayOutput{})
	pulumi.RegisterOutputType(SsoSettingsMapOutput{})
}