// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesEmail;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesMicrosoftTeams;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesPhoneCall;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesSlack;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesSms;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesTelegram;
import com.pulumi.grafana.outputs.OncallIntegrationTemplatesWeb;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OncallIntegrationTemplates {
    /**
     * @return Template for sending a signal to acknowledge the Incident.
     * 
     */
    private @Nullable String acknowledgeSignal;
    /**
     * @return Templates for Email.
     * 
     */
    private @Nullable OncallIntegrationTemplatesEmail email;
    /**
     * @return Template for the key by which alerts are grouped.
     * 
     */
    private @Nullable String groupingKey;
    /**
     * @return Templates for Microsoft Teams.
     * 
     */
    private @Nullable OncallIntegrationTemplatesMicrosoftTeams microsoftTeams;
    /**
     * @return Templates for Phone Call.
     * 
     */
    private @Nullable OncallIntegrationTemplatesPhoneCall phoneCall;
    /**
     * @return Template for sending a signal to resolve the Incident.
     * 
     */
    private @Nullable String resolveSignal;
    /**
     * @return Templates for Slack.
     * 
     */
    private @Nullable OncallIntegrationTemplatesSlack slack;
    /**
     * @return Templates for SMS.
     * 
     */
    private @Nullable OncallIntegrationTemplatesSms sms;
    /**
     * @return Template for a source link.
     * 
     */
    private @Nullable String sourceLink;
    /**
     * @return Templates for Telegram.
     * 
     */
    private @Nullable OncallIntegrationTemplatesTelegram telegram;
    /**
     * @return Templates for Web.
     * 
     */
    private @Nullable OncallIntegrationTemplatesWeb web;

    private OncallIntegrationTemplates() {}
    /**
     * @return Template for sending a signal to acknowledge the Incident.
     * 
     */
    public Optional<String> acknowledgeSignal() {
        return Optional.ofNullable(this.acknowledgeSignal);
    }
    /**
     * @return Templates for Email.
     * 
     */
    public Optional<OncallIntegrationTemplatesEmail> email() {
        return Optional.ofNullable(this.email);
    }
    /**
     * @return Template for the key by which alerts are grouped.
     * 
     */
    public Optional<String> groupingKey() {
        return Optional.ofNullable(this.groupingKey);
    }
    /**
     * @return Templates for Microsoft Teams.
     * 
     */
    public Optional<OncallIntegrationTemplatesMicrosoftTeams> microsoftTeams() {
        return Optional.ofNullable(this.microsoftTeams);
    }
    /**
     * @return Templates for Phone Call.
     * 
     */
    public Optional<OncallIntegrationTemplatesPhoneCall> phoneCall() {
        return Optional.ofNullable(this.phoneCall);
    }
    /**
     * @return Template for sending a signal to resolve the Incident.
     * 
     */
    public Optional<String> resolveSignal() {
        return Optional.ofNullable(this.resolveSignal);
    }
    /**
     * @return Templates for Slack.
     * 
     */
    public Optional<OncallIntegrationTemplatesSlack> slack() {
        return Optional.ofNullable(this.slack);
    }
    /**
     * @return Templates for SMS.
     * 
     */
    public Optional<OncallIntegrationTemplatesSms> sms() {
        return Optional.ofNullable(this.sms);
    }
    /**
     * @return Template for a source link.
     * 
     */
    public Optional<String> sourceLink() {
        return Optional.ofNullable(this.sourceLink);
    }
    /**
     * @return Templates for Telegram.
     * 
     */
    public Optional<OncallIntegrationTemplatesTelegram> telegram() {
        return Optional.ofNullable(this.telegram);
    }
    /**
     * @return Templates for Web.
     * 
     */
    public Optional<OncallIntegrationTemplatesWeb> web() {
        return Optional.ofNullable(this.web);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OncallIntegrationTemplates defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String acknowledgeSignal;
        private @Nullable OncallIntegrationTemplatesEmail email;
        private @Nullable String groupingKey;
        private @Nullable OncallIntegrationTemplatesMicrosoftTeams microsoftTeams;
        private @Nullable OncallIntegrationTemplatesPhoneCall phoneCall;
        private @Nullable String resolveSignal;
        private @Nullable OncallIntegrationTemplatesSlack slack;
        private @Nullable OncallIntegrationTemplatesSms sms;
        private @Nullable String sourceLink;
        private @Nullable OncallIntegrationTemplatesTelegram telegram;
        private @Nullable OncallIntegrationTemplatesWeb web;
        public Builder() {}
        public Builder(OncallIntegrationTemplates defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acknowledgeSignal = defaults.acknowledgeSignal;
    	      this.email = defaults.email;
    	      this.groupingKey = defaults.groupingKey;
    	      this.microsoftTeams = defaults.microsoftTeams;
    	      this.phoneCall = defaults.phoneCall;
    	      this.resolveSignal = defaults.resolveSignal;
    	      this.slack = defaults.slack;
    	      this.sms = defaults.sms;
    	      this.sourceLink = defaults.sourceLink;
    	      this.telegram = defaults.telegram;
    	      this.web = defaults.web;
        }

        @CustomType.Setter
        public Builder acknowledgeSignal(@Nullable String acknowledgeSignal) {
            this.acknowledgeSignal = acknowledgeSignal;
            return this;
        }
        @CustomType.Setter
        public Builder email(@Nullable OncallIntegrationTemplatesEmail email) {
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder groupingKey(@Nullable String groupingKey) {
            this.groupingKey = groupingKey;
            return this;
        }
        @CustomType.Setter
        public Builder microsoftTeams(@Nullable OncallIntegrationTemplatesMicrosoftTeams microsoftTeams) {
            this.microsoftTeams = microsoftTeams;
            return this;
        }
        @CustomType.Setter
        public Builder phoneCall(@Nullable OncallIntegrationTemplatesPhoneCall phoneCall) {
            this.phoneCall = phoneCall;
            return this;
        }
        @CustomType.Setter
        public Builder resolveSignal(@Nullable String resolveSignal) {
            this.resolveSignal = resolveSignal;
            return this;
        }
        @CustomType.Setter
        public Builder slack(@Nullable OncallIntegrationTemplatesSlack slack) {
            this.slack = slack;
            return this;
        }
        @CustomType.Setter
        public Builder sms(@Nullable OncallIntegrationTemplatesSms sms) {
            this.sms = sms;
            return this;
        }
        @CustomType.Setter
        public Builder sourceLink(@Nullable String sourceLink) {
            this.sourceLink = sourceLink;
            return this;
        }
        @CustomType.Setter
        public Builder telegram(@Nullable OncallIntegrationTemplatesTelegram telegram) {
            this.telegram = telegram;
            return this;
        }
        @CustomType.Setter
        public Builder web(@Nullable OncallIntegrationTemplatesWeb web) {
            this.web = web;
            return this;
        }
        public OncallIntegrationTemplates build() {
            final var o = new OncallIntegrationTemplates();
            o.acknowledgeSignal = acknowledgeSignal;
            o.email = email;
            o.groupingKey = groupingKey;
            o.microsoftTeams = microsoftTeams;
            o.phoneCall = phoneCall;
            o.resolveSignal = resolveSignal;
            o.slack = slack;
            o.sms = sms;
            o.sourceLink = sourceLink;
            o.telegram = telegram;
            o.web = web;
            return o;
        }
    }
}
