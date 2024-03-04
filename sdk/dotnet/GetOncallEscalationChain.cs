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
    public static class GetOncallEscalationChain
    {
        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Grafana.GetOncallEscalationChain.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOncallEscalationChainResult> InvokeAsync(GetOncallEscalationChainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOncallEscalationChainResult>("grafana:index/getOncallEscalationChain:getOncallEscalationChain", args ?? new GetOncallEscalationChainArgs(), options.WithDefaults());

        /// <summary>
        /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Grafana.GetOncallEscalationChain.Invoke(new()
        ///     {
        ///         Name = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOncallEscalationChainResult> Invoke(GetOncallEscalationChainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOncallEscalationChainResult>("grafana:index/getOncallEscalationChain:getOncallEscalationChain", args ?? new GetOncallEscalationChainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOncallEscalationChainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetOncallEscalationChainArgs()
        {
        }
        public static new GetOncallEscalationChainArgs Empty => new GetOncallEscalationChainArgs();
    }

    public sealed class GetOncallEscalationChainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetOncallEscalationChainInvokeArgs()
        {
        }
        public static new GetOncallEscalationChainInvokeArgs Empty => new GetOncallEscalationChainInvokeArgs();
    }


    [OutputType]
    public sealed class GetOncallEscalationChainResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The escalation chain name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetOncallEscalationChainResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
