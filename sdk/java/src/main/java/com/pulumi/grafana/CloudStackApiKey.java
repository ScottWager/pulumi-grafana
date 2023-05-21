// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.CloudStackApiKeyArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.CloudStackApiKeyState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="grafana:index/cloudStackApiKey:CloudStackApiKey")
public class CloudStackApiKey extends com.pulumi.resources.CustomResource {
    @Export(name="expiration", refs={String.class}, tree="[0]")
    private Output<String> expiration;

    public Output<String> expiration() {
        return this.expiration;
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
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }
    @Export(name="secondsToLive", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> secondsToLive;

    public Output<Optional<Integer>> secondsToLive() {
        return Codegen.optional(this.secondsToLive);
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
    public CloudStackApiKey(String name) {
        this(name, CloudStackApiKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudStackApiKey(String name, CloudStackApiKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudStackApiKey(String name, CloudStackApiKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/cloudStackApiKey:CloudStackApiKey", name, args == null ? CloudStackApiKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudStackApiKey(String name, Output<String> id, @Nullable CloudStackApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/cloudStackApiKey:CloudStackApiKey", name, state, makeResourceOptions(options, id));
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
    public static CloudStackApiKey get(String name, Output<String> id, @Nullable CloudStackApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudStackApiKey(name, id, state, options);
    }
}
