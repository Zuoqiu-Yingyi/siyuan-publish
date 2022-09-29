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
            'publish-dom',
        ],
        CALL: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.DOM = this.context.meta.get('DOM');

        this.context.meta.set('theme', this.theme);
    }

    async call() {
        /* 使用服务器配置设置主题模式 */
        switch (this.context.siyuan.config.appearance.mode) {
            case 0:
            default:
                this.DOM.themeDefaultStyle.href = this.context.publish.config.theme.light.default;
                this.DOM.themeStyle.href = this.context.publish.config.theme.light.theme;
                this.DOM.themeCustomStyle.href = this.context.publish.config.theme.light.custom;
                break;
            case 1:
                this.DOM.themeDefaultStyle.href = this.context.publish.config.theme.dark.default;
                this.DOM.themeStyle.href = this.context.publish.config.theme.dark.theme;
                this.DOM.themeCustomStyle.href = this.context.publish.config.theme.dark.custom;
                break;
        }

        /* 移除默认颜色 */
        this.context.document.body.style.color = null;
        this.context.document.body.style.backgroundColor = null;
    }
}
