// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@pulumiverse/grafana";
 *
 * const testFolderFolder = new grafana.oss.Folder("testFolderFolder", {title: "Terraform Test Folder"});
 * const testFolderDashboard = new grafana.oss.Dashboard("testFolderDashboard", {
 *     folder: testFolderFolder.id,
 *     configJson: `{
 *   "title": "Dashboard in folder",
 *   "uid": "dashboard-in-folder"
 * }
 * `,
 * });
 * const testFolderWithUid = new grafana.oss.Folder("testFolderWithUid", {
 *     uid: "test-folder-uid",
 *     title: "Terraform Test Folder With UID",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import grafana:index/folder:Folder name "{{ uid }}"
 * ```
 *
 * ```sh
 * $ pulumi import grafana:index/folder:Folder name "{{ orgID }}:{{ uid }}"
 * ```
 *
 * @deprecated grafana.index/folder.Folder has been deprecated in favor of grafana.oss/folder.Folder
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState, opts?: pulumi.CustomResourceOptions): Folder {
        pulumi.log.warn("Folder is deprecated: grafana.index/folder.Folder has been deprecated in favor of grafana.oss/folder.Folder")
        return new Folder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/folder:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Folder.__pulumiType;
    }

    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
     */
    public readonly parentFolderUid!: pulumi.Output<string | undefined>;
    /**
     * Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). This feature requires Grafana 10.2 or later. Defaults to `false`.
     */
    public readonly preventDestroyIfNotEmpty!: pulumi.Output<boolean | undefined>;
    /**
     * The title of the folder.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Unique identifier.
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * The full URL of the folder.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated grafana.index/folder.Folder has been deprecated in favor of grafana.oss/folder.Folder */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated grafana.index/folder.Folder has been deprecated in favor of grafana.oss/folder.Folder */
    constructor(name: string, argsOrState?: FolderArgs | FolderState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Folder is deprecated: grafana.index/folder.Folder has been deprecated in favor of grafana.oss/folder.Folder")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FolderState | undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["parentFolderUid"] = state ? state.parentFolderUid : undefined;
            resourceInputs["preventDestroyIfNotEmpty"] = state ? state.preventDestroyIfNotEmpty : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as FolderArgs | undefined;
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["parentFolderUid"] = args ? args.parentFolderUid : undefined;
            resourceInputs["preventDestroyIfNotEmpty"] = args ? args.preventDestroyIfNotEmpty : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["uid"] = args ? args.uid : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "grafana:index/folder:Folder" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Folder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Folder resources.
 */
export interface FolderState {
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
     */
    parentFolderUid?: pulumi.Input<string>;
    /**
     * Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). This feature requires Grafana 10.2 or later. Defaults to `false`.
     */
    preventDestroyIfNotEmpty?: pulumi.Input<boolean>;
    /**
     * The title of the folder.
     */
    title?: pulumi.Input<string>;
    /**
     * Unique identifier.
     */
    uid?: pulumi.Input<string>;
    /**
     * The full URL of the folder.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The uid of the parent folder. If set, the folder will be nested. If not set, the folder will be created in the root folder. Note: This requires the nestedFolders feature flag to be enabled on your Grafana instance.
     */
    parentFolderUid?: pulumi.Input<string>;
    /**
     * Prevent deletion of the folder if it is not empty (contains dashboards or alert rules). This feature requires Grafana 10.2 or later. Defaults to `false`.
     */
    preventDestroyIfNotEmpty?: pulumi.Input<boolean>;
    /**
     * The title of the folder.
     */
    title: pulumi.Input<string>;
    /**
     * Unique identifier.
     */
    uid?: pulumi.Input<string>;
}
