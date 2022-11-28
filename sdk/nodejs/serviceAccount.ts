// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * **Note:** This resource is available only with Grafana 9.1+.
 *
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 *
 * const admin = new grafana.ServiceAccount("admin", {
 *     isDisabled: false,
 *     role: "Admin",
 * });
 * ```
 */
export class ServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountState, opts?: pulumi.CustomResourceOptions): ServiceAccount {
        return new ServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/serviceAccount:ServiceAccount';

    /**
     * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccount.__pulumiType;
    }

    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    public readonly isDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the service account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The basic role of the service account in the organization.
     */
    public readonly role!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceAccountArgs | ServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceAccountState | undefined;
            resourceInputs["isDisabled"] = state ? state.isDisabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ServiceAccountArgs | undefined;
            resourceInputs["isDisabled"] = args ? args.isDisabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccount resources.
 */
export interface ServiceAccountState {
    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    isDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The basic role of the service account in the organization.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceAccount resource.
 */
export interface ServiceAccountArgs {
    /**
     * The disabled status for the service account. Defaults to `false`.
     */
    isDisabled?: pulumi.Input<boolean>;
    /**
     * The name of the service account.
     */
    name?: pulumi.Input<string>;
    /**
     * The basic role of the service account in the organization.
     */
    role?: pulumi.Input<string>;
}
