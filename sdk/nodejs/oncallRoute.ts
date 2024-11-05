// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/routes/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const exampleSlackChannel = grafana.onCall.getSlackChannel({
 *     name: "example_slack_channel",
 * });
 * const _default = new grafana.oncall.EscalationChain("default", {name: "default"});
 * const exampleIntegration = new grafana.oncall.Integration("example_integration", {
 *     name: "Grafana Integration",
 *     type: "grafana",
 *     defaultRoute: {},
 * });
 * const exampleRoute = new grafana.oncall.Route("example_route", {
 *     integrationId: exampleIntegration.id,
 *     escalationChainId: _default.id,
 *     routingRegex: "us-(east|west)",
 *     position: 0,
 *     slack: {
 *         channelId: exampleSlackChannel.then(exampleSlackChannel => exampleSlackChannel.slackId),
 *         enabled: true,
 *     },
 *     telegram: {
 *         id: "ONCALLTELEGRAMID",
 *         enabled: true,
 *     },
 *     msteams: {
 *         id: "ONCALLMSTEAMSID",
 *         enabled: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/oncallRoute:OncallRoute name "{{ id }}"
 * ```
 *
 * @deprecated grafana.index/oncallroute.OncallRoute has been deprecated in favor of grafana.oncall/route.Route
 */
export class OncallRoute extends pulumi.CustomResource {
    /**
     * Get an existing OncallRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OncallRouteState, opts?: pulumi.CustomResourceOptions): OncallRoute {
        pulumi.log.warn("OncallRoute is deprecated: grafana.index/oncallroute.OncallRoute has been deprecated in favor of grafana.oncall/route.Route")
        return new OncallRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/oncallRoute:OncallRoute';

    /**
     * Returns true if the given object is an instance of OncallRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OncallRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OncallRoute.__pulumiType;
    }

    /**
     * The ID of the escalation chain.
     */
    public readonly escalationChainId!: pulumi.Output<string>;
    /**
     * The ID of the integration.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * MS teams-specific settings for a route.
     */
    public readonly msteams!: pulumi.Output<outputs.OncallRouteMsteams | undefined>;
    /**
     * The position of the route (starts from 0).
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
     */
    public readonly routingRegex!: pulumi.Output<string>;
    /**
     * The type of route. Can be jinja2, regex Defaults to `regex`.
     */
    public readonly routingType!: pulumi.Output<string | undefined>;
    /**
     * Slack-specific settings for a route.
     */
    public readonly slack!: pulumi.Output<outputs.OncallRouteSlack | undefined>;
    /**
     * Telegram-specific settings for a route.
     */
    public readonly telegram!: pulumi.Output<outputs.OncallRouteTelegram | undefined>;

    /**
     * Create a OncallRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/oncallroute.OncallRoute has been deprecated in favor of grafana.oncall/route.Route */
    constructor(name: string, args: OncallRouteArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/oncallroute.OncallRoute has been deprecated in favor of grafana.oncall/route.Route */
    constructor(name: string, argsOrState?: OncallRouteArgs | OncallRouteState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("OncallRoute is deprecated: grafana.index/oncallroute.OncallRoute has been deprecated in favor of grafana.oncall/route.Route")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OncallRouteState | undefined;
            resourceInputs["escalationChainId"] = state ? state.escalationChainId : undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["msteams"] = state ? state.msteams : undefined;
            resourceInputs["position"] = state ? state.position : undefined;
            resourceInputs["routingRegex"] = state ? state.routingRegex : undefined;
            resourceInputs["routingType"] = state ? state.routingType : undefined;
            resourceInputs["slack"] = state ? state.slack : undefined;
            resourceInputs["telegram"] = state ? state.telegram : undefined;
        } else {
            const args = argsOrState as OncallRouteArgs | undefined;
            if ((!args || args.escalationChainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'escalationChainId'");
            }
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            if ((!args || args.position === undefined) && !opts.urn) {
                throw new Error("Missing required property 'position'");
            }
            if ((!args || args.routingRegex === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routingRegex'");
            }
            resourceInputs["escalationChainId"] = args ? args.escalationChainId : undefined;
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["msteams"] = args ? args.msteams : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["routingRegex"] = args ? args.routingRegex : undefined;
            resourceInputs["routingType"] = args ? args.routingType : undefined;
            resourceInputs["slack"] = args ? args.slack : undefined;
            resourceInputs["telegram"] = args ? args.telegram : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OncallRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OncallRoute resources.
 */
export interface OncallRouteState {
    /**
     * The ID of the escalation chain.
     */
    escalationChainId?: pulumi.Input<string>;
    /**
     * The ID of the integration.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * MS teams-specific settings for a route.
     */
    msteams?: pulumi.Input<inputs.OncallRouteMsteams>;
    /**
     * The position of the route (starts from 0).
     */
    position?: pulumi.Input<number>;
    /**
     * Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
     */
    routingRegex?: pulumi.Input<string>;
    /**
     * The type of route. Can be jinja2, regex Defaults to `regex`.
     */
    routingType?: pulumi.Input<string>;
    /**
     * Slack-specific settings for a route.
     */
    slack?: pulumi.Input<inputs.OncallRouteSlack>;
    /**
     * Telegram-specific settings for a route.
     */
    telegram?: pulumi.Input<inputs.OncallRouteTelegram>;
}

/**
 * The set of arguments for constructing a OncallRoute resource.
 */
export interface OncallRouteArgs {
    /**
     * The ID of the escalation chain.
     */
    escalationChainId: pulumi.Input<string>;
    /**
     * The ID of the integration.
     */
    integrationId: pulumi.Input<string>;
    /**
     * MS teams-specific settings for a route.
     */
    msteams?: pulumi.Input<inputs.OncallRouteMsteams>;
    /**
     * The position of the route (starts from 0).
     */
    position: pulumi.Input<number>;
    /**
     * Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
     */
    routingRegex: pulumi.Input<string>;
    /**
     * The type of route. Can be jinja2, regex Defaults to `regex`.
     */
    routingType?: pulumi.Input<string>;
    /**
     * Slack-specific settings for a route.
     */
    slack?: pulumi.Input<inputs.OncallRouteSlack>;
    /**
     * Telegram-specific settings for a route.
     */
    telegram?: pulumi.Input<inputs.OncallRouteTelegram>;
}
