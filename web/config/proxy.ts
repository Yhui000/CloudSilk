/**
 * 在生产环境 代理是无法生效的，所以这里没有生产环境的配置
 * -------------------------------
 * The agent cannot take effect in the production environment
 * so there is no configuration of the production environment
 * For details, please see
 * https://pro.ant.design/docs/deploy
 */
export default {
  prod:{
    "/api/":{
      target: 'http://yunsuo.zh.city:50015',
      changeOrigin: true,
      pathRewrite: { '^': '' },
    },
  },
  test:{
    "/api/":{
      target: 'http://127.0.0.1:48089',
      changeOrigin: true,
      pathRewrite: { '^': '' },
    },
  },
};
