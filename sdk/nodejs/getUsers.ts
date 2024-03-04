// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
 *
 * This data source uses Grafana's admin APIs for reading users which
 * does not currently work with API Tokens. You must use basic auth.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testAllUsers = new grafana.User("testAllUsers", {
 *     email: "all_users@example.com",
 *     login: "test-grafana-users",
 *     password: "my-password",
 * });
 * const allUsers = grafana.getUsers({});
 * ```
 */
export function getUsers(opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:index/getUsers:getUsers", {
    }, opts);
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Grafana instance's users.
     */
    readonly users: outputs.GetUsersUser[];
}
/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
 *
 * This data source uses Grafana's admin APIs for reading users which
 * does not currently work with API Tokens. You must use basic auth.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testAllUsers = new grafana.User("testAllUsers", {
 *     email: "all_users@example.com",
 *     login: "test-grafana-users",
 *     password: "my-password",
 * });
 * const allUsers = grafana.getUsers({});
 * ```
 */
export function getUsersOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(getUsers(opts))
}
