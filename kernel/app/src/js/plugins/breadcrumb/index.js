import { Base } from './../base/index.js';

export {
    Breadcrumb as Plugin,
};

class Breadcrumb extends Base {
    static META = {
        NAME: 'publish-breadcrumb',
        UUID: '7DC055D5-6644-4E37-B6B7-E8BDBC4E982D',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.1',
        DESCRIPTION: '面包屑渲染',
        DEPENDENCY: [
            'publish-dom',
            'publish-icon',
        ],
        BEFORE: {
            async: true,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.DOM = this.context.meta.get('DOM');
        this.createIcon = this.context.hand.get('createIcon');

        this.breadcrumb = this.DOM.breadcrumb;

        this.context.hand.set('createBreadcrumbItem', this.createBreadcrumbItem);
    }

    async before() {
        if (this.breadcrumb) {
            const paths = this.context.publish.document.path.split('/');
            const hpaths = this.context.publish.document.hpath.split('/');
            const depth = Math.min(paths.length, hpaths.length) - 1; // 当前文档的深度
            for (let i = 0; i <= depth; i++) {
                breadcrumb.appendChild(this.createIcon(
                    '#iconRight',
                    'protyle-breadcrumb__arrow',
                ));
                breadcrumb.appendChild(this.createBreadcrumbItem(
                    hpaths[i],
                    paths[i],
                    '#iconFile',
                    i === depth,
                ));
            }
        }
    }


    /**
     * 创建面包屑项
     * @params {string} text: 面包屑项的文本
     * @params {string} id: 面包屑项的 ID
     * @params {string} icon: 面包屑项的图标
     * @params {boolean} active: 面包屑项是否激活
     * @returns {HTMLElement} 面包屑项
     */
    createBreadcrumbItem(text, id = null, icon = '#iconFile', active = false) {
        const breadcrumb__item = this.context.document.createElement('span');
        breadcrumb__item.classList.add('protyle-breadcrumb__item');
        if (active) breadcrumb__item.classList.add('protyle-breadcrumb__item--active');
        if (id) breadcrumb__item.dataset.nodeId = id;

        const breadcrumb__icon = this.createIcon(icon, 'popover__block', id);
        const breadcrumb__text = this.context.document.createElement('span');
        breadcrumb__text.classList.add(
            'protyle-breadcrumb__text',
            'protyle-breadcrumb__text--ellipsis',
        );
        breadcrumb__text.innerText = text;
        breadcrumb__item.title = text;

        breadcrumb__item.appendChild(breadcrumb__icon);
        breadcrumb__item.appendChild(breadcrumb__text);
        return breadcrumb__item;
    }
}
