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
    /// * [Official documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#stacks/)
    /// 
    /// Required access policy scopes:
    /// 
    /// * stacks:read
    /// * stacks:write
    /// * stacks:delete
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
    ///     var test = new Grafana.Cloud.Stack("test", new()
    ///     {
    ///         Description = "Test Grafana Cloud Stack",
    ///         RegionSlug = "eu",
    ///         Slug = "gcloudstacktest",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/cloudStack:CloudStack name "{{ stackSlugOrID }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/cloudstack.CloudStack has been deprecated in favor of grafana.cloud/stack.Stack")]
    [GrafanaResourceType("grafana:index/cloudStack:CloudStack")]
    public partial class CloudStack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Alertmanager instance configured for this stack.
        /// </summary>
        [Output("alertmanagerName")]
        public Output<string> AlertmanagerName { get; private set; } = null!;

        /// <summary>
        /// Status of the Alertmanager instance configured for this stack.
        /// </summary>
        [Output("alertmanagerStatus")]
        public Output<string> AlertmanagerStatus { get; private set; } = null!;

        /// <summary>
        /// Base URL of the Alertmanager instance configured for this stack.
        /// </summary>
        [Output("alertmanagerUrl")]
        public Output<string> AlertmanagerUrl { get; private set; } = null!;

        /// <summary>
        /// User ID of the Alertmanager instance configured for this stack.
        /// </summary>
        [Output("alertmanagerUserId")]
        public Output<int> AlertmanagerUserId { get; private set; } = null!;

        /// <summary>
        /// Description of stack.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("graphiteName")]
        public Output<string> GraphiteName { get; private set; } = null!;

        [Output("graphiteStatus")]
        public Output<string> GraphiteStatus { get; private set; } = null!;

        [Output("graphiteUrl")]
        public Output<string> GraphiteUrl { get; private set; } = null!;

        [Output("graphiteUserId")]
        public Output<int> GraphiteUserId { get; private set; } = null!;

        /// <summary>
        /// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheus_user_id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
        /// </summary>
        [Output("influxUrl")]
        public Output<string> InfluxUrl { get; private set; } = null!;

        /// <summary>
        /// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("logsName")]
        public Output<string> LogsName { get; private set; } = null!;

        [Output("logsStatus")]
        public Output<string> LogsStatus { get; private set; } = null!;

        [Output("logsUrl")]
        public Output<string> LogsUrl { get; private set; } = null!;

        [Output("logsUserId")]
        public Output<int> LogsUserId { get; private set; } = null!;

        /// <summary>
        /// Name of stack. Conventionally matches the url of the instance (e.g. `&lt;stack_slug&gt;.grafana.net`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Organization id to assign to this stack.
        /// </summary>
        [Output("orgId")]
        public Output<int> OrgId { get; private set; } = null!;

        /// <summary>
        /// Organization name to assign to this stack.
        /// </summary>
        [Output("orgName")]
        public Output<string> OrgName { get; private set; } = null!;

        /// <summary>
        /// Organization slug to assign to this stack.
        /// </summary>
        [Output("orgSlug")]
        public Output<string> OrgSlug { get; private set; } = null!;

        /// <summary>
        /// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
        /// </summary>
        [Output("otlpUrl")]
        public Output<string> OtlpUrl { get; private set; } = null!;

        [Output("profilesName")]
        public Output<string> ProfilesName { get; private set; } = null!;

        [Output("profilesStatus")]
        public Output<string> ProfilesStatus { get; private set; } = null!;

        [Output("profilesUrl")]
        public Output<string> ProfilesUrl { get; private set; } = null!;

        [Output("profilesUserId")]
        public Output<int> ProfilesUserId { get; private set; } = null!;

        /// <summary>
        /// Prometheus name for this instance.
        /// </summary>
        [Output("prometheusName")]
        public Output<string> PrometheusName { get; private set; } = null!;

        /// <summary>
        /// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
        /// </summary>
        [Output("prometheusRemoteEndpoint")]
        public Output<string> PrometheusRemoteEndpoint { get; private set; } = null!;

        /// <summary>
        /// Use this URL to send prometheus metrics to Grafana cloud
        /// </summary>
        [Output("prometheusRemoteWriteEndpoint")]
        public Output<string> PrometheusRemoteWriteEndpoint { get; private set; } = null!;

        /// <summary>
        /// Prometheus status for this instance.
        /// </summary>
        [Output("prometheusStatus")]
        public Output<string> PrometheusStatus { get; private set; } = null!;

        /// <summary>
        /// Prometheus url for this instance.
        /// </summary>
        [Output("prometheusUrl")]
        public Output<string> PrometheusUrl { get; private set; } = null!;

        /// <summary>
        /// Prometheus user ID. Used for e.g. remote_write.
        /// </summary>
        [Output("prometheusUserId")]
        public Output<int> PrometheusUserId { get; private set; } = null!;

        /// <summary>
        /// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Output("regionSlug")]
        public Output<string?> RegionSlug { get; private set; } = null!;

        /// <summary>
        /// Subdomain that the Grafana instance will be available at. Setting slug to `&lt;stack_slug&gt;` will make the instance available at `https://&lt;stack_slug&gt;.grafana.net`.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// Status of the stack.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tracesName")]
        public Output<string> TracesName { get; private set; } = null!;

        [Output("tracesStatus")]
        public Output<string> TracesStatus { get; private set; } = null!;

        /// <summary>
        /// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
        /// </summary>
        [Output("tracesUrl")]
        public Output<string> TracesUrl { get; private set; } = null!;

        [Output("tracesUserId")]
        public Output<int> TracesUserId { get; private set; } = null!;

        /// <summary>
        /// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        /// <summary>
        /// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
        /// </summary>
        [Output("waitForReadiness")]
        public Output<bool?> WaitForReadiness { get; private set; } = null!;

        /// <summary>
        /// How long to wait for readiness (if enabled). Defaults to `5m0s`.
        /// </summary>
        [Output("waitForReadinessTimeout")]
        public Output<string?> WaitForReadinessTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a CloudStack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudStack(string name, CloudStackArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/cloudStack:CloudStack", name, args ?? new CloudStackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudStack(string name, Input<string> id, CloudStackState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/cloudStack:CloudStack", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/cloudStack:CloudStack" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudStack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudStack Get(string name, Input<string> id, CloudStackState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudStack(name, id, state, options);
        }
    }

    public sealed class CloudStackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of stack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of stack. Conventionally matches the url of the instance (e.g. `&lt;stack_slug&gt;.grafana.net`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Input("regionSlug")]
        public Input<string>? RegionSlug { get; set; }

        /// <summary>
        /// Subdomain that the Grafana instance will be available at. Setting slug to `&lt;stack_slug&gt;` will make the instance available at `https://&lt;stack_slug&gt;.grafana.net`.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        /// <summary>
        /// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
        /// </summary>
        [Input("waitForReadiness")]
        public Input<bool>? WaitForReadiness { get; set; }

        /// <summary>
        /// How long to wait for readiness (if enabled). Defaults to `5m0s`.
        /// </summary>
        [Input("waitForReadinessTimeout")]
        public Input<string>? WaitForReadinessTimeout { get; set; }

        public CloudStackArgs()
        {
        }
        public static new CloudStackArgs Empty => new CloudStackArgs();
    }

    public sealed class CloudStackState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Alertmanager instance configured for this stack.
        /// </summary>
        [Input("alertmanagerName")]
        public Input<string>? AlertmanagerName { get; set; }

        /// <summary>
        /// Status of the Alertmanager instance configured for this stack.
        /// </summary>
        [Input("alertmanagerStatus")]
        public Input<string>? AlertmanagerStatus { get; set; }

        /// <summary>
        /// Base URL of the Alertmanager instance configured for this stack.
        /// </summary>
        [Input("alertmanagerUrl")]
        public Input<string>? AlertmanagerUrl { get; set; }

        /// <summary>
        /// User ID of the Alertmanager instance configured for this stack.
        /// </summary>
        [Input("alertmanagerUserId")]
        public Input<int>? AlertmanagerUserId { get; set; }

        /// <summary>
        /// Description of stack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("graphiteName")]
        public Input<string>? GraphiteName { get; set; }

        [Input("graphiteStatus")]
        public Input<string>? GraphiteStatus { get; set; }

        [Input("graphiteUrl")]
        public Input<string>? GraphiteUrl { get; set; }

        [Input("graphiteUserId")]
        public Input<int>? GraphiteUserId { get; set; }

        /// <summary>
        /// Base URL of the InfluxDB instance configured for this stack. The username is the same as the metrics' (`prometheus_user_id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/metrics/metrics-influxdb/push-from-telegraf/ for docs on how to use this.
        /// </summary>
        [Input("influxUrl")]
        public Input<string>? InfluxUrl { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A map of labels to assign to the stack. Label keys and values must match the following regexp: "^[a-zA-Z0-9/\-.]+$" and stacks cannot have more than 10 labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("logsName")]
        public Input<string>? LogsName { get; set; }

        [Input("logsStatus")]
        public Input<string>? LogsStatus { get; set; }

        [Input("logsUrl")]
        public Input<string>? LogsUrl { get; set; }

        [Input("logsUserId")]
        public Input<int>? LogsUserId { get; set; }

        /// <summary>
        /// Name of stack. Conventionally matches the url of the instance (e.g. `&lt;stack_slug&gt;.grafana.net`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organization id to assign to this stack.
        /// </summary>
        [Input("orgId")]
        public Input<int>? OrgId { get; set; }

        /// <summary>
        /// Organization name to assign to this stack.
        /// </summary>
        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        /// <summary>
        /// Organization slug to assign to this stack.
        /// </summary>
        [Input("orgSlug")]
        public Input<string>? OrgSlug { get; set; }

        /// <summary>
        /// Base URL of the OTLP instance configured for this stack. The username is the stack's ID (`id` attribute of this resource). See https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/ for docs on how to use this.
        /// </summary>
        [Input("otlpUrl")]
        public Input<string>? OtlpUrl { get; set; }

        [Input("profilesName")]
        public Input<string>? ProfilesName { get; set; }

        [Input("profilesStatus")]
        public Input<string>? ProfilesStatus { get; set; }

        [Input("profilesUrl")]
        public Input<string>? ProfilesUrl { get; set; }

        [Input("profilesUserId")]
        public Input<int>? ProfilesUserId { get; set; }

        /// <summary>
        /// Prometheus name for this instance.
        /// </summary>
        [Input("prometheusName")]
        public Input<string>? PrometheusName { get; set; }

        /// <summary>
        /// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
        /// </summary>
        [Input("prometheusRemoteEndpoint")]
        public Input<string>? PrometheusRemoteEndpoint { get; set; }

        /// <summary>
        /// Use this URL to send prometheus metrics to Grafana cloud
        /// </summary>
        [Input("prometheusRemoteWriteEndpoint")]
        public Input<string>? PrometheusRemoteWriteEndpoint { get; set; }

        /// <summary>
        /// Prometheus status for this instance.
        /// </summary>
        [Input("prometheusStatus")]
        public Input<string>? PrometheusStatus { get; set; }

        /// <summary>
        /// Prometheus url for this instance.
        /// </summary>
        [Input("prometheusUrl")]
        public Input<string>? PrometheusUrl { get; set; }

        /// <summary>
        /// Prometheus user ID. Used for e.g. remote_write.
        /// </summary>
        [Input("prometheusUserId")]
        public Input<int>? PrometheusUserId { get; set; }

        /// <summary>
        /// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
        /// </summary>
        [Input("regionSlug")]
        public Input<string>? RegionSlug { get; set; }

        /// <summary>
        /// Subdomain that the Grafana instance will be available at. Setting slug to `&lt;stack_slug&gt;` will make the instance available at `https://&lt;stack_slug&gt;.grafana.net`.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// Status of the stack.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tracesName")]
        public Input<string>? TracesName { get; set; }

        [Input("tracesStatus")]
        public Input<string>? TracesStatus { get; set; }

        /// <summary>
        /// Base URL of the Traces instance configured for this stack. To use this in the Tempo data source in Grafana, append `/tempo` to the URL.
        /// </summary>
        [Input("tracesUrl")]
        public Input<string>? TracesUrl { get; set; }

        [Input("tracesUserId")]
        public Input<int>? TracesUserId { get; set; }

        /// <summary>
        /// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
        /// </summary>
        [Input("waitForReadiness")]
        public Input<bool>? WaitForReadiness { get; set; }

        /// <summary>
        /// How long to wait for readiness (if enabled). Defaults to `5m0s`.
        /// </summary>
        [Input("waitForReadinessTimeout")]
        public Input<string>? WaitForReadinessTimeout { get; set; }

        public CloudStackState()
        {
        }
        public static new CloudStackState Empty => new CloudStackState();
    }
}
