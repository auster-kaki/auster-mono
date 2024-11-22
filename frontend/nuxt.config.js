import colors from 'vuetify/es5/util/colors'

export default {
  target: 'static',
  router: {
    base: '/auster-mono/'
  },
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: '%s - auster-frontend',
    title: 'auster-frontend',
    htmlAttrs: {
      lang: 'ja'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
      { name: 'apple-mobile-web-app-capable', content: 'yes' },// for pwa
      { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' } // for pwa
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'apple-touch-icon', href: '/icon.png' } // iOS用アイコン
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['~/assets/style/global/index.scss'],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
    '@nuxtjs/composition-api/module',
    '@pinia/nuxt'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/dotenv',
    '@nuxtjs/pwa'
  ],
  pwa: {
    manifest: {
      name: 'My PWA App', // アプリの名前
      short_name: 'MyApp', // アイコン下に表示される短縮名
      description: 'A Progressive Web App built with Nuxt.js',
      display: 'standalone', // フルスクリーンモードを指定
      start_url: '/', // 起動時のURL
      background_color: '#ffffff', // スプラッシュスクリーンの背景色
      theme_color: '#3CB371', // ステータスバーやナビゲーションバーの色
      lang: 'ja', // アプリの言語設定
      icons: [
        {
          src: '/pwa_icon_192x192.png',
          sizes: '192x192',
          type: 'image/png'
        },
        {
          src: '/pwa_icon_512x512.png',
          sizes: '512x512',
          type: 'image/png'
        }
      ]
    },
    workbox: {
      enabled: true // 開発モードでも有効にする場合は true
    }
  },
  env: {
    baseURL: process.env.BASE_URL || 'http://localhost:8080'
  },

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/style/variables/index.scss'],
    theme: {
      dark: false,
      themes: {
        light: {
          primary: '#4E97E0',
          accent: '#3CB371',
          secondary: '#3CB371',
          info: '#03A9F4',
          warning: '#FFEB3B',
          error: '#F44336',
          success: '#4CAF50'
        }
      }
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {}
}
