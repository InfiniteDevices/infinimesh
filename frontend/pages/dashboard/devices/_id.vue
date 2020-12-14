<template>
  <div id="device">
    <a-spin
      :spinning="!device || !device.name"
      style="min-height: 15rem"
      size="large"
      :tip="$t('device.loading') + '...'"
    >
      <a-row style="padding-top: 10px">
        <a-col
          :xs="{ span: 21 }"
          :sm="{ span: 20, offset: 1 }"
          :md="{ span: 20, offset: 1 }"
          :lg="{ span: 20, offset: 1 }"
          :xl="{ span: 16, offset: 1 }"
          :xxl="{ span: 12, offset: 1 }"
        >
          <transition name="fade">
            <h1 class="lead" v-if="device && !active_edit">
              {{ device.name }}
              <span class="muted">{{ device.id }}</span>
            </h1>
          </transition>
          <a-input
            v-if="device && active_edit"
            :placeholder="$t('device.name_placeholder')"
            class="device-name-input"
            v-model="device.name"
          />
        </a-col>
        <a-col
          :xs="1"
          :sm="1"
          :md="1"
          :xl="{ span: 1, offset: 1 }"
          :xxl="{ span: 1, offset: 1 }"
        >
          <a-row type="flex" justify="end">
            <a-tooltip
              :title="$t('device.state', { not: device.enabled ? '' : 'not ' })"
              placement="right"
            >
              <transition name="fade">
                <a-icon
                  v-show="device"
                  type="bulb"
                  class="device-state-bulb"
                  :style="{ color: deviceStateBulbColor }"
                  theme="filled"
                />
              </transition>
            </a-tooltip>
          </a-row>
        </a-col>
      </a-row>
      <a-row>
        <a-col
          :sm="{ span: 22, offset: 1 }"
          :md="{ span: 22, offset: 1 }"
          :lg="{ span: 22, offset: 1 }"
          :xl="{ span: 18, offset: 1 }"
          :xxl="{ span: 14, offset: 1 }"
        >
          <transition-group name="slide">
            <a-card key="details" v-if="device" hoverable>
              <a-row slot="title" type="flex" justify="space-between">
                <a-col :span="3"> {{ $t("device.details") }} </a-col>
                <a-col :xxl="5" :lg="6" :md="9" :sm="10" v-if="active_edit">
                  <a-space>
                    <a-button
                      type="primary"
                      icon="close"
                      @click="active_edit = false"
                      >{{ $t("generics.cancel") }}</a-button
                    >
                    <a-button type="success" icon="save" @click="patchDevice">{{
                      $t("generics.save")
                    }}</a-button>
                  </a-space>
                </a-col>
                <a-col :lg="2" :md="3" :sm="4" v-else>
                  <a-button
                    type="primary"
                    icon="edit"
                    @click="active_edit = true"
                    >{{ $t("generics.edit") }}</a-button
                  >
                </a-col>
              </a-row>
              <template v-if="active_edit">
                <a-select
                  mode="tags"
                  :token-separators="[',']"
                  v-model="device.tags"
                  style="min-width: 50%; margin: 15px 0"
                  :placeholder="$t('device.tags_placeholder')"
                />
              </template>
              <template v-else>
                <a-row v-if="device.tags && device.tags.length">
                  <a-col :span="2">{{ $t("device.tags") }}:</a-col>
                  <a-col :span="22">
                    <a-tag
                      v-for="tag in device.tags"
                      :key="tag"
                      style="margin-bottom: 5px"
                      >{{ tag }}</a-tag
                    >
                  </a-col>
                </a-row>
                <a-row v-else type="flex" justify="center" class="muted">
                  <p>{{ $t("device.no_tags") }}</p>
                </a-row>
              </template>
              <template>
                <p>
                  {{ $t("device.namespace") }}:
                  <u>{{ device.namespace }}</u>
                </p>
              </template>
            </a-card>
            <a-card
              :title="$t('device.actions')"
              key="actions"
              v-if="device"
              hoverable
            >
              <device-actions
                :device="device"
                @delete="handleDeviceDelete"
                @toogle="handleToogleDevice"
              />
            </a-card>
            <a-card
              :title="$t('device.state_slug')"
              key="state"
              v-if="device && device.state"
              hoverable
            >
              <a-row>
                <a-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12" :xxl="12">
                  <device-state
                    :title="$t('device.state_reported')"
                    :state="device.state.shadow.reported"
                  />
                </a-col>
                <a-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12" :xxl="12">
                  <device-state
                    :title="$t('device.state_desired')"
                    :state="device.state.shadow.desired"
                    :editable="true"
                    @update="handleStateUpdate"
                  />
                </a-col>
              </a-row>
            </a-card>
          </transition-group>
        </a-col>
      </a-row>
    </a-spin>
  </div>
</template>

<script>
import DeviceState from "@/components/device/State";
import DeviceActions from "@/components/device/Actions";

import deviceControlMixin from "@/mixins/device-control";

export default {
  /**
   * Represents Device as both object and component
   * @displayName Device
   */
  components: { DeviceState, DeviceActions },
  mixins: [deviceControlMixin],
  props: {
    /**
     * Device ID - not required if component is mounted via Router _id
     */
    deviceId: {
      required: false,
    },
  },
  data() {
    return {
      deviceObject: false,
      active_edit: false,
    };
  },
  computed: {
    device: {
      get() {
        return this.deviceObject;
      },
      set(obj) {
        this.deviceObject = { ...this.deviceObject, ...obj };
      },
    },
    deviceStateBulbColor() {
      if (!(this.device && this.device.enabled !== undefined)) {
        return "black";
      } else if (this.device.enabled) {
        return "#52c41a";
      } else {
        return "#eb2f96";
      }
    },
  },
  mounted() {
    this.device = {
      id: this.deviceId || this.$route.params.id,
    };
    // Getting Device data from API
    this.refresh();
    this.$store.commit("window/setTopAction", {
      icon: "undo",
      callback: this.refresh,
    });
  },
  beforeDestroy() {
    this.$store.commit("window/unsetTopAction");
  },
  methods: {
    patchDevice() {
      this.active_edit = false;
      this.handleDeviceUpdate(
        {
          name: this.device.name,
          tags: this.device.tags,
        },
        {
          refresh: true,
          success: () => {
            this.$message.success(this.$t("device_control.update_success"));
          },
          error: (e) => {
            this.$notification.error({
              message: this.$t("device_control.update_error"),
              description: e.response.data.message,
            });
          },
        }
      );
    },
    validate({ params }) {
      return /0[xX][0-9a-fA-F]+/.test(params.id);
    },
  },
};
</script>

<style scoped>
.device-name-input {
  width: 50%;
  margin-bottom: 16px;
  font-size: 1.7rem;
  padding: 0 !important;
  line-height: normal !important;
}
</style>
<style>
#device {
  overflow: hidden;
  font-family: Exo;
  font-weight: 500;
}

.slide-leave-active,
.slide-enter-active {
  transition: 1s;
}
.slide-enter {
  transform: translate(100%, 0);
}
.slide-leave-to {
  transform: translate(-100%, 0);
}
.device-state-bulb {
  font-size: 1.5rem;
  padding-top: 0.8rem;
}

.ant-card + .ant-card {
  margin-top: 1rem;
}
</style>
