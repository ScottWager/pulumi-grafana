// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.grafana.OncallEscalationArgs;
import com.pulumi.grafana.Utilities;
import com.pulumi.grafana.inputs.OncallEscalationState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * * [Official documentation](https://grafana.com/docs/oncall/latest/escalation-policies/)
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_policies/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.grafana.OncallEscalationChain;
 * import com.pulumi.grafana.OncallEscalationChainArgs;
 * import com.pulumi.grafana.GrafanaFunctions;
 * import com.pulumi.grafana.inputs.GetOncallUserArgs;
 * import com.pulumi.grafana.OncallEscalation;
 * import com.pulumi.grafana.OncallEscalationArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var default_ = new OncallEscalationChain(&#34;default&#34;, OncallEscalationChainArgs.Empty, CustomResourceOptions.builder()
 *             .provider(grafana.oncall())
 *             .build());
 * 
 *         final var alex = GrafanaFunctions.getOncallUser(GetOncallUserArgs.builder()
 *             .username(&#34;alex&#34;)
 *             .build());
 * 
 *         var exampleNotifyStepOncallEscalation = new OncallEscalation(&#34;exampleNotifyStepOncallEscalation&#34;, OncallEscalationArgs.builder()        
 *             .escalationChainId(default_.id())
 *             .type(&#34;notify_persons&#34;)
 *             .personsToNotifies(alex.applyValue(getOncallUserResult -&gt; getOncallUserResult.id()))
 *             .position(0)
 *             .build());
 * 
 *         var exampleNotifyStepIndex_oncallEscalationOncallEscalation = new OncallEscalation(&#34;exampleNotifyStepIndex/oncallEscalationOncallEscalation&#34;, OncallEscalationArgs.builder()        
 *             .escalationChainId(default_.id())
 *             .type(&#34;wait&#34;)
 *             .duration(300)
 *             .position(1)
 *             .build());
 * 
 *         var exampleNotifyStepGrafanaIndex_oncallEscalationOncallEscalation = new OncallEscalation(&#34;exampleNotifyStepGrafanaIndex/oncallEscalationOncallEscalation&#34;, OncallEscalationArgs.builder()        
 *             .escalationChainId(default_.id())
 *             .type(&#34;notify_persons&#34;)
 *             .important(true)
 *             .personsToNotifies(alex.applyValue(getOncallUserResult -&gt; getOncallUserResult.id()))
 *             .position(0)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import grafana:index/oncallEscalation:OncallEscalation escalation_name {{escalation_id}}
 * ```
 * 
 */
@ResourceType(type="grafana:index/oncallEscalation:OncallEscalation")
public class OncallEscalation extends com.pulumi.resources.CustomResource {
    /**
     * The ID of an Action for trigger_action type step.
     * 
     */
    @Export(name="actionToTrigger", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> actionToTrigger;

    /**
     * @return The ID of an Action for trigger_action type step.
     * 
     */
    public Output<Optional<String>> actionToTrigger() {
        return Codegen.optional(this.actionToTrigger);
    }
    /**
     * The duration of delay for wait type step.
     * 
     */
    @Export(name="duration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> duration;

    /**
     * @return The duration of delay for wait type step.
     * 
     */
    public Output<Optional<Integer>> duration() {
        return Codegen.optional(this.duration);
    }
    /**
     * The ID of the escalation chain.
     * 
     */
    @Export(name="escalationChainId", refs={String.class}, tree="[0]")
    private Output<String> escalationChainId;

    /**
     * @return The ID of the escalation chain.
     * 
     */
    public Output<String> escalationChainId() {
        return this.escalationChainId;
    }
    /**
     * The ID of a User Group for notify*user*group type step.
     * 
     */
    @Export(name="groupToNotify", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupToNotify;

    /**
     * @return The ID of a User Group for notify*user*group type step.
     * 
     */
    public Output<Optional<String>> groupToNotify() {
        return Codegen.optional(this.groupToNotify);
    }
    /**
     * Will activate &#34;important&#34; personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
     * 
     */
    @Export(name="important", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> important;

    /**
     * @return Will activate &#34;important&#34; personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user_group
     * 
     */
    public Output<Optional<Boolean>> important() {
        return Codegen.optional(this.important);
    }
    /**
     * The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
     * 
     */
    @Export(name="notifyIfTimeFrom", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notifyIfTimeFrom;

    /**
     * @return The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
     * 
     */
    public Output<Optional<String>> notifyIfTimeFrom() {
        return Codegen.optional(this.notifyIfTimeFrom);
    }
    /**
     * The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
     * 
     */
    @Export(name="notifyIfTimeTo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notifyIfTimeTo;

    /**
     * @return The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
     * 
     */
    public Output<Optional<String>> notifyIfTimeTo() {
        return Codegen.optional(this.notifyIfTimeTo);
    }
    /**
     * ID of a Schedule for notify*on*call*from*schedule type step.
     * 
     */
    @Export(name="notifyOnCallFromSchedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notifyOnCallFromSchedule;

    /**
     * @return ID of a Schedule for notify*on*call*from*schedule type step.
     * 
     */
    public Output<Optional<String>> notifyOnCallFromSchedule() {
        return Codegen.optional(this.notifyOnCallFromSchedule);
    }
    /**
     * The list of ID&#39;s of users for notify_persons type step.
     * 
     */
    @Export(name="personsToNotifies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> personsToNotifies;

    /**
     * @return The list of ID&#39;s of users for notify_persons type step.
     * 
     */
    public Output<Optional<List<String>>> personsToNotifies() {
        return Codegen.optional(this.personsToNotifies);
    }
    /**
     * The list of ID&#39;s of users for notify*person*next*each*time type step.
     * 
     */
    @Export(name="personsToNotifyNextEachTimes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> personsToNotifyNextEachTimes;

    /**
     * @return The list of ID&#39;s of users for notify*person*next*each*time type step.
     * 
     */
    public Output<Optional<List<String>>> personsToNotifyNextEachTimes() {
        return Codegen.optional(this.personsToNotifyNextEachTimes);
    }
    /**
     * The position of the escalation step (starts from 0).
     * 
     */
    @Export(name="position", refs={Integer.class}, tree="[0]")
    private Output<Integer> position;

    /**
     * @return The position of the escalation step (starts from 0).
     * 
     */
    public Output<Integer> position() {
        return this.position;
    }
    /**
     * The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*action, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat_escalation
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OncallEscalation(String name) {
        this(name, OncallEscalationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OncallEscalation(String name, OncallEscalationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OncallEscalation(String name, OncallEscalationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/oncallEscalation:OncallEscalation", name, args == null ? OncallEscalationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OncallEscalation(String name, Output<String> id, @Nullable OncallEscalationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("grafana:index/oncallEscalation:OncallEscalation", name, state, makeResourceOptions(options, id));
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
    public static OncallEscalation get(String name, Output<String> id, @Nullable OncallEscalationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OncallEscalation(name, id, state, options);
    }
}
