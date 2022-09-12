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

    public sealed class FolderPermissionPermissionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// Manage permissions for `Viewer` or `Editor` roles.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// ID of the team to manage permissions for. Defaults to `0`.
        /// </summary>
        [Input("teamId")]
        public Input<int>? TeamId { get; set; }

        /// <summary>
        /// ID of the user to manage permissions for. Defaults to `0`.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public FolderPermissionPermissionGetArgs()
        {
        }
        public static new FolderPermissionPermissionGetArgs Empty => new FolderPermissionPermissionGetArgs();
    }
}