// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get details about a Grafana Datasource querying by either name, uid or ID
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const prometheus = new grafana.oss.DataSource("prometheus", {
 *     type: "prometheus",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const fromName = grafana.oss.getDataSourceOutput({
 *     name: prometheus.name,
 * });
 * const fromUid = grafana.oss.getDataSourceOutput({
 *     uid: prometheus.uid,
 * });
 * ```
 */
/** @deprecated grafana.index/getdatasource.getDataSource has been deprecated in favor of grafana.oss/getdatasource.getDataSource */
export function getDataSource(args?: GetDataSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceResult> {
    pulumi.log.warn("getDataSource is deprecated: grafana.index/getdatasource.getDataSource has been deprecated in favor of grafana.oss/getdatasource.getDataSource")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getDataSource:getDataSource", {
        "name": args.name,
        "orgId": args.orgId,
        "uid": args.uid,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataSource.
 */
export interface GetDataSourceArgs {
    name?: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: string;
    uid?: string;
}

/**
 * A collection of values returned by getDataSource.
 */
export interface GetDataSourceResult {
    /**
     * The method by which Grafana will access the data source: `proxy` or `direct`.
     */
    readonly accessMode: string;
    /**
     * Whether to enable basic auth for the data source.
     */
    readonly basicAuthEnabled: boolean;
    /**
     * Basic auth username.
     */
    readonly basicAuthUsername: string;
    /**
     * (Required by some data source types) The name of the database to use on the selected data source server.
     */
    readonly databaseName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether to set the data source as default. This should only be `true` to a single data source.
     */
    readonly isDefault: boolean;
    /**
     * Serialized JSON string containing the json data. This attribute can be used to pass configuration options to the data source. To figure out what options a datasource has available, see its docs or inspect the network data when saving it from the Grafana UI. Note that keys in this map are usually camelCased.
     */
    readonly jsonDataEncoded: string;
    readonly name: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    readonly orgId?: string;
    /**
     * The data source type. Must be one of the supported data source keywords.
     */
    readonly type: string;
    readonly uid: string;
    /**
     * The URL for the data source. The type of URL required varies depending on the chosen data source type.
     */
    readonly url: string;
    /**
     * (Required by some data source types) The username to use to authenticate to the data source.
     */
    readonly username: string;
}
/**
 * Get details about a Grafana Datasource querying by either name, uid or ID
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const prometheus = new grafana.oss.DataSource("prometheus", {
 *     type: "prometheus",
 *     uid: "prometheus-ds-test-uid",
 *     url: "https://my-instance.com",
 *     basicAuthEnabled: true,
 *     basicAuthUsername: "username",
 *     jsonDataEncoded: JSON.stringify({
 *         httpMethod: "POST",
 *         prometheusType: "Mimir",
 *         prometheusVersion: "2.4.0",
 *     }),
 *     secureJsonDataEncoded: JSON.stringify({
 *         basicAuthPassword: "password",
 *     }),
 * });
 * const fromName = grafana.oss.getDataSourceOutput({
 *     name: prometheus.name,
 * });
 * const fromUid = grafana.oss.getDataSourceOutput({
 *     uid: prometheus.uid,
 * });
 * ```
 */
/** @deprecated grafana.index/getdatasource.getDataSource has been deprecated in favor of grafana.oss/getdatasource.getDataSource */
export function getDataSourceOutput(args?: GetDataSourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceResult> {
    return pulumi.output(args).apply((a: any) => getDataSource(a, opts))
}

/**
 * A collection of arguments for invoking getDataSource.
 */
export interface GetDataSourceOutputArgs {
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    uid?: pulumi.Input<string>;
}
