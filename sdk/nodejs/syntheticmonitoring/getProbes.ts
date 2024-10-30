// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for retrieving all probes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.syntheticMonitoring.getProbes({});
 * ```
 */
export function getProbes(args?: GetProbesArgs, opts?: pulumi.InvokeOptions): Promise<GetProbesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:syntheticMonitoring/getProbes:getProbes", {
        "filterDeprecated": args.filterDeprecated,
    }, opts);
}

/**
 * A collection of arguments for invoking getProbes.
 */
export interface GetProbesArgs {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    filterDeprecated?: boolean;
}

/**
 * A collection of values returned by getProbes.
 */
export interface GetProbesResult {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    readonly filterDeprecated?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Map of probes with their names as keys and IDs as values.
     */
    readonly probes: {[key: string]: number};
}
/**
 * Data source for retrieving all probes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.syntheticMonitoring.getProbes({});
 * ```
 */
export function getProbesOutput(args?: GetProbesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProbesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:syntheticMonitoring/getProbes:getProbes", {
        "filterDeprecated": args.filterDeprecated,
    }, opts);
}

/**
 * A collection of arguments for invoking getProbes.
 */
export interface GetProbesOutputArgs {
    /**
     * If true, only probes that are not deprecated will be returned. Defaults to `true`.
     */
    filterDeprecated?: pulumi.Input<boolean>;
}
