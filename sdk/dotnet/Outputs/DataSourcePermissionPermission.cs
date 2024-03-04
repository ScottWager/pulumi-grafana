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
    public sealed class DataSourcePermissionPermission
    {
        /// <summary>
        /// Name of the basic role to manage permissions for. Options: `Viewer`, `Editor` or `Admin`. Can only be set from Grafana v9.2.3+. Defaults to ``.
        /// </summary>
        public readonly string? BuiltInRole;
        /// <summary>
        /// Permission to associate with item. Options: `Query` or `Edit` (`Edit` can only be used with Grafana v9.2.3+).
        /// </summary>
        public readonly string Permission;
        /// <summary>
        /// ID of the team to manage permissions for. Defaults to `0`.
        /// </summary>
        public readonly string? TeamId;
        /// <summary>
        /// ID of the user or service account to manage permissions for. Defaults to `0`.
        /// </summary>
        public readonly string? UserId;

        [OutputConstructor]
        private DataSourcePermissionPermission(
            string? builtInRole,

            string permission,

            string? teamId,

            string? userId)
        {
            BuiltInRole = builtInRole;
            Permission = permission;
            TeamId = teamId;
            UserId = userId;
        }
    }
}
