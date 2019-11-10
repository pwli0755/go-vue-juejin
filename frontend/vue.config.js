module.exports = {
  devServer: {
    port: 8080,
    proxy: {
      '^/api/v1/user': {
        target: 'http://localhost:3000',
        ws: true,
        changeOrigin: true,
      },
    },
  },
};
