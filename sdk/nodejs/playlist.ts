// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/playlist/)
 */
export class Playlist extends pulumi.CustomResource {
    /**
     * Get an existing Playlist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlaylistState, opts?: pulumi.CustomResourceOptions): Playlist {
        return new Playlist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/playlist:Playlist';

    /**
     * Returns true if the given object is an instance of Playlist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Playlist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Playlist.__pulumiType;
    }

    public readonly interval!: pulumi.Output<string>;
    public readonly items!: pulumi.Output<outputs.PlaylistItem[]>;
    /**
     * The name of the playlist.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;

    /**
     * Create a Playlist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlaylistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlaylistArgs | PlaylistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PlaylistState | undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["items"] = state ? state.items : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
        } else {
            const args = argsOrState as PlaylistArgs | undefined;
            if ((!args || args.interval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interval'");
            }
            if ((!args || args.items === undefined) && !opts.urn) {
                throw new Error("Missing required property 'items'");
            }
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["items"] = args ? args.items : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Playlist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Playlist resources.
 */
export interface PlaylistState {
    interval?: pulumi.Input<string>;
    items?: pulumi.Input<pulumi.Input<inputs.PlaylistItem>[]>;
    /**
     * The name of the playlist.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Playlist resource.
 */
export interface PlaylistArgs {
    interval: pulumi.Input<string>;
    items: pulumi.Input<pulumi.Input<inputs.PlaylistItem>[]>;
    /**
     * The name of the playlist.
     */
    name?: pulumi.Input<string>;
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
}
