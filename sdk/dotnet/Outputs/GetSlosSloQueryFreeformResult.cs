// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Outputs
{

    [OutputType]
    public sealed class GetSlosSloQueryFreeformResult
    {
        /// <summary>
        /// Freeform Query Field
        /// </summary>
        public readonly string Query;

        [OutputConstructor]
        private GetSlosSloQueryFreeformResult(string query)
        {
            Query = query;
        }
    }
}
