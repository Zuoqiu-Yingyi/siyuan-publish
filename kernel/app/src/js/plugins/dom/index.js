import { Base } from './../base/index.js';

export {
    DOM as Plugin,
};

class DOM extends Base {
    static META = {
        NAME: 'publish-dom',
        UUID: '0B8E96C8-885B-41E6-BA4D-B24A7C2C7666',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.2',
        DESCRIPTION: '页面 DOM 元素',
        DEPENDENCY: [],
    };

    constructor(context) {
        super(context);
        this.DOM = {
            editorFontSize: this.context.document.getElementById('editorFontSize'),

            baseStyle: this.context.document.getElementById('baseStyle'),
            themeDefaultStyle: this.context.document.getElementById('themeDefaultStyle'),
            themeStyle: this.context.document.getElementById('themeStyle'),
            themeCustomStyle: this.context.document.getElementById('themeCustomStyle'),
            themeDefaultStyle: this.context.document.getElementById('themeDefaultStyle'),
            themeStyle: this.context.document.getElementById('themeStyle'),
            themeCustomStyle: this.context.document.getElementById('themeCustomStyle'),
            publishStyle: this.context.document.getElementById('publishStyle'),

            callScript: this.context.document.getElementById('callScript'),
            luteScript: this.context.document.getElementById('luteScript'),
            htmlScript: this.context.document.getElementById('htmlScript'),
            protyleScript: this.context.document.getElementById('protyleScript'),
            iconScript: this.context.document.getElementById('iconScript'),
            emojiScript: this.context.document.getElementById('emojiScript'),
            renderScript: this.context.document.getElementById('renderScript'),
            afterScript: this.context.document.getElementById('afterScript'),

            background: this.context.document.getElementById('background'),
            breadcrumb: this.context.document.getElementById('breadcrumb'),
            home: this.context.document.getElementById('home'),
            title: this.context.document.getElementById('title'),
            titleImg: this.context.document.getElementById('title-img'),
            titleTags: this.context.document.getElementById('title-tags'),
            titleIcon: this.context.document.getElementById('title-icon'),
            editor: this.context.document.getElementById('editor'),
            protyle: this.context.document.getElementById('protyle'),
        };

        this.context.meta.set('DOM', this.DOM);
    }
}
