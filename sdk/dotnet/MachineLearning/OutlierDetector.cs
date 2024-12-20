// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.MachineLearning
{
    /// <summary>
    /// An outlier detector monitors the results of a query and reports when its values are outside normal bands.
    /// 
    /// The normal band is configured by choice of algorithm, its sensitivity and other configuration.
    /// 
    /// Visit https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for more details.
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:machineLearning/outlierDetector:OutlierDetector name "{{ id }}"
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:machineLearning/outlierDetector:OutlierDetector")]
    public partial class OutlierDetector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        /// </summary>
        [Output("algorithm")]
        public Output<Outputs.OutlierDetectorAlgorithm> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Output("datasourceType")]
        public Output<string> DatasourceType { get; private set; } = null!;

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Output("datasourceUid")]
        public Output<string> DatasourceUid { get; private set; } = null!;

        /// <summary>
        /// A description of the outlier detector.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The data interval in seconds to monitor. Defaults to `300`.
        /// </summary>
        [Output("interval")]
        public Output<int?> Interval { get; private set; } = null!;

        /// <summary>
        /// The metric used to query the outlier detector results.
        /// </summary>
        [Output("metric")]
        public Output<string> Metric { get; private set; } = null!;

        /// <summary>
        /// The name of the outlier detector.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        [Output("queryParams")]
        public Output<ImmutableDictionary<string, string>> QueryParams { get; private set; } = null!;


        /// <summary>
        /// Create a OutlierDetector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutlierDetector(string name, OutlierDetectorArgs args, CustomResourceOptions? options = null)
            : base("grafana:machineLearning/outlierDetector:OutlierDetector", name, args ?? new OutlierDetectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OutlierDetector(string name, Input<string> id, OutlierDetectorState? state = null, CustomResourceOptions? options = null)
            : base("grafana:machineLearning/outlierDetector:OutlierDetector", name, state, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OutlierDetector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutlierDetector Get(string name, Input<string> id, OutlierDetectorState? state = null, CustomResourceOptions? options = null)
        {
            return new OutlierDetector(name, id, state, options);
        }
    }

    public sealed class OutlierDetectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<Inputs.OutlierDetectorAlgorithmArgs> Algorithm { get; set; } = null!;

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Input("datasourceType", required: true)]
        public Input<string> DatasourceType { get; set; } = null!;

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Input("datasourceUid", required: true)]
        public Input<string> DatasourceUid { get; set; } = null!;

        /// <summary>
        /// A description of the outlier detector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The data interval in seconds to monitor. Defaults to `300`.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The metric used to query the outlier detector results.
        /// </summary>
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        /// <summary>
        /// The name of the outlier detector.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queryParams", required: true)]
        private InputMap<string>? _queryParams;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        public InputMap<string> QueryParams
        {
            get => _queryParams ?? (_queryParams = new InputMap<string>());
            set => _queryParams = value;
        }

        public OutlierDetectorArgs()
        {
        }
        public static new OutlierDetectorArgs Empty => new OutlierDetectorArgs();
    }

    public sealed class OutlierDetectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm to use and its configuration. See https://grafana.com/docs/grafana-cloud/machine-learning/outlier-detection/ for details.
        /// </summary>
        [Input("algorithm")]
        public Input<Inputs.OutlierDetectorAlgorithmGetArgs>? Algorithm { get; set; }

        /// <summary>
        /// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
        /// </summary>
        [Input("datasourceType")]
        public Input<string>? DatasourceType { get; set; }

        /// <summary>
        /// The uid of the datasource to query.
        /// </summary>
        [Input("datasourceUid")]
        public Input<string>? DatasourceUid { get; set; }

        /// <summary>
        /// A description of the outlier detector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The data interval in seconds to monitor. Defaults to `300`.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The metric used to query the outlier detector results.
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        /// <summary>
        /// The name of the outlier detector.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queryParams")]
        private InputMap<string>? _queryParams;

        /// <summary>
        /// An object representing the query params to query Grafana with.
        /// </summary>
        public InputMap<string> QueryParams
        {
            get => _queryParams ?? (_queryParams = new InputMap<string>());
            set => _queryParams = value;
        }

        public OutlierDetectorState()
        {
        }
        public static new OutlierDetectorState Empty => new OutlierDetectorState();
    }
}
