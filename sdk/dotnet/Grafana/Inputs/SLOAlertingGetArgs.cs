// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana.Inputs
{

    public sealed class SLOAlertingGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<Inputs.SLOAlertingAnnotationGetArgs>? _annotations;

        /// <summary>
        /// Annotations will be attached to all alerts generated by any of these rules.
        /// </summary>
        public InputList<Inputs.SLOAlertingAnnotationGetArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.SLOAlertingAnnotationGetArgs>());
            set => _annotations = value;
        }

        [Input("fastburns")]
        private InputList<Inputs.SLOAlertingFastburnGetArgs>? _fastburns;

        /// <summary>
        /// Alerting Rules generated for Fast Burn alerts
        /// </summary>
        public InputList<Inputs.SLOAlertingFastburnGetArgs> Fastburns
        {
            get => _fastburns ?? (_fastburns = new InputList<Inputs.SLOAlertingFastburnGetArgs>());
            set => _fastburns = value;
        }

        [Input("labels")]
        private InputList<Inputs.SLOAlertingLabelGetArgs>? _labels;

        /// <summary>
        /// Labels will be attached to all alerts generated by any of these rules.
        /// </summary>
        public InputList<Inputs.SLOAlertingLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SLOAlertingLabelGetArgs>());
            set => _labels = value;
        }

        [Input("slowburns")]
        private InputList<Inputs.SLOAlertingSlowburnGetArgs>? _slowburns;

        /// <summary>
        /// Alerting Rules generated for Slow Burn alerts
        /// </summary>
        public InputList<Inputs.SLOAlertingSlowburnGetArgs> Slowburns
        {
            get => _slowburns ?? (_slowburns = new InputList<Inputs.SLOAlertingSlowburnGetArgs>());
            set => _slowburns = value;
        }

        public SLOAlertingGetArgs()
        {
        }
        public static new SLOAlertingGetArgs Empty => new SLOAlertingGetArgs();
    }
}
