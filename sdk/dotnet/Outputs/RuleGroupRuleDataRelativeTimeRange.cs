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
    public sealed class RuleGroupRuleDataRelativeTimeRange
    {
        /// <summary>
        /// The number of seconds in the past, relative to when the rule is evaluated, at which the time range begins.
        /// </summary>
        public readonly int From;
        /// <summary>
        /// The number of seconds in the past, relative to when the rule is evaluated, at which the time range ends.
        /// </summary>
        public readonly int To;

        [OutputConstructor]
        private RuleGroupRuleDataRelativeTimeRange(
            int from,

            int to)
        {
            From = from;
            To = to;
        }
    }
}
