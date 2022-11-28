// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Synthetic Monitoring checks are tests that run on selected probes at defined
 * intervals and report metrics and logs back to your Grafana Cloud account. The
 * target for checks can be a domain name, a server, or a website, depending on
 * what information you would like to gather about your endpoint. You can define
 * multiple checks for a single endpoint to check different capabilities.
 *
 * * [Official documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/checks/)
 *
 * ## Example Usage
 * ### DNS Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const dns = new grafana.SyntheticMonitoringCheck("dns", {
 *     job: "DNS Defaults",
 *     target: "grafana.com",
 *     enabled: false,
 *     probes: [main.then(main => main.probes?.Atlanta)],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         dns: {},
 *     },
 * });
 * ```
 * ### DNS Complex
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const dns = new grafana.SyntheticMonitoringCheck("dns", {
 *     job: "DNS Updated",
 *     target: "grafana.net",
 *     enabled: false,
 *     probes: [
 *         main.then(main => main.probes?.Frankfurt),
 *         main.then(main => main.probes?.London),
 *     ],
 *     labels: {
 *         foo: "baz",
 *     },
 *     settings: {
 *         dns: {
 *             ipVersion: "Any",
 *             server: "8.8.4.4",
 *             port: 8600,
 *             recordType: "CNAME",
 *             protocol: "TCP",
 *             validRCodes: [
 *                 "NOERROR",
 *                 "NOTAUTH",
 *             ],
 *             validateAnswerRrs: {
 *                 failIfMatchesRegexps: [".+-bad-stuff*"],
 *                 failIfNotMatchesRegexps: [".+-good-stuff*"],
 *             },
 *             validateAuthorityRrs: {
 *                 failIfMatchesRegexps: [".+-bad-stuff*"],
 *                 failIfNotMatchesRegexps: [".+-good-stuff*"],
 *             },
 *             validateAdditionalRrs: [{
 *                 failIfMatchesRegexps: [".+-bad-stuff*"],
 *                 failIfNotMatchesRegexps: [".+-good-stuff*"],
 *             }],
 *         },
 *     },
 * });
 * ```
 * ### HTTP Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const http = new grafana.SyntheticMonitoringCheck("http", {
 *     job: "HTTP Defaults",
 *     target: "https://grafana.com",
 *     enabled: false,
 *     probes: [main.then(main => main.probes?.Atlanta)],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         http: {},
 *     },
 * });
 * ```
 * ### HTTP Complex
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const http = new grafana.SyntheticMonitoringCheck("http", {
 *     job: "HTTP Defaults",
 *     target: "https://grafana.org",
 *     enabled: false,
 *     probes: [
 *         main.then(main => main.probes?.Bangalore),
 *         main.then(main => main.probes?.Mumbai),
 *     ],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         http: {
 *             ipVersion: "V6",
 *             method: "TRACE",
 *             body: "and spirit",
 *             noFollowRedirects: true,
 *             bearerToken: "asdfjkl;",
 *             proxyUrl: "https://almost-there",
 *             failIfSsl: true,
 *             failIfNotSsl: true,
 *             cacheBustingQueryParamName: "pineapple",
 *             tlsConfig: {
 *                 serverName: "grafana.org",
 *                 clientCert: `-----BEGIN CERTIFICATE-----
 * MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
 * RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
 * MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
 * 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
 * h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
 * BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
 * iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
 * a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
 * FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
 * qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
 * FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
 * Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
 * 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
 * UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
 * yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
 * e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
 * XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
 * tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
 * QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
 * tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
 * prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
 * 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
 * l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
 * 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
 * vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
 * -----END CERTIFICATE-----
 * `,
 *             },
 *             headers: ["Content-Type: multipart/form-data; boundary=something"],
 *             basicAuth: {
 *                 username: "open",
 *                 password: "sesame",
 *             },
 *             validStatusCodes: [
 *                 200,
 *                 201,
 *             ],
 *             validHttpVersions: [
 *                 "HTTP/1.0",
 *                 "HTTP/1.1",
 *                 "HTTP/2",
 *             ],
 *             failIfBodyMatchesRegexps: ["*bad stuff*"],
 *             failIfBodyNotMatchesRegexps: ["*good stuff*"],
 *             failIfHeaderMatchesRegexps: [{
 *                 header: "Content-Type",
 *                 regexp: "application/soap*",
 *                 allowMissing: true,
 *             }],
 *         },
 *     },
 * });
 * ```
 * ### Ping Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const ping = new grafana.SyntheticMonitoringCheck("ping", {
 *     job: "Ping Defaults",
 *     target: "grafana.com",
 *     enabled: false,
 *     probes: [main.then(main => main.probes?.Atlanta)],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         ping: {},
 *     },
 * });
 * ```
 * ### Ping Complex
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const ping = new grafana.SyntheticMonitoringCheck("ping", {
 *     job: "Ping Updated",
 *     target: "grafana.net",
 *     enabled: false,
 *     probes: [
 *         main.then(main => main.probes?.Frankfurt),
 *         main.then(main => main.probes?.London),
 *     ],
 *     labels: {
 *         foo: "baz",
 *     },
 *     settings: {
 *         ping: {
 *             ipVersion: "Any",
 *             payloadSize: 20,
 *             dontFragment: true,
 *         },
 *     },
 * });
 * ```
 * ### TCP Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const tcp = new grafana.SyntheticMonitoringCheck("tcp", {
 *     job: "TCP Defaults",
 *     target: "grafana.com:80",
 *     enabled: false,
 *     probes: [main.then(main => main.probes?.Atlanta)],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         tcp: {},
 *     },
 * });
 * ```
 * ### TCP Complex
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const tcp = new grafana.SyntheticMonitoringCheck("tcp", {
 *     job: "TCP Defaults",
 *     target: "grafana.com:443",
 *     enabled: false,
 *     probes: [
 *         main.then(main => main.probes?.Frankfurt),
 *         main.then(main => main.probes?.London),
 *     ],
 *     labels: {
 *         foo: "baz",
 *     },
 *     settings: {
 *         tcp: {
 *             ipVersion: "V6",
 *             tls: true,
 *             queryResponses: [
 *                 {
 *                     send: "howdy",
 *                     expect: "hi",
 *                 },
 *                 {
 *                     send: "like this",
 *                     expect: "like that",
 *                     startTls: true,
 *                 },
 *             ],
 *             tlsConfig: {
 *                 serverName: "grafana.com",
 *                 caCert: `-----BEGIN CERTIFICATE-----
 * MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
 * RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
 * MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
 * 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
 * h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
 * BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
 * iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
 * a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
 * FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
 * qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
 * FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
 * Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
 * 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
 * UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
 * yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
 * e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
 * XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
 * tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
 * QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
 * tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
 * prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
 * 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
 * l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
 * 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
 * vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
 * -----END CERTIFICATE-----
 * `,
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Traceroute Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const traceroute = new grafana.SyntheticMonitoringCheck("traceroute", {
 *     job: "Traceroute defaults",
 *     target: "grafana.com",
 *     enabled: false,
 *     frequency: 120000,
 *     timeout: 30000,
 *     probes: [main.then(main => main.probes?.Atlanta)],
 *     labels: {
 *         foo: "bar",
 *     },
 *     settings: {
 *         traceroute: {},
 *     },
 * });
 * ```
 * ### Traceroute Complex
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as grafana from "@lbrlabs/pulumi-grafana";
 * import * as grafana from "@pulumi/grafana";
 *
 * const main = grafana.getSyntheticMonitoringProbes({});
 * const traceroute = new grafana.SyntheticMonitoringCheck("traceroute", {
 *     job: "Traceroute complex",
 *     target: "grafana.net",
 *     enabled: false,
 *     frequency: 120000,
 *     timeout: 30000,
 *     probes: [
 *         main.then(main => main.probes?.Frankfurt),
 *         main.then(main => main.probes?.London),
 *     ],
 *     labels: {
 *         foo: "baz",
 *     },
 *     settings: {
 *         traceroute: {
 *             maxHops: 25,
 *             maxUnknownHops: 10,
 *             ptrLookup: false,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck check {{check-id}}
 * ```
 */
export class SyntheticMonitoringCheck extends pulumi.CustomResource {
    /**
     * Get an existing SyntheticMonitoringCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyntheticMonitoringCheckState, opts?: pulumi.CustomResourceOptions): SyntheticMonitoringCheck {
        return new SyntheticMonitoringCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck';

    /**
     * Returns true if the given object is an instance of SyntheticMonitoringCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyntheticMonitoringCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyntheticMonitoringCheck.__pulumiType;
    }

    /**
     * Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
     */
    public readonly alertSensitivity!: pulumi.Output<string | undefined>;
    /**
     * Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
     */
    public readonly basicMetricsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable the check. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
     */
    public readonly frequency!: pulumi.Output<number | undefined>;
    /**
     * Name used for job label.
     */
    public readonly job!: pulumi.Output<string>;
    /**
     * Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List of probe location IDs where this target will be checked from.
     */
    public readonly probes!: pulumi.Output<number[]>;
    /**
     * Check settings. Should contain exactly one nested block.
     */
    public readonly settings!: pulumi.Output<outputs.SyntheticMonitoringCheckSettings>;
    /**
     * Hostname to ping.
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * The tenant ID of the check.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<number>;
    /**
     * Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a SyntheticMonitoringCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyntheticMonitoringCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyntheticMonitoringCheckArgs | SyntheticMonitoringCheckState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyntheticMonitoringCheckState | undefined;
            resourceInputs["alertSensitivity"] = state ? state.alertSensitivity : undefined;
            resourceInputs["basicMetricsOnly"] = state ? state.basicMetricsOnly : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["job"] = state ? state.job : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["probes"] = state ? state.probes : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as SyntheticMonitoringCheckArgs | undefined;
            if ((!args || args.job === undefined) && !opts.urn) {
                throw new Error("Missing required property 'job'");
            }
            if ((!args || args.probes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'probes'");
            }
            if ((!args || args.settings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settings'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["alertSensitivity"] = args ? args.alertSensitivity : undefined;
            resourceInputs["basicMetricsOnly"] = args ? args.basicMetricsOnly : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["job"] = args ? args.job : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["probes"] = args ? args.probes : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["tenantId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SyntheticMonitoringCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyntheticMonitoringCheck resources.
 */
export interface SyntheticMonitoringCheckState {
    /**
     * Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
     */
    alertSensitivity?: pulumi.Input<string>;
    /**
     * Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
     */
    basicMetricsOnly?: pulumi.Input<boolean>;
    /**
     * Whether to enable the check. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
     */
    frequency?: pulumi.Input<number>;
    /**
     * Name used for job label.
     */
    job?: pulumi.Input<string>;
    /**
     * Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of probe location IDs where this target will be checked from.
     */
    probes?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Check settings. Should contain exactly one nested block.
     */
    settings?: pulumi.Input<inputs.SyntheticMonitoringCheckSettings>;
    /**
     * Hostname to ping.
     */
    target?: pulumi.Input<string>;
    /**
     * The tenant ID of the check.
     */
    tenantId?: pulumi.Input<number>;
    /**
     * Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SyntheticMonitoringCheck resource.
 */
export interface SyntheticMonitoringCheckArgs {
    /**
     * Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
     */
    alertSensitivity?: pulumi.Input<string>;
    /**
     * Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
     */
    basicMetricsOnly?: pulumi.Input<boolean>;
    /**
     * Whether to enable the check. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
     */
    frequency?: pulumi.Input<number>;
    /**
     * Name used for job label.
     */
    job: pulumi.Input<string>;
    /**
     * Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of probe location IDs where this target will be checked from.
     */
    probes: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Check settings. Should contain exactly one nested block.
     */
    settings: pulumi.Input<inputs.SyntheticMonitoringCheckSettings>;
    /**
     * Hostname to ping.
     */
    target: pulumi.Input<string>;
    /**
     * Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
     */
    timeout?: pulumi.Input<number>;
}
