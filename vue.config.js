const { defineConfig } = require('@vue/cli-service')
const messages = require('./src/locales/messages');

module.exports = defineConfig({
  transpileDependencies: true,
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].i18n = messages.en; // Use the English messages for the noscript tag
        return args;
      })
  }
})