(() => {

    function then() {
        if (!window.publish.plugin.loaded) setTimeout(then, 0);
        else {
            /* 激活渲染后处理插件 */
            const plugins = window.publish.plugin.plugins;
            plugins.activate(window.publish.plugin.after, 'after');
        }
    }
    then();

    const POPOVER_TRIGGER = "popover-trigger"; // 可悬浮预览元素的类名
    const POPOVER_SIZE = { // 悬浮预览元素的尺寸
        width: window.top.publish.config.popover.width,
        height: window.top.publish.config.popover.height,
    };
})();
