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
        VERSION: '0.0.2',
        DESCRIPTION: 'URL 参数解析',
        DEPENDENCY: [],
    };

    constructor(context) {
        super(context);
        this.url = new URL(this.context.location.href);
        this.root = new URL(this.url);
        this.root.pathname = "/block";

        this.context.meta.set('URL', {
            url: this.url,
            root: this.root,
        });
    }
}
