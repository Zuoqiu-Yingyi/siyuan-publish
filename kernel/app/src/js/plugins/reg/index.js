import { Base } from './../base/index.js';

export {
    REG as Plugin,
};

class REG extends Base {
    static META = {
        NAME: 'publish-reg',
        UUID: '7DB8FEC0-D1C9-481E-9EC8-063599F8344E',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '正则表达式',
        DEPENDENCY: [],
    };

    constructor(context) {
        super(context);
        this.REG = {
            id: /^\d{14}\-[0-9a-z]{7}$/,
            url: /^siyuan:\/\/blocks\/(\d{14}\-[0-9a-z]{7})/,
        };

        this.context.meta.set('REG', this.REG);
    }
}
