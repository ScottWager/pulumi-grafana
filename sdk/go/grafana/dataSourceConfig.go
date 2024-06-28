// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/internal"
)

// * [Official documentation](https://grafana.com/docs/grafana/latest/datasources/)
// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/data_source/)
//
// The required arguments for this resource vary depending on the type of data
// source selected (via the 'type' argument).
//
// Use this resource for configuring multiple datasources, when that configuration (`jsonDataEncoded` field) requires circular references like in the example below.
//
// > When using the `oss.DataSourceConfig` resource, the corresponding `oss.DataSource` resources must have the `jsonDataEncoded` and `httpHeaders` fields ignored. Otherwise, an infinite update loop will occur. See the example below.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-grafana/sdk/go/grafana/oss"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lokiDataSource, err := oss.NewDataSource(ctx, "lokiDataSource", &oss.DataSourceArgs{
//				Type: pulumi.String("loki"),
//				Url:  pulumi.String("http://localhost:3100"),
//			})
//			if err != nil {
//				return err
//			}
//			tempoDataSource, err := oss.NewDataSource(ctx, "tempoDataSource", &oss.DataSourceArgs{
//				Type: pulumi.String("tempo"),
//				Url:  pulumi.String("http://localhost:3200"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewDataSourceConfig(ctx, "lokiDataSourceConfig", &oss.DataSourceConfigArgs{
//				Uid: lokiDataSource.Uid,
//				JsonDataEncoded: tempoDataSource.Uid.ApplyT(func(uid string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON0, err := json.Marshal(map[string]interface{}{
//						"derivedFields": []map[string]interface{}{
//							map[string]interface{}{
//								"datasourceUid": uid,
//								"matcherRegex":  "[tT]race_?[iI][dD]\"?[:=]\"?(\\w+)",
//								"matcherType":   "regex",
//								"name":          "traceID",
//								"url":           "${__value.raw}",
//							},
//						},
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json0 := string(tmpJSON0)
//					return pulumi.String(json0), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewDataSourceConfig(ctx, "tempoDataSourceConfig", &oss.DataSourceConfigArgs{
//				Uid: tempoDataSource.Uid,
//				JsonDataEncoded: lokiDataSource.Uid.ApplyT(func(uid string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON1, err := json.Marshal(map[string]interface{}{
//						"tracesToLogsV2": map[string]interface{}{
//							"customQuery":     true,
//							"datasourceUid":   uid,
//							"filterBySpanID":  false,
//							"filterByTraceID": false,
//							"query":           "|=\"${__trace.traceId}\" | json",
//						},
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json1 := string(tmpJSON1)
//					return pulumi.String(json1), nil
//				}).(pulumi.StringOutput),
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
// $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ uid }}"
// ```
//
// ```sh
// $ pulumi import grafana:index/dataSourceConfig:DataSourceConfig name "{{ orgID }}:{{ uid }}"
// ```
//
// Deprecated: grafana.index/datasourceconfig.DataSourceConfig has been deprecated in favor of grafana.oss/datasourceconfig.DataSourceConfig
type DataSourceConfig struct {
	pulumi.CustomResourceState

	// Custom HTTP headers
	HttpHeaders pulumi.StringMapOutput `pulumi:"httpHeaders"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumi.StringPtrOutput `pulumi:"jsonDataEncoded"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumi.StringPtrOutput `pulumi:"secureJsonDataEncoded"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewDataSourceConfig registers a new resource with the given unique name, arguments, and options.
func NewDataSourceConfig(ctx *pulumi.Context,
	name string, args *DataSourceConfigArgs, opts ...pulumi.ResourceOption) (*DataSourceConfig, error) {
	if args == nil {
		args = &DataSourceConfigArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("grafana:index/dataSourceConfig:DataSourceConfig"),
		},
	})
	opts = append(opts, aliases)
	if args.HttpHeaders != nil {
		args.HttpHeaders = pulumi.ToSecret(args.HttpHeaders).(pulumi.StringMapInput)
	}
	if args.SecureJsonDataEncoded != nil {
		args.SecureJsonDataEncoded = pulumi.ToSecret(args.SecureJsonDataEncoded).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"httpHeaders",
		"secureJsonDataEncoded",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataSourceConfig
	err := ctx.RegisterResource("grafana:index/dataSourceConfig:DataSourceConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSourceConfig gets an existing DataSourceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSourceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceConfigState, opts ...pulumi.ResourceOption) (*DataSourceConfig, error) {
	var resource DataSourceConfig
	err := ctx.ReadResource("grafana:index/dataSourceConfig:DataSourceConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSourceConfig resources.
type dataSourceConfigState struct {
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid *string `pulumi:"uid"`
}

type DataSourceConfigState struct {
	// Custom HTTP headers
	HttpHeaders pulumi.StringMapInput
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumi.StringPtrInput
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringPtrInput
}

func (DataSourceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceConfigState)(nil)).Elem()
}

type dataSourceConfigArgs struct {
	// Custom HTTP headers
	HttpHeaders map[string]string `pulumi:"httpHeaders"`
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded *string `pulumi:"jsonDataEncoded"`
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId *string `pulumi:"orgId"`
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded *string `pulumi:"secureJsonDataEncoded"`
	// Unique identifier. If unset, this will be automatically generated.
	Uid *string `pulumi:"uid"`
}

// The set of arguments for constructing a DataSourceConfig resource.
type DataSourceConfigArgs struct {
	// Custom HTTP headers
	HttpHeaders pulumi.StringMapInput
	// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	JsonDataEncoded pulumi.StringPtrInput
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgId pulumi.StringPtrInput
	// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
	SecureJsonDataEncoded pulumi.StringPtrInput
	// Unique identifier. If unset, this will be automatically generated.
	Uid pulumi.StringPtrInput
}

func (DataSourceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceConfigArgs)(nil)).Elem()
}

type DataSourceConfigInput interface {
	pulumi.Input

	ToDataSourceConfigOutput() DataSourceConfigOutput
	ToDataSourceConfigOutputWithContext(ctx context.Context) DataSourceConfigOutput
}

func (*DataSourceConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceConfig)(nil)).Elem()
}

func (i *DataSourceConfig) ToDataSourceConfigOutput() DataSourceConfigOutput {
	return i.ToDataSourceConfigOutputWithContext(context.Background())
}

func (i *DataSourceConfig) ToDataSourceConfigOutputWithContext(ctx context.Context) DataSourceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceConfigOutput)
}

// DataSourceConfigArrayInput is an input type that accepts DataSourceConfigArray and DataSourceConfigArrayOutput values.
// You can construct a concrete instance of `DataSourceConfigArrayInput` via:
//
//	DataSourceConfigArray{ DataSourceConfigArgs{...} }
type DataSourceConfigArrayInput interface {
	pulumi.Input

	ToDataSourceConfigArrayOutput() DataSourceConfigArrayOutput
	ToDataSourceConfigArrayOutputWithContext(context.Context) DataSourceConfigArrayOutput
}

type DataSourceConfigArray []DataSourceConfigInput

func (DataSourceConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourceConfig)(nil)).Elem()
}

func (i DataSourceConfigArray) ToDataSourceConfigArrayOutput() DataSourceConfigArrayOutput {
	return i.ToDataSourceConfigArrayOutputWithContext(context.Background())
}

func (i DataSourceConfigArray) ToDataSourceConfigArrayOutputWithContext(ctx context.Context) DataSourceConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceConfigArrayOutput)
}

// DataSourceConfigMapInput is an input type that accepts DataSourceConfigMap and DataSourceConfigMapOutput values.
// You can construct a concrete instance of `DataSourceConfigMapInput` via:
//
//	DataSourceConfigMap{ "key": DataSourceConfigArgs{...} }
type DataSourceConfigMapInput interface {
	pulumi.Input

	ToDataSourceConfigMapOutput() DataSourceConfigMapOutput
	ToDataSourceConfigMapOutputWithContext(context.Context) DataSourceConfigMapOutput
}

type DataSourceConfigMap map[string]DataSourceConfigInput

func (DataSourceConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourceConfig)(nil)).Elem()
}

func (i DataSourceConfigMap) ToDataSourceConfigMapOutput() DataSourceConfigMapOutput {
	return i.ToDataSourceConfigMapOutputWithContext(context.Background())
}

func (i DataSourceConfigMap) ToDataSourceConfigMapOutputWithContext(ctx context.Context) DataSourceConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceConfigMapOutput)
}

type DataSourceConfigOutput struct{ *pulumi.OutputState }

func (DataSourceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceConfig)(nil)).Elem()
}

func (o DataSourceConfigOutput) ToDataSourceConfigOutput() DataSourceConfigOutput {
	return o
}

func (o DataSourceConfigOutput) ToDataSourceConfigOutputWithContext(ctx context.Context) DataSourceConfigOutput {
	return o
}

// Custom HTTP headers
func (o DataSourceConfigOutput) HttpHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataSourceConfig) pulumi.StringMapOutput { return v.HttpHeaders }).(pulumi.StringMapOutput)
}

// Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
func (o DataSourceConfigOutput) JsonDataEncoded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourceConfig) pulumi.StringPtrOutput { return v.JsonDataEncoded }).(pulumi.StringPtrOutput)
}

// The Organization ID. If not set, the Org ID defined in the provider block will be used.
func (o DataSourceConfigOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourceConfig) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Serialized JSON string containing the secure json data. This attribute can be used to pass secure configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
func (o DataSourceConfigOutput) SecureJsonDataEncoded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSourceConfig) pulumi.StringPtrOutput { return v.SecureJsonDataEncoded }).(pulumi.StringPtrOutput)
}

// Unique identifier. If unset, this will be automatically generated.
func (o DataSourceConfigOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSourceConfig) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type DataSourceConfigArrayOutput struct{ *pulumi.OutputState }

func (DataSourceConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSourceConfig)(nil)).Elem()
}

func (o DataSourceConfigArrayOutput) ToDataSourceConfigArrayOutput() DataSourceConfigArrayOutput {
	return o
}

func (o DataSourceConfigArrayOutput) ToDataSourceConfigArrayOutputWithContext(ctx context.Context) DataSourceConfigArrayOutput {
	return o
}

func (o DataSourceConfigArrayOutput) Index(i pulumi.IntInput) DataSourceConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataSourceConfig {
		return vs[0].([]*DataSourceConfig)[vs[1].(int)]
	}).(DataSourceConfigOutput)
}

type DataSourceConfigMapOutput struct{ *pulumi.OutputState }

func (DataSourceConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSourceConfig)(nil)).Elem()
}

func (o DataSourceConfigMapOutput) ToDataSourceConfigMapOutput() DataSourceConfigMapOutput {
	return o
}

func (o DataSourceConfigMapOutput) ToDataSourceConfigMapOutputWithContext(ctx context.Context) DataSourceConfigMapOutput {
	return o
}

func (o DataSourceConfigMapOutput) MapIndex(k pulumi.StringInput) DataSourceConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataSourceConfig {
		return vs[0].(map[string]*DataSourceConfig)[vs[1].(string)]
	}).(DataSourceConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceConfigInput)(nil)).Elem(), &DataSourceConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceConfigArrayInput)(nil)).Elem(), DataSourceConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceConfigMapInput)(nil)).Elem(), DataSourceConfigMap{})
	pulumi.RegisterOutputType(DataSourceConfigOutput{})
	pulumi.RegisterOutputType(DataSourceConfigArrayOutput{})
	pulumi.RegisterOutputType(DataSourceConfigMapOutput{})
}
