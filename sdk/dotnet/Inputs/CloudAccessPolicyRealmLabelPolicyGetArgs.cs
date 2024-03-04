// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Inputs
{

    public sealed class CloudAccessPolicyRealmLabelPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
        /// </summary>
        [Input("selector", required: true)]
        public Input<string> Selector { get; set; } = null!;

        public CloudAccessPolicyRealmLabelPolicyGetArgs()
        {
        }
        public static new CloudAccessPolicyRealmLabelPolicyGetArgs Empty => new CloudAccessPolicyRealmLabelPolicyGetArgs();
    }
}
