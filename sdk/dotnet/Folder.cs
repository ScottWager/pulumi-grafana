// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testFolderFolder = new Grafana.Folder("testFolderFolder", new()
    ///     {
    ///         Title = "Terraform Test Folder",
    ///     });
    /// 
    ///     var testFolderDashboard = new Grafana.Dashboard("testFolderDashboard", new()
    ///     {
    ///         Folder = testFolderFolder.Id,
    ///         ConfigJson = @"{
    ///   ""title"": ""Dashboard in folder"",
    ///   ""uid"": ""dashboard-in-folder""
    /// }
    /// ",
    ///     });
    /// 
    ///     var testFolderWithUid = new Grafana.Folder("testFolderWithUid", new()
    ///     {
    ///         Uid = "test-folder-uid",
    ///         Title = "Terraform Test Folder With UID",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/folder:Folder by_integer_id {{folder_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/folder:Folder by_uid {{folder_uid}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/folder:Folder")]
    public partial class Folder : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
        /// </summary>
        [Output("preventDestroyIfNotEmpty")]
        public Output<bool?> PreventDestroyIfNotEmpty { get; private set; } = null!;

        /// <summary>
        /// The title of the folder.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Unique identifier.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The full URL of the folder.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Folder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Folder(string name, FolderArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/folder:Folder", name, args ?? new FolderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Folder(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/folder:Folder", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Folder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Folder Get(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
        {
            return new Folder(name, id, state, options);
        }
    }

    public sealed class FolderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
        /// </summary>
        [Input("preventDestroyIfNotEmpty")]
        public Input<bool>? PreventDestroyIfNotEmpty { get; set; }

        /// <summary>
        /// The title of the folder.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Unique identifier.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public FolderArgs()
        {
        }
        public static new FolderArgs Empty => new FolderArgs();
    }

    public sealed class FolderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). Defaults to `false`.
        /// </summary>
        [Input("preventDestroyIfNotEmpty")]
        public Input<bool>? PreventDestroyIfNotEmpty { get; set; }

        /// <summary>
        /// The title of the folder.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Unique identifier.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The full URL of the folder.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FolderState()
        {
        }
        public static new FolderState Empty => new FolderState();
    }
}
