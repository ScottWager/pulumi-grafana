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
    public static class GetSlos
    {
        /// <summary>
        /// Datasource for retrieving all SLOs.
        /// 		
        /// * [Official documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/)
        /// * [API documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/api/)
        /// * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)
        /// </summary>
        public static Task<GetSlosResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlosResult>("grafana:index/getSlos:getSlos", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Datasource for retrieving all SLOs.
        /// 		
        /// * [Official documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/)
        /// * [API documentation](https://grafana.com/docs/grafana-cloud/alerting-and-irm/slo/api/)
        /// * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)
        /// </summary>
        public static Output<GetSlosResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlosResult>("grafana:index/getSlos:getSlos", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetSlosResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Returns a list of all SLOs"
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlosSloResult> Slos;

        [OutputConstructor]
        private GetSlosResult(
            string id,

            ImmutableArray<Outputs.GetSlosSloResult> slos)
        {
            Id = id;
            Slos = slos;
        }
    }
}