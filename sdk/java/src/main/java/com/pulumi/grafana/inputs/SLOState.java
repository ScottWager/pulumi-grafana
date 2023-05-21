// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.grafana.inputs.SLOAlertingArgs;
import com.pulumi.grafana.inputs.SLOLabelArgs;
import com.pulumi.grafana.inputs.SLOObjectiveArgs;
import com.pulumi.grafana.inputs.SLOQueryArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SLOState extends com.pulumi.resources.ResourceArgs {

    public static final SLOState Empty = new SLOState();

    /**
     * Configures the alerting rules that will be generated for each
     * 			time window associated with the SLO. Grafana SLOs can generate
     * 			alerts when the short-term error budget burn is very high, the
     * 			long-term error budget burn rate is high, or when the remaining
     * 			error budget is below a certain threshold.
     * 
     */
    @Import(name="alertings")
    private @Nullable Output<List<SLOAlertingArgs>> alertings;

    /**
     * @return Configures the alerting rules that will be generated for each
     * 			time window associated with the SLO. Grafana SLOs can generate
     * 			alerts when the short-term error budget burn is very high, the
     * 			long-term error budget burn rate is high, or when the remaining
     * 			error budget is below a certain threshold.
     * 
     */
    public Optional<Output<List<SLOAlertingArgs>>> alertings() {
        return Optional.ofNullable(this.alertings);
    }

    /**
     * A reference to a dashboard that the plugin has installed in Grafana based on this SLO. This field is read-only, it is generated by the Grafana SLO Plugin.
     * 
     */
    @Import(name="dashboardUid")
    private @Nullable Output<String> dashboardUid;

    /**
     * @return A reference to a dashboard that the plugin has installed in Grafana based on this SLO. This field is read-only, it is generated by the Grafana SLO Plugin.
     * 
     */
    public Optional<Output<String>> dashboardUid() {
        return Optional.ofNullable(this.dashboardUid);
    }

    /**
     * Description is a free-text field that can provide more context to an SLO.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description is a free-text field that can provide more context to an SLO.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<SLOLabelArgs>> labels;

    /**
     * @return Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand.
     * 
     */
    public Optional<Output<List<SLOLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Name should be a short description of your indicator. Consider names like &#34;API Availability&#34;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name should be a short description of your indicator. Consider names like &#34;API Availability&#34;
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
     * 
     */
    @Import(name="objectives")
    private @Nullable Output<List<SLOObjectiveArgs>> objectives;

    /**
     * @return Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
     * 
     */
    public Optional<Output<List<SLOObjectiveArgs>>> objectives() {
        return Optional.ofNullable(this.objectives);
    }

    /**
     * Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
     * 
     */
    @Import(name="queries")
    private @Nullable Output<List<SLOQueryArgs>> queries;

    /**
     * @return Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
     * 
     */
    public Optional<Output<List<SLOQueryArgs>>> queries() {
        return Optional.ofNullable(this.queries);
    }

    private SLOState() {}

    private SLOState(SLOState $) {
        this.alertings = $.alertings;
        this.dashboardUid = $.dashboardUid;
        this.description = $.description;
        this.labels = $.labels;
        this.name = $.name;
        this.objectives = $.objectives;
        this.queries = $.queries;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SLOState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SLOState $;

        public Builder() {
            $ = new SLOState();
        }

        public Builder(SLOState defaults) {
            $ = new SLOState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertings Configures the alerting rules that will be generated for each
         * 			time window associated with the SLO. Grafana SLOs can generate
         * 			alerts when the short-term error budget burn is very high, the
         * 			long-term error budget burn rate is high, or when the remaining
         * 			error budget is below a certain threshold.
         * 
         * @return builder
         * 
         */
        public Builder alertings(@Nullable Output<List<SLOAlertingArgs>> alertings) {
            $.alertings = alertings;
            return this;
        }

        /**
         * @param alertings Configures the alerting rules that will be generated for each
         * 			time window associated with the SLO. Grafana SLOs can generate
         * 			alerts when the short-term error budget burn is very high, the
         * 			long-term error budget burn rate is high, or when the remaining
         * 			error budget is below a certain threshold.
         * 
         * @return builder
         * 
         */
        public Builder alertings(List<SLOAlertingArgs> alertings) {
            return alertings(Output.of(alertings));
        }

        /**
         * @param alertings Configures the alerting rules that will be generated for each
         * 			time window associated with the SLO. Grafana SLOs can generate
         * 			alerts when the short-term error budget burn is very high, the
         * 			long-term error budget burn rate is high, or when the remaining
         * 			error budget is below a certain threshold.
         * 
         * @return builder
         * 
         */
        public Builder alertings(SLOAlertingArgs... alertings) {
            return alertings(List.of(alertings));
        }

        /**
         * @param dashboardUid A reference to a dashboard that the plugin has installed in Grafana based on this SLO. This field is read-only, it is generated by the Grafana SLO Plugin.
         * 
         * @return builder
         * 
         */
        public Builder dashboardUid(@Nullable Output<String> dashboardUid) {
            $.dashboardUid = dashboardUid;
            return this;
        }

        /**
         * @param dashboardUid A reference to a dashboard that the plugin has installed in Grafana based on this SLO. This field is read-only, it is generated by the Grafana SLO Plugin.
         * 
         * @return builder
         * 
         */
        public Builder dashboardUid(String dashboardUid) {
            return dashboardUid(Output.of(dashboardUid));
        }

        /**
         * @param description Description is a free-text field that can provide more context to an SLO.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description is a free-text field that can provide more context to an SLO.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param labels Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<SLOLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand.
         * 
         * @return builder
         * 
         */
        public Builder labels(List<SLOLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand.
         * 
         * @return builder
         * 
         */
        public Builder labels(SLOLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param name Name should be a short description of your indicator. Consider names like &#34;API Availability&#34;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name should be a short description of your indicator. Consider names like &#34;API Availability&#34;
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param objectives Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
         * 
         * @return builder
         * 
         */
        public Builder objectives(@Nullable Output<List<SLOObjectiveArgs>> objectives) {
            $.objectives = objectives;
            return this;
        }

        /**
         * @param objectives Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
         * 
         * @return builder
         * 
         */
        public Builder objectives(List<SLOObjectiveArgs> objectives) {
            return objectives(Output.of(objectives));
        }

        /**
         * @param objectives Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
         * 
         * @return builder
         * 
         */
        public Builder objectives(SLOObjectiveArgs... objectives) {
            return objectives(List.of(objectives));
        }

        /**
         * @param queries Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
         * 
         * @return builder
         * 
         */
        public Builder queries(@Nullable Output<List<SLOQueryArgs>> queries) {
            $.queries = queries;
            return this;
        }

        /**
         * @param queries Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
         * 
         * @return builder
         * 
         */
        public Builder queries(List<SLOQueryArgs> queries) {
            return queries(Output.of(queries));
        }

        /**
         * @param queries Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
         * 
         * @return builder
         * 
         */
        public Builder queries(SLOQueryArgs... queries) {
            return queries(List.of(queries));
        }

        public SLOState build() {
            return $;
        }
    }

}
