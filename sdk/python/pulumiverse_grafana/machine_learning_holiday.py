# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MachineLearningHolidayArgs', 'MachineLearningHoliday']

@pulumi.input_type
class MachineLearningHolidayArgs:
    def __init__(__self__, *,
                 custom_periods: Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ical_timezone: Optional[pulumi.Input[str]] = None,
                 ical_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MachineLearningHoliday resource.
        :param pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]] custom_periods: A list of custom periods for the holiday.
        :param pulumi.Input[str] description: A description of the holiday.
        :param pulumi.Input[str] ical_timezone: The timezone to use for events in the iCal file pointed to by ical_url.
        :param pulumi.Input[str] ical_url: A URL to an iCal file containing all occurrences of the holiday.
        :param pulumi.Input[str] name: The name of the holiday.
        """
        if custom_periods is not None:
            pulumi.set(__self__, "custom_periods", custom_periods)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ical_timezone is not None:
            pulumi.set(__self__, "ical_timezone", ical_timezone)
        if ical_url is not None:
            pulumi.set(__self__, "ical_url", ical_url)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="customPeriods")
    def custom_periods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]]:
        """
        A list of custom periods for the holiday.
        """
        return pulumi.get(self, "custom_periods")

    @custom_periods.setter
    def custom_periods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]]):
        pulumi.set(self, "custom_periods", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the holiday.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="icalTimezone")
    def ical_timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The timezone to use for events in the iCal file pointed to by ical_url.
        """
        return pulumi.get(self, "ical_timezone")

    @ical_timezone.setter
    def ical_timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ical_timezone", value)

    @property
    @pulumi.getter(name="icalUrl")
    def ical_url(self) -> Optional[pulumi.Input[str]]:
        """
        A URL to an iCal file containing all occurrences of the holiday.
        """
        return pulumi.get(self, "ical_url")

    @ical_url.setter
    def ical_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ical_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the holiday.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _MachineLearningHolidayState:
    def __init__(__self__, *,
                 custom_periods: Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ical_timezone: Optional[pulumi.Input[str]] = None,
                 ical_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MachineLearningHoliday resources.
        :param pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]] custom_periods: A list of custom periods for the holiday.
        :param pulumi.Input[str] description: A description of the holiday.
        :param pulumi.Input[str] ical_timezone: The timezone to use for events in the iCal file pointed to by ical_url.
        :param pulumi.Input[str] ical_url: A URL to an iCal file containing all occurrences of the holiday.
        :param pulumi.Input[str] name: The name of the holiday.
        """
        if custom_periods is not None:
            pulumi.set(__self__, "custom_periods", custom_periods)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ical_timezone is not None:
            pulumi.set(__self__, "ical_timezone", ical_timezone)
        if ical_url is not None:
            pulumi.set(__self__, "ical_url", ical_url)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="customPeriods")
    def custom_periods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]]:
        """
        A list of custom periods for the holiday.
        """
        return pulumi.get(self, "custom_periods")

    @custom_periods.setter
    def custom_periods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MachineLearningHolidayCustomPeriodArgs']]]]):
        pulumi.set(self, "custom_periods", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the holiday.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="icalTimezone")
    def ical_timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The timezone to use for events in the iCal file pointed to by ical_url.
        """
        return pulumi.get(self, "ical_timezone")

    @ical_timezone.setter
    def ical_timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ical_timezone", value)

    @property
    @pulumi.getter(name="icalUrl")
    def ical_url(self) -> Optional[pulumi.Input[str]]:
        """
        A URL to an iCal file containing all occurrences of the holiday.
        """
        return pulumi.get(self, "ical_url")

    @ical_url.setter
    def ical_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ical_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the holiday.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


warnings.warn("""grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday""", DeprecationWarning)


