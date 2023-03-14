import { Base } from './../base/index.js';

export {
    Link as Plugin,
};

class Link extends Base {
    static META = {
        NAME: 'publish-link',
        UUID: 'FF3EE694-A68D-4090-BC24-36A552304EF7',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.4',
        DESCRIPTION: '超链接处理',
        DEPENDENCY: [
            'publish-url',
            'publish-dom',
            'publish-icon',
            'publish-selected',
        ],
        CALL: {
            async: false,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.URL = this.context.meta.get('URL');
        this.REG = this.context.meta.get('REG');
        this.DOM = this.context.meta.get('DOM');
        this.CLASS_NAME_LINK_TRIGGER = 'link-trigger';
        this.CLASS_NAME_LINK_COPY = 'link-copy';
        this.TYPE_ICON_MAP = this.context.meta.get('TYPE_ICON_MAP');
        this.SELECTED = this.context.meta.get('SELECTED');

        this.publish_selected = this.context.obj.get('publish-selected');

        this.home = this.DOM.home;
        this.preview = this.DOM.preview;
        this.breadcrumb = this.DOM.breadcrumb;

        this.context.meta.set('CLASS_NAME_LINK_TRIGGER', this.CLASS_NAME_LINK_TRIGGER);
    }

    async call() {
        this.ref2link();
        this.a2link();
        this.home2link();
        this.breadcrumb2link();
        this.blockLink();
    }

    ref2link() {
        /* 将块引用转化为超链接 */
        this.preview.querySelectorAll(`span[data-type="block-ref"][data-id]`).forEach(item => {
            const id = item.dataset.id;
            const a = this.context.document.createElement("a");
            a.classList.add(this.CLASS_NAME_LINK_TRIGGER);
            this.URL.root.searchParams.set("id", id);
            a.href = this.URL.root.href.replace(this.URL.root.origin, "");
            // a.target = "_blank";
            item.parentElement.replaceChild(a, item);
            a.appendChild(item);
        });
    }

    a2link() {
        /* 将链接转化为超链接 */
        this.preview.querySelectorAll(`span[data-type="a"][data-href]`).forEach(item => {
            const a = this.context.document.createElement("a");
            a.classList.add(this.CLASS_NAME_LINK_TRIGGER)
            let href = item.dataset.href;
            if (this.REG.url.test(href)) { // 思源块超链接转化为站点超链接
                const id = this.REG.url.exec(href)[1];
                this.URL.root.searchParams.set("id", id);
                href = this.URL.root.href.replace(this.URL.root.origin, "");
            }
            a.href = href;
            // a.target = "_blank";
            item.parentElement.replaceChild(a, item);
            a.appendChild(item);
        });
    }

    home2link() {
        /* 文档首页转化为超链接 */
        if (this.home) {
            const a = this.context.document.createElement("a");
            a.href = window.publish.home.url;
            a.title = this.context.publish.i18n['home'];

            /* 为图标设置鼠标悬浮预览属性 */
            const icon = home.querySelector(".popover__block");
            if (icon) icon.classList.add(this.CLASS_NAME_LINK_TRIGGER);

            home.parentElement.replaceChild(a, home);
            a.appendChild(this.home);
        }
    }

    breadcrumb2link() {
        /* 将面包屑转化为超链接 */
        this.breadcrumb.querySelectorAll(`.protyle-breadcrumb__item[data-node-id]`).forEach(item => {
            const id = item.dataset.nodeId;
            const a = this.context.document.createElement("a");
            this.URL.root.searchParams.set("id", id);
            a.href = this.URL.root.href.replace(this.URL.root.origin, "");

            /* 为图标设置鼠标悬浮预览属性 */
            const icon = item.querySelector(".popover__block");
            if (icon) icon.classList.add(this.CLASS_NAME_LINK_TRIGGER);

            item.parentElement.replaceChild(a, item);
            a.appendChild(item);
        });
    }

    blockLink() {
        /* 为所有块添加悬浮复制超链接 */
        this.preview.querySelectorAll(`[data-node-id]`).forEach(item => {
            this.URL.root.searchParams.set("id", item.dataset.nodeId);
            const icon = typeof this.TYPE_ICON_MAP[item.dataset.type] === 'string'
                ? this.TYPE_ICON_MAP[item.dataset.type]
                : this.TYPE_ICON_MAP[item.dataset.type][item.dataset.subtype];
            const a = this.context.document.createElement("a");
            a.classList.add(this.CLASS_NAME_LINK_COPY);
            a.href = this.URL.root.href.replace(this.URL.root.origin, "");
            a.title = this.URL.root.href;
            a.innerHTML = `<svg style="height: 1rem; width: 1rem"><use xlink:href="${icon}"></use></svg>`

            /* 鼠标悬浮超链接时高亮对应的块 */
            a.addEventListener("mouseenter", () => {
                this.publish_selected.clearSelected(); // 取消所有块的高亮
                item.classList.add(this.SELECTED); // 高亮当前块
            });
            a.addEventListener("mouseleave", _ => this.publish_selected.clearSelected());

            item.appendChild(a);
            // item.parentElement.insertBefore(a, item);
        });
    }
}
