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

// Manages Grafana SSO Settings for OAuth2 and SAML. Support for SAML is currently in preview, it will be available in Grafana Enterprise starting with v11.1.
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
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Configure SSO for GitHub using OAuth2
//			_, err := oss.NewSsoSettings(ctx, "github_sso_settings", &oss.SsoSettingsArgs{
//				ProviderName: pulumi.String("github"),
//				Oauth2Settings: &oss.SsoSettingsOauth2SettingsArgs{
//					Name:                 pulumi.String("Github"),
//					ClientId:             pulumi.String("<your GitHub app client id>"),
//					ClientSecret:         pulumi.String("<your GitHub app client secret>"),
//					AllowSignUp:          pulumi.Bool(true),
//					AutoLogin:            pulumi.Bool(false),
//					Scopes:               pulumi.String("user:email,read:org"),
//					TeamIds:              pulumi.String("150,300"),
//					AllowedOrganizations: pulumi.String("[\"My Organization\", \"Octocats\"]"),
//					AllowedDomains:       pulumi.String("mycompany.com mycompany.org"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// Configure SSO using generic OAuth2
//			_, err = oss.NewSsoSettings(ctx, "generic_sso_settings", &oss.SsoSettingsArgs{
//				ProviderName: pulumi.String("generic_oauth"),
//				Oauth2Settings: &oss.SsoSettingsOauth2SettingsArgs{
//					Name:            pulumi.String("Auth0"),
//					AuthUrl:         pulumi.String("https://<domain>/authorize"),
//					TokenUrl:        pulumi.String("https://<domain>/oauth/token"),
//					ApiUrl:          pulumi.String("https://<domain>/userinfo"),
//					ClientId:        pulumi.String("<client id>"),
//					ClientSecret:    pulumi.String("<client secret>"),
//					AllowSignUp:     pulumi.Bool(true),
//					AutoLogin:       pulumi.Bool(false),
//					Scopes:          pulumi.String("openid profile email offline_access"),
//					UsePkce:         pulumi.Bool(true),
//					UseRefreshToken: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// Configure SSO using SAML
//			_, err = oss.NewSsoSettings(ctx, "saml_sso_settings", &oss.SsoSettingsArgs{
//				ProviderName: pulumi.String("saml"),
//				SamlSettings: &oss.SsoSettingsSamlSettingsArgs{
//					AllowSignUp:             pulumi.Bool(true),
//					CertificatePath:         pulumi.String("devenv/docker/blocks/auth/saml-enterprise/cert.crt"),
//					PrivateKeyPath:          pulumi.String("devenv/docker/blocks/auth/saml-enterprise/key.pem"),
//					IdpMetadataUrl:          pulumi.String("https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml"),
//					SignatureAlgorithm:      pulumi.String("rsa-sha256"),
//					AssertionAttributeLogin: pulumi.String("login"),
//					AssertionAttributeEmail: pulumi.String("email"),
//					NameIdFormat:            pulumi.String("urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"),
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
// $ pulumi import grafana:index/ssoSettings:SsoSettings name "{{ provider }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/ssoSettings:SsoSettings name "{{ orgID }}:{{ provider }}"
// ```
//
// Deprecated: grafana.index/ssosettings.SsoSettings has been deprecated in favor of grafana.oss/ssosettings.SsoSettings
type SsoSettings struct {
	pulumi.CustomResourceState

	// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
	Oauth2Settings SsoSettingsOauth2SettingsPtrOutput `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	// The SAML settings set. Required for the saml provider.
	SamlSettings SsoSettingsSamlSettingsPtrOutput `pulumi:"samlSettings"`
}

// NewSsoSettings registers a new resource with the given unique name, arguments, and options.
func NewSsoSettings(ctx *pulumi.Context,
	name string, args *SsoSettingsArgs, opts ...pulumi.ResourceOption) (*SsoSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/ssoSettings:SsoSettings"),
		},
	})
	opts = append(opts, aliases)
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
	// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
	Oauth2Settings *SsoSettingsOauth2Settings `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
	ProviderName *string `pulumi:"providerName"`
	// The SAML settings set. Required for the saml provider.
	SamlSettings *SsoSettingsSamlSettings `pulumi:"samlSettings"`
}

type SsoSettingsState struct {
	// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
	Oauth2Settings SsoSettingsOauth2SettingsPtrInput
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
	ProviderName pulumi.StringPtrInput
	// The SAML settings set. Required for the saml provider.
	SamlSettings SsoSettingsSamlSettingsPtrInput
}

func (SsoSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ssoSettingsState)(nil)).Elem()
}

type ssoSettingsArgs struct {
	// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
	Oauth2Settings *SsoSettingsOauth2Settings `pulumi:"oauth2Settings"`
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
	ProviderName string `pulumi:"providerName"`
	// The SAML settings set. Required for the saml provider.
	SamlSettings *SsoSettingsSamlSettings `pulumi:"samlSettings"`
}

// The set of arguments for constructing a SsoSettings resource.
type SsoSettingsArgs struct {
	// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
	Oauth2Settings SsoSettingsOauth2SettingsPtrInput
	// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
	ProviderName pulumi.StringInput
	// The SAML settings set. Required for the saml provider.
	SamlSettings SsoSettingsSamlSettingsPtrInput
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

// The OAuth2 settings set. Required for github, gitlab, google, azuread, okta, generic*oauth providers.
func (o SsoSettingsOutput) Oauth2Settings() SsoSettingsOauth2SettingsPtrOutput {
	return o.ApplyT(func(v *SsoSettings) SsoSettingsOauth2SettingsPtrOutput { return v.Oauth2Settings }).(SsoSettingsOauth2SettingsPtrOutput)
}

// The name of the SSO provider. Supported values: github, gitlab, google, azuread, okta, generic_oauth, saml.
func (o SsoSettingsOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *SsoSettings) pulumi.StringOutput { return v.ProviderName }).(pulumi.StringOutput)
}

// The SAML settings set. Required for the saml provider.
func (o SsoSettingsOutput) SamlSettings() SsoSettingsSamlSettingsPtrOutput {
	return o.ApplyT(func(v *SsoSettings) SsoSettingsSamlSettingsPtrOutput { return v.SamlSettings }).(SsoSettingsSamlSettingsPtrOutput)
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
