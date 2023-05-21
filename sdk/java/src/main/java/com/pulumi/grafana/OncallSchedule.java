// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.OncallScheduleArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.OncallScheduleState;
import com.pulumi.grafana.outputs.OncallScheduleSlack;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/schedules/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.GrafanaFunctions;
 * import com.pulumi.grafana.inputs.GetOnCallSlackChannelArgs;
 * import com.pulumi.grafana.inputs.GetOncallUserGroupArgs;
 * import com.pulumi.grafana.OncallSchedule;
 * import com.pulumi.grafana.OncallScheduleArgs;
 * import com.pulumi.grafana.inputs.OncallScheduleSlackArgs;
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
 *         final var exampleSlackChannel = GrafanaFunctions.getOnCallSlackChannel(GetOnCallSlackChannelArgs.builder()
 *             .name(&#34;example_slack_channel&#34;)
 *             .build());
 * 
 *         final var exampleUserGroup = GrafanaFunctions.getOncallUserGroup(GetOncallUserGroupArgs.builder()
 *             .slackHandle(&#34;example_slack_handle&#34;)
 *             .build());
 * 
 *         var exampleScheduleOncallSchedule = new OncallSchedule(&#34;exampleScheduleOncallSchedule&#34;, OncallScheduleArgs.builder()        
 *             .type(&#34;ical&#34;)
 *             .icalUrlPrimary(&#34;https://example.com/example_ical.ics&#34;)
 *             .icalUrlOverrides(&#34;https://example.com/example_overrides_ical.ics&#34;)
 *             .slack(OncallScheduleSlackArgs.builder()
 *                 .channelId(exampleSlackChannel.applyValue(getOnCallSlackChannelResult -&gt; getOnCallSlackChannelResult.slackId()))
 *                 .userGroupId(exampleUserGroup.applyValue(getOncallUserGroupResult -&gt; getOncallUserGroupResult.slackId()))
 *                 .build())
 *             .build());
 * 
 *         var exampleScheduleIndex_oncallScheduleOncallSchedule = new OncallSchedule(&#34;exampleScheduleIndex/oncallScheduleOncallSchedule&#34;, OncallScheduleArgs.builder()        
 *             .type(&#34;calendar&#34;)
 *             .timeZone(&#34;America/New_York&#34;)
 *             .shifts()
 *             .icalUrlOverrides(&#34;https://example.com/example_overrides_ical.ics&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/oncallSchedule:OncallSchedule schedule_name {{schedule_id}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/oncallSchedule:OncallSchedule")
public class OncallSchedule extends com.pulumi.resources.CustomResource {
    /**
     * The URL of external iCal calendar which override primary events.
     * 
     */
    @Export(name="icalUrlOverrides", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icalUrlOverrides;

    /**
     * @return The URL of external iCal calendar which override primary events.
     * 
     */
    public Output<Optional<String>> icalUrlOverrides() {
        return Codegen.optional(this.icalUrlOverrides);
    }
    /**
     * The URL of the external calendar iCal file.
     * 
     */
    @Export(name="icalUrlPrimary", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icalUrlPrimary;

    /**
     * @return The URL of the external calendar iCal file.
     * 
     */
    public Output<Optional<String>> icalUrlPrimary() {
        return Codegen.optional(this.icalUrlPrimary);
    }
    /**
     * The schedule&#39;s name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The schedule&#39;s name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The list of ID&#39;s of on-call shifts.
     * 
     */
    @Export(name="shifts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> shifts;

    /**
     * @return The list of ID&#39;s of on-call shifts.
     * 
     */
    public Output<Optional<List<String>>> shifts() {
        return Codegen.optional(this.shifts);
    }
    /**
     * The Slack-specific settings for a schedule.
     * 
     */
    @Export(name="slack", refs={OncallScheduleSlack.class}, tree="[0]")
    private Output</* @Nullable */ OncallScheduleSlack> slack;

    /**
     * @return The Slack-specific settings for a schedule.
     * 
     */
    public Output<Optional<OncallScheduleSlack>> slack() {
        return Codegen.optional(this.slack);
    }
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     * 
     */
    @Export(name="teamId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> teamId;

    /**
     * @return The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     * 
     */
    public Output<Optional<String>> teamId() {
        return Codegen.optional(this.teamId);
    }
    /**
     * The schedule&#39;s time zone.
     * 
     */
    @Export(name="timeZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timeZone;

    /**
     * @return The schedule&#39;s time zone.
     * 
     */
    public Output<Optional<String>> timeZone() {
        return Codegen.optional(this.timeZone);
    }
    /**
     * The schedule&#39;s type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The schedule&#39;s type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OncallSchedule(String name) {
        this(name, OncallScheduleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OncallSchedule(String name, OncallScheduleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OncallSchedule(String name, OncallScheduleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/oncallSchedule:OncallSchedule", name, args == null ? OncallScheduleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OncallSchedule(String name, Output<String> id, @Nullable OncallScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/oncallSchedule:OncallSchedule", name, state, makeResourceOptions(options, id));
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
    public static OncallSchedule get(String name, Output<String> id, @Nullable OncallScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OncallSchedule(name, id, state, options);
    }
}
