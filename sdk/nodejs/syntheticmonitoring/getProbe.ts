// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for retrieving a single probe by name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const atlanta = grafana.syntheticMonitoring.getProbe({
 *     name: "Atlanta",
 * });
 * ```
 */
export function getProbe(args: GetProbeArgs, opts?: pulumi.InvokeOptions): Promise<GetProbeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:syntheticMonitoring/getProbe:getProbe", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProbe.
 */
export interface GetProbeArgs {
    /**
     * Name of the probe.
     */
    name: string;
}

/**
 * A collection of values returned by getProbe.
 */
export interface GetProbeResult {
    /**
     * Disables scripted checks for this probe.
     */
    readonly disableScriptedChecks: boolean;
    /**
     * The ID of the probe.
     */
    readonly id: string;
    /**
     * Custom labels to be included with collected metrics and logs.
     */
    readonly labels: {[key: string]: string};
    /**
     * Latitude coordinates.
     */
    readonly latitude: number;
    /**
     * Longitude coordinates.
     */
    readonly longitude: number;
    /**
     * Name of the probe.
     */
    readonly name: string;
    /**
     * Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`.
     */
    readonly public: boolean;
    /**
     * Region of the probe.
     */
    readonly region: string;
    /**
     * The tenant ID of the probe.
     */
    readonly tenantId: number;
}
/**
 * Data source for retrieving a single probe by name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const atlanta = grafana.syntheticMonitoring.getProbe({
 *     name: "Atlanta",
 * });
 * ```
 */
export function getProbeOutput(args: GetProbeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProbeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:syntheticMonitoring/getProbe:getProbe", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProbe.
 */
export interface GetProbeOutputArgs {
    /**
     * Name of the probe.
     */
    name: pulumi.Input<string>;
}
