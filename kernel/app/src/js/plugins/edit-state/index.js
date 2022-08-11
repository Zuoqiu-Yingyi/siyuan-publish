import { Base } from './../base/index.js';

export {
    EditState as Plugin,
};

class EditState extends Base {
    static META = {
        NAME: 'publish-edit-state',
        UUID: 'CE73DFE8-6C07-403F-B656-F7D7AAF01CDD',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '编辑状态设置',
        DEPENDENCY: [
        ],
        CALL: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
    }

    async call() {
        /* 设置文档编辑状态 */
        this.setDocEditState(
            this.context.siyuan.config.editor.contenteditable,
            this.context.siyuan.config.editor.spellcheck,
        );
    }


    /**
     * 设置文档内容编辑状态
     * @params {boolean} contenteditable: 是否可编辑
     * @params {boolean} spellcheck: 是否开启拼写检查
     */
    setDocEditState(contenteditable = false, spellcheck = false) {
        document.querySelectorAll('[contenteditable][spellcheck]').forEach(item => {
            item.contentEditable = contenteditable;
            switch (true) {
                case item.classList.contains('hljs'): // 代码块禁用拼写检查
                    item.spellcheck = false;
                    break;
                default:
                    item.spellcheck = spellcheck;
                    break;
            }
        });
    }
}
