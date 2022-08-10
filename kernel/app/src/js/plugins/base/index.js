export {
    Base,
};

class Base {
    static META = {
        NAME: 'publish-base', // 插件名称
        UUID: '073A5F57-40EA-494D-A80A-C12DEAA41150', // 插件唯一标识符
        REPO: '', // 插件仓库地址
        AUTHOR: 'siyuan-publish', // 插件作者
        VERSION: '0.0.1', // 插件版本
        DESCRIPTION: '基础插件', // 插件描述
        DEPENDENCY: [], // 插件依赖(插件名称列表)
        BEFORE: { // 渲染前配置
            async: false, // 是否异步执行
            defer: false, // 是否延迟执行
        },
        AFTER: { // 渲染后配置
            async: false, // 是否异步执行
            defer: false, // 是否延迟执行
        },
    };
    static init() {
        console.log(`plugin ${this.META.NAME} v${this.META.VERSION} loaded.`);
    }

    /* 构造函数 */
    constructor(context) {
        this.context = context;
        this.constructor.init();
    }

    async before() { }
    async after() { }
}
