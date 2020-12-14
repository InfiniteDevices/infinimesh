export default {
  generics: {
    username_cap: "Username",
    username: "username",
    password_cap: "Password",
    login_cap: "Login",
    oops: "Oops",
    redirect: "Redirecting...",
    select_all: "Select All",
    deselect_all: "Deselect All",
    save: "Save",
    cancel: "Cancel",
    edit: "Edit",
    enable: "Enable",
    disable: "Disable",
    delete: "Delete",
    admin: "Admin",
    enabled: "Enabled",
    actions: "Actions",
    title: "Title",
    reset_password: "Reset password",
    reset_password_success: "Password changed successfuly",
    reset_password_error: "Reset password failed"
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
    response: "Response: {message}",
    response_multiple: "Result: success - {success}, failed: {fail}."
  },
  account_control: {
    create_success: "Account successfuly created!",
    create_error: "Failed to create an account",
    delete_success: "Account successfuly deleted!",
    delete_error: "Error deleting account {name}",
    toogle_success: "Account successfuly {res}!",
    toogle_error: "Error while {res} account"
  },
  accounts_page: {
    accounts: "Accounts",
    create_account: "Create Account",
    user_rights_b: "User",
    user_rights_e: "admin rights",
    user_rights_def: "has",
    user_rights_def_neg: "has no",
    toogle_rights: "Make {not}Admin",
    toogle_rights_success: "User {name} is now {q}Admin",
    toogle_rights_error: "Failed to make user {name} {q}Admin",
    user: "User",
    admin: "Admin"
  },
  device_control: {
    not_found: "Device wasn't found",
    no_access: "You have no access to this device",
    create_success: "",
    create_error: "Failed to create the device",
    update_success: "Device successfuly updated!",
    update_error: "Error updating device!",
    delete_success: "Device successfuly deleted!",
    delete_error: "Error deleting device!",
    toogle_multiple: "{res} {len} devices!",
    toogle_success: "Device successfuly {res}!",
    toogle_error: "Error {res} device!"
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
  device: {
    namespace: "Namespace",
    loading: "Loading device",
    name_placeholder: "Enter new device name",
    state: "Device is {not}enabled",
    tags: "Tags",
    no_tags: "No tags were provided",
    tags_placeholder: "Enter a comma-separated list of tags, e.g. tag1, tag2",
    state_reported: "Reported",
    state_desired: "Desired",
    state_slug: "State",
    details: "Details",
    actions: "Actions"
  },
  search_options: {
    all: "Everywhere",
    name: "Names",
    id: "IDs",
    tags: "Tags",
    namespace: "Namespace"
  },
  namespaces: {
    title: "Namespaces",
    create_ns: "Create Namespace",
    ns_name_placeholder: "Enter new name",
    delete_ns_hint:
      "Namespace and its devices won't be deleted immeadeatly, but after two weeks",
    deleted_warn_msg: "Going to be deleted on {date}, click to restore",

    ns_create_success: "Namespace successfuly created!",
    ns_create_error: "Failed to create a namespace!",

    ns_rename_success: "Namespace successfuly renamed!",
    ns_rename_error: "Error renaming namespace {name}!",

    ns_delete_success: "Namespace successfuly deleted!",
    ns_delete_error: "Error deleting namespace {name}!",

    ns_restore_success: "Namespace successfuly restored!",
    ns_restore_error: "Error restoring namespace {name}!"
  }
};
