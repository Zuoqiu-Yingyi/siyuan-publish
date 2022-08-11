import { Base } from './../base/index.js';

export {
    Selected as Plugin,
};

class Selected extends Base {
    static META = {
        NAME: 'publish-selected',
        UUID: '979103CC-F25E-4BAC-8174-741DB7BA9B34',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '块选择插件',
        DEPENDENCY: [
        ],
    };

    constructor(context) {
        super(context);
        this.context.obj.set('selected', this);

        this.SELECTED = 'protyle-wysiwyg--select';

        this.context.meta.set('SELECTED', this.SELECTED);
    }

    /**
     * 取消所有块的选中状态
     */
    clearSelected() {
        Array.from(this.context.document.getElementsByClassName(this.SELECTED)).forEach(item => item.classList.remove(this.SELECTED));
    }

}
