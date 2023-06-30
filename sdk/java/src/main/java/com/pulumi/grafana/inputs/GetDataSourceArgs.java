// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDataSourceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDataSourceArgs Empty = new GetDataSourceArgs();

    /**
     * The ID of this resource.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The ID of this resource.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="uid")
    private @Nullable Output<String> uid;

    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    private GetDataSourceArgs() {}

    private GetDataSourceArgs(GetDataSourceArgs $) {
        this.id = $.id;
        this.name = $.name;
        this.uid = $.uid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDataSourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDataSourceArgs $;

        public Builder() {
            $ = new GetDataSourceArgs();
        }

        public Builder(GetDataSourceArgs defaults) {
            $ = new GetDataSourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        public GetDataSourceArgs build() {
            return $;
        }
    }

}