# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOncallUserGroupResult',
    'AwaitableGetOncallUserGroupResult',
    'get_oncall_user_group',
    'get_oncall_user_group_output',
]

@pulumi.output_type
class GetOncallUserGroupResult:
    """
    A collection of values returned by getOncallUserGroup.
    """
    def __init__(__self__, id=None, slack_handle=None, slack_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if slack_handle and not isinstance(slack_handle, str):
            raise TypeError("Expected argument 'slack_handle' to be a str")
        pulumi.set(__self__, "slack_handle", slack_handle)
        if slack_id and not isinstance(slack_id, str):
            raise TypeError("Expected argument 'slack_id' to be a str")
        pulumi.set(__self__, "slack_id", slack_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="slackHandle")
    def slack_handle(self) -> str:
        return pulumi.get(self, "slack_handle")

    @property
    @pulumi.getter(name="slackId")
    def slack_id(self) -> str:
        return pulumi.get(self, "slack_id")


class AwaitableGetOncallUserGroupResult(GetOncallUserGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOncallUserGroupResult(
            id=self.id,
            slack_handle=self.slack_handle,
            slack_id=self.slack_id)


def get_oncall_user_group(slack_handle: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOncallUserGroupResult:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/user_groups/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_user_group = grafana.get_oncall_user_group(slack_handle="example_slack_handle")
    ```
    """
    __args__ = dict()
    __args__['slackHandle'] = slack_handle
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOncallUserGroup:getOncallUserGroup', __args__, opts=opts, typ=GetOncallUserGroupResult).value

    return AwaitableGetOncallUserGroupResult(
        id=pulumi.get(__ret__, 'id'),
        slack_handle=pulumi.get(__ret__, 'slack_handle'),
        slack_id=pulumi.get(__ret__, 'slack_id'))


@_utilities.lift_output_func(get_oncall_user_group)
def get_oncall_user_group_output(slack_handle: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOncallUserGroupResult]:
    """
    * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/user_groups/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_user_group = grafana.get_oncall_user_group(slack_handle="example_slack_handle")
    ```
    """
    ...
