<template>
  <a-modal
    :visible="active"
    :title="$t('reset_password.reset_account_base', account)"
    :okText="$t('generics.reset')"
    @ok="handleSubmit"
    @cancel="() => $emit('cancel')"
  >
    <a-form-model :model="model" :rules="rules" ref="resetAccountPasswordForm">
      <a-form-model-item prop="password" label="Password">
        <a-input-password
          :placeholder="$t('reset_password.password_placeholder')"
          v-model="model.password"
        />
      </a-form-model-item>
      <a-form-model-item prop="confirm_password" label="Confirm">
        <a-input-password
          :placeholder="$t('reset_password.confirm_password_placeholder')"
          v-model="model.confirm_password"
        />
      </a-form-model-item>
    </a-form-model>
  </a-modal>
</template>

<script>
export default {
  props: {
    active: {
      required: true,
    },
    account: {
      required: true,
    },
  },
  data() {
    return {
      model: {},
      rules: {
        password: [
          {
            required: true,
            message: this.$t("reset_password.password_required_msg"),
          },
        ],
        confirm_password: [
          {
            required: true,
            message: this.$t("reset_password.confirm_required_msg"),
          },
          {
            validator: (rule, value, raise) => {
              if (value != this.model.password) raise("oh noo");
            },
            message: this.$t("reset_password.dont_match_msg"),
          },
        ],
      },
    };
  },
  watch: {
    active: "setDefault",
  },
  mounted() {
    this.setDefault();
  },
  methods: {
    setDefault() {
      this.model = {
        password: "",
        confirm_password: "",
      };
    },
    handleSubmit() {
      let form = this.$refs.resetAccountPasswordForm;
      let errors = [];
      form.validateField(Object.keys(this.model), (err) => {
        if (err) {
          errors.push(err);
        }
      });
      if (errors.length === 0) this.$emit("reset", this.model.password);
    },
  },
};
</script>