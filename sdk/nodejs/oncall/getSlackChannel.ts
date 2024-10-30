// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/slack_channels/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const exampleSlackChannel = grafana.onCall.getSlackChannel({
 *     name: "example_slack_channel",
 * });
 * ```
 */
export function getSlackChannel(args: GetSlackChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetSlackChannelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("grafana:onCall/getSlackChannel:getSlackChannel", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlackChannel.
 */
export interface GetSlackChannelArgs {
    /**
     * The Slack channel name.
     */
    name: string;
}

/**
 * A collection of values returned by getSlackChannel.
 */
export interface GetSlackChannelResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Slack channel name.
     */
    readonly name: string;
    /**
     * The Slack ID of the channel.
     */
    readonly slackId: string;
}
/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/slack_channels/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const exampleSlackChannel = grafana.onCall.getSlackChannel({
 *     name: "example_slack_channel",
 * });
 * ```
 */
export function getSlackChannelOutput(args: GetSlackChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlackChannelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("grafana:onCall/getSlackChannel:getSlackChannel", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlackChannel.
 */
export interface GetSlackChannelOutputArgs {
    /**
     * The Slack channel name.
     */
    name: pulumi.Input<string>;
}
