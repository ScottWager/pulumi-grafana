// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana
{
    /// <summary>
    /// Manages Grafana Alerting mute timings.
    /// 
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/alerting/notifications/mute-timings/)
    /// * [HTTP API](https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#mute-timings)
    /// 
    /// This resource requires Grafana 9.1.0 or later.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Grafana = Lbrlabs.PulumiPackage.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myMuteTiming = new Grafana.MuteTiming("myMuteTiming", new()
    ///     {
    ///         Intervals = new[]
    ///         {
    ///             new Grafana.Inputs.MuteTimingIntervalArgs
    ///             {
    ///                 DaysOfMonths = new[]
    ///                 {
    ///                     "1:7",
    ///                     "-1",
    ///                 },
    ///                 Months = new[]
    ///                 {
    ///                     "1:3",
    ///                     "december",
    ///                 },
    ///                 Times = new[]
    ///                 {
    ///                     new Grafana.Inputs.MuteTimingIntervalTimeArgs
    ///                     {
    ///                         End = "14:17",
    ///                         Start = "04:56",
    ///                     },
    ///                 },
    ///                 Weekdays = new[]
    ///                 {
    ///                     "monday",
    ///                     "tuesday:thursday",
    ///                 },
    ///                 Years = new[]
    ///                 {
    ///                     "2030",
    ///                     "2025:2026",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/muteTiming:MuteTiming mute_timing_name {{mute_timing_name}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/muteTiming:MuteTiming")]
    public partial class MuteTiming : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time intervals at which to mute notifications.
        /// </summary>
        [Output("intervals")]
        public Output<ImmutableArray<Outputs.MuteTimingInterval>> Intervals { get; private set; } = null!;

        /// <summary>
        /// The name of the mute timing.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MuteTiming resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MuteTiming(string name, MuteTimingArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/muteTiming:MuteTiming", name, args ?? new MuteTimingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MuteTiming(string name, Input<string> id, MuteTimingState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/muteTiming:MuteTiming", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MuteTiming resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MuteTiming Get(string name, Input<string> id, MuteTimingState? state = null, CustomResourceOptions? options = null)
        {
            return new MuteTiming(name, id, state, options);
        }
    }

    public sealed class MuteTimingArgs : global::Pulumi.ResourceArgs
    {
        [Input("intervals")]
        private InputList<Inputs.MuteTimingIntervalArgs>? _intervals;

        /// <summary>
        /// The time intervals at which to mute notifications.
        /// </summary>
        public InputList<Inputs.MuteTimingIntervalArgs> Intervals
        {
            get => _intervals ?? (_intervals = new InputList<Inputs.MuteTimingIntervalArgs>());
            set => _intervals = value;
        }

        /// <summary>
        /// The name of the mute timing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MuteTimingArgs()
        {
        }
        public static new MuteTimingArgs Empty => new MuteTimingArgs();
    }

    public sealed class MuteTimingState : global::Pulumi.ResourceArgs
    {
        [Input("intervals")]
        private InputList<Inputs.MuteTimingIntervalGetArgs>? _intervals;

        /// <summary>
        /// The time intervals at which to mute notifications.
        /// </summary>
        public InputList<Inputs.MuteTimingIntervalGetArgs> Intervals
        {
            get => _intervals ?? (_intervals = new InputList<Inputs.MuteTimingIntervalGetArgs>());
            set => _intervals = value;
        }

        /// <summary>
        /// The name of the mute timing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MuteTimingState()
        {
        }
        public static new MuteTimingState Empty => new MuteTimingState();
    }
}
