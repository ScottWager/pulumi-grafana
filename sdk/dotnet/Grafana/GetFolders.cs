// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana
{
    public static class GetFolders
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/dashboard-folders/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
        /// </summary>
        public static Task<GetFoldersResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFoldersResult>("grafana:index/getFolders:getFolders", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetFoldersResult
    {
        /// <summary>
        /// The Grafana instance's folders.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFoldersFolderResult> Folders;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFoldersResult(
            ImmutableArray<Outputs.GetFoldersFolderResult> folders,

            string id)
        {
            Folders = folders;
            Id = id;
        }
    }
}
