// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.GetSlosSloQueryFreeform;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSlosSloQuery {
    private GetSlosSloQueryFreeform freeform;
    private String type;

    private GetSlosSloQuery() {}
    public GetSlosSloQueryFreeform freeform() {
        return this.freeform;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSlosSloQuery defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetSlosSloQueryFreeform freeform;
        private String type;
        public Builder() {}
        public Builder(GetSlosSloQuery defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.freeform = defaults.freeform;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder freeform(GetSlosSloQueryFreeform freeform) {
            this.freeform = Objects.requireNonNull(freeform);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetSlosSloQuery build() {
            final var o = new GetSlosSloQuery();
            o.freeform = freeform;
            o.type = type;
            return o;
        }
    }
}
