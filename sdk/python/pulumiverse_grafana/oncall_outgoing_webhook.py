# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OncallOutgoingWebhookArgs', 'OncallOutgoingWebhook']

@pulumi.input_type
class OncallOutgoingWebhookArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str],
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OncallOutgoingWebhook resource.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Forwards whole payload of the alert to the webhook's url as POST data.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        :param pulumi.Input[str] user: The auth data of the webhook. Used for Basic authentication.
        """
        pulumi.set(__self__, "url", url)
        if authorization_header is not None:
            pulumi.set(__self__, "authorization_header", authorization_header)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if forward_whole_payload is not None:
            pulumi.set(__self__, "forward_whole_payload", forward_whole_payload)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @authorization_header.setter
    def authorization_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_header", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> Optional[pulumi.Input[bool]]:
        """
        Forwards whole payload of the alert to the webhook's url as POST data.
        """
        return pulumi.get(self, "forward_whole_payload")

    @forward_whole_payload.setter
    def forward_whole_payload(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forward_whole_payload", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _OncallOutgoingWebhookState:
    def __init__(__self__, *,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OncallOutgoingWebhook resources.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Forwards whole payload of the alert to the webhook's url as POST data.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: The auth data of the webhook. Used for Basic authentication.
        """
        if authorization_header is not None:
            pulumi.set(__self__, "authorization_header", authorization_header)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if forward_whole_payload is not None:
            pulumi.set(__self__, "forward_whole_payload", forward_whole_payload)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @authorization_header.setter
    def authorization_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_header", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> Optional[pulumi.Input[bool]]:
        """
        Forwards whole payload of the alert to the webhook's url as POST data.
        """
        return pulumi.get(self, "forward_whole_payload")

    @forward_whole_payload.setter
    def forward_whole_payload(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forward_whole_payload", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The auth data of the webhook. Used for Basic authentication.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class OncallOutgoingWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_acc_outgoing_webhook = grafana.OncallOutgoingWebhook("test-acc-outgoingWebhook", url="https://example.com/",
        opts=pulumi.ResourceOptions(provider=grafana["oncall"]))
        ```

        ## Import

        ```sh
         $ pulumi import grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook outgoing_webhook_name {{outgoing_webhook_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Forwards whole payload of the alert to the webhook's url as POST data.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: The auth data of the webhook. Used for Basic authentication.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OncallOutgoingWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [HTTP API](https://grafana.com/docs/oncall/latest/oncall-api-reference/outgoing_webhooks/)

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test_acc_outgoing_webhook = grafana.OncallOutgoingWebhook("test-acc-outgoingWebhook", url="https://example.com/",
        opts=pulumi.ResourceOptions(provider=grafana["oncall"]))
        ```

        ## Import

        ```sh
         $ pulumi import grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook outgoing_webhook_name {{outgoing_webhook_id}}
        ```

        :param str resource_name: The name of the resource.
        :param OncallOutgoingWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OncallOutgoingWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_header: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 forward_whole_payload: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OncallOutgoingWebhookArgs.__new__(OncallOutgoingWebhookArgs)

            __props__.__dict__["authorization_header"] = None if authorization_header is None else pulumi.Output.secret(authorization_header)
            __props__.__dict__["data"] = data
            __props__.__dict__["forward_whole_payload"] = forward_whole_payload
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["team_id"] = team_id
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["user"] = user
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["authorizationHeader", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(OncallOutgoingWebhook, __self__).__init__(
            'grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_header: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[str]] = None,
            forward_whole_payload: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'OncallOutgoingWebhook':
        """
        Get an existing OncallOutgoingWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_header: The auth data of the webhook. Used in Authorization header instead of user/password auth.
        :param pulumi.Input[str] data: The data of the webhook.
        :param pulumi.Input[bool] forward_whole_payload: Forwards whole payload of the alert to the webhook's url as POST data.
        :param pulumi.Input[str] name: The name of the outgoing webhook.
        :param pulumi.Input[str] password: The auth data of the webhook. Used for Basic authentication
        :param pulumi.Input[str] team_id: The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        :param pulumi.Input[str] url: The webhook URL.
        :param pulumi.Input[str] user: The auth data of the webhook. Used for Basic authentication.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OncallOutgoingWebhookState.__new__(_OncallOutgoingWebhookState)

        __props__.__dict__["authorization_header"] = authorization_header
        __props__.__dict__["data"] = data
        __props__.__dict__["forward_whole_payload"] = forward_whole_payload
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["url"] = url
        __props__.__dict__["user"] = user
        return OncallOutgoingWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationHeader")
    def authorization_header(self) -> pulumi.Output[Optional[str]]:
        """
        The auth data of the webhook. Used in Authorization header instead of user/password auth.
        """
        return pulumi.get(self, "authorization_header")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[str]]:
        """
        The data of the webhook.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="forwardWholePayload")
    def forward_whole_payload(self) -> pulumi.Output[Optional[bool]]:
        """
        Forwards whole payload of the alert to the webhook's url as POST data.
        """
        return pulumi.get(self, "forward_whole_payload")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the outgoing webhook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The auth data of the webhook. Used for Basic authentication
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `get_oncall_team` datasource.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The webhook URL.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        The auth data of the webhook. Used for Basic authentication.
        """
        return pulumi.get(self, "user")
