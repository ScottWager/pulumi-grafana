// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Inputs
{

    public sealed class OncallIntegrationTemplatesMicrosoftTeamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        public OncallIntegrationTemplatesMicrosoftTeamsArgs()
        {
        }
        public static new OncallIntegrationTemplatesMicrosoftTeamsArgs Empty => new OncallIntegrationTemplatesMicrosoftTeamsArgs();
    }
}
