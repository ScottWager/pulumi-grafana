// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.TeamArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.TeamState;
import com.pulumi.grafana.outputs.TeamPreferences;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
 * * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.Team;
 * import com.pulumi.grafana.TeamArgs;
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
 *         var test_team = new Team(&#34;test-team&#34;, TeamArgs.builder()        
 *             .email(&#34;teamemail@example.com&#34;)
 *             .members(&#34;viewer-01@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="grafana:index/team:Team")
public class Team extends com.pulumi.resources.CustomResource {
    /**
     * An email address for the team.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> email;

    /**
     * @return An email address for the team.
     * 
     */
    public Output<Optional<String>> email() {
        return Codegen.optional(this.email);
    }
    /**
     * Ignores team members that have been added to team by [Team Sync](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/).
     * Team Sync can be provisioned using grafana*team*external_group resource.
     * Defaults to `true`.
     * 
     */
    @Export(name="ignoreExternallySyncedMembers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreExternallySyncedMembers;

    /**
     * @return Ignores team members that have been added to team by [Team Sync](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/).
     * Team Sync can be provisioned using grafana*team*external_group resource.
     * Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> ignoreExternallySyncedMembers() {
        return Codegen.optional(this.ignoreExternallySyncedMembers);
    }
    /**
     * A set of email addresses corresponding to users who should be given membership
     * to the team. Note: users specified here must already exist in Grafana.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> members;

    /**
     * @return A set of email addresses corresponding to users who should be given membership
     * to the team. Note: users specified here must already exist in Grafana.
     * 
     */
    public Output<Optional<List<String>>> members() {
        return Codegen.optional(this.members);
    }
    /**
     * The display name for the Grafana team created.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The display name for the Grafana team created.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="preferences", refs={TeamPreferences.class}, tree="[0]")
    private Output</* @Nullable */ TeamPreferences> preferences;

    public Output<Optional<TeamPreferences>> preferences() {
        return Codegen.optional(this.preferences);
    }
    /**
     * The team id assigned to this team by Grafana.
     * 
     */
    @Export(name="teamId", refs={Integer.class}, tree="[0]")
    private Output<Integer> teamId;

    /**
     * @return The team id assigned to this team by Grafana.
     * 
     */
    public Output<Integer> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Team(String name) {
        this(name, TeamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Team(String name, @Nullable TeamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Team(String name, @Nullable TeamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/team:Team", name, args == null ? TeamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Team(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/team:Team", name, state, makeResourceOptions(options, id));
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
    public static Team get(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Team(name, id, state, options);
    }
}
