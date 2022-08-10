import { Base } from './../base/index.js';

export {
    Url as Plugin,
};

class Url extends Base {
    static META = {
        NAME: 'publish-url',
        UUID: '6FF7DD8E-AAC9-4521-88F2-45E997E29992',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: 'URL 参数解析',
        DEPENDENCY: [],
    };

    constructor(context) {
        super(context);
        this.url = new URL(this.context.location.href);
        this.context.meta.set('url', this.url);
    }
}
