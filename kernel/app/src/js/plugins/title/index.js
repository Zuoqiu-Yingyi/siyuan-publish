import { Base } from './../base/index.js';

export {
    Title as Plugin,
};

class Title extends Base {
    static META = {
        NAME: 'publish-title',
        UUID: '40175F78-450B-460E-9CDE-F4DF241F7CFF',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '题头渲染',
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
        this.url = this.context.meta.get('URL');
        this.DOM = this.context.meta.get('DOM');
        this.background = this.DOM.background;
        this.titleImg = this.DOM.titleImg;

        this.title = {
            height: this.context.publish.config.title.height,
            img: this.context.publish.config.title.img,
        };

        this.context.meta.set('title', this.title);
    }

    async call() {
        /* 设置题头 */
        if (this.background) {
            this.background.style.minHeight = this.title.height;
        }
        if (this.titleImg && this.title.img) {
            const reg = /background-image:url\((.*)\);/
            if (reg.test(this.title.img)) { // 使用了自定义背景图片
                const src = reg.exec(this.title.img)[1];
                const style = this.title.img.replace(/background-image:url\((.*)\);/, "");
                this.titleImg.src = src;
                this.titleImg.style = style;
            } else { // 使用了随机背景图片
                this.titleImg.style = this.title.img;
            }
        }
    }
}
