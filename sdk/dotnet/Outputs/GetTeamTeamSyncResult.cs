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
    public sealed class GetTeamTeamSyncResult
    {
        public readonly ImmutableArray<string> Groups;

        [OutputConstructor]
        private GetTeamTeamSyncResult(ImmutableArray<string> groups)
        {
            Groups = groups;
        }
    }
}
