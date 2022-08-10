import { Base } from './../base/index.js';

export {
    Locate as Plugin,
};

class Locate extends Base {
    static META = {
        NAME: 'publish-locate',
        UUID: 'CE73DFE8-6C07-403F-B656-F7D7AAF01CDD',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '定位目标块',
        DEPENDENCY: [
            'publish-url',
            'publish-selected',
        ],
        AFTER: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.URL = this.context.meta.get('URL');
        this.SELECTED = this.context.meta.get('SELECTED');;

        this.id = this.URL.url.searchParams.get("id");

        this.context.meta.set('id', this.id);
    }

    async after() {
        if (this.id) {
            const block = this.context.document.querySelector(`[data-node-id="${this.id}"]`);
            if (block) {
                block.classList.add(this.SELECTED); // 高亮指定的块
                block.scrollIntoView(true); // 滚动到指定的块
            }
        }
    }
}
