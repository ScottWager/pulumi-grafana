// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const _default = new grafana.OncallEscalationChain("default", {}, {
 *     provider: grafana.oncall,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/oncallEscalationChain:OncallEscalationChain escalation_chain_name {{escalation_chain_id}}
 * ```
 */
export class OncallEscalationChain extends pulumi.CustomResource {
    /**
     * Get an existing OncallEscalationChain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OncallEscalationChainState, opts?: pulumi.CustomResourceOptions): OncallEscalationChain {
        return new OncallEscalationChain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/oncallEscalationChain:OncallEscalationChain';

    /**
     * Returns true if the given object is an instance of OncallEscalationChain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OncallEscalationChain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OncallEscalationChain.__pulumiType;
    }

    /**
     * The name of the escalation chain.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    public readonly teamId!: pulumi.Output<string | undefined>;

    /**
     * Create a OncallEscalationChain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OncallEscalationChainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OncallEscalationChainArgs | OncallEscalationChainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OncallEscalationChainState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as OncallEscalationChainArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OncallEscalationChain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OncallEscalationChain resources.
 */
export interface OncallEscalationChainState {
    /**
     * The name of the escalation chain.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OncallEscalationChain resource.
 */
export interface OncallEscalationChainArgs {
    /**
     * The name of the escalation chain.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    teamId?: pulumi.Input<string>;
}
