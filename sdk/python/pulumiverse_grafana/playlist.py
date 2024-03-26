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

__all__ = ['PlaylistArgs', 'Playlist']

@pulumi.input_type
class PlaylistArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[str],
                 items: pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]],
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Playlist resource.
        :param pulumi.Input[str] name: The name of the playlist.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "items", items)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[str]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[str]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def items(self) -> pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]]:
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]]):
        pulumi.set(self, "items", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the playlist.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _PlaylistState:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Playlist resources.
        :param pulumi.Input[str] name: The name of the playlist.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if items is not None:
            pulumi.set(__self__, "items", items)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]]]:
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PlaylistItemArgs']]]]):
        pulumi.set(self, "items", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the playlist.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


class Playlist(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interval: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaylistItemArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/playlist/)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.Playlist("test",
            interval="5m",
            items=[
                grafana.PlaylistItemArgs(
                    order=2,
                    title="Terraform Dashboard By Tag",
                    type="dashboard_by_tag",
                    value="terraform",
                ),
                grafana.PlaylistItemArgs(
                    order=1,
                    title="Terraform Dashboard By ID",
                    type="dashboard_by_id",
                    value="3",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the playlist.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PlaylistArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/)
        * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/playlist/)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_grafana as grafana

        test = grafana.Playlist("test",
            interval="5m",
            items=[
                grafana.PlaylistItemArgs(
                    order=2,
                    title="Terraform Dashboard By Tag",
                    type="dashboard_by_tag",
                    value="terraform",
                ),
                grafana.PlaylistItemArgs(
                    order=1,
                    title="Terraform Dashboard By ID",
                    type="dashboard_by_id",
                    value="3",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param PlaylistArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PlaylistArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 interval: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaylistItemArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PlaylistArgs.__new__(PlaylistArgs)

            if interval is None and not opts.urn:
                raise TypeError("Missing required property 'interval'")
            __props__.__dict__["interval"] = interval
            if items is None and not opts.urn:
                raise TypeError("Missing required property 'items'")
            __props__.__dict__["items"] = items
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
        super(Playlist, __self__).__init__(
            'grafana:index/playlist:Playlist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            interval: Optional[pulumi.Input[str]] = None,
            items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaylistItemArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None) -> 'Playlist':
        """
        Get an existing Playlist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the playlist.
        :param pulumi.Input[str] org_id: The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PlaylistState.__new__(_PlaylistState)

        __props__.__dict__["interval"] = interval
        __props__.__dict__["items"] = items
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        return Playlist(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def items(self) -> pulumi.Output[Sequence['outputs.PlaylistItem']]:
        return pulumi.get(self, "items")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the playlist.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Organization ID. If not set, the Org ID defined in the provider block will be used.
        """
        return pulumi.get(self, "org_id")

