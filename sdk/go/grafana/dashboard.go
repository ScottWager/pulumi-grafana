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

// Manages Grafana dashboards.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testFolder, err := grafana.NewFolder(ctx, "testFolder", &grafana.FolderArgs{
//				Title: pulumi.String("My Folder"),
//				Uid:   pulumi.String("my-folder-uid"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"title": "My Dashboard",
//				"uid":   "my-dashboard-uid",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = grafana.NewDashboard(ctx, "testDashboard", &grafana.DashboardArgs{
//				Folder:     testFolder.Uid,
//				ConfigJson: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ```sh
// $ pulumi import grafana:index/dashboard:Dashboard dashboard_name {{dashboard_uid}} # To use the default provider org
// ```
//
// ```sh
// $ pulumi import grafana:index/dashboard:Dashboard dashboard_name {{org_id}}:{{dashboard_uid}} # When "org_id" is set on the resource
// ```
type Dashboard struct {
	pulumi.CustomResourceState

	// The complete dashboard model JSON.
	ConfigJson pulumi.StringOutput `pulumi:"configJson"`
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId pulumi.IntOutput `pulumi:"dashboardId"`
	// The id or UID of the folder to save the dashboard in.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Set a commit message for the version history.
	Message pulumi.StringPtrOutput `pulumi:"message"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite pulumi.BoolPtrOutput `pulumi:"overwrite"`
	// The unique identifier of a dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The full URL of the dashboard.
	Url pulumi.StringOutput `pulumi:"url"`
	// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigJson == nil {
		return nil, errors.New("invalid value for required argument 'ConfigJson'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("grafana:index/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("grafana:index/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// The complete dashboard model JSON.
	ConfigJson *string `pulumi:"configJson"`
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId *int `pulumi:"dashboardId"`
	// The id or UID of the folder to save the dashboard in.
	Folder *string `pulumi:"folder"`
	// Set a commit message for the version history.
	Message *string `pulumi:"message"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite *bool `pulumi:"overwrite"`
	// The unique identifier of a dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
	Uid *string `pulumi:"uid"`
	// The full URL of the dashboard.
	Url *string `pulumi:"url"`
	// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
	Version *int `pulumi:"version"`
}

type DashboardState struct {
	// The complete dashboard model JSON.
	ConfigJson pulumi.StringPtrInput
	// The numeric ID of the dashboard computed by Grafana.
	DashboardId pulumi.IntPtrInput
	// The id or UID of the folder to save the dashboard in.
	Folder pulumi.StringPtrInput
	// Set a commit message for the version history.
	Message pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite pulumi.BoolPtrInput
	// The unique identifier of a dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
	Uid pulumi.StringPtrInput
	// The full URL of the dashboard.
	Url pulumi.StringPtrInput
	// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
	Version pulumi.IntPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The complete dashboard model JSON.
	ConfigJson string `pulumi:"configJson"`
	// The id or UID of the folder to save the dashboard in.
	Folder *string `pulumi:"folder"`
	// Set a commit message for the version history.
	Message *string `pulumi:"message"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite *bool `pulumi:"overwrite"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The complete dashboard model JSON.
	ConfigJson pulumi.StringInput
	// The id or UID of the folder to save the dashboard in.
	Folder pulumi.StringPtrInput
	// Set a commit message for the version history.
	Message pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite pulumi.BoolPtrInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//	DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//	DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// The complete dashboard model JSON.
func (o DashboardOutput) ConfigJson() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.ConfigJson }).(pulumi.StringOutput)
}

// The numeric ID of the dashboard computed by Grafana.
func (o DashboardOutput) DashboardId() pulumi.IntOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.IntOutput { return v.DashboardId }).(pulumi.IntOutput)
}

// The id or UID of the folder to save the dashboard in.
func (o DashboardOutput) Folder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.Folder }).(pulumi.StringPtrOutput)
}

// Set a commit message for the version history.
func (o DashboardOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DashboardOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
func (o DashboardOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.BoolPtrOutput { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

// The unique identifier of a dashboard. This is used to construct its URL. It's automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
func (o DashboardOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The full URL of the dashboard.
func (o DashboardOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
func (o DashboardOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].([]*Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].(map[string]*Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardArrayInput)(nil)).Elem(), DashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardMapInput)(nil)).Elem(), DashboardMap{})
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}
