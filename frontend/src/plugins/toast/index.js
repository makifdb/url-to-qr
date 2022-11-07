import Toast from './toast.vue'
import Api from './api.js'


const Plugin = (app, options = {}) => {
    let methods = Api(options)
    app.$toast = methods
    app.provide("toast", app.$toast)
    app.config.globalProperties.$toast = methods
}

Toast.install = Plugin

export default Toast
export { Toast }