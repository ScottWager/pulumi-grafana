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

    public sealed class ContactPointSnGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// AWS access key ID used to authenticate with Amazon SNS.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the role to assume to send notifications to Amazon SNS.
        /// </summary>
        [Input("assumeRoleArn")]
        public Input<string>? AssumeRoleArn { get; set; }

        /// <summary>
        /// The authentication provider to use. Valid values are `default`, `arn` and `keys`. Default is `default`. Defaults to `default`.
        /// </summary>
        [Input("authProvider")]
        public Input<string>? AuthProvider { get; set; }

        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Whether to disable sending resolve messages. Defaults to `false`.
        /// </summary>
        [Input("disableResolveMessage")]
        public Input<bool>? DisableResolveMessage { get; set; }

        /// <summary>
        /// The external ID to use when assuming the role.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The format of the message to send. Valid values are `text`, `body` and `json`. Default is `text`. Defaults to `text`.
        /// </summary>
        [Input("messageFormat")]
        public Input<string>? MessageFormat { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// AWS secret access key used to authenticate with Amazon SNS.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// The Amazon SNS topic to send notifications to.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        /// <summary>
        /// The UID of the contact point.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public ContactPointSnGetArgs()
        {
        }
        public static new ContactPointSnGetArgs Empty => new ContactPointSnGetArgs();
    }
}