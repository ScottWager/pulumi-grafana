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

    public sealed class SyntheticMonitoringCheckSettingsHttpBasicAuthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic auth password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Basic auth username.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SyntheticMonitoringCheckSettingsHttpBasicAuthArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsHttpBasicAuthArgs Empty => new SyntheticMonitoringCheckSettingsHttpBasicAuthArgs();
    }
}
