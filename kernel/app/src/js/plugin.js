import { Context } from "./context.js";

export {
    plugins,
};

class Plugins {
    constructor() {
        this.list = new Array(); // 插件名称列表
        this.index = new Map(); // 插件名称 => 序号映射
        this.class = new Map(); // 已注册插件名称 => 插件类映射
        this.object = new Map(); // 已加载插件名称 => 插件对象映射
        this.context = new Context(); // 插件上下文
    }
    /* 注册插件 */
    register(Plugin) {
        if (!this.class.has(Plugin.META.NAME)) {
            this.class.set(Plugin.META.NAME, Plugin);
            this.index.set(Plugin.META.NAME, this.list.length);
            this.list.push(Plugin);
        }
        return Plugin.META.NAME;
    }
    /* 解析插件依赖 */
    resolve() {
        const resolver = new Kahn(this.list);
        this.list.forEach(Plugin => {
            Plugin.META.DEPENDENCY.forEach(dependency => {
                let dependency_index = this.index.get(dependency);
                if (dependency_index !== undefined) {
                    resolver.addEdge(
                        this.index.get(Plugin.META.NAME),
                        dependency_index,
                    )
                }
                else {
                    throw new Error(`plugin ${Plugin.META.NAME} depends on ${dependency} but it is not registered.`);
                }
            });
        });
        this.list = resolver.topoSort();
    }
    /* 加载插件 */
    load() {
        this.list.forEach(Plugin => {
            this.object.set(Plugin.META.NAME, new Plugin(this.context));
        });
    }
    /* 以指定模式激活插件组 */
    async activate(plugins) {
        for (let plugin of plugins) {
            plugin = this.object.get(plugin);
            const config = plugin.constructor.META.CALL;
            if (config) {
                switch (true) {
                    case config.async === true && config.defer === true:
                        setTimeout(() => plugin.call(), 0);
                        break;
                    case config.async === true && config.defer === false:
                        plugin.call();
                        break;
                    case config.async === false && config.defer === true:
                        setTimeout(async () => await plugin.call(), 0);
                        break;
                    case config.async === false && config.defer === false:
                        await plugin.call();
                        break;
                }
            }
        }
    }
}

/**Kahn 算法求解拓扑排序
 * REF [JavaScript 拓扑排序的实现](https://blog.csdn.net/krysliang/article/details/122328212)
 * 算法思想：
 * 如果 s 需要先于 t 执行，就添加一条s指向t的边。也就是说 a 先执行，b 依赖于 a 的执行，那么 a->b
 * 如果当前顶点没有任何边指向它，证明它可以先执行了
 * 所以算法的第一步，就是找出入度为 0 的顶点，然后输出，然后把所有他指向的顶点中的入度减1.
 * 然后不断的找入度为0的顶点，输出，其他顶点入度-1.
 */
class Kahn {
    constructor(nodes) {
        let list = new Array(nodes.length); // 邻接表 每个顶点后面紧接着是他指向别的顶点 顶点的集合
        for (let i = 0; i < nodes.length; i++) {
            list[i] = new Array(); // 顶点指向其他顶点的集合
        }

        this.nodes = nodes; // 节点数组
        this.list = list; // 邻接表
        this.result = []; // 结果
    }
    addEdge(s, t) {
        // s 依赖 t，t 先于 s 执行, t->s
        this.list[t].push(s);
    }
    topoSort() {
        // 统计每个顶点的入度(入边数)
        const indegrees = new Array(this.nodes.length).fill(0);
        for (let i = 0; i < this.nodes.length; i++) {
            for (let j = 0; j < this.list[i].length; j++) {
                indegrees[this.list[i][j]]++;
            }
        }

        const queue = []; // 入度为 0 的节点序号
        for (let i = 0; i < this.nodes.length; i++) {
            if (indegrees[i] == 0) queue.push(i); // 入度为 0 的节点入列
        }

        while (queue.length > 0) {
            let i = queue.shift(); // 队头出队
            this.result.push(this.nodes[i]); // 结果入列
            for (let j = 0; j < this.list[i].length; j++) {
                //对 i->k 的节点 k，全部入度 -1
                let k = this.list[i][j];
                indegrees[k]--;
                if (indegrees[k] == 0) queue.push(k);
            }
        }
        return this.result;
    }
}

const plugins = new Plugins();
