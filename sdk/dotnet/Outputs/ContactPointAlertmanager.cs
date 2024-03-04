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
    public sealed class ContactPointAlertmanager
    {
        /// <summary>
        /// The password component of the basic auth credentials to use.
        /// </summary>
        public readonly string? BasicAuthPassword;
        /// <summary>
        /// The username component of the basic auth credentials to use.
        /// </summary>
        public readonly string? BasicAuthUser;
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        public readonly bool? DisableResolveMessage;
        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Settings;
        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        public readonly string? Uid;
        /// <summary>
        /// The URL of the Alertmanager instance.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private ContactPointAlertmanager(
            string? basicAuthPassword,

            string? basicAuthUser,

            bool? disableResolveMessage,

            ImmutableDictionary<string, string>? settings,

            string? uid,

            string url)
        {
            BasicAuthPassword = basicAuthPassword;
            BasicAuthUser = basicAuthUser;
            DisableResolveMessage = disableResolveMessage;
            Settings = settings;
            Uid = uid;
            Url = url;
        }
    }
}
