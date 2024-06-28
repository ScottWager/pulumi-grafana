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
    [Obsolete(@"grafana.index/getsyntheticmonitoringprobes.getSyntheticMonitoringProbes has been deprecated in favor of grafana.syntheticmonitoring/getprobes.getProbes")]
    public static class GetSyntheticMonitoringProbes
    {
        /// <summary>
        /// Data source for retrieving all probes.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSyntheticMonitoringProbesResult> InvokeAsync(GetSyntheticMonitoringProbesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSyntheticMonitoringProbesResult>("grafana:index/getSyntheticMonitoringProbes:getSyntheticMonitoringProbes", args ?? new GetSyntheticMonitoringProbesArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for retrieving all probes.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Grafana.SyntheticMonitoring.GetProbes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSyntheticMonitoringProbesResult> Invoke(GetSyntheticMonitoringProbesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyntheticMonitoringProbesResult>("grafana:index/getSyntheticMonitoringProbes:getSyntheticMonitoringProbes", args ?? new GetSyntheticMonitoringProbesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSyntheticMonitoringProbesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        [Input("filterDeprecated")]
        public bool? FilterDeprecated { get; set; }

        public GetSyntheticMonitoringProbesArgs()
        {
        }
        public static new GetSyntheticMonitoringProbesArgs Empty => new GetSyntheticMonitoringProbesArgs();
    }

    public sealed class GetSyntheticMonitoringProbesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        [Input("filterDeprecated")]
        public Input<bool>? FilterDeprecated { get; set; }

        public GetSyntheticMonitoringProbesInvokeArgs()
        {
        }
        public static new GetSyntheticMonitoringProbesInvokeArgs Empty => new GetSyntheticMonitoringProbesInvokeArgs();
    }


    [OutputType]
    public sealed class GetSyntheticMonitoringProbesResult
    {
        /// <summary>
        /// If true, only probes that are not deprecated will be returned. Defaults to `true`.
        /// </summary>
        public readonly bool? FilterDeprecated;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Map of probes with their names as keys and IDs as values.
        /// </summary>
        public readonly ImmutableDictionary<string, int> Probes;

        [OutputConstructor]
        private GetSyntheticMonitoringProbesResult(
            bool? filterDeprecated,

            string id,

            ImmutableDictionary<string, int> probes)
        {
            FilterDeprecated = filterDeprecated;
            Id = id;
            Probes = probes;
        }
    }
}
