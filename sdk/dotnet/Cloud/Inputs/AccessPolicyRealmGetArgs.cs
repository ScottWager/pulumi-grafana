// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Cloud.Inputs
{

    public sealed class AccessPolicyRealmGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        [Input("labelPolicies")]
        private InputList<Inputs.AccessPolicyRealmLabelPolicyGetArgs>? _labelPolicies;
        public InputList<Inputs.AccessPolicyRealmLabelPolicyGetArgs> LabelPolicies
        {
            get => _labelPolicies ?? (_labelPolicies = new InputList<Inputs.AccessPolicyRealmLabelPolicyGetArgs>());
            set => _labelPolicies = value;
        }

        /// <summary>
        /// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AccessPolicyRealmGetArgs()
        {
        }
        public static new AccessPolicyRealmGetArgs Empty => new AccessPolicyRealmGetArgs();
    }
}