export default {
  methods: {
    getAccountsPool() {
      const vm = this;
      vm.loading = true;
      vm.$axios
        .get("/api/accounts")
        .then(res => (vm.accounts = res.data.accounts))
        .catch(e => {
          if (e.response.status === 403) {
            vm.$notification.error({
              message: vm.$t("generics.oops"),
              description: e.response.data.message
            });
            vm.$store.commit("window/noAccess", "dashboard-accounts");
            vm.$router.push({ name: "dashboard-devices" });
          } else if (e.response.status === 500) {
            vm.$notification.error({
              message: vm.$t("internal.server_error"),
              description: e.response.data.message
            });
          }
        })
        .then(() => (vm.loading = false));
    },
    deleteAccount(account) {
      const vm = this;
      this.$axios({
        url: `/api/accounts/${account.uid}`,
        method: "delete"
      })
        .then(() => {
          vm.$message.success(vm.$t("account_control.delete_success"));
          vm.getAccountsPool();
        })
        .catch(e => {
          vm.$notification.error({
            message: vm.$t("account_control.delete_error", account.name),
            description: e.response.data.message
          });
        });
    },
    toogleAccount(account) {
      this.updateAccount(
        account.uid,
        {
          enabled: !account.enabled
        },
        this.$t("account_control.toogle_success", {
          res: account.enabled ? "disabled" : "enabled"
        }),
        this.$t("account_control.toogle_error", {
          res: account.enabled ? "disabling" : "enabling"
        })
      );
    },
    handleAccountAdd(account) {
      const vm = this;
      account.account.owner = vm.user.uid;
      vm.$axios({
        method: "post",
        url: "/api/accounts",
        data: account
      })
        .then(() => {
          vm.$notification.success({
            message: vm.$t("account_control.create_success")
          });
          vm.createAccountDrawerVisible = false;
          vm.getAccountsPool();
        })
        .catch(err => {
          this.$notification.error({
            message: vm.$t("account_control.create_error"),
            description: vm.$t("internal.response", err.response.data),
            duration: 10
          });
        });
    },
    updateAccount(id, data, success, error, success_callback, error_callback) {
      const vm = this;
      vm.loading = true;
      vm.$axios({
        method: "patch",
        url: `/api/accounts/${id}`,
        data: data
      })
        .then(
          success_callback
            ? success_callback
            : () => {
                vm.$message.success(success);
                vm.getAccountsPool();
              }
        )
        .catch(
          error_callback
            ? error_callback(e)
            : e => {
                vm.$notification.error({
                  message: error,
                  description: e.response.data.message
                });
              }
        )
        .then(() => (vm.loading = false));
    }
  }
};
