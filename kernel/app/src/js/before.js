import { merge } from './utils.js';
import { plugins } from './plugin.js';


(() => {
    /* 覆盖 ./app/templates/config.html 配置文件中定义的字段 */
    merge(window.siyuan, {});
    merge(window.publish, {});

    window.publish.plugin.plugins = plugins;

    /**异步并发请求
     * REF [使用promise实现多请求并发](https://blog.csdn.net/djh052900/article/details/124617912)
     */
    const modules = new Map();
    /* 并发请求 */
    function concurrentRequest(path) {
        const url = `${window.publish.plugin.path}/${path}/index.js`;
        return import(url);
    }
    [
        ...window.publish.plugin.load,
        ...window.publish.plugin.before,
        ...window.publish.plugin.after,
    ].forEach(path => {
        if (!modules.has(path)) modules.set(path, concurrentRequest(path));
    });

    Promise.all(modules).then(res => {
        // console.log(res);
        // console.log(modules);
        const path2plugin = new Map(); // 插件目录 => 插件类映射关系
        res.forEach(async item => {
            // console.log(item);
            const module = await item[1]
            // console.log(module);
            path2plugin.set(item[0], module.Plugin);
            plugins.register(module.Plugin);
        });
        function then() {
            // console.log("before-then");
            if (path2plugin.size !== modules.size) setTimeout(then, 0);
            else {
                window.publish.plugin.loaded = true;
                function path2name(paths) {
                    return paths.filter(path => path2plugin.has(path)).map(path => path2plugin.get(path).META.NAME);
                }
                window.publish.plugin.load = path2name(window.publish.plugin.load);
                window.publish.plugin.before = path2name(window.publish.plugin.before);
                window.publish.plugin.after = path2name(window.publish.plugin.after);

                plugins.resolve(); // 解析插件依赖
                plugins.load(); // 加载插件

                /* 激活渲染前处理插件 */
                plugins.activate(window.publish.plugin.before);
            }
        }
        then();
    }).catch((err) => {
        console.warn(err);
    });
})();
