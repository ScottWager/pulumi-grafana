// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.DashboardArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.DashboardState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Grafana dashboards.
 * 
 * * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/dashboard/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.Dashboard;
 * import com.pulumi.grafana.DashboardArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var metrics = new Dashboard(&#34;metrics&#34;, DashboardArgs.builder()        
 *             .configJson(Files.readString(Paths.get(&#34;grafana-dashboard.json&#34;)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/dashboard:Dashboard dashboard_name {{dashboard_uid}} # To use the default provider org
 * ```
 * 
 * ```sh
 *  $ pulumi import grafana:index/dashboard:Dashboard dashboard_name {{org_id}}:{{dashboard_uid}} # When &#34;org_id&#34; is set on the resource
 * ```
 * 
 */
@ResourceType(type="grafana:index/dashboard:Dashboard")
public class Dashboard extends com.pulumi.resources.CustomResource {
    /**
     * The complete dashboard model JSON.
     * 
     */
    @Export(name="configJson", refs={String.class}, tree="[0]")
    private Output<String> configJson;

    /**
     * @return The complete dashboard model JSON.
     * 
     */
    public Output<String> configJson() {
        return this.configJson;
    }
    /**
     * The numeric ID of the dashboard computed by Grafana.
     * 
     */
    @Export(name="dashboardId", refs={Integer.class}, tree="[0]")
    private Output<Integer> dashboardId;

    /**
     * @return The numeric ID of the dashboard computed by Grafana.
     * 
     */
    public Output<Integer> dashboardId() {
        return this.dashboardId;
    }
    /**
     * The id or UID of the folder to save the dashboard in.
     * 
     */
    @Export(name="folder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> folder;

    /**
     * @return The id or UID of the folder to save the dashboard in.
     * 
     */
    public Output<Optional<String>> folder() {
        return Codegen.optional(this.folder);
    }
    /**
     * Set a commit message for the version history.
     * 
     */
    @Export(name="message", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> message;

    /**
     * @return Set a commit message for the version history.
     * 
     */
    public Output<Optional<String>> message() {
        return Codegen.optional(this.message);
    }
    /**
     * The Organization ID. If not set, the Org ID defined in the provider block will be used.
     * 
     */
    @Export(name="orgId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> orgId;

    /**
     * @return The Organization ID. If not set, the Org ID defined in the provider block will be used.
     * 
     */
    public Output<Optional<String>> orgId() {
        return Codegen.optional(this.orgId);
    }
    /**
     * Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
     * 
     */
    @Export(name="overwrite", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> overwrite;

    /**
     * @return Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
     * 
     */
    public Output<Optional<Boolean>> overwrite() {
        return Codegen.optional(this.overwrite);
    }
    /**
     * URL friendly version of the dashboard title. This field is deprecated, please use `uid` instead.
     * 
     * @deprecated
     * Use `uid` instead.
     * 
     */
    @Deprecated /* Use `uid` instead. */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return URL friendly version of the dashboard title. This field is deprecated, please use `uid` instead.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }
    /**
     * The unique identifier of a dashboard. This is used to construct its URL. It&#39;s automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The unique identifier of a dashboard. This is used to construct its URL. It&#39;s automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * The full URL of the dashboard.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The full URL of the dashboard.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
     * 
     */
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    /**
     * @return Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
     * 
     */
    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dashboard(String name) {
        this(name, DashboardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dashboard(String name, DashboardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dashboard(String name, DashboardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/dashboard:Dashboard", name, args == null ? DashboardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Dashboard(String name, Output<String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/dashboard:Dashboard", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Dashboard get(String name, Output<String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dashboard(name, id, state, options);
    }
}
