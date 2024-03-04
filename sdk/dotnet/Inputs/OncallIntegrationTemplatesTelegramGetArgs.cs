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

    public sealed class OncallIntegrationTemplatesTelegramGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template for Alert image url.
        /// </summary>
        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        /// <summary>
        /// Template for Alert message.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Template for Alert title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public OncallIntegrationTemplatesTelegramGetArgs()
        {
        }
        public static new OncallIntegrationTemplatesTelegramGetArgs Empty => new OncallIntegrationTemplatesTelegramGetArgs();
    }
}
