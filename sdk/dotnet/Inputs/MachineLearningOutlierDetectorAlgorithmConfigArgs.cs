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

    public sealed class MachineLearningOutlierDetectorAlgorithmConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the epsilon parameter (positive float)
        /// </summary>
        [Input("epsilon", required: true)]
        public Input<double> Epsilon { get; set; } = null!;

        public MachineLearningOutlierDetectorAlgorithmConfigArgs()
        {
        }
        public static new MachineLearningOutlierDetectorAlgorithmConfigArgs Empty => new MachineLearningOutlierDetectorAlgorithmConfigArgs();
    }
}
