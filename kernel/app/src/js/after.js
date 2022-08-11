(() => {
    function then() {
        // console.log("after-then");
        if (!window.publish.plugin.loaded) setTimeout(then, 0);
        else {
            /* 激活渲染后处理插件 */
            const plugins = window.publish.plugin.plugins;
            plugins.activate(window.publish.plugin.after);
        }
    }
    then();
})();
