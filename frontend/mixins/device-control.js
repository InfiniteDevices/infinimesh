export default {
  methods: {
    refresh() {
      this.$axios
        .get(`/api/devices/${this.device.id}`)
        .then(res => {
          this.device = res.data.device;
          this.socket = new WebSocket(
            `wss://${this.$config.baseURL.replace("https://", "")}/devices/${
              this.device.id
            }/state/stream`,
            this.$auth
              .getToken("local")
              .replace("bearer", "Bearer")
              .split(" ")
          );
          this.socket.onmessage = msg => {
            let response = JSON.parse(msg.data).result;
            if (response)
              this.device.state.shadow.reported = response.reportedState;
          };
          window.addEventListener("beforeunload", function(event) {
            this.socket.close();
          });
        })
        .catch(res => {
          if (res.response.status == 404) {
            this.$notification.error({
              message: this.$t("device_control.not_found"),
              description: this.$t("generics.redirect")
            });
          } else if (res.response.status == 403) {
            this.$notification.error({
              message: this.$t("device_control.no_access"),
              description: this.$t("generics.redirect")
            });
          }
          this.$router.push({ name: "dashboard-devices" });
        });
      this.deviceStateGet();
    },
    /**
     * Obtains device state(s) (desired and reported) and merges them into deviceObject
     */
    async deviceStateGet() {
      await this.$axios
        .get(`/api/devices/${this.device.id}/state`)
        .then(res => {
          this.device = {
            ...this.device,
            state: res.data
          };
        });
    },
    /**
     * Performs PATCH /device/id and changes desired state to given.
     * Used by device-state component so it invokes callback after patch response handling is done
     * @param {String} state, the desired state as in JSON String
     */
    handleStateUpdate(state, callback) {
      this.$axios({
        url: `/api/devices/${this.device.id}/state`,
        method: "patch",
        data: state
      })
        .then(res => {
          this.deviceStateGet();
        })
        .catch(res => {
          console.error(res);
        })
        .then(() => {
          callback();
        });
    },
    handleDeviceDelete() {
      this.$axios({
        url: `/api/devices/${this.device.id}`,
        method: "delete"
      })
        .then(() => {
          this.$message.success(this.$t("device_control.create_success"));
          this.$store.dispatch("devices/get");
          this.$router.push({ name: "dashboard-devices" });
        })
        .catch(e => {
          this.$notification.error({
            message: this.$t("device_control.create_error"),
            description: e.response.data.message
          });
        });
    },
    handleToogleDevice(refresh = true) {
      this.handleDeviceUpdate(
        {
          enabled: !this.device.enabled
        },
        {
          refresh: refresh,
          success: () => {
            this.$message.success(
              this.$t("device_control.toogle_success", [device.enabled])
            );
          },
          error: () => {
            this.$notification.error({
              message: this.$t("device_control.toogle_error", [device.enabled]),
              description: e.response.data.message
            });
          }
        }
      );
    },
    handleDeviceUpdate(data, { refresh, success, error }) {
      this.$axios({
        url: `/api/devices/${this.device.id}`,
        method: "patch",
        data: data
      })
        .then(() => {
          if (success) success();
        })
        .catch(e => {
          if (error) error(e);
        })
        .then(() => {
          if (refresh) this.refresh();
          this.$store.dispatch("devices/get");
        });
    }
  }
};
