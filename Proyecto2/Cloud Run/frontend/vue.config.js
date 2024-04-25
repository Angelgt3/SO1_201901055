const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/logs':{ 
        target:'http://localhost:3000'
      }
    }
  }
});
