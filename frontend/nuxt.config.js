module.exports = {
  /*
   ** Headers of the page
   */
  head: {
    title: "infinimesh | Open Source IoT Platform",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: "Console infinimesh UI"
      }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      { rel: "icon", type: "image/png", href: "/favicon.png" }
    ]
  },
  plugins: ["@/plugins/ant-design-vue", "@/plugins/typeface-exo"],
  /*
   ** Customize the progress bar color
   */
  loading: { color: "#104E83" },
  /*
   ** Build configuration
   */
  build: {
    /*
     ** Run ESLint on save
     */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/
        });
      }
    },
    loaders: {
      less: { lessOptions: { javascriptEnabled: true } }
    }
  },
  buildModules: [
    "@nuxt/typescript-build",
    "@nuxtjs/style-resources",
    "@nuxtjs/color-mode",
    "@nuxtjs/vuetify"
  ],
  modules: [
    "@nuxtjs/axios",
    "@nuxtjs/auth",
    "nuxt-i18n",
    ["@nuxtjs/pwa", { meta: false, icon: false, manifest: false }]
  ],
  css: ["@/assets/main.css"],
  axios: {
    proxy: true
  },
  proxy: {
    "/api": {
      target: process.env.APISERVER_URL, // This will be defined on start
      pathRewrite: {
        "^/api": "/"
      }
    }
  },
  auth: {
    strategies: {
      local: {
        endpoints: {
          login: {
            url: "api/account/token",
            method: "post",
            propertyName: "token"
          },
          user: { url: "api/account", method: "get", propertyName: false },
          logout: false
        },
        tokenType: "bearer"
      }
    },
    redirect: {
      logout: "/login"
    }
  },
  router: {
    middleware: ["auth"]
  },
  server: {
    host: process.env.NODE_ENV == "production" ? "0.0.0.0" : "localhost",
    port: process.env.NODE_ENV == "production" ? 80 : 3000
  },

  publicRuntimeConfig: {
    baseURL: process.env.APISERVER_URL
  },

  i18n: {
    defaultLocale: "en",
    vueI18n: {
      fallbackLocale: "en"
    },
    lazy: true,
    langDir: "locale/",
    locales: [
      {
        code: "en",
        file: "en.js"
      }
    ],
    strategy: "no_prefix",
    detectBrowserLanguage: {
      useCookie: true,
      cookieKey: "inf_console_lang",
      onlyOnRoot: false
    }
  }
};
