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
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/annotations/)
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
    ///     var test = new Grafana.Oss.Annotation("test", new()
    ///     {
    ///         Text = "basic text",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/annotation:Annotation name "{{ id }}"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/annotation:Annotation name "{{ orgID }}:{{ id }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/annotation.Annotation has been deprecated in favor of grafana.oss/annotation.Annotation")]
    [GrafanaResourceType("grafana:index/annotation:Annotation")]
    public partial class Annotation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The UID of the dashboard on which to create the annotation.
        /// </summary>
        [Output("dashboardUid")]
        public Output<string?> DashboardUid { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// The ID of the dashboard panel on which to create the annotation.
        /// </summary>
        [Output("panelId")]
        public Output<int?> PanelId { get; private set; } = null!;

        /// <summary>
        /// The tags to associate with the annotation.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The text to associate with the annotation.
        /// </summary>
        [Output("text")]
        public Output<string> Text { get; private set; } = null!;

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's time.
        /// </summary>
        [Output("time")]
        public Output<string> Time { get; private set; } = null!;

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's end time.
        /// </summary>
        [Output("timeEnd")]
        public Output<string> TimeEnd { get; private set; } = null!;


        /// <summary>
        /// Create a Annotation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Annotation(string name, AnnotationArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/annotation:Annotation", name, args ?? new AnnotationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Annotation(string name, Input<string> id, AnnotationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/annotation:Annotation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "grafana:index/annotation:Annotation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Annotation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Annotation Get(string name, Input<string> id, AnnotationState? state = null, CustomResourceOptions? options = null)
        {
            return new Annotation(name, id, state, options);
        }
    }

    public sealed class AnnotationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the dashboard on which to create the annotation.
        /// </summary>
        [Input("dashboardUid")]
        public Input<string>? DashboardUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The ID of the dashboard panel on which to create the annotation.
        /// </summary>
        [Input("panelId")]
        public Input<int>? PanelId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags to associate with the annotation.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The text to associate with the annotation.
        /// </summary>
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's time.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's end time.
        /// </summary>
        [Input("timeEnd")]
        public Input<string>? TimeEnd { get; set; }

        public AnnotationArgs()
        {
        }
        public static new AnnotationArgs Empty => new AnnotationArgs();
    }

    public sealed class AnnotationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UID of the dashboard on which to create the annotation.
        /// </summary>
        [Input("dashboardUid")]
        public Input<string>? DashboardUid { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The ID of the dashboard panel on which to create the annotation.
        /// </summary>
        [Input("panelId")]
        public Input<int>? PanelId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags to associate with the annotation.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The text to associate with the annotation.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's time.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        /// <summary>
        /// The RFC 3339-formatted time string indicating the annotation's end time.
        /// </summary>
        [Input("timeEnd")]
        public Input<string>? TimeEnd { get; set; }

        public AnnotationState()
        {
        }
        public static new AnnotationState Empty => new AnnotationState();
    }
}
