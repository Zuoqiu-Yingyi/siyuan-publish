import { Base } from './../base/index.js';

export {
    Theme as Plugin,
};

class Theme extends Base {
    static META = {
        NAME: 'publish-theme',
        UUID: '51B1E8AC-A3B6-4581-BBFF-DA8BE72C2BFD',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '主题渲染',
        DEPENDENCY: [
            'publish-url',
            'publish-dom',
        ],
        CALL: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.URL = this.context.meta.get('URL');
        this.DOM = this.context.meta.get('DOM');

        this.theme = this.URL.url.searchParams.get('theme');

        this.context.meta.set('theme', this.theme);
    }

    async call() {
        /* 使用 URL 参数 theme 设置主题模式 */
        switch (this.theme) {
            case 'light':
                this.setThemeMode(0);
                break;
            case 'dark':
                this.setThemeMode(1);
                break;
            case 'auto':
            default:
                /* 使用服务器配置设置主题模式 */
                switch (this.context.siyuan.config.appearance.mode) {
                    case 0:
                    default:
                        this.setThemeMode(0);
                        break;
                    case 1:
                        this.setThemeMode(1);
                        break;
                    case 2:
                        /* 使用浏览器配置设置主题模式 */
                        switch (true) {
                            case this.context.top.matchMedia('(prefers-color-scheme: light)').matches:
                                this.setThemeMode(0);
                                break;
                            case this.context.top.matchMedia('(prefers-color-scheme: dark)').matches:
                                this.setThemeMode(1);
                                break;
                            default:
                                this.setThemeMode(0);
                                break;
                        }
                        break;
                }
        }
    }

    /* 设置主题模式 */
    setThemeMode(mode) {
        switch (mode) {
            case 0:
                this.DOM.themeDefaultStyle.href = this.context.publish.config.theme.light.default;
                this.DOM.themeStyle.href = this.context.publish.config.theme.light.theme;
                this.DOM.themeCustomStyle.href = this.context.publish.config.theme.light.custom;
                this.context.siyuan.config.appearance.mode = 0;
                break;
            case 1:
                this.DOM.themeDefaultStyle.href = this.context.publish.config.theme.dark.default;
                this.DOM.themeStyle.href = this.context.publish.config.theme.dark.theme;
                this.DOM.themeCustomStyle.href = this.context.publish.config.theme.dark.custom;
                this.context.siyuan.config.appearance.mode = 1;
                break;
        }
    }
}
