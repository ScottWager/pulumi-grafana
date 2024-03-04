// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/on_call_shifts/)
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import grafana:index/oncallOnCallShift:OncallOnCallShift on_call_shift_name {{on_call_shift_id}}
 * ```
 */
export class OncallOnCallShift extends pulumi.CustomResource {
    /**
     * Get an existing OncallOnCallShift resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OncallOnCallShiftState, opts?: pulumi.CustomResourceOptions): OncallOnCallShift {
        return new OncallOnCallShift(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/oncallOnCallShift:OncallOnCallShift';

    /**
     * Returns true if the given object is an instance of OncallOnCallShift.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OncallOnCallShift {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OncallOnCallShift.__pulumiType;
    }

    /**
     * This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    public readonly byDays!: pulumi.Output<string[] | undefined>;
    /**
     * This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
     */
    public readonly byMonthdays!: pulumi.Output<number[] | undefined>;
    /**
     * This parameter takes a list of months. Valid values are 1 to 12
     */
    public readonly byMonths!: pulumi.Output<number[] | undefined>;
    /**
     * The duration of the event.
     */
    public readonly duration!: pulumi.Output<number>;
    /**
     * The frequency of the event. Can be daily, weekly, monthly
     */
    public readonly frequency!: pulumi.Output<string | undefined>;
    /**
     * The positive integer representing at which intervals the recurrence rule repeats.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * The priority level. The higher the value, the higher the priority.
     */
    public readonly level!: pulumi.Output<number | undefined>;
    /**
     * The shift's name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of lists with on-call users (for rollingUsers event type)
     */
    public readonly rollingUsers!: pulumi.Output<string[][] | undefined>;
    /**
     * The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
     */
    public readonly start!: pulumi.Output<string>;
    /**
     * The index of the list of users in rolling_users, from which on-call rotation starts.
     */
    public readonly startRotationFromUserIndex!: pulumi.Output<number | undefined>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    public readonly teamId!: pulumi.Output<string | undefined>;
    /**
     * The shift's timezone.  Overrides schedule's timezone.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;
    /**
     * The shift's type. Can be rolling*users, recurrent*event, single_event
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The list of on-call users (for single*event and recurrent*event event type).
     */
    public readonly users!: pulumi.Output<string[] | undefined>;
    /**
     * Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    public readonly weekStart!: pulumi.Output<string | undefined>;

    /**
     * Create a OncallOnCallShift resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OncallOnCallShiftArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OncallOnCallShiftArgs | OncallOnCallShiftState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OncallOnCallShiftState | undefined;
            resourceInputs["byDays"] = state ? state.byDays : undefined;
            resourceInputs["byMonthdays"] = state ? state.byMonthdays : undefined;
            resourceInputs["byMonths"] = state ? state.byMonths : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rollingUsers"] = state ? state.rollingUsers : undefined;
            resourceInputs["start"] = state ? state.start : undefined;
            resourceInputs["startRotationFromUserIndex"] = state ? state.startRotationFromUserIndex : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["weekStart"] = state ? state.weekStart : undefined;
        } else {
            const args = argsOrState as OncallOnCallShiftArgs | undefined;
            if ((!args || args.duration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'duration'");
            }
            if ((!args || args.start === undefined) && !opts.urn) {
                throw new Error("Missing required property 'start'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["byDays"] = args ? args.byDays : undefined;
            resourceInputs["byMonthdays"] = args ? args.byMonthdays : undefined;
            resourceInputs["byMonths"] = args ? args.byMonths : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rollingUsers"] = args ? args.rollingUsers : undefined;
            resourceInputs["start"] = args ? args.start : undefined;
            resourceInputs["startRotationFromUserIndex"] = args ? args.startRotationFromUserIndex : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["weekStart"] = args ? args.weekStart : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OncallOnCallShift.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OncallOnCallShift resources.
 */
export interface OncallOnCallShiftState {
    /**
     * This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    byDays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
     */
    byMonthdays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * This parameter takes a list of months. Valid values are 1 to 12
     */
    byMonths?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The duration of the event.
     */
    duration?: pulumi.Input<number>;
    /**
     * The frequency of the event. Can be daily, weekly, monthly
     */
    frequency?: pulumi.Input<string>;
    /**
     * The positive integer representing at which intervals the recurrence rule repeats.
     */
    interval?: pulumi.Input<number>;
    /**
     * The priority level. The higher the value, the higher the priority.
     */
    level?: pulumi.Input<number>;
    /**
     * The shift's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of lists with on-call users (for rollingUsers event type)
     */
    rollingUsers?: pulumi.Input<pulumi.Input<pulumi.Input<string>[]>[]>;
    /**
     * The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
     */
    start?: pulumi.Input<string>;
    /**
     * The index of the list of users in rolling_users, from which on-call rotation starts.
     */
    startRotationFromUserIndex?: pulumi.Input<number>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The shift's timezone.  Overrides schedule's timezone.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The shift's type. Can be rolling*users, recurrent*event, single_event
     */
    type?: pulumi.Input<string>;
    /**
     * The list of on-call users (for single*event and recurrent*event event type).
     */
    users?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    weekStart?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OncallOnCallShift resource.
 */
export interface OncallOnCallShiftArgs {
    /**
     * This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    byDays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
     */
    byMonthdays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * This parameter takes a list of months. Valid values are 1 to 12
     */
    byMonths?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The duration of the event.
     */
    duration: pulumi.Input<number>;
    /**
     * The frequency of the event. Can be daily, weekly, monthly
     */
    frequency?: pulumi.Input<string>;
    /**
     * The positive integer representing at which intervals the recurrence rule repeats.
     */
    interval?: pulumi.Input<number>;
    /**
     * The priority level. The higher the value, the higher the priority.
     */
    level?: pulumi.Input<number>;
    /**
     * The shift's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of lists with on-call users (for rollingUsers event type)
     */
    rollingUsers?: pulumi.Input<pulumi.Input<pulumi.Input<string>[]>[]>;
    /**
     * The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
     */
    start: pulumi.Input<string>;
    /**
     * The index of the list of users in rolling_users, from which on-call rotation starts.
     */
    startRotationFromUserIndex?: pulumi.Input<number>;
    /**
     * The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana.getOncallTeam` datasource.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The shift's timezone.  Overrides schedule's timezone.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The shift's type. Can be rolling*users, recurrent*event, single_event
     */
    type: pulumi.Input<string>;
    /**
     * The list of on-call users (for single*event and recurrent*event event type).
     */
    users?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
     */
    weekStart?: pulumi.Input<string>;
}
