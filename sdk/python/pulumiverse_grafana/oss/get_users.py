# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, id=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        The Grafana instance's users.
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            id=self.id,
            users=self.users)


def get_users(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)

    This data source uses Grafana's admin APIs for reading users which
    does not currently work with API Tokens. You must use basic auth.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test_all_users = grafana.oss.User("testAllUsers",
        email="all_users@example.com",
        login="test-grafana-users",
        password="my-password")
    all_users = grafana.oss.get_users()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('grafana:oss/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        id=pulumi.get(__ret__, 'id'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_users)
def get_users_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
    * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)

    This data source uses Grafana's admin APIs for reading users which
    does not currently work with API Tokens. You must use basic auth.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_grafana as grafana
    import pulumiverse_grafana as grafana

    test_all_users = grafana.oss.User("testAllUsers",
        email="all_users@example.com",
        login="test-grafana-users",
        password="my-password")
    all_users = grafana.oss.get_users()
    ```
    """
    ...