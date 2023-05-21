// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudStackServiceAccountState extends com.pulumi.resources.ResourceArgs {

    public static final CloudStackServiceAccountState Empty = new CloudStackServiceAccountState();

    /**
     * The disabled status for the service account. Defaults to `false`.
     * 
     */
    @Import(name="isDisabled")
    private @Nullable Output<Boolean> isDisabled;

    /**
     * @return The disabled status for the service account. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> isDisabled() {
        return Optional.ofNullable(this.isDisabled);
    }

    /**
     * The name of the service account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the service account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The basic role of the service account in the organization.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The basic role of the service account in the organization.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="stackSlug")
    private @Nullable Output<String> stackSlug;

    public Optional<Output<String>> stackSlug() {
        return Optional.ofNullable(this.stackSlug);
    }

    private CloudStackServiceAccountState() {}

    private CloudStackServiceAccountState(CloudStackServiceAccountState $) {
        this.isDisabled = $.isDisabled;
        this.name = $.name;
        this.role = $.role;
        this.stackSlug = $.stackSlug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudStackServiceAccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudStackServiceAccountState $;

        public Builder() {
            $ = new CloudStackServiceAccountState();
        }

        public Builder(CloudStackServiceAccountState defaults) {
            $ = new CloudStackServiceAccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param isDisabled The disabled status for the service account. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isDisabled(@Nullable Output<Boolean> isDisabled) {
            $.isDisabled = isDisabled;
            return this;
        }

        /**
         * @param isDisabled The disabled status for the service account. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder isDisabled(Boolean isDisabled) {
            return isDisabled(Output.of(isDisabled));
        }

        /**
         * @param name The name of the service account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the service account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param role The basic role of the service account in the organization.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The basic role of the service account in the organization.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder stackSlug(@Nullable Output<String> stackSlug) {
            $.stackSlug = stackSlug;
            return this;
        }

        public Builder stackSlug(String stackSlug) {
            return stackSlug(Output.of(stackSlug));
        }

        public CloudStackServiceAccountState build() {
            return $;
        }
    }

}
