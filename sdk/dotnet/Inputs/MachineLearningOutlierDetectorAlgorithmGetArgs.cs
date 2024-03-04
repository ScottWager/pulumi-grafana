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

    public sealed class MachineLearningOutlierDetectorAlgorithmGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For DBSCAN only, specify the configuration map
        /// </summary>
        [Input("config")]
        public Input<Inputs.MachineLearningOutlierDetectorAlgorithmConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// The name of the algorithm to use ('mad' or 'dbscan').
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specify the sensitivity of the detector (in range [0,1]).
        /// </summary>
        [Input("sensitivity", required: true)]
        public Input<double> Sensitivity { get; set; } = null!;

        public MachineLearningOutlierDetectorAlgorithmGetArgs()
        {
        }
        public static new MachineLearningOutlierDetectorAlgorithmGetArgs Empty => new MachineLearningOutlierDetectorAlgorithmGetArgs();
    }
}