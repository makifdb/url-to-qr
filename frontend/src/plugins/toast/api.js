import Toast from './toast.vue'
import mount from './mount'

const Api = (globalOptions = {}) => {
    return {
        show(message, duration, options = {}) {
            let localOptions = { message, duration, ...options }
            const c = mount(Toast, {
                props: { ...globalOptions, ...localOptions }
            })
            return c
        },

        success(message, duration, options = {}) {
            options.type = 'success'
            return this.show(message, duration, options)
        },
        error(message, duration, options = {}) {
            options.type = 'error'
            return this.show(message, duration, options)
        },
        info(message, duration, options = {}) {
            options.type = 'info'
            return this.show(message, duration, options)
        },
        warning(message, duration, options = {}) {
            options.type = 'warning'
            return this.show(message, duration, options)
        }
    }
}

export default Api