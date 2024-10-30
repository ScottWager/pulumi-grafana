// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 *         * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const adminServiceAccount = new grafana.oss.ServiceAccount("admin", {
 *     name: "admin sa",
 *     role: "Admin",
 *     isDisabled: false,
 * });
 * const admin = grafana.oss.getServiceAccountOutput({
 *     name: adminServiceAccount.name,
 * });
 * ```
 */
export function getServiceAccount(args: GetServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:oss/getServiceAccount:getServiceAccount", {
        "name": args.name,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccount.
 */
export interface GetServiceAccountArgs {
    /**
     * The name of the Service Account.
     */
    name: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: string;
}

/**
 * A collection of values returned by getServiceAccount.
 */
export interface GetServiceAccountResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The disabled status for the service account.
     */
    readonly isDisabled: boolean;
    /**
     * The name of the Service Account.
     */
    readonly name: string;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    readonly orgId?: string;
    /**
     * The basic role of the service account in the organization.
     */
    readonly role: string;
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 *         * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const adminServiceAccount = new grafana.oss.ServiceAccount("admin", {
 *     name: "admin sa",
 *     role: "Admin",
 *     isDisabled: false,
 * });
 * const admin = grafana.oss.getServiceAccountOutput({
 *     name: adminServiceAccount.name,
 * });
 * ```
 */
export function getServiceAccountOutput(args: GetServiceAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:oss/getServiceAccount:getServiceAccount", {
        "name": args.name,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccount.
 */
export interface GetServiceAccountOutputArgs {
    /**
     * The name of the Service Account.
     */
    name: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
}
