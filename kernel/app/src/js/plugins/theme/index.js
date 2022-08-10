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
        DESCRIPTION: '页面主题设置',
        DEPENDENCY: [
            'publish-url',
            'publish-dom',
        ],
    };

    constructor(context) {
        super(context);
        this.url = this.context.meta.get('url');
        this.dom = this.context.meta.get('dom');

        this.theme = this.url.searchParams.get('theme');

        this.context.meta.set('theme', this.theme);
    }

    async before() {
        /* 使用 URL 参数 theme 设置主题模式 */
        switch (this.theme) {
            case 'light':
                setThemeMode(0);
                break;
            case 'dark':
                setThemeMode(1);
                break;
            case 'auto':
            default:
                /* 使用服务器配置设置主题模式 */
                switch (this.context.siyuan.config.appearance.mode) {
                    case 0:
                    default:
                        setThemeMode(0);
                        break;
                    case 1:
                        setThemeMode(1);
                        break;
                    case 2:
                        /* 使用浏览器配置设置主题模式 */
                        switch (true) {
                            case window.matchMedia('(prefers-color-scheme: light)').matches:
                                setThemeMode(0);
                                break;
                            case window.matchMedia('(prefers-color-scheme: dark)').matches:
                                setThemeMode(1);
                                break;
                            default:
                                setThemeMode(0);
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
                this.dom.themeDefaultStyle.href = this.context.publish.config.theme.light.default;
                this.dom.themeStyle.href = this.context.publish.config.theme.light.theme;
                this.dom.themeCustomStyle.href = this.context.publish.config.theme.light.custom;
                this.context.siyuan.config.appearance.mode = 0;
                break;
            case 1:
                this.dom.themeDefaultStyle.href = this.context.publish.config.theme.dark.default;
                this.dom.themeStyle.href = this.context.publish.config.theme.dark.theme;
                this.dom.themeCustomStyle.href = this.context.publish.config.theme.dark.custom;
                this.context.siyuan.config.appearance.mode = 1;
                break;
        }
    }
}
