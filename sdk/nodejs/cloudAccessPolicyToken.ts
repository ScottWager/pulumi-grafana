// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/account-management/authentication-and-permissions/access-policies/)
 * * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-a-token)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumi/grafana";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const current = grafana.getCloudOrganization({
 *     slug: "<your org slug>",
 * });
 * const testCloudAccessPolicy = new grafana.CloudAccessPolicy("testCloudAccessPolicy", {
 *     region: "us",
 *     displayName: "My Policy",
 *     scopes: [
 *         "metrics:read",
 *         "logs:read",
 *     ],
 *     realms: [{
 *         type: "org",
 *         identifier: current.then(current => current.id),
 *         labelPolicies: [{
 *             selector: "{namespace=\"default\"}",
 *         }],
 *     }],
 * });
 * const testCloudAccessPolicyToken = new grafana.CloudAccessPolicyToken("testCloudAccessPolicyToken", {
 *     region: "us",
 *     accessPolicyId: testCloudAccessPolicy.policyId,
 *     displayName: "My Policy Token",
 *     expiresAt: "2023-01-01T00:00:00Z",
 * });
 * ```
 */
export class CloudAccessPolicyToken extends pulumi.CustomResource {
    /**
     * Get an existing CloudAccessPolicyToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudAccessPolicyTokenState, opts?: pulumi.CustomResourceOptions): CloudAccessPolicyToken {
        return new CloudAccessPolicyToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken';

    /**
     * Returns true if the given object is an instance of CloudAccessPolicyToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudAccessPolicyToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudAccessPolicyToken.__pulumiType;
    }

    /**
     * ID of the access policy for which to create a token.
     */
    public readonly accessPolicyId!: pulumi.Output<string>;
    /**
     * Creation date of the access policy token.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Display name of the access policy token. Defaults to the name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Expiration date of the access policy token. Does not expire by default.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * Name of the access policy token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Last update date of the access policy token.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a CloudAccessPolicyToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudAccessPolicyTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudAccessPolicyTokenArgs | CloudAccessPolicyTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudAccessPolicyTokenState | undefined;
            resourceInputs["accessPolicyId"] = state ? state.accessPolicyId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as CloudAccessPolicyTokenArgs | undefined;
            if ((!args || args.accessPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPolicyId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["accessPolicyId"] = args ? args.accessPolicyId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CloudAccessPolicyToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudAccessPolicyToken resources.
 */
export interface CloudAccessPolicyTokenState {
    /**
     * ID of the access policy for which to create a token.
     */
    accessPolicyId?: pulumi.Input<string>;
    /**
     * Creation date of the access policy token.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Display name of the access policy token. Defaults to the name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Expiration date of the access policy token. Does not expire by default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Name of the access policy token.
     */
    name?: pulumi.Input<string>;
    /**
     * Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    region?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
    /**
     * Last update date of the access policy token.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudAccessPolicyToken resource.
 */
export interface CloudAccessPolicyTokenArgs {
    /**
     * ID of the access policy for which to create a token.
     */
    accessPolicyId: pulumi.Input<string>;
    /**
     * Display name of the access policy token. Defaults to the name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Expiration date of the access policy token. Does not expire by default.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Name of the access policy token.
     */
    name?: pulumi.Input<string>;
    /**
     * Region of the access policy. Should be set to the same region as the access policy. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
     */
    region: pulumi.Input<string>;
}
