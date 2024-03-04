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

    public sealed class OncallIntegrationDefaultRouteMsteamsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable notification in MS teams. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public OncallIntegrationDefaultRouteMsteamsGetArgs()
        {
        }
        public static new OncallIntegrationDefaultRouteMsteamsGetArgs Empty => new OncallIntegrationDefaultRouteMsteamsGetArgs();
    }
}
