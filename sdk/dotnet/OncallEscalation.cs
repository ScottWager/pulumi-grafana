// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/oncall/latest/configure/escalation-chains-and-routes/)
    /// * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_policies/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumi.Grafana;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Grafana.OnCall.EscalationChain("default", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = grafana.Oncall,
    ///     });
    /// 
    ///     var alex = Grafana.OnCall.GetUser.Invoke(new()
    ///     {
    ///         Username = "alex",
    ///     });
    /// 
    ///     // Notify step
    ///     var exampleNotifyStepEscalation = new Grafana.OnCall.Escalation("exampleNotifyStepEscalation", new()
    ///     {
    ///         EscalationChainId = @default.Id,
    ///         Type = "notify_persons",
    ///         PersonsToNotifies = new[]
    ///         {
    ///             alex.Apply(getUserResult =&gt; getUserResult.Id),
    ///         },
    ///         Position = 0,
    ///     });
    /// 
    ///     // Wait step
    ///     var exampleNotifyStepOnCall_escalationEscalation = new Grafana.OnCall.Escalation("exampleNotifyStepOnCall/escalationEscalation", new()
    ///     {
    ///         EscalationChainId = @default.Id,
    ///         Type = "wait",
    ///         Duration = 300,
    ///         Position = 1,
    ///     });
    /// 
    ///     // Important step
    ///     var exampleNotifyStepGrafanaOnCall_escalationEscalation = new Grafana.OnCall.Escalation("exampleNotifyStepGrafanaOnCall/escalationEscalation", new()
    ///     {
    ///         EscalationChainId = @default.Id,
    ///         Type = "notify_persons",
    ///         Important = true,
    ///         PersonsToNotifies = new[]
    ///         {
    ///             alex.Apply(getUserResult =&gt; getUserResult.Id),
    ///         },
    ///         Position = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import grafana:index/oncallEscalation:OncallEscalation name "{{ id }}"
    /// ```
    /// </summary>
    [Obsolete(@"grafana.index/oncallescalation.OncallEscalation has been deprecated in favor of grafana.oncall/escalation.Escalation")]
    [GrafanaResourceType("grafana:index/oncallEscalation:OncallEscalation")]
    public partial class OncallEscalation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of an Action for trigger_webhook type step.
        /// </summary>
        [Output("actionToTrigger")]
        public Output<string?> ActionToTrigger { get; private set; } = null!;

        /// <summary>
        /// The duration of delay for wait type step.
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// The ID of the escalation chain.
        /// </summary>
        [Output("escalationChainId")]
        public Output<string> EscalationChainId { get; private set; } = null!;

        /// <summary>
        /// The ID of a User Group for notify*user*group type step.
        /// </summary>
        [Output("groupToNotify")]
        public Output<string?> GroupToNotify { get; private set; } = null!;

        /// <summary>
        /// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
        /// </summary>
        [Output("important")]
        public Output<bool?> Important { get; private set; } = null!;

        /// <summary>
        /// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
        /// </summary>
        [Output("notifyIfTimeFrom")]
        public Output<string?> NotifyIfTimeFrom { get; private set; } = null!;

        /// <summary>
        /// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
        /// </summary>
        [Output("notifyIfTimeTo")]
        public Output<string?> NotifyIfTimeTo { get; private set; } = null!;

        /// <summary>
        /// ID of a Schedule for notify*on*call*from*schedule type step.
        /// </summary>
        [Output("notifyOnCallFromSchedule")]
        public Output<string?> NotifyOnCallFromSchedule { get; private set; } = null!;

        /// <summary>
        /// The ID of a Team for a notify*team*members type step.
        /// </summary>
        [Output("notifyToTeamMembers")]
        public Output<string?> NotifyToTeamMembers { get; private set; } = null!;

        /// <summary>
        /// The list of ID's of users for notify_persons type step.
        /// </summary>
        [Output("personsToNotifies")]
        public Output<ImmutableArray<string>> PersonsToNotifies { get; private set; } = null!;

        /// <summary>
        /// The list of ID's of users for notify*person*next*each*time type step.
        /// </summary>
        [Output("personsToNotifyNextEachTimes")]
        public Output<ImmutableArray<string>> PersonsToNotifyNextEachTimes { get; private set; } = null!;

        /// <summary>
        /// The position of the escalation step (starts from 0).
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a OncallEscalation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OncallEscalation(string name, OncallEscalationArgs args, CustomResourceOptions? options = null)
            : base("grafana:index/oncallEscalation:OncallEscalation", name, args ?? new OncallEscalationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OncallEscalation(string name, Input<string> id, OncallEscalationState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/oncallEscalation:OncallEscalation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "grafana:index/oncallEscalation:OncallEscalation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OncallEscalation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OncallEscalation Get(string name, Input<string> id, OncallEscalationState? state = null, CustomResourceOptions? options = null)
        {
            return new OncallEscalation(name, id, state, options);
        }
    }

    public sealed class OncallEscalationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of an Action for trigger_webhook type step.
        /// </summary>
        [Input("actionToTrigger")]
        public Input<string>? ActionToTrigger { get; set; }

        /// <summary>
        /// The duration of delay for wait type step.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The ID of the escalation chain.
        /// </summary>
        [Input("escalationChainId", required: true)]
        public Input<string> EscalationChainId { get; set; } = null!;

        /// <summary>
        /// The ID of a User Group for notify*user*group type step.
        /// </summary>
        [Input("groupToNotify")]
        public Input<string>? GroupToNotify { get; set; }

        /// <summary>
        /// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
        /// </summary>
        [Input("important")]
        public Input<bool>? Important { get; set; }

        /// <summary>
        /// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
        /// </summary>
        [Input("notifyIfTimeFrom")]
        public Input<string>? NotifyIfTimeFrom { get; set; }

        /// <summary>
        /// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
        /// </summary>
        [Input("notifyIfTimeTo")]
        public Input<string>? NotifyIfTimeTo { get; set; }

        /// <summary>
        /// ID of a Schedule for notify*on*call*from*schedule type step.
        /// </summary>
        [Input("notifyOnCallFromSchedule")]
        public Input<string>? NotifyOnCallFromSchedule { get; set; }

        /// <summary>
        /// The ID of a Team for a notify*team*members type step.
        /// </summary>
        [Input("notifyToTeamMembers")]
        public Input<string>? NotifyToTeamMembers { get; set; }

        [Input("personsToNotifies")]
        private InputList<string>? _personsToNotifies;

        /// <summary>
        /// The list of ID's of users for notify_persons type step.
        /// </summary>
        public InputList<string> PersonsToNotifies
        {
            get => _personsToNotifies ?? (_personsToNotifies = new InputList<string>());
            set => _personsToNotifies = value;
        }

        [Input("personsToNotifyNextEachTimes")]
        private InputList<string>? _personsToNotifyNextEachTimes;

        /// <summary>
        /// The list of ID's of users for notify*person*next*each*time type step.
        /// </summary>
        public InputList<string> PersonsToNotifyNextEachTimes
        {
            get => _personsToNotifyNextEachTimes ?? (_personsToNotifyNextEachTimes = new InputList<string>());
            set => _personsToNotifyNextEachTimes = value;
        }

        /// <summary>
        /// The position of the escalation step (starts from 0).
        /// </summary>
        [Input("position", required: true)]
        public Input<int> Position { get; set; } = null!;

        /// <summary>
        /// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public OncallEscalationArgs()
        {
        }
        public static new OncallEscalationArgs Empty => new OncallEscalationArgs();
    }

    public sealed class OncallEscalationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of an Action for trigger_webhook type step.
        /// </summary>
        [Input("actionToTrigger")]
        public Input<string>? ActionToTrigger { get; set; }

        /// <summary>
        /// The duration of delay for wait type step.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The ID of the escalation chain.
        /// </summary>
        [Input("escalationChainId")]
        public Input<string>? EscalationChainId { get; set; }

        /// <summary>
        /// The ID of a User Group for notify*user*group type step.
        /// </summary>
        [Input("groupToNotify")]
        public Input<string>? GroupToNotify { get; set; }

        /// <summary>
        /// Will activate "important" personal notification rules. Actual for steps: notify*persons, notify*on*call*from*schedule and notify*user*group,notify*team_members
        /// </summary>
        [Input("important")]
        public Input<bool>? Important { get; set; }

        /// <summary>
        /// The beginning of the time interval for notify*if*time*from*to type step in UTC (for example 08:00:00Z).
        /// </summary>
        [Input("notifyIfTimeFrom")]
        public Input<string>? NotifyIfTimeFrom { get; set; }

        /// <summary>
        /// The end of the time interval for notify*if*time*from*to type step in UTC (for example 18:00:00Z).
        /// </summary>
        [Input("notifyIfTimeTo")]
        public Input<string>? NotifyIfTimeTo { get; set; }

        /// <summary>
        /// ID of a Schedule for notify*on*call*from*schedule type step.
        /// </summary>
        [Input("notifyOnCallFromSchedule")]
        public Input<string>? NotifyOnCallFromSchedule { get; set; }

        /// <summary>
        /// The ID of a Team for a notify*team*members type step.
        /// </summary>
        [Input("notifyToTeamMembers")]
        public Input<string>? NotifyToTeamMembers { get; set; }

        [Input("personsToNotifies")]
        private InputList<string>? _personsToNotifies;

        /// <summary>
        /// The list of ID's of users for notify_persons type step.
        /// </summary>
        public InputList<string> PersonsToNotifies
        {
            get => _personsToNotifies ?? (_personsToNotifies = new InputList<string>());
            set => _personsToNotifies = value;
        }

        [Input("personsToNotifyNextEachTimes")]
        private InputList<string>? _personsToNotifyNextEachTimes;

        /// <summary>
        /// The list of ID's of users for notify*person*next*each*time type step.
        /// </summary>
        public InputList<string> PersonsToNotifyNextEachTimes
        {
            get => _personsToNotifyNextEachTimes ?? (_personsToNotifyNextEachTimes = new InputList<string>());
            set => _personsToNotifyNextEachTimes = value;
        }

        /// <summary>
        /// The position of the escalation step (starts from 0).
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// The type of escalation policy. Can be wait, notify*persons, notify*person*next*each*time, notify*on*call*from*schedule, trigger*webhook, notify*user*group, resolve, notify*whole*channel, notify*if*time*from*to, repeat*escalation, notify*team_members
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OncallEscalationState()
        {
        }
        public static new OncallEscalationState Empty => new OncallEscalationState();
    }
}
