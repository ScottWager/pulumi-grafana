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

    public sealed class ContactPointTeamGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        /// <summary>
        /// The templated message content to send.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The templated subtitle for each message section.
        /// </summary>
        [Input("sectionTitle")]
        public Input<string>? SectionTitle { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Additional custom properties to attach to the notifier. Defaults to `map[]`.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _settings = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The templated title of the message.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        [Input("url", required: true)]
        private Input<string>? _url;

        /// <summary>
        /// A Teams webhook URL.
        /// </summary>
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ContactPointTeamGetArgs()
        {
        }
        public static new ContactPointTeamGetArgs Empty => new ContactPointTeamGetArgs();
    }
}
