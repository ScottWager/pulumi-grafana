# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert_notification import *
from .annotation import *
from .api_key import *
from .builtin_role_assignment import *
from .cloud_access_policy import *
from .cloud_access_policy_token import *
from .cloud_api_key import *
from .cloud_plugin_installation import *
from .cloud_stack import *
from .cloud_stack_api_key import *
from .cloud_stack_service_account import *
from .cloud_stack_service_account_token import *
from .contact_point import *
from .dashboard import *
from .dashboard_permission import *
from .data_source import *
from .data_source_permission import *
from .folder import *
from .folder_permission import *
from .get_cloud_ips import *
from .get_cloud_organization import *
from .get_cloud_stack import *
from .get_dashboard import *
from .get_dashboards import *
from .get_data_source import *
from .get_folder import *
from .get_folders import *
from .get_library_panel import *
from .get_on_call_slack_channel import *
from .get_oncall_action import *
from .get_oncall_escalation_chain import *
from .get_oncall_outgoing_webhook import *
from .get_oncall_schedule import *
from .get_oncall_team import *
from .get_oncall_user import *
from .get_oncall_user_group import *
from .get_organization import *
from .get_organization_preferences import *
from .get_slos import *
from .get_synthetic_monitoring_probe import *
from .get_synthetic_monitoring_probes import *
from .get_team import *
from .get_user import *
from .get_users import *
from .library_panel import *
from .machine_learning_holiday import *
from .machine_learning_job import *
from .machine_learning_outlier_detector import *
from .message_template import *
from .mute_timing import *
from .notification_policy import *
from .oncall_escalation import *
from .oncall_escalation_chain import *
from .oncall_integration import *
from .oncall_on_call_shift import *
from .oncall_outgoing_webhook import *
from .oncall_route import *
from .oncall_schedule import *
from .organization import *
from .organization_preference import *
from .playlist import *
from .provider import *
from .report import *
from .role import *
from .role_assignment import *
from .rule_group import *
from .service_account import *
from .service_account_permission import *
from .service_account_token import *
from .slo import *
from .synthetic_monitoring_check import *
from .synthetic_monitoring_installation import *
from .synthetic_monitoring_probe import *
from .team import *
from .team_external_group import *
from .team_preferences import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import lbrlabs_pulumi_grafana.config as __config
    config = __config
else:
    config = _utilities.lazy_import('lbrlabs_pulumi_grafana.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "grafana",
  "mod": "index/alertNotification",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/alertNotification:AlertNotification": "AlertNotification"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/annotation",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/annotation:Annotation": "Annotation"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/apiKey",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/apiKey:ApiKey": "ApiKey"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/builtinRoleAssignment",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/builtinRoleAssignment:BuiltinRoleAssignment": "BuiltinRoleAssignment"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudAccessPolicy",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudAccessPolicy:CloudAccessPolicy": "CloudAccessPolicy"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudAccessPolicyToken",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudAccessPolicyToken:CloudAccessPolicyToken": "CloudAccessPolicyToken"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudApiKey",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudApiKey:CloudApiKey": "CloudApiKey"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudPluginInstallation",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudPluginInstallation:CloudPluginInstallation": "CloudPluginInstallation"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudStack",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudStack:CloudStack": "CloudStack"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudStackApiKey",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudStackApiKey:CloudStackApiKey": "CloudStackApiKey"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudStackServiceAccount",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudStackServiceAccount:CloudStackServiceAccount": "CloudStackServiceAccount"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/cloudStackServiceAccountToken",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/cloudStackServiceAccountToken:CloudStackServiceAccountToken": "CloudStackServiceAccountToken"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/contactPoint",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/contactPoint:ContactPoint": "ContactPoint"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/dashboard",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/dashboard:Dashboard": "Dashboard"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/dashboardPermission",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/dashboardPermission:DashboardPermission": "DashboardPermission"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/dataSource",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/dataSource:DataSource": "DataSource"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/dataSourcePermission",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/dataSourcePermission:DataSourcePermission": "DataSourcePermission"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/folder",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/folder:Folder": "Folder"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/folderPermission",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/folderPermission:FolderPermission": "FolderPermission"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/libraryPanel",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/libraryPanel:LibraryPanel": "LibraryPanel"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/machineLearningHoliday",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/machineLearningHoliday:MachineLearningHoliday": "MachineLearningHoliday"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/machineLearningJob",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/machineLearningJob:MachineLearningJob": "MachineLearningJob"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/machineLearningOutlierDetector",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/machineLearningOutlierDetector:MachineLearningOutlierDetector": "MachineLearningOutlierDetector"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/messageTemplate",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/messageTemplate:MessageTemplate": "MessageTemplate"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/muteTiming",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/muteTiming:MuteTiming": "MuteTiming"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/notificationPolicy",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/notificationPolicy:NotificationPolicy": "NotificationPolicy"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallEscalation",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallEscalation:OncallEscalation": "OncallEscalation"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallEscalationChain",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallEscalationChain:OncallEscalationChain": "OncallEscalationChain"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallIntegration",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallIntegration:OncallIntegration": "OncallIntegration"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallOnCallShift",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallOnCallShift:OncallOnCallShift": "OncallOnCallShift"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallOutgoingWebhook",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallOutgoingWebhook:OncallOutgoingWebhook": "OncallOutgoingWebhook"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallRoute",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallRoute:OncallRoute": "OncallRoute"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/oncallSchedule",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/oncallSchedule:OncallSchedule": "OncallSchedule"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/organization",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/organization:Organization": "Organization"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/organizationPreference",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/organizationPreference:OrganizationPreference": "OrganizationPreference"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/playlist",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/playlist:Playlist": "Playlist"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/report",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/report:Report": "Report"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/role",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/role:Role": "Role"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/roleAssignment",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/roleAssignment:RoleAssignment": "RoleAssignment"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/ruleGroup",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/ruleGroup:RuleGroup": "RuleGroup"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/sLO",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/sLO:SLO": "SLO"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/serviceAccount",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/serviceAccount:ServiceAccount": "ServiceAccount"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/serviceAccountPermission",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/serviceAccountPermission:ServiceAccountPermission": "ServiceAccountPermission"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/serviceAccountToken",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/serviceAccountToken:ServiceAccountToken": "ServiceAccountToken"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/syntheticMonitoringCheck",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck": "SyntheticMonitoringCheck"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/syntheticMonitoringInstallation",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/syntheticMonitoringInstallation:SyntheticMonitoringInstallation": "SyntheticMonitoringInstallation"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/syntheticMonitoringProbe",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/syntheticMonitoringProbe:SyntheticMonitoringProbe": "SyntheticMonitoringProbe"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/team",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/team:Team": "Team"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/teamExternalGroup",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/teamExternalGroup:TeamExternalGroup": "TeamExternalGroup"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/teamPreferences",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/teamPreferences:TeamPreferences": "TeamPreferences"
  }
 },
 {
  "pkg": "grafana",
  "mod": "index/user",
  "fqn": "lbrlabs_pulumi_grafana",
  "classes": {
   "grafana:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "grafana",
  "token": "pulumi:providers:grafana",
  "fqn": "lbrlabs_pulumi_grafana",
  "class": "Provider"
 }
]
"""
)
