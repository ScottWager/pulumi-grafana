// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Grafana
{
    public static class GetOrganizationPreferences
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/manage-organizations/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/preferences/#get-current-org-prefs)
        /// </summary>
        public static Task<GetOrganizationPreferencesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationPreferencesResult>("grafana:index/getOrganizationPreferences:getOrganizationPreferences", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationPreferencesResult
    {
        /// <summary>
        /// The Organization home dashboard ID.
        /// </summary>
        public readonly int HomeDashboardId;
        /// <summary>
        /// The Organization home dashboard UID.
        /// </summary>
        public readonly string HomeDashboardUid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Organization theme. Available values are `light`, `dark`, or an empty string for the default.
        /// </summary>
        public readonly string Theme;
        /// <summary>
        /// The Organization timezone. Available values are `utc`, `browser`, or an empty string for the default.
        /// </summary>
        public readonly string Timezone;
        /// <summary>
        /// The Organization week start.
        /// </summary>
        public readonly string WeekStart;

        [OutputConstructor]
        private GetOrganizationPreferencesResult(
            int homeDashboardId,

            string homeDashboardUid,

            string id,

            string theme,

            string timezone,

            string weekStart)
        {
            HomeDashboardId = homeDashboardId;
            HomeDashboardUid = homeDashboardUid;
            Id = id;
            Theme = theme;
            Timezone = timezone;
            WeekStart = weekStart;
        }
    }
}