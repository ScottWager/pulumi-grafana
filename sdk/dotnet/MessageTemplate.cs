// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// Manages Grafana Alerting message templates.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/manage-notifications/template-notifications/create-notification-templates/)
    /// * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#templates)
    /// 
    /// This resource requires Grafana 9.1.0 or later.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myTemplate = new Grafana.MessageTemplate("myTemplate", new()
    ///     {
    ///         Template = @"{{define ""My Reusable Template"" }}
    ///  template content
    /// {{ end }}
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/messageTemplate:MessageTemplate message_template_name {{message_template_name}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/messageTemplate:MessageTemplate")]
    public partial class MessageTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the message template.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The content of the message template.
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;


        /// <summary>
        /// Create a MessageTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MessageTemplate(string name, MessageTemplateArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/messageTemplate:MessageTemplate", name, args ?? new MessageTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MessageTemplate(string name, Input<string> id, MessageTemplateState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/messageTemplate:MessageTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MessageTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MessageTemplate Get(string name, Input<string> id, MessageTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MessageTemplate(name, id, state, options);
        }
    }

    public sealed class MessageTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the message template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The content of the message template.
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public MessageTemplateArgs()
        {
        }
        public static new MessageTemplateArgs Empty => new MessageTemplateArgs();
    }

    public sealed class MessageTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the message template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The content of the message template.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public MessageTemplateState()
        {
        }
        public static new MessageTemplateState Empty => new MessageTemplateState();
    }
}
