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
    /// Use the `team_sync` attribute of the `grafana.Team` resource instead.
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
    ///     var test_team_group = new Grafana.TeamExternalGroup("test-team-group", new()
    ///     {
    ///         Groups = new[]
    ///         {
    ///             "test-group-1",
    ///             "test-group-2",
    ///         },
    ///         TeamId = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/teamExternalGroup:TeamExternalGroup main {{team_id}}
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/teamExternalGroup:TeamExternalGroup")]
    public partial class TeamExternalGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The team external groups list
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// The Team ID
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a TeamExternalGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamExternalGroup(string name, TeamExternalGroupArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/teamExternalGroup:TeamExternalGroup", name, args ?? new TeamExternalGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamExternalGroup(string name, Input<string> id, TeamExternalGroupState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/teamExternalGroup:TeamExternalGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamExternalGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamExternalGroup Get(string name, Input<string> id, TeamExternalGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamExternalGroup(name, id, state, options);
        }
    }

    public sealed class TeamExternalGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups", required: true)]
        private InputList<string>? _groups;

        /// <summary>
        /// The team external groups list
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The Team ID
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        public TeamExternalGroupArgs()
        {
        }
        public static new TeamExternalGroupArgs Empty => new TeamExternalGroupArgs();
    }

    public sealed class TeamExternalGroupState : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The team external groups list
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The Team ID
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public TeamExternalGroupState()
        {
        }
        public static new TeamExternalGroupState Empty => new TeamExternalGroupState();
    }
}
