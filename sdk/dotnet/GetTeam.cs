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
    [Obsolete(@"grafana.index/getteam.getTeam has been deprecated in favor of grafana.oss/getteam.getTeam")]
    public static class GetTeam
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// using Grafana = Pulumiverse.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = new Grafana.Oss.Team("test", new()
        ///     {
        ///         Email = "test-team-email@test.com",
        ///         Preferences = new Grafana.Oss.Inputs.TeamPreferencesArgs
        ///         {
        ///             Theme = "dark",
        ///             Timezone = "utc",
        ///         },
        ///     });
        /// 
        ///     var fromName = Grafana.Oss.GetTeam.Invoke(new()
        ///     {
        ///         Name = test.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("grafana:index/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// using Grafana = Pulumiverse.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = new Grafana.Oss.Team("test", new()
        ///     {
        ///         Email = "test-team-email@test.com",
        ///         Preferences = new Grafana.Oss.Inputs.TeamPreferencesArgs
        ///         {
        ///             Theme = "dark",
        ///             Timezone = "utc",
        ///         },
        ///     });
        /// 
        ///     var fromName = Grafana.Oss.GetTeam.Invoke(new()
        ///     {
        ///         Name = test.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("grafana:index/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("orgId")]
        public string? OrgId { get; set; }

        [Input("readTeamSync")]
        public bool? ReadTeamSync { get; set; }

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("readTeamSync")]
        public Input<bool>? ReadTeamSync { get; set; }

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Members;
        public readonly string Name;
        public readonly string? OrgId;
        public readonly ImmutableArray<Outputs.GetTeamPreferenceResult> Preferences;
        public readonly bool? ReadTeamSync;
        public readonly int TeamId;
        public readonly ImmutableArray<Outputs.GetTeamTeamSyncResult> TeamSyncs;

        [OutputConstructor]
        private GetTeamResult(
            string email,

            string id,

            ImmutableArray<string> members,

            string name,

            string? orgId,

            ImmutableArray<Outputs.GetTeamPreferenceResult> preferences,

            bool? readTeamSync,

            int teamId,

            ImmutableArray<Outputs.GetTeamTeamSyncResult> teamSyncs)
        {
            Email = email;
            Id = id;
            Members = members;
            Name = name;
            OrgId = orgId;
            Preferences = preferences;
            ReadTeamSync = readTeamSync;
            TeamId = teamId;
            TeamSyncs = teamSyncs;
        }
    }
}
