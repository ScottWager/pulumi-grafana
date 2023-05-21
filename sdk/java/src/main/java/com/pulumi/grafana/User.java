// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.UserArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.UserState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
 * 
 * This resource represents an instance-scoped resource and uses Grafana&#39;s admin APIs.
 * It does not work with API tokens or service accounts which are org-scoped.
 * You must use basic auth.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.User;
 * import com.pulumi.grafana.UserArgs;
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
 *         var staff = new User(&#34;staff&#34;, UserArgs.builder()        
 *             .email(&#34;staff.name@example.com&#34;)
 *             .isAdmin(false)
 *             .login(&#34;staff&#34;)
 *             .password(&#34;my-password&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/user:User user_name {{user_id}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * The email address of the Grafana user.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return The email address of the Grafana user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * Whether to make user an admin. Defaults to `false`.
     * 
     */
    @Export(name="isAdmin", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isAdmin;

    /**
     * @return Whether to make user an admin. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> isAdmin() {
        return Codegen.optional(this.isAdmin);
    }
    /**
     * The username for the Grafana user.
     * 
     */
    @Export(name="login", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> login;

    /**
     * @return The username for the Grafana user.
     * 
     */
    public Output<Optional<String>> login() {
        return Codegen.optional(this.login);
    }
    /**
     * The display name for the Grafana user.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The display name for the Grafana user.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password for the Grafana user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The password for the Grafana user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The numerical ID of the Grafana user.
     * 
     */
    @Export(name="userId", refs={Integer.class}, tree="[0]")
    private Output<Integer> userId;

    /**
     * @return The numerical ID of the Grafana user.
     * 
     */
    public Output<Integer> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
