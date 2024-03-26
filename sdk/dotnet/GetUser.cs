// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    public static class GetUser
    {
        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
        /// 
        /// This data source uses Grafana's admin APIs for reading users which
        /// does not currently work with API Tokens. You must use basic auth.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// using Grafana = Pulumiverse.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = new Grafana.User("test", new()
        ///     {
        ///         Email = "test.datasource@example.com",
        ///         Login = "test-datasource",
        ///         Password = "my-password",
        ///         IsAdmin = true,
        ///     });
        /// 
        ///     var fromId = Grafana.GetUser.Invoke(new()
        ///     {
        ///         UserId = test.UserId,
        ///     });
        /// 
        ///     var fromEmail = Grafana.GetUser.Invoke(new()
        ///     {
        ///         Email = test.Email,
        ///     });
        /// 
        ///     var fromLogin = Grafana.GetUser.Invoke(new()
        ///     {
        ///         Login = test.Login,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("grafana:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/user-management/server-user-management/)
        /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/user/)
        /// 
        /// This data source uses Grafana's admin APIs for reading users which
        /// does not currently work with API Tokens. You must use basic auth.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Grafana = Pulumi.Grafana;
        /// using Grafana = Pulumiverse.Grafana;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = new Grafana.User("test", new()
        ///     {
        ///         Email = "test.datasource@example.com",
        ///         Login = "test-datasource",
        ///         Password = "my-password",
        ///         IsAdmin = true,
        ///     });
        /// 
        ///     var fromId = Grafana.GetUser.Invoke(new()
        ///     {
        ///         UserId = test.UserId,
        ///     });
        /// 
        ///     var fromEmail = Grafana.GetUser.Invoke(new()
        ///     {
        ///         Email = test.Email,
        ///     });
        /// 
        ///     var fromLogin = Grafana.GetUser.Invoke(new()
        ///     {
        ///         Login = test.Login,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("grafana:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email address of the Grafana user. Defaults to ``.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        /// <summary>
        /// The username for the Grafana user. Defaults to ``.
        /// </summary>
        [Input("login")]
        public string? Login { get; set; }

        /// <summary>
        /// The numerical ID of the Grafana user. Defaults to `-1`.
        /// </summary>
        [Input("userId")]
        public int? UserId { get; set; }

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email address of the Grafana user. Defaults to ``.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The username for the Grafana user. Defaults to ``.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        /// <summary>
        /// The numerical ID of the Grafana user. Defaults to `-1`.
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The email address of the Grafana user. Defaults to ``.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the user is an admin.
        /// </summary>
        public readonly bool IsAdmin;
        /// <summary>
        /// The username for the Grafana user. Defaults to ``.
        /// </summary>
        public readonly string? Login;
        /// <summary>
        /// The display name for the Grafana user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The numerical ID of the Grafana user. Defaults to `-1`.
        /// </summary>
        public readonly int? UserId;

        [OutputConstructor]
        private GetUserResult(
            string? email,

            string id,

            bool isAdmin,

            string? login,

            string name,

            int? userId)
        {
            Email = email;
            Id = id;
            IsAdmin = isAdmin;
            Login = login;
            Name = name;
            UserId = userId;
        }
    }
}
