import { merge } from './utils.js';
import { plugins } from './plugin.js';

/* 覆盖 ./app/templates/config.html 配置文件中定义的字段 */
merge(window.siyuan, {});
merge(window.publish, {});

/* 注册一组插件 */
async function registerPluginGroup(arr) {
    const result = [];
    for (let i = 0; i < arr.length; i++) {
        const url = `${window.publish.plugin.path}/${arr[i]}/index.js`;
        try {
            const module = await import(url);
            if (module) {
                plugins.register(module.Plugin);
                result.push(module.Plugin.META.NAME)
            }
        } catch (err) {
            console.error(`plugin file ${url} can't load.`);
        }
    }
    return result;
}
window.publish.plugin.load = await registerPluginGroup(window.publish.plugin.load);
window.publish.plugin.before = await registerPluginGroup(window.publish.plugin.before);
window.publish.plugin.after = await registerPluginGroup(window.publish.plugin.after);

plugins.resolve(); // 解析插件依赖
plugins.load(); // 加载插件

/* 激活渲染前处理插件 */
plugins.activate(window.publish.plugin.before, 'before');

window.publish.plugin.plugins = plugins;

/**
 * 创建图标
 * @params {string} icon: 图标名
 * @params {string} className: 类名
 * @params {string} id: 块 ID
 * @returns {HTMLElement} 图标
 */
window.publish.createIcon = (icon, className, id = null) => {
    const element = document.createElement('div');
    element.innerHTML = `
        <svg class="${className}">
            <use xlink:href="${icon}"></use>
        </svg>
    `;
    if (id) element.firstElementChild.dataset.nodeId = id;
    return element.firstElementChild;
}

/**
 * 创建面包屑项
 * @params {string} text: 面包屑项的文本
 * @params {string} id: 面包屑项的 ID
 * @params {string} icon: 面包屑项的图标
 * @params {boolean} active: 面包屑项是否激活
 * @returns {HTMLElement} 面包屑项
 */
window.publish.createBreadcrumbItem = (text, id = null, icon = '#iconFile', active = false) => {
    const breadcrumb__item = document.createElement('span');
    breadcrumb__item.classList.add('protyle-breadcrumb__item');
    if (active) breadcrumb__item.classList.add('protyle-breadcrumb__item--active');
    if (id) breadcrumb__item.dataset.nodeId = id;

    const breadcrumb__icon = window.publish.createIcon(icon, 'popover__block', id);
    const breadcrumb__text = document.createElement('span');
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

/**
 * 设置文档内容编辑状态
 * @params {boolean} contenteditable: 是否可编辑
 * @params {boolean} spellcheck: 是否开启拼写检查
 */
window.publish.setDocEditState = (contenteditable = false, spellcheck = false) => {
    document.querySelectorAll('[contenteditable][spellcheck]').forEach(item => {
        item.contentEditable = contenteditable;
        switch (true) {
            case item.classList.contains('hljs'): // 代码块禁用拼写检查
                break;
            default:
                item.spellcheck = spellcheck;
        }
    });
}

(() => {
    const url = new URL(window.location.href);

    /* 设置题头 */
    const background = document.querySelector(".protyle-background");
    if (background) {
        background.style.minHeight = window.publish.config.title.height;
    }
    const title_img = window.publish.config.title.img
    if (title_img) {
        const img = document.getElementById('title-img');
        const reg = /background-image:url\((.*)\);/
        if (reg.test(title_img)) { // 使用了自定义背景图片
            const src = reg.exec(title_img)[1];
            const style = title_img.replace(/background-image:url\((.*)\);/, "");
            img.src = src;
            img.style = style;
        } else { // 使用了随机背景图片
            img.style = title_img;
        }
    }

    /* 添加面包屑 */
    const breadcrumb = document.getElementById('breadcrumb');
    if (breadcrumb) {
        const paths = window.publish.document.path.split('/');
        const hpaths = window.publish.document.hpath.split('/');
        const depth = Math.min(paths.length, hpaths.length) - 1; // 当前文档的深度
        for (let i = 0; i <= depth; i++) {
            breadcrumb.appendChild(window.publish.createIcon(
                '#iconRight',
                'protyle-breadcrumb__arrow',
            ));
            breadcrumb.appendChild(window.publish.createBreadcrumbItem(
                hpaths[i],
                paths[i],
                '#iconFile',
                i === depth,
            ));
        }
    }

    /* 设置文档编辑状态 */
    window.publish.setDocEditState(
        window.siyuan.config.editor.contenteditable,
        window.siyuan.config.editor.spellcheck,
    );
})()
