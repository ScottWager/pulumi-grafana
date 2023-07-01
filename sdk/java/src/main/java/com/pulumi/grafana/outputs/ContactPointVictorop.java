// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContactPointVictorop {
    /**
     * @return Templated description of the message.
     * 
     */
    private @Nullable String description;
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    private @Nullable Boolean disableResolveMessage;
    /**
     * @return The VictorOps alert state - typically either `CRITICAL` or `RECOVERY`.
     * 
     */
    private @Nullable String messageType;
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    private @Nullable Map<String,String> settings;
    /**
     * @return Templated title to display.
     * 
     */
    private @Nullable String title;
    /**
     * @return The UID of the contact point.
     * 
     */
    private @Nullable String uid;
    /**
     * @return The VictorOps webhook URL.
     * 
     */
    private String url;

    private ContactPointVictorop() {}
    /**
     * @return Templated description of the message.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Whether to disable sending resolve messages. Defaults to `false`.
     * 
     */
    public Optional<Boolean> disableResolveMessage() {
        return Optional.ofNullable(this.disableResolveMessage);
    }
    /**
     * @return The VictorOps alert state - typically either `CRITICAL` or `RECOVERY`.
     * 
     */
    public Optional<String> messageType() {
        return Optional.ofNullable(this.messageType);
    }
    /**
     * @return Additional custom properties to attach to the notifier. Defaults to `map[]`.
     * 
     */
    public Map<String,String> settings() {
        return this.settings == null ? Map.of() : this.settings;
    }
    /**
     * @return Templated title to display.
     * 
     */
    public Optional<String> title() {
        return Optional.ofNullable(this.title);
    }
    /**
     * @return The UID of the contact point.
     * 
     */
    public Optional<String> uid() {
        return Optional.ofNullable(this.uid);
    }
    /**
     * @return The VictorOps webhook URL.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContactPointVictorop defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable Boolean disableResolveMessage;
        private @Nullable String messageType;
        private @Nullable Map<String,String> settings;
        private @Nullable String title;
        private @Nullable String uid;
        private String url;
        public Builder() {}
        public Builder(ContactPointVictorop defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.disableResolveMessage = defaults.disableResolveMessage;
    	      this.messageType = defaults.messageType;
    	      this.settings = defaults.settings;
    	      this.title = defaults.title;
    	      this.uid = defaults.uid;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disableResolveMessage(@Nullable Boolean disableResolveMessage) {
            this.disableResolveMessage = disableResolveMessage;
            return this;
        }
        @CustomType.Setter
        public Builder messageType(@Nullable String messageType) {
            this.messageType = messageType;
            return this;
        }
        @CustomType.Setter
        public Builder settings(@Nullable Map<String,String> settings) {
            this.settings = settings;
            return this;
        }
        @CustomType.Setter
        public Builder title(@Nullable String title) {
            this.title = title;
            return this;
        }
        @CustomType.Setter
        public Builder uid(@Nullable String uid) {
            this.uid = uid;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public ContactPointVictorop build() {
            final var o = new ContactPointVictorop();
            o.description = description;
            o.disableResolveMessage = disableResolveMessage;
            o.messageType = messageType;
            o.settings = settings;
            o.title = title;
            o.uid = uid;
            o.url = url;
            return o;
        }
    }
}
