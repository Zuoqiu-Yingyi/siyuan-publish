/* 上下文 */
export {
    Context,
};

class Context {
    static async activate(plugin, mode) {
        var config;
        switch (mode) {
            case 'before':
                config = plugin.constructor.META.BEFORE;
                if (config) {
                    switch (true) {
                        case config.async === true && config.defer === true:
                            setTimeout(() => plugin.before(), 0);
                            break;
                        case config.async === true && config.defer === false:
                            plugin.before();
                            break;
                        case config.async === false && config.defer === true:
                            setTimeout(async () => await plugin.before(), 0);
                            break;
                        case config.async === false && config.defer === false:
                            await plugin.before();
                            break;
                    }
                }
                break;
            case 'after':
                config = plugin.constructor.META.AFTER;
                if (config) {
                    switch (true) {
                        case config.async === true && config.defer === true:
                            setTimeout(plugin.after, 0);
                            break;
                        case config.async === true && config.defer === false:
                            plugin.after();
                            break;
                        case config.async === false && config.defer === true:
                            setTimeout(async () => await plugin.after(), 0);
                            break;
                        case config.async === false && config.defer === false:
                            await plugin.after();
                            break;
                    }
                }
                break;
            default:
                return;
        }
    }

    constructor(meta = new Map(), data = new Map()) {
        this.meta = meta;
        this.data = data;
        this.document = window.document;
        this.siyuan = window.siyuan;
        this.publish = window.publish;
    }
}
