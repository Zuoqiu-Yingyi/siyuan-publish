/* 上下文 */
export {
    Context,
};

class Context {
    constructor(obj = new Map(), meta = new Map(), data = new Map()) {
        this.obj = obj;
        this.meta = meta;
        this.data = data;

        this.window = window;

        this.top = window.top;
        this.location = window.location;
        this.document = window.document;

        this.siyuan = window.siyuan;
        this.publish = window.publish;
    }
}
