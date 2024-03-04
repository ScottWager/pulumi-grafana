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

    public sealed class NotificationPolicyPolicyPolicyPolicyMatcherGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the label to match against.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The operator to apply when matching values of the given label. Allowed operators are `=` for equality, `!=` for negated equality, `=~` for regex equality, and `!~` for negated regex equality.
        /// </summary>
        [Input("match", required: true)]
        public Input<string> Match { get; set; } = null!;

        /// <summary>
        /// The label value to match against.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NotificationPolicyPolicyPolicyPolicyMatcherGetArgs()
        {
        }
        public static new NotificationPolicyPolicyPolicyPolicyMatcherGetArgs Empty => new NotificationPolicyPolicyPolicyPolicyMatcherGetArgs();
    }
}
