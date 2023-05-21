// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Grafana Alerting message templates.
//
// * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/manage-notifications/template-notifications/create-notification-templates/)
// * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#templates)
//
// This resource requires Grafana 9.1.0 or later.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := grafana.NewMessageTemplate(ctx, "myTemplate", &grafana.MessageTemplateArgs{
//				Template: pulumi.String("{{define \"My Reusable Template\" }}\n template content\n{{ end }}\n"),
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
//	$ pulumi import grafana:index/messageTemplate:MessageTemplate message_template_name {{message_template_name}}
//
// ```
type MessageTemplate struct {
	pulumi.CustomResourceState

	// The name of the message template.
	Name pulumi.StringOutput `pulumi:"name"`
	// The content of the message template.
	Template pulumi.StringOutput `pulumi:"template"`
}

// NewMessageTemplate registers a new resource with the given unique name, arguments, and options.
func NewMessageTemplate(ctx *pulumi.Context,
	name string, args *MessageTemplateArgs, opts ...pulumi.ResourceOption) (*MessageTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MessageTemplate
	err := ctx.RegisterResource("grafana:index/messageTemplate:MessageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMessageTemplate gets an existing MessageTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMessageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MessageTemplateState, opts ...pulumi.ResourceOption) (*MessageTemplate, error) {
	var resource MessageTemplate
	err := ctx.ReadResource("grafana:index/messageTemplate:MessageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MessageTemplate resources.
type messageTemplateState struct {
	// The name of the message template.
	Name *string `pulumi:"name"`
	// The content of the message template.
	Template *string `pulumi:"template"`
}

type MessageTemplateState struct {
	// The name of the message template.
	Name pulumi.StringPtrInput
	// The content of the message template.
	Template pulumi.StringPtrInput
}

func (MessageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*messageTemplateState)(nil)).Elem()
}

type messageTemplateArgs struct {
	// The name of the message template.
	Name *string `pulumi:"name"`
	// The content of the message template.
	Template string `pulumi:"template"`
}

// The set of arguments for constructing a MessageTemplate resource.
type MessageTemplateArgs struct {
	// The name of the message template.
	Name pulumi.StringPtrInput
	// The content of the message template.
	Template pulumi.StringInput
}

func (MessageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*messageTemplateArgs)(nil)).Elem()
}

type MessageTemplateInput interface {
	pulumi.Input

	ToMessageTemplateOutput() MessageTemplateOutput
	ToMessageTemplateOutputWithContext(ctx context.Context) MessageTemplateOutput
}

func (*MessageTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageTemplate)(nil)).Elem()
}

func (i *MessageTemplate) ToMessageTemplateOutput() MessageTemplateOutput {
	return i.ToMessageTemplateOutputWithContext(context.Background())
}

func (i *MessageTemplate) ToMessageTemplateOutputWithContext(ctx context.Context) MessageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageTemplateOutput)
}

// MessageTemplateArrayInput is an input type that accepts MessageTemplateArray and MessageTemplateArrayOutput values.
// You can construct a concrete instance of `MessageTemplateArrayInput` via:
//
//	MessageTemplateArray{ MessageTemplateArgs{...} }
type MessageTemplateArrayInput interface {
	pulumi.Input

	ToMessageTemplateArrayOutput() MessageTemplateArrayOutput
	ToMessageTemplateArrayOutputWithContext(context.Context) MessageTemplateArrayOutput
}

type MessageTemplateArray []MessageTemplateInput

func (MessageTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MessageTemplate)(nil)).Elem()
}

func (i MessageTemplateArray) ToMessageTemplateArrayOutput() MessageTemplateArrayOutput {
	return i.ToMessageTemplateArrayOutputWithContext(context.Background())
}

func (i MessageTemplateArray) ToMessageTemplateArrayOutputWithContext(ctx context.Context) MessageTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageTemplateArrayOutput)
}

// MessageTemplateMapInput is an input type that accepts MessageTemplateMap and MessageTemplateMapOutput values.
// You can construct a concrete instance of `MessageTemplateMapInput` via:
//
//	MessageTemplateMap{ "key": MessageTemplateArgs{...} }
type MessageTemplateMapInput interface {
	pulumi.Input

	ToMessageTemplateMapOutput() MessageTemplateMapOutput
	ToMessageTemplateMapOutputWithContext(context.Context) MessageTemplateMapOutput
}

type MessageTemplateMap map[string]MessageTemplateInput

func (MessageTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MessageTemplate)(nil)).Elem()
}

func (i MessageTemplateMap) ToMessageTemplateMapOutput() MessageTemplateMapOutput {
	return i.ToMessageTemplateMapOutputWithContext(context.Background())
}

func (i MessageTemplateMap) ToMessageTemplateMapOutputWithContext(ctx context.Context) MessageTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageTemplateMapOutput)
}

type MessageTemplateOutput struct{ *pulumi.OutputState }

func (MessageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageTemplate)(nil)).Elem()
}

func (o MessageTemplateOutput) ToMessageTemplateOutput() MessageTemplateOutput {
	return o
}

func (o MessageTemplateOutput) ToMessageTemplateOutputWithContext(ctx context.Context) MessageTemplateOutput {
	return o
}

// The name of the message template.
func (o MessageTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MessageTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The content of the message template.
func (o MessageTemplateOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v *MessageTemplate) pulumi.StringOutput { return v.Template }).(pulumi.StringOutput)
}

type MessageTemplateArrayOutput struct{ *pulumi.OutputState }

func (MessageTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MessageTemplate)(nil)).Elem()
}

func (o MessageTemplateArrayOutput) ToMessageTemplateArrayOutput() MessageTemplateArrayOutput {
	return o
}

func (o MessageTemplateArrayOutput) ToMessageTemplateArrayOutputWithContext(ctx context.Context) MessageTemplateArrayOutput {
	return o
}

func (o MessageTemplateArrayOutput) Index(i pulumi.IntInput) MessageTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MessageTemplate {
		return vs[0].([]*MessageTemplate)[vs[1].(int)]
	}).(MessageTemplateOutput)
}

type MessageTemplateMapOutput struct{ *pulumi.OutputState }

func (MessageTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MessageTemplate)(nil)).Elem()
}

func (o MessageTemplateMapOutput) ToMessageTemplateMapOutput() MessageTemplateMapOutput {
	return o
}

func (o MessageTemplateMapOutput) ToMessageTemplateMapOutputWithContext(ctx context.Context) MessageTemplateMapOutput {
	return o
}

func (o MessageTemplateMapOutput) MapIndex(k pulumi.StringInput) MessageTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MessageTemplate {
		return vs[0].(map[string]*MessageTemplate)[vs[1].(string)]
	}).(MessageTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MessageTemplateInput)(nil)).Elem(), &MessageTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessageTemplateArrayInput)(nil)).Elem(), MessageTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessageTemplateMapInput)(nil)).Elem(), MessageTemplateMap{})
	pulumi.RegisterOutputType(MessageTemplateOutput{})
	pulumi.RegisterOutputType(MessageTemplateArrayOutput{})
	pulumi.RegisterOutputType(MessageTemplateMapOutput{})
}
