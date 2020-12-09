export default {
  generics: {
    username_cap: "Username",
    username: "username",
    password_cap: "Password",
    login_cap: "Login",
    oops: "Oops"
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
    response: ctx => `Response: ${ctx.list(0)}`
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
  }
};
