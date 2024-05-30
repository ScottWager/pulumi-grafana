# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OncallEscalationArgs', 'OncallEscalation']

@pulumi.input_type
class OncallEscalationArgs:
    def __init__(__self__, *,
                 escalation_chain_id: pulumi.Input[str],
                 position: pulumi.Input[int],
                 action_to_trigger: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 group_to_notify: Optional[pulumi.Input[str]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 notify_if_time_from: Optional[pulumi.Input[str]] = None,
                 notify_if_time_to: Optional[pulumi.Input[str]] = None,
                 notify_on_call_from_schedule: Optional[pulumi.Input[str]] = None,
                 notify_to_team_members: Optional[pulumi.Input[str]] = None,
                 persons_to_notifies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 persons_to_notify_next_each_times: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OncallEscalation resource.
        :param pulumi.Input[str] escalation_chain_id: The ID of the escalation chain.
        :param pulumi.Input[int] position: The position of the escalation step (starts from 0).
        :param pulumi.Input[str] action_to_trigger: The ID of an Action for trigger_webhook type step.
        :param pulumi.Input[int] duration: The duration of delay for wait type step.
        :param pulumi.Input[str] group_to_notify: The ID of a User Group for notify_user_group type step.
        :param pulumi.Input[bool] important: Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
               and notify_user_group,notify_team_members
        :param pulumi.Input[str] notify_if_time_from: The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        :param pulumi.Input[str] notify_if_time_to: The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        :param pulumi.Input[str] notify_on_call_from_schedule: ID of a Schedule for notify_on_call_from_schedule type step.
        :param pulumi.Input[str] notify_to_team_members: The ID of a Team for a notify_team_members type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notifies: The list of ID's of users for notify_persons type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notify_next_each_times: The list of ID's of users for notify_person_next_each_time type step.
        :param pulumi.Input[str] type: The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
               trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
               notify_team_members
        """
        pulumi.set(__self__, "escalation_chain_id", escalation_chain_id)
        pulumi.set(__self__, "position", position)
        if action_to_trigger is not None:
            pulumi.set(__self__, "action_to_trigger", action_to_trigger)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if group_to_notify is not None:
            pulumi.set(__self__, "group_to_notify", group_to_notify)
        if important is not None:
            pulumi.set(__self__, "important", important)
        if notify_if_time_from is not None:
            pulumi.set(__self__, "notify_if_time_from", notify_if_time_from)
        if notify_if_time_to is not None:
            pulumi.set(__self__, "notify_if_time_to", notify_if_time_to)
        if notify_on_call_from_schedule is not None:
            pulumi.set(__self__, "notify_on_call_from_schedule", notify_on_call_from_schedule)
        if notify_to_team_members is not None:
            pulumi.set(__self__, "notify_to_team_members", notify_to_team_members)
        if persons_to_notifies is not None:
            pulumi.set(__self__, "persons_to_notifies", persons_to_notifies)
        if persons_to_notify_next_each_times is not None:
            pulumi.set(__self__, "persons_to_notify_next_each_times", persons_to_notify_next_each_times)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="escalationChainId")
    def escalation_chain_id(self) -> pulumi.Input[str]:
        """
        The ID of the escalation chain.
        """
        return pulumi.get(self, "escalation_chain_id")

    @escalation_chain_id.setter
    def escalation_chain_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "escalation_chain_id", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[int]:
        """
        The position of the escalation step (starts from 0).
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[int]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="actionToTrigger")
    def action_to_trigger(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of an Action for trigger_webhook type step.
        """
        return pulumi.get(self, "action_to_trigger")

    @action_to_trigger.setter
    def action_to_trigger(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_to_trigger", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The duration of delay for wait type step.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="groupToNotify")
    def group_to_notify(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a User Group for notify_user_group type step.
        """
        return pulumi.get(self, "group_to_notify")

    @group_to_notify.setter
    def group_to_notify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_to_notify", value)

    @property
    @pulumi.getter
    def important(self) -> Optional[pulumi.Input[bool]]:
        """
        Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
        and notify_user_group,notify_team_members
        """
        return pulumi.get(self, "important")

    @important.setter
    def important(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "important", value)

    @property
    @pulumi.getter(name="notifyIfTimeFrom")
    def notify_if_time_from(self) -> Optional[pulumi.Input[str]]:
        """
        The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_from")

    @notify_if_time_from.setter
    def notify_if_time_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_if_time_from", value)

    @property
    @pulumi.getter(name="notifyIfTimeTo")
    def notify_if_time_to(self) -> Optional[pulumi.Input[str]]:
        """
        The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_to")

    @notify_if_time_to.setter
    def notify_if_time_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_if_time_to", value)

    @property
    @pulumi.getter(name="notifyOnCallFromSchedule")
    def notify_on_call_from_schedule(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a Schedule for notify_on_call_from_schedule type step.
        """
        return pulumi.get(self, "notify_on_call_from_schedule")

    @notify_on_call_from_schedule.setter
    def notify_on_call_from_schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_on_call_from_schedule", value)

    @property
    @pulumi.getter(name="notifyToTeamMembers")
    def notify_to_team_members(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a Team for a notify_team_members type step.
        """
        return pulumi.get(self, "notify_to_team_members")

    @notify_to_team_members.setter
    def notify_to_team_members(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_to_team_members", value)

    @property
    @pulumi.getter(name="personsToNotifies")
    def persons_to_notifies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of ID's of users for notify_persons type step.
        """
        return pulumi.get(self, "persons_to_notifies")

    @persons_to_notifies.setter
    def persons_to_notifies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "persons_to_notifies", value)

    @property
    @pulumi.getter(name="personsToNotifyNextEachTimes")
    def persons_to_notify_next_each_times(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of ID's of users for notify_person_next_each_time type step.
        """
        return pulumi.get(self, "persons_to_notify_next_each_times")

    @persons_to_notify_next_each_times.setter
    def persons_to_notify_next_each_times(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "persons_to_notify_next_each_times", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
        trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
        notify_team_members
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _OncallEscalationState:
    def __init__(__self__, *,
                 action_to_trigger: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 escalation_chain_id: Optional[pulumi.Input[str]] = None,
                 group_to_notify: Optional[pulumi.Input[str]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 notify_if_time_from: Optional[pulumi.Input[str]] = None,
                 notify_if_time_to: Optional[pulumi.Input[str]] = None,
                 notify_on_call_from_schedule: Optional[pulumi.Input[str]] = None,
                 notify_to_team_members: Optional[pulumi.Input[str]] = None,
                 persons_to_notifies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 persons_to_notify_next_each_times: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OncallEscalation resources.
        :param pulumi.Input[str] action_to_trigger: The ID of an Action for trigger_webhook type step.
        :param pulumi.Input[int] duration: The duration of delay for wait type step.
        :param pulumi.Input[str] escalation_chain_id: The ID of the escalation chain.
        :param pulumi.Input[str] group_to_notify: The ID of a User Group for notify_user_group type step.
        :param pulumi.Input[bool] important: Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
               and notify_user_group,notify_team_members
        :param pulumi.Input[str] notify_if_time_from: The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        :param pulumi.Input[str] notify_if_time_to: The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        :param pulumi.Input[str] notify_on_call_from_schedule: ID of a Schedule for notify_on_call_from_schedule type step.
        :param pulumi.Input[str] notify_to_team_members: The ID of a Team for a notify_team_members type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notifies: The list of ID's of users for notify_persons type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notify_next_each_times: The list of ID's of users for notify_person_next_each_time type step.
        :param pulumi.Input[int] position: The position of the escalation step (starts from 0).
        :param pulumi.Input[str] type: The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
               trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
               notify_team_members
        """
        if action_to_trigger is not None:
            pulumi.set(__self__, "action_to_trigger", action_to_trigger)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if escalation_chain_id is not None:
            pulumi.set(__self__, "escalation_chain_id", escalation_chain_id)
        if group_to_notify is not None:
            pulumi.set(__self__, "group_to_notify", group_to_notify)
        if important is not None:
            pulumi.set(__self__, "important", important)
        if notify_if_time_from is not None:
            pulumi.set(__self__, "notify_if_time_from", notify_if_time_from)
        if notify_if_time_to is not None:
            pulumi.set(__self__, "notify_if_time_to", notify_if_time_to)
        if notify_on_call_from_schedule is not None:
            pulumi.set(__self__, "notify_on_call_from_schedule", notify_on_call_from_schedule)
        if notify_to_team_members is not None:
            pulumi.set(__self__, "notify_to_team_members", notify_to_team_members)
        if persons_to_notifies is not None:
            pulumi.set(__self__, "persons_to_notifies", persons_to_notifies)
        if persons_to_notify_next_each_times is not None:
            pulumi.set(__self__, "persons_to_notify_next_each_times", persons_to_notify_next_each_times)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="actionToTrigger")
    def action_to_trigger(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of an Action for trigger_webhook type step.
        """
        return pulumi.get(self, "action_to_trigger")

    @action_to_trigger.setter
    def action_to_trigger(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_to_trigger", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The duration of delay for wait type step.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="escalationChainId")
    def escalation_chain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the escalation chain.
        """
        return pulumi.get(self, "escalation_chain_id")

    @escalation_chain_id.setter
    def escalation_chain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "escalation_chain_id", value)

    @property
    @pulumi.getter(name="groupToNotify")
    def group_to_notify(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a User Group for notify_user_group type step.
        """
        return pulumi.get(self, "group_to_notify")

    @group_to_notify.setter
    def group_to_notify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_to_notify", value)

    @property
    @pulumi.getter
    def important(self) -> Optional[pulumi.Input[bool]]:
        """
        Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
        and notify_user_group,notify_team_members
        """
        return pulumi.get(self, "important")

    @important.setter
    def important(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "important", value)

    @property
    @pulumi.getter(name="notifyIfTimeFrom")
    def notify_if_time_from(self) -> Optional[pulumi.Input[str]]:
        """
        The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_from")

    @notify_if_time_from.setter
    def notify_if_time_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_if_time_from", value)

    @property
    @pulumi.getter(name="notifyIfTimeTo")
    def notify_if_time_to(self) -> Optional[pulumi.Input[str]]:
        """
        The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_to")

    @notify_if_time_to.setter
    def notify_if_time_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_if_time_to", value)

    @property
    @pulumi.getter(name="notifyOnCallFromSchedule")
    def notify_on_call_from_schedule(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a Schedule for notify_on_call_from_schedule type step.
        """
        return pulumi.get(self, "notify_on_call_from_schedule")

    @notify_on_call_from_schedule.setter
    def notify_on_call_from_schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_on_call_from_schedule", value)

    @property
    @pulumi.getter(name="notifyToTeamMembers")
    def notify_to_team_members(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of a Team for a notify_team_members type step.
        """
        return pulumi.get(self, "notify_to_team_members")

    @notify_to_team_members.setter
    def notify_to_team_members(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_to_team_members", value)

    @property
    @pulumi.getter(name="personsToNotifies")
    def persons_to_notifies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of ID's of users for notify_persons type step.
        """
        return pulumi.get(self, "persons_to_notifies")

    @persons_to_notifies.setter
    def persons_to_notifies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "persons_to_notifies", value)

    @property
    @pulumi.getter(name="personsToNotifyNextEachTimes")
    def persons_to_notify_next_each_times(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of ID's of users for notify_person_next_each_time type step.
        """
        return pulumi.get(self, "persons_to_notify_next_each_times")

    @persons_to_notify_next_each_times.setter
    def persons_to_notify_next_each_times(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "persons_to_notify_next_each_times", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        The position of the escalation step (starts from 0).
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
        trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
        notify_team_members
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class OncallEscalation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_to_trigger: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 escalation_chain_id: Optional[pulumi.Input[str]] = None,
                 group_to_notify: Optional[pulumi.Input[str]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 notify_if_time_from: Optional[pulumi.Input[str]] = None,
                 notify_if_time_to: Optional[pulumi.Input[str]] = None,
                 notify_on_call_from_schedule: Optional[pulumi.Input[str]] = None,
                 notify_to_team_members: Optional[pulumi.Input[str]] = None,
                 persons_to_notifies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 persons_to_notify_next_each_times: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a OncallEscalation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action_to_trigger: The ID of an Action for trigger_webhook type step.
        :param pulumi.Input[int] duration: The duration of delay for wait type step.
        :param pulumi.Input[str] escalation_chain_id: The ID of the escalation chain.
        :param pulumi.Input[str] group_to_notify: The ID of a User Group for notify_user_group type step.
        :param pulumi.Input[bool] important: Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
               and notify_user_group,notify_team_members
        :param pulumi.Input[str] notify_if_time_from: The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        :param pulumi.Input[str] notify_if_time_to: The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        :param pulumi.Input[str] notify_on_call_from_schedule: ID of a Schedule for notify_on_call_from_schedule type step.
        :param pulumi.Input[str] notify_to_team_members: The ID of a Team for a notify_team_members type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notifies: The list of ID's of users for notify_persons type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notify_next_each_times: The list of ID's of users for notify_person_next_each_time type step.
        :param pulumi.Input[int] position: The position of the escalation step (starts from 0).
        :param pulumi.Input[str] type: The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
               trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
               notify_team_members
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OncallEscalationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a OncallEscalation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param OncallEscalationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OncallEscalationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_to_trigger: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 escalation_chain_id: Optional[pulumi.Input[str]] = None,
                 group_to_notify: Optional[pulumi.Input[str]] = None,
                 important: Optional[pulumi.Input[bool]] = None,
                 notify_if_time_from: Optional[pulumi.Input[str]] = None,
                 notify_if_time_to: Optional[pulumi.Input[str]] = None,
                 notify_on_call_from_schedule: Optional[pulumi.Input[str]] = None,
                 notify_to_team_members: Optional[pulumi.Input[str]] = None,
                 persons_to_notifies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 persons_to_notify_next_each_times: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OncallEscalationArgs.__new__(OncallEscalationArgs)

            __props__.__dict__["action_to_trigger"] = action_to_trigger
            __props__.__dict__["duration"] = duration
            if escalation_chain_id is None and not opts.urn:
                raise TypeError("Missing required property 'escalation_chain_id'")
            __props__.__dict__["escalation_chain_id"] = escalation_chain_id
            __props__.__dict__["group_to_notify"] = group_to_notify
            __props__.__dict__["important"] = important
            __props__.__dict__["notify_if_time_from"] = notify_if_time_from
            __props__.__dict__["notify_if_time_to"] = notify_if_time_to
            __props__.__dict__["notify_on_call_from_schedule"] = notify_on_call_from_schedule
            __props__.__dict__["notify_to_team_members"] = notify_to_team_members
            __props__.__dict__["persons_to_notifies"] = persons_to_notifies
            __props__.__dict__["persons_to_notify_next_each_times"] = persons_to_notify_next_each_times
            if position is None and not opts.urn:
                raise TypeError("Missing required property 'position'")
            __props__.__dict__["position"] = position
            __props__.__dict__["type"] = type
        super(OncallEscalation, __self__).__init__(
            'grafana:index/oncallEscalation:OncallEscalation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_to_trigger: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[int]] = None,
            escalation_chain_id: Optional[pulumi.Input[str]] = None,
            group_to_notify: Optional[pulumi.Input[str]] = None,
            important: Optional[pulumi.Input[bool]] = None,
            notify_if_time_from: Optional[pulumi.Input[str]] = None,
            notify_if_time_to: Optional[pulumi.Input[str]] = None,
            notify_on_call_from_schedule: Optional[pulumi.Input[str]] = None,
            notify_to_team_members: Optional[pulumi.Input[str]] = None,
            persons_to_notifies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            persons_to_notify_next_each_times: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            position: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'OncallEscalation':
        """
        Get an existing OncallEscalation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action_to_trigger: The ID of an Action for trigger_webhook type step.
        :param pulumi.Input[int] duration: The duration of delay for wait type step.
        :param pulumi.Input[str] escalation_chain_id: The ID of the escalation chain.
        :param pulumi.Input[str] group_to_notify: The ID of a User Group for notify_user_group type step.
        :param pulumi.Input[bool] important: Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
               and notify_user_group,notify_team_members
        :param pulumi.Input[str] notify_if_time_from: The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        :param pulumi.Input[str] notify_if_time_to: The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        :param pulumi.Input[str] notify_on_call_from_schedule: ID of a Schedule for notify_on_call_from_schedule type step.
        :param pulumi.Input[str] notify_to_team_members: The ID of a Team for a notify_team_members type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notifies: The list of ID's of users for notify_persons type step.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] persons_to_notify_next_each_times: The list of ID's of users for notify_person_next_each_time type step.
        :param pulumi.Input[int] position: The position of the escalation step (starts from 0).
        :param pulumi.Input[str] type: The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
               trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
               notify_team_members
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OncallEscalationState.__new__(_OncallEscalationState)

        __props__.__dict__["action_to_trigger"] = action_to_trigger
        __props__.__dict__["duration"] = duration
        __props__.__dict__["escalation_chain_id"] = escalation_chain_id
        __props__.__dict__["group_to_notify"] = group_to_notify
        __props__.__dict__["important"] = important
        __props__.__dict__["notify_if_time_from"] = notify_if_time_from
        __props__.__dict__["notify_if_time_to"] = notify_if_time_to
        __props__.__dict__["notify_on_call_from_schedule"] = notify_on_call_from_schedule
        __props__.__dict__["notify_to_team_members"] = notify_to_team_members
        __props__.__dict__["persons_to_notifies"] = persons_to_notifies
        __props__.__dict__["persons_to_notify_next_each_times"] = persons_to_notify_next_each_times
        __props__.__dict__["position"] = position
        __props__.__dict__["type"] = type
        return OncallEscalation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionToTrigger")
    def action_to_trigger(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of an Action for trigger_webhook type step.
        """
        return pulumi.get(self, "action_to_trigger")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[int]]:
        """
        The duration of delay for wait type step.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="escalationChainId")
    def escalation_chain_id(self) -> pulumi.Output[str]:
        """
        The ID of the escalation chain.
        """
        return pulumi.get(self, "escalation_chain_id")

    @property
    @pulumi.getter(name="groupToNotify")
    def group_to_notify(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a User Group for notify_user_group type step.
        """
        return pulumi.get(self, "group_to_notify")

    @property
    @pulumi.getter
    def important(self) -> pulumi.Output[Optional[bool]]:
        """
        Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule
        and notify_user_group,notify_team_members
        """
        return pulumi.get(self, "important")

    @property
    @pulumi.getter(name="notifyIfTimeFrom")
    def notify_if_time_from(self) -> pulumi.Output[Optional[str]]:
        """
        The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_from")

    @property
    @pulumi.getter(name="notifyIfTimeTo")
    def notify_if_time_to(self) -> pulumi.Output[Optional[str]]:
        """
        The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
        """
        return pulumi.get(self, "notify_if_time_to")

    @property
    @pulumi.getter(name="notifyOnCallFromSchedule")
    def notify_on_call_from_schedule(self) -> pulumi.Output[Optional[str]]:
        """
        ID of a Schedule for notify_on_call_from_schedule type step.
        """
        return pulumi.get(self, "notify_on_call_from_schedule")

    @property
    @pulumi.getter(name="notifyToTeamMembers")
    def notify_to_team_members(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of a Team for a notify_team_members type step.
        """
        return pulumi.get(self, "notify_to_team_members")

    @property
    @pulumi.getter(name="personsToNotifies")
    def persons_to_notifies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of ID's of users for notify_persons type step.
        """
        return pulumi.get(self, "persons_to_notifies")

    @property
    @pulumi.getter(name="personsToNotifyNextEachTimes")
    def persons_to_notify_next_each_times(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of ID's of users for notify_person_next_each_time type step.
        """
        return pulumi.get(self, "persons_to_notify_next_each_times")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[int]:
        """
        The position of the escalation step (starts from 0).
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule,
        trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation,
        notify_team_members
        """
        return pulumi.get(self, "type")