class MachineLearningHoliday(pulumi.CustomResource):
    warnings.warn("""grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_periods: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MachineLearningHolidayCustomPeriodArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ical_timezone: Optional[pulumi.Input[str]] = None,
                 ical_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A holiday describes time periods where a time series is expected to behave differently to normal.

        To use a holiday in a job, use its id in the `holidays` attribute of a `machineLearning.Job`:

        ## Import

        ```sh
        $ pulumi import grafana:index/machineLearningHoliday:MachineLearningHoliday name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MachineLearningHolidayCustomPeriodArgs']]]] custom_periods: A list of custom periods for the holiday.
        :param pulumi.Input[str] description: A description of the holiday.
        :param pulumi.Input[str] ical_timezone: The timezone to use for events in the iCal file pointed to by ical_url.
        :param pulumi.Input[str] ical_url: A URL to an iCal file containing all occurrences of the holiday.
        :param pulumi.Input[str] name: The name of the holiday.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MachineLearningHolidayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A holiday describes time periods where a time series is expected to behave differently to normal.

        To use a holiday in a job, use its id in the `holidays` attribute of a `machineLearning.Job`:

        ## Import

        ```sh
        $ pulumi import grafana:index/machineLearningHoliday:MachineLearningHoliday name "{{ id }}"
        ```

        :param str resource_name: The name of the resource.
        :param MachineLearningHolidayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MachineLearningHolidayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_periods: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MachineLearningHolidayCustomPeriodArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ical_timezone: Optional[pulumi.Input[str]] = None,
                 ical_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""MachineLearningHoliday is deprecated: grafana.index/machinelearningholiday.MachineLearningHoliday has been deprecated in favor of grafana.machinelearning/holiday.Holiday""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MachineLearningHolidayArgs.__new__(MachineLearningHolidayArgs)

            __props__.__dict__["custom_periods"] = custom_periods
            __props__.__dict__["description"] = description
            __props__.__dict__["ical_timezone"] = ical_timezone
            __props__.__dict__["ical_url"] = ical_url
            __props__.__dict__["name"] = name
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="grafana:index/machineLearningHoliday:MachineLearningHoliday")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MachineLearningHoliday, __self__).__init__(
            'grafana:index/machineLearningHoliday:MachineLearningHoliday',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_periods: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MachineLearningHolidayCustomPeriodArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ical_timezone: Optional[pulumi.Input[str]] = None,
            ical_url: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'MachineLearningHoliday':
        """
        Get an existing MachineLearningHoliday resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MachineLearningHolidayCustomPeriodArgs']]]] custom_periods: A list of custom periods for the holiday.
        :param pulumi.Input[str] description: A description of the holiday.
        :param pulumi.Input[str] ical_timezone: The timezone to use for events in the iCal file pointed to by ical_url.
        :param pulumi.Input[str] ical_url: A URL to an iCal file containing all occurrences of the holiday.
        :param pulumi.Input[str] name: The name of the holiday.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MachineLearningHolidayState.__new__(_MachineLearningHolidayState)

        __props__.__dict__["custom_periods"] = custom_periods
        __props__.__dict__["description"] = description
        __props__.__dict__["ical_timezone"] = ical_timezone
        __props__.__dict__["ical_url"] = ical_url
        __props__.__dict__["name"] = name
        return MachineLearningHoliday(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customPeriods")
    def custom_periods(self) -> pulumi.Output[Optional[Sequence['outputs.MachineLearningHolidayCustomPeriod']]]:
        """
        A list of custom periods for the holiday.
        """
        return pulumi.get(self, "custom_periods")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the holiday.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="icalTimezone")
    def ical_timezone(self) -> pulumi.Output[Optional[str]]:
        """
        The timezone to use for events in the iCal file pointed to by ical_url.
        """
        return pulumi.get(self, "ical_timezone")

    @property
    @pulumi.getter(name="icalUrl")
    def ical_url(self) -> pulumi.Output[Optional[str]]:
        """
        A URL to an iCal file containing all occurrences of the holiday.
        """
        return pulumi.get(self, "ical_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the holiday.
        """
        return pulumi.get(self, "name")

