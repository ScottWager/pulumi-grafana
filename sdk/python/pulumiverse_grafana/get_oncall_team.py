# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetOncallTeamResult',
    'AwaitableGetOncallTeamResult',
    'get_oncall_team',
    'get_oncall_team_output',
]

warnings.warn("""grafana.index/getoncallteam.getOncallTeam has been deprecated in favor of grafana.oncall/getteam.getTeam""", DeprecationWarning)

@pulumi.output_type
class GetOncallTeamResult:
    """
    A collection of values returned by getOncallTeam.
    """
    def __init__(__self__, avatar_url=None, email=None, id=None, name=None):
        if avatar_url and not isinstance(avatar_url, str):
            raise TypeError("Expected argument 'avatar_url' to be a str")
        pulumi.set(__self__, "avatar_url", avatar_url)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="avatarUrl")
    def avatar_url(self) -> str:
        return pulumi.get(self, "avatar_url")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The team name.
        """
        return pulumi.get(self, "name")


class AwaitableGetOncallTeamResult(GetOncallTeamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOncallTeamResult(
            avatar_url=self.avatar_url,
            email=self.email,
            id=self.id,
            name=self.name)


def get_oncall_team(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOncallTeamResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_team = grafana.onCall.get_team(name="example_team")
    ```


    :param str name: The team name.
    """
    pulumi.log.warn("""get_oncall_team is deprecated: grafana.index/getoncallteam.getOncallTeam has been deprecated in favor of grafana.oncall/getteam.getTeam""")
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:index/getOncallTeam:getOncallTeam', __args__, opts=opts, typ=GetOncallTeamResult).value

    return AwaitableGetOncallTeamResult(
        avatar_url=pulumi.get(__ret__, 'avatar_url'),
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_oncall_team_output(name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOncallTeamResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana

    example_team = grafana.onCall.get_team(name="example_team")
    ```


    :param str name: The team name.
    """
    pulumi.log.warn("""get_oncall_team is deprecated: grafana.index/getoncallteam.getOncallTeam has been deprecated in favor of grafana.oncall/getteam.getTeam""")
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('grafana:index/getOncallTeam:getOncallTeam', __args__, opts=opts, typ=GetOncallTeamResult)
    return __ret__.apply(lambda __response__: GetOncallTeamResult(
        avatar_url=pulumi.get(__response__, 'avatar_url'),
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
