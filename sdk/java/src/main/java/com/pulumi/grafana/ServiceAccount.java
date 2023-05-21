// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.ServiceAccountArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.ServiceAccountState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * **Note:** This resource is available only with Grafana 9.1+.
 * 
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/service-accounts/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.ServiceAccount;
 * import com.pulumi.grafana.ServiceAccountArgs;
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
 *         var admin = new ServiceAccount(&#34;admin&#34;, ServiceAccountArgs.builder()        
 *             .isDisabled(false)
 *             .role(&#34;Admin&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="grafana:index/serviceAccount:ServiceAccount")
public class ServiceAccount extends com.pulumi.resources.CustomResource {
    /**
     * The disabled status for the service account. Defaults to `false`.
     * 
     */
    @Export(name="isDisabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isDisabled;

    /**
     * @return The disabled status for the service account. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> isDisabled() {
        return Codegen.optional(this.isDisabled);
    }
    /**
     * The name of the service account.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the service account.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The basic role of the service account in the organization.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return The basic role of the service account in the organization.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceAccount(String name) {
        this(name, ServiceAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceAccount(String name, @Nullable ServiceAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceAccount(String name, @Nullable ServiceAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/serviceAccount:ServiceAccount", name, args == null ? ServiceAccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceAccount(String name, Output<String> id, @Nullable ServiceAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/serviceAccount:ServiceAccount", name, state, makeResourceOptions(options, id));
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
    public static ServiceAccount get(String name, Output<String> id, @Nullable ServiceAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceAccount(name, id, state, options);
    }
}
