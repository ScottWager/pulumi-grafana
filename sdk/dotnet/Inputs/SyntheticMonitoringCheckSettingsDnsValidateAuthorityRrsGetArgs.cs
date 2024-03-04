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

    public sealed class SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("failIfMatchesRegexps")]
        private InputList<string>? _failIfMatchesRegexps;

        /// <summary>
        /// Fail if value matches regex.
        /// </summary>
        public InputList<string> FailIfMatchesRegexps
        {
            get => _failIfMatchesRegexps ?? (_failIfMatchesRegexps = new InputList<string>());
            set => _failIfMatchesRegexps = value;
        }

        [Input("failIfNotMatchesRegexps")]
        private InputList<string>? _failIfNotMatchesRegexps;

        /// <summary>
        /// Fail if value does not match regex.
        /// </summary>
        public InputList<string> FailIfNotMatchesRegexps
        {
            get => _failIfNotMatchesRegexps ?? (_failIfNotMatchesRegexps = new InputList<string>());
            set => _failIfNotMatchesRegexps = value;
        }

        public SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrsGetArgs()
        {
        }
        public static new SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrsGetArgs Empty => new SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrsGetArgs();
    }
}
