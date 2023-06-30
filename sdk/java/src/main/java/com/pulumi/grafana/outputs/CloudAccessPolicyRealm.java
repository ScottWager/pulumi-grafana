// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.CloudAccessPolicyRealmLabelPolicy;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class CloudAccessPolicyRealm {
    /**
     * @return The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
     * 
     */
    private String identifier;
    private @Nullable List<CloudAccessPolicyRealmLabelPolicy> labelPolicies;
    /**
     * @return Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
     * 
     */
    private String type;

    private CloudAccessPolicyRealm() {}
    /**
     * @return The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
     * 
     */
    public String identifier() {
        return this.identifier;
    }
    public List<CloudAccessPolicyRealmLabelPolicy> labelPolicies() {
        return this.labelPolicies == null ? List.of() : this.labelPolicies;
    }
    /**
     * @return Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CloudAccessPolicyRealm defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String identifier;
        private @Nullable List<CloudAccessPolicyRealmLabelPolicy> labelPolicies;
        private String type;
        public Builder() {}
        public Builder(CloudAccessPolicyRealm defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.identifier = defaults.identifier;
    	      this.labelPolicies = defaults.labelPolicies;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder identifier(String identifier) {
            this.identifier = Objects.requireNonNull(identifier);
            return this;
        }
        @CustomType.Setter
        public Builder labelPolicies(@Nullable List<CloudAccessPolicyRealmLabelPolicy> labelPolicies) {
            this.labelPolicies = labelPolicies;
            return this;
        }
        public Builder labelPolicies(CloudAccessPolicyRealmLabelPolicy... labelPolicies) {
            return labelPolicies(List.of(labelPolicies));
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public CloudAccessPolicyRealm build() {
            final var o = new CloudAccessPolicyRealm();
            o.identifier = identifier;
            o.labelPolicies = labelPolicies;
            o.type = type;
            return o;
        }
    }
}