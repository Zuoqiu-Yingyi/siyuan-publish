import { merge } from './utils.js';
import { Context } from './context.js';

/* 覆盖 ./app/templates/config.html 配置文件中定义的字段 */
merge(window.siyuan, {});
merge(window.publish, {});

/* 加载插件 */
const context = new Context();
for (const plugin of [...window.publish.plugin.before, ...window.publish.plugin.after]) {
    try {
        // 合并配置文件 custom.js
        const module = await import(`${window.publish.plugin.path}/${plugin}/index.js`);
        if (plugin) {
            const plugin_obj = new module.Plugin(context);
            window.publish.plugin.plugins.set(plugin, plugin_obj)
        }
    } catch (err) {
        console.warn(err);
    }
}

/* 激活渲染前处理插件 */
for (const plugin of window.publish.plugin.before) {
    if (window.publish.plugin.plugins.has(plugin)) {
        const plugin_obj = window.publish.plugin.plugins.get(plugin);
        await Context.activate(plugin_obj, 'before');
    }
}

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

    /* 设置主题模式 */
    function setThemeMode(mode) {
        switch (mode) {
            case 0:
                document.getElementById('themeDefaultStyle').href = window.publish.config.theme.light.default;
                document.getElementById('themeStyle').href = window.publish.config.theme.light.theme;
                document.getElementById('themeCustomStyle').href = window.publish.config.theme.light.custom;
                window.siyuan.config.appearance.mode = 0;
                break;
            case 1:
                document.getElementById('themeDefaultStyle').href = window.publish.config.theme.dark.default;
                document.getElementById('themeStyle').href = window.publish.config.theme.dark.theme;
                document.getElementById('themeCustomStyle').href = window.publish.config.theme.dark.custom;
                window.siyuan.config.appearance.mode = 1;
                break;
        }
    }

    /* 使用 URL 参数 theme 设置主题模式 */
    switch (url.searchParams.get('theme')) {
        case 'light':
            setThemeMode(0);
            break;
        case 'dark':
            setThemeMode(1);
            break;
        case 'auto':
        default:
            /* 使用服务器配置设置主题模式 */
            switch (window.siyuan.config.appearance.mode) {
                case 0:
                default:
                    setThemeMode(0);
                    break;
                case 1:
                    setThemeMode(1);
                    break;
                case 2:
                    /* 使用浏览器配置设置主题模式 */
                    switch (true) {
                        case window.matchMedia('(prefers-color-scheme: light)').matches:
                            setThemeMode(0);
                            break;
                        case window.matchMedia('(prefers-color-scheme: dark)').matches:
                            setThemeMode(1);
                            break;
                        default:
                            setThemeMode(0);
                            break;
                    }
                    break;
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
