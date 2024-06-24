// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource is deprecated and will be removed in a future release. Please use grafana.cloud.AccessPolicy instead.
 *
 * Manages a single API key on the Grafana Cloud portal (on the organization level)
 * * [API documentation](https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#api-keys)
 *
 * Required access policy scopes:
 *
 * * api-keys:read
 * * api-keys:write
 * * api-keys:delete
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const test = new grafana.cloud.ApiKey("test", {
 *     cloudOrgSlug: "myorg",
 *     role: "Admin",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:cloud/apiKey:ApiKey name "{{ orgSlug }}:{{ apiKeyName }}"
 * ```
 */
export class ApiKey extends pulumi.CustomResource {
    /**
     * Get an existing ApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiKeyState, opts?: pulumi.CustomResourceOptions): ApiKey {
        return new ApiKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:cloud/apiKey:ApiKey';

    /**
     * Returns true if the given object is an instance of ApiKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiKey.__pulumiType;
    }

    /**
     * The slug of the organization to create the API key in. This is the same slug as the organization name in the URL.
     */
    public readonly cloudOrgSlug!: pulumi.Output<string>;
    /**
     * The generated API key.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * Name of the API key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Role of the API key. Should be one of [Viewer Editor Admin MetricsPublisher PluginPublisher]. See https://grafana.com/docs/grafana-cloud/api/#create-api-key for details.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiKeyArgs | ApiKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiKeyState | undefined;
            resourceInputs["cloudOrgSlug"] = state ? state.cloudOrgSlug : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ApiKeyArgs | undefined;
            if ((!args || args.cloudOrgSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudOrgSlug'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["cloudOrgSlug"] = args ? args.cloudOrgSlug : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["key"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/cloudApiKey:CloudApiKey" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApiKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiKey resources.
 */
export interface ApiKeyState {
    /**
     * The slug of the organization to create the API key in. This is the same slug as the organization name in the URL.
     */
    cloudOrgSlug?: pulumi.Input<string>;
    /**
     * The generated API key.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the API key.
     */
    name?: pulumi.Input<string>;
    /**
     * Role of the API key. Should be one of [Viewer Editor Admin MetricsPublisher PluginPublisher]. See https://grafana.com/docs/grafana-cloud/api/#create-api-key for details.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiKey resource.
 */
export interface ApiKeyArgs {
    /**
     * The slug of the organization to create the API key in. This is the same slug as the organization name in the URL.
     */
    cloudOrgSlug: pulumi.Input<string>;
    /**
     * Name of the API key.
     */
    name?: pulumi.Input<string>;
    /**
     * Role of the API key. Should be one of [Viewer Editor Admin MetricsPublisher PluginPublisher]. See https://grafana.com/docs/grafana-cloud/api/#create-api-key for details.
     */
    role: pulumi.Input<string>;
}