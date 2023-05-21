// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OncallIntegrationTemplatesTelegramArgs extends com.pulumi.resources.ResourceArgs {

    public static final OncallIntegrationTemplatesTelegramArgs Empty = new OncallIntegrationTemplatesTelegramArgs();

    @Import(name="imageUrl")
    private @Nullable Output<String> imageUrl;

    public Optional<Output<String>> imageUrl() {
        return Optional.ofNullable(this.imageUrl);
    }

    @Import(name="message")
    private @Nullable Output<String> message;

    public Optional<Output<String>> message() {
        return Optional.ofNullable(this.message);
    }

    @Import(name="title")
    private @Nullable Output<String> title;

    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    private OncallIntegrationTemplatesTelegramArgs() {}

    private OncallIntegrationTemplatesTelegramArgs(OncallIntegrationTemplatesTelegramArgs $) {
        this.imageUrl = $.imageUrl;
        this.message = $.message;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OncallIntegrationTemplatesTelegramArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OncallIntegrationTemplatesTelegramArgs $;

        public Builder() {
            $ = new OncallIntegrationTemplatesTelegramArgs();
        }

        public Builder(OncallIntegrationTemplatesTelegramArgs defaults) {
            $ = new OncallIntegrationTemplatesTelegramArgs(Objects.requireNonNull(defaults));
        }

        public Builder imageUrl(@Nullable Output<String> imageUrl) {
            $.imageUrl = imageUrl;
            return this;
        }

        public Builder imageUrl(String imageUrl) {
            return imageUrl(Output.of(imageUrl));
        }

        public Builder message(@Nullable Output<String> message) {
            $.message = message;
            return this;
        }

        public Builder message(String message) {
            return message(Output.of(message));
        }

        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        public Builder title(String title) {
            return title(Output.of(title));
        }

        public OncallIntegrationTemplatesTelegramArgs build() {
            return $;
        }
    }

}
