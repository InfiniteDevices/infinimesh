<template>
  <div id="accountsTable">
    <a-row>
      <a-col :span="21" :offset="1">
        <a-row type="flex" align="middle" justify="space-between">
          <a-col>
            <h1 class="lead">{{ $t("accounts_page.accounts") }}</h1>
          </a-col>
          <a-col>
            <a-row type="flex" justify="end">
              <a-button
                type="primary"
                icon="plus"
                @click="createAccountDrawerVisible = true"
                >{{ $t("accounts_page.create_account") }}</a-button
              >
            </a-row>
            <account-add
              :active="createAccountDrawerVisible"
              @cancel="createAccountDrawerVisible = false"
              @add="handleAccountAdd"
            />
          </a-col>
        </a-row>
      </a-col>
    </a-row>
    <a-row style="margin-top: 10px">
      <a-col :span="21" :offset="1">
        <a-table
          :columns="columns"
          :data-source="accounts"
          :loading="loading"
          rowKey="uid"
          class="accounts-table"
          :show-header="false"
          :scroll="{ x: true }"
        >
          <span slot="name" slot-scope="name">
            <b>{{ name }}</b>
          </span>
          <span
            slot="uid"
            slot-scope="uid"
            v-if="user.is_admin || user.is_root"
          >
            <b class="muted">{{ uid }}</b>
          </span>
          <span slot="is_admin" slot-scope="is_admin">
            <a-row type="flex" justify="space-around">
              <a-tooltip>
                <span slot="title">
                  {{ $t("accounts_page.user_rights_b") }}
                  <u>{{
                    $t(
                      is_admin
                        ? "accounts_page.user_rights_def"
                        : "accounts_page.user_rights_def_neg"
                    )
                  }}</u>
                  {{ $t("accounts_page.user_rights_e") }}
                </span>
                {{
                  is_admin
                    ? $t("accounts_page.admin")
                    : $t("accounts_page.user")
                }}
              </a-tooltip>
            </a-row>
          </span>
          <span slot="enabled" slot-scope="enabled">
            <a-row type="flex" justify="space-around">
              <a-icon
                type="bulb"
                :style="{ color: enabled ? 'green' : 'red', fontSize: '24px' }"
              />
            </a-row>
          </span>
          <span slot="actions" slot-scope="text, account">
            <a-dropdown :trigger="['click']">
              <a-button type="link" icon="menu" />
              <a-menu slot="overlay">
                <a-menu-item>
                  <a-button
                    type="link"
                    @click="resetAccountPassword(account)"
                    >{{ $t("generics.reset_password") }}</a-button
                  >
                </a-menu-item>
                <a-menu-item @click="toogleAdmin(account)">
                  <a-button type="link">
                    {{
                      $t("accounts_page.toogle_rights", {
                        not: account.is_admin ? "not " : "",
                      })
                    }}
                  </a-button>
                </a-menu-item>
                <a-menu-item>
                  <a-button type="link" @click="toogleAccount(account)">{{
                    account.enabled
                      ? $t("generics.disable")
                      : $t("generics.enable")
                  }}</a-button>
                </a-menu-item>
                <a-menu-item>
                  <a-button type="link" @click="deleteAccount(account)">
                    {{ $t("generics.delete") }}
                  </a-button>
                </a-menu-item>
              </a-menu>
            </a-dropdown>
            <account-reset-password
              v-if="selectedAccount"
              :active="resetAccountPasswordVisible"
              :account="selectedAccount"
              @cancel="resetAccountPasswordVisible = false"
              @reset="handleResetAccountPassword"
            />
          </span>
        </a-table>
      </a-col>
    </a-row>
  </div>
</template>

<script>
import AccountAdd from "@/components/account/Add.vue";
import AccountResetPassword from "@/components/account/ResetPassword.vue";

import AccountControlMixin from "@/mixins/account-control";

export default {
  mixins: [AccountControlMixin],
  components: {
    AccountAdd,
    AccountResetPassword,
  },
  computed: {
    user() {
      return this.$store.getters.loggedInUser;
    },
  },
  data() {
    const columns = [
      {
        title: this.$t("generics.username_cap"),
        dataIndex: "name",
        sorter: true,
        scopedSlots: { customRender: "name" },
      },
      {
        title: "ID",
        dataIndex: "uid",
        sorter: true,
        scopedSlots: { customRender: "uid" },
      },
      {
        title: this.$t("generics.admin"),
        dataIndex: "is_admin",
        sorter: true,
        scopedSlots: { customRender: "is_admin" },
      },
      {
        title: this.$t("generics.enabled"),
        dataIndex: "enabled",
        sorter: true,
        scopedSlots: { customRender: "enabled" },
      },
      {
        title: this.$t("generics.actions"),
        key: "actions",
        fixed: "right",
        scopedSlots: { customRender: "actions" },
      },
    ];
    return {
      columns,
      accounts: [],
      loading: false,

      createAccountDrawerVisible: false,

      resetAccountPasswordVisible: false,
      selectedAccount: null,
    };
  },
  mounted() {
    this.getAccountsPool();
  },
  methods: {
    toogleAdmin(account) {
      const vm = this;
      vm.updateAccount(
        account.uid,
        { is_admin: !account.is_admin },
        vm.$t("accounts_page.toogle_rights_success", {
          name: account.name,
          q: account.is_admin ? "not " : "",
        }),
        vm.$t("accounts_page.toogle_rights_error", {
          name: account.name,
          q: account.is_admin ? "" : "not ",
        })
      );
    },
    resetAccountPassword(account) {
      this.selectedAccount = account;
      this.resetAccountPasswordVisible = true;
    },
    handleResetAccountPassword(password) {
      this.resetAccountPasswordVisible = false;
      this.updateAccount(
        this.selectedAccount.uid,
        { password: password },
        this.$t("generics.reset_password_success"),
        this.$t("generics.reset_password_error")
      );
    },
  },
};
</script>

<style>
.accounts-table .ant-table td {
  white-space: nowrap;
}
.ant-empty-description {
  color: lightgrey !important;
}
table.accounts-table {
  border-collapse: collapse;
}
.accounts-table > table,
th,
td {
  border-bottom: 1px solid var(--primary-color) !important;
  color: black;
}
</style>
