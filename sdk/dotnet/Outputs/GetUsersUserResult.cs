// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        public readonly string Email;
        public readonly int Id;
        public readonly bool IsAdmin;
        public readonly string Login;
        public readonly string Name;

        [OutputConstructor]
        private GetUsersUserResult(
            string email,

            int id,

            bool isAdmin,

            string login,

            string name)
        {
            Email = email;
            Id = id;
            IsAdmin = isAdmin;
            Login = login;
            Name = name;
        }
    }
}
