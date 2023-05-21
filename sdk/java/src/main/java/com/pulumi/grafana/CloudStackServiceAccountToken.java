// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.CloudStackServiceAccountTokenArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.CloudStackServiceAccountTokenState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * **Note:** This resource is available only with Grafana 9.1+.
 * 
 * Manages service account tokens of a Grafana Cloud stack using the Cloud API
 * This can be used to bootstrap a management service account token for a new stack
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
 * import com.pulumi.grafana.CloudStackServiceAccount;
 * import com.pulumi.grafana.CloudStackServiceAccountArgs;
 * import com.pulumi.grafana.CloudStackServiceAccountToken;
 * import com.pulumi.grafana.CloudStackServiceAccountTokenArgs;
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
 *         var cloudSa = new CloudStackServiceAccount(&#34;cloudSa&#34;, CloudStackServiceAccountArgs.builder()        
 *             .stackSlug(&#34;&lt;your stack slug&gt;&#34;)
 *             .role(&#34;Admin&#34;)
 *             .isDisabled(false)
 *             .build());
 * 
 *         var foo = new CloudStackServiceAccountToken(&#34;foo&#34;, CloudStackServiceAccountTokenArgs.builder()        
 *             .serviceAccountId(cloudSa.id())
 *             .build());
 * 
 *         ctx.export(&#34;serviceAccountTokenFooKey&#34;, foo.key());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken")
public class CloudStackServiceAccountToken extends com.pulumi.resources.CustomResource {
    @Export(name="expiration", refs={String.class}, tree="[0]")
    private Output<String> expiration;

    public Output<String> expiration() {
        return this.expiration;
    }
    @Export(name="hasExpired", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hasExpired;

    public Output<Boolean> hasExpired() {
        return this.hasExpired;
    }
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="secondsToLive", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> secondsToLive;

    public Output<Optional<Integer>> secondsToLive() {
        return Codegen.optional(this.secondsToLive);
    }
    @Export(name="serviceAccountId", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountId;

    public Output<String> serviceAccountId() {
        return this.serviceAccountId;
    }
    @Export(name="stackSlug", refs={String.class}, tree="[0]")
    private Output<String> stackSlug;

    public Output<String> stackSlug() {
        return this.stackSlug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudStackServiceAccountToken(String name) {
        this(name, CloudStackServiceAccountTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudStackServiceAccountToken(String name, CloudStackServiceAccountTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudStackServiceAccountToken(String name, CloudStackServiceAccountTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken", name, args == null ? CloudStackServiceAccountTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudStackServiceAccountToken(String name, Output<String> id, @Nullable CloudStackServiceAccountTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "key"
            ))
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
    public static CloudStackServiceAccountToken get(String name, Output<String> id, @Nullable CloudStackServiceAccountTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudStackServiceAccountToken(name, id, state, options);
    }
}
