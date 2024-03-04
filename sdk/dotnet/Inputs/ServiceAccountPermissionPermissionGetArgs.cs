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

    public sealed class ServiceAccountPermissionPermissionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permission to associate with item. Must be `Edit` or `Admin`.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// ID of the team to manage permissions for. Specify either this or `user_id`. Defaults to `0`.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// ID of the user or service account to manage permissions for. Specify either this or `team_id`. Defaults to `0`.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public ServiceAccountPermissionPermissionGetArgs()
        {
        }
        public static new ServiceAccountPermissionPermissionGetArgs Empty => new ServiceAccountPermissionPermissionGetArgs();
    }
}
