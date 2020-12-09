export default {
  generics: {
    username_cap: "Username",
    username: "username",
    password_cap: "Password",
    login_cap: "Login",
    oops: "Oops",
    redirect: "Redirecting...",
    select_all: "Select All",
    deselect_all: "Deselect All"
  },

  login_page: {
    welcome_msg:
      "Welcome to infinimesh. Log in with your username and password.",
    input_email_placeholder: "Please input your E-Mail",
    input_pass_placeholder: "Please input your password",
    no_account_msg: "No account yet? Please contact us! Click here"
  },
  internal: {
    server_error: "Internal Server Error",
    response: ctx => `Response: ${ctx.list(0)}`,
    reponse_multiple: ctx =>
      `Result: success - ${ctx.list(0)}, failed: ${ctx.list(1)}.`
  },
  account_control: {
    create_success: "Account successfuly created!",
    create_error: "Failed to create an account",
    delete_success: "Account successfuly deleted!",
    delete_error: ctx => `Error deleting account ${ctx.list(0)}`,
    toogle_success: ctx =>
      `Account successfuly ${ctx.list(0) ? "disabled" : "enabled"}!`,
    toogle_error: ctx =>
      `Error ${ctx.list(0) ? "disabling" : "enabling"} account`
  },
  device_control: {
    not_found: "Device wasn't found",
    no_access: "You have no access to this device",
    create_success: "",
    create_error: "Failed to create the device",
    delete_success: "Device successfuly deleted!",
    delete_error: "Error deleting device!",
    toogle_multiple: ctx =>
      `${ctx.list(0) ? "Enabled" : "Disabled"} ${ctx.list(1)} devices!`,
    toogle_success: ctx =>
      `Device successfuly ${ctx.list(0) ? "disabled" : "enabled"}!`,
    toogle_error: ctx =>
      `Error ${ctx.list(0) ? "disabling" : "enabling"} device!`
  },
  devices_page: {
    search_placeholder: "Search device",
    group_by_tags: "Group by Tags",
    whole_registry: "Whole registry",

    enable_all: "Enable All",
    disable_all: "Disable All",
    root_ns_msg:
      "You can't create devices in your root namespace, switch to another one to perform device create.",
    create_or_switch_ns_msg:
      "Click here to create new namespace, or switch namespace on top of the page."
  },
  search_options: {
    all: "Everywhere",
    name: "Names",
    id: "IDs",
    tags: "Tags",
    namespace: "Namespace"
  }
};
