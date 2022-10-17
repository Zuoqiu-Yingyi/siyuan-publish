import { Base } from './../base/index.js';

export {
    Popover as Plugin,
};

class Popover extends Base {
    static META = {
        NAME: 'publish-popover',
        UUID: '2B17F759-5E17-4A00-B588-E3C447C8F73D',
        REPO: '',
        AUTHOR: 'siyuan-publish',
        VERSION: '0.0.3',
        DESCRIPTION: '鼠标悬浮预览',
        DEPENDENCY: [
            'publish-url'
        ],
        CALL: {
            async: false,
            defer: false,
        },
    };

    constructor(context) {
        super(context);
        this.URL = this.context.meta.get('URL')
        this.POPOVER_TRIGGER = 'popover-trigger';
        this.POPOVER_SIZE = { // 悬浮预览元素的尺寸
            width: this.context.top.publish.config.popover.width,
            height: this.context.top.publish.config.popover.height,
        };

        this.context.meta.set('POPOVER_TRIGGER', this.POPOVER_TRIGGER);
        this.context.meta.set('POPOVER_SIZE', this.POPOVER_SIZE);
    }

    async call() {
        /* 超链接鼠标悬浮预览 */
        if (this.context.top === this.context.window) { // 只在顶层窗口执行
            /* 注册多窗口共用的属性与方法 */
            this.context.publish.mouse = {
                position: { // 当前鼠标位置
                    x: 0,
                    y: 0,
                },
            };
            /* 获得鼠标位置 */
            document.addEventListener("mousemove", e => {
                /* 不能取消其他默认事件处理, 不然无法划选 */
                // e.preventDefault();
                // e.stopPropagation();

                this.context.publish.mouse.position.x = e.pageX;
                this.context.publish.mouse.position.y = e.pageY;
            });
            this.context.publish.popover = {
                drag: { // 拖拽
                    position: { // 拖动点相对于拖拽元素的位置
                        x: 0,
                        y: 0,
                    },
                    size: { // 拖拽前的尺寸
                        width: 0,
                        height: 0,
                    },
                },
                timeout: null, // 定时器
                mouse_position: { x: 0, y: 0 }, // 鼠标位置
                z_index: 1024, // 当前最高层级
                handler: (element) => { // 鼠标悬浮事件处理
                    const href = element.href
                        || (element.dataset.nodeId
                            ? (this.URL.root.searchParams.set("id", element.dataset.nodeId), this.URL.root.href)
                            : '#');
                    const doc = this.context.top.document; // 顶层窗口的 document
                    const popover = this.context.top.publish.popover;
                    const position = this.context.top.publish.mouse.position;
                    const TIMEOUT = {
                        SHOW_MIN: 125, // 悬浮预览元素的显示最小时间
                        CLOSE: 0, // 悬浮预览元素的关闭延时
                    }
                    // console.log(element);
                    const block__popover = doc.createElement("div"); // 悬浮预览显示元素
                    block__popover.classList.add("block__popover", "block__popover--move", "block__popover--open");
                    block__popover.style.zIndex = popover.z_index++;
                    block__popover.style.width = this.POPOVER_SIZE.width;
                    block__popover.style.height = this.POPOVER_SIZE.height;
                    const midline = { // 窗口中线
                        x: doc.documentElement.clientWidth / 2,
                        y: doc.documentElement.clientHeight / 2,
                    };
                    // console.log(position, midline);
                    switch (true) { // 判断当前鼠标在屏幕哪个象限中
                        case position.x <= midline.x
                            && position.y <= midline.y:
                            // 左上象限
                            block__popover.style.left = `${position.x}px`;
                            block__popover.style.top = `${position.y}px`;
                            break;
                        case position.x > midline.x
                            && position.y < midline.y:
                            // 右上象限
                            // block__popover.style.left = `${position.x - this.POPOVER_SIZE.width}px`;
                            block__popover.style.left = `calc(${position.x}px - ${this.POPOVER_SIZE.width})`;
                            block__popover.style.top = `${position.y}px`;
                            break;
                        case position.x < midline.x
                            && position.y > midline.y:
                            // 左下象限
                            block__popover.style.left = `${position.x}px`;
                            // block__popover.style.top = `${position.y - this.POPOVER_SIZE.height}px`;
                            block__popover.style.top = `calc(${position.y}px - ${this.POPOVER_SIZE.height})`;
                            break;
                        case position.x >= midline.x
                            && position.y >= midline.y:
                            // 右下象限
                            // block__popover.style.left = `${position.x - this.POPOVER_SIZE.width}px`;
                            // block__popover.style.top = `${position.y - this.POPOVER_SIZE.height}px`;
                            block__popover.style.left = `calc(${position.x}px - ${this.POPOVER_SIZE.width})`;
                            block__popover.style.top = `calc(${position.y}px - ${this.POPOVER_SIZE.height})`;
                            break;
                    }
                    block__popover.innerHTML = `
                    <div class="block__icons block__icons--border">
                        <span class="fn__space fn__flex-1"></span>
                        <span data-type="open-page" class="block__icon b3-tooltips b3-tooltips__sw" title="${this.context.publish.i18n['new-page']}">
                            <svg>
                                <use xlink:href="#iconExport"></use>
                            </svg>
                        </span>
                        <span class="fn__space"></span>
                        <span data-type="open-window" class="block__icon b3-tooltips b3-tooltips__sw" title="${this.context.publish.i18n['new-window']}">
                            <svg class="ft__secondary">
                                <use xlink:href="#iconExport"></use>
                            </svg>
                        </span>
                        <span class="fn__space"></span>
                        <span data-type="pin" class="block__icon b3-tooltips b3-tooltips__sw" title="${this.context.publish.i18n['pin']}">
                            <svg>
                                <use xlink:href="#iconPin"></use>
                            </svg>
                        </span>
                        <span class="fn__space"></span>
                        <span data-type="close" class="block__icon b3-tooltips b3-tooltips__sw" title="${this.context.publish.i18n['close']}">
                            <svg style="width: 10px">
                                <use xlink:href="#iconClose"></use>
                            </svg>
                        </span>
                    </div>
                    <div class="block__content">
                        <iframe src="${href}" border="0" frameborder="no" framespacing="0" allowfullscreen="true" class="fn__flex-1"></iframe>
                    </div>
                    <div class="block__nwse"></div>
                    <div class="block__ew"></div>
                    <div class="block__ns"></div>`;

                    /* 拖动功能 */
                    const border = block__popover.querySelector(".block__icons--border");
                    const iframe = block__popover.querySelector("iframe");
                    const width_handle = block__popover.querySelector(".block__ew");
                    const height_handle = block__popover.querySelector(".block__ns");
                    const size_handle = block__popover.querySelector(".block__nwse");

                    // 鼠标移动时状态
                    var flag_popover_dragging = false; // 悬浮预览窗口是否正在拖动

                    /* 标题栏拖动功能 */
                    /* 窗口拖动注册 */
                    const that = this;
                    function dragRegister(element, mousemoveHandler) {
                        element.addEventListener("mousedown", e => {
                            /* 取消其他默认事件处理 */
                            e.preventDefault();
                            e.stopPropagation();

                            flag_popover_dragging = true; // 正在拖拽
                            block__popover.style.zIndex = popover.z_index++; // 将当前窗口置顶
                            /* 记录鼠标与窗口的相对位置 */
                            popover.drag.position.x = e.clientX - block__popover.offsetLeft; // 鼠标相对于子窗口左上角的横向偏移量(鼠标横坐标 - popover 的 左侧偏移量)
                            popover.drag.position.y = e.clientY - block__popover.offsetTop; // 鼠标相对于子窗口左上角的纵向偏移量(鼠标纵坐标 - popover 的 上侧偏移量)
                            popover.drag.size.width = block__popover.offsetWidth; // 窗口宽度
                            popover.drag.size.height = block__popover.offsetHeight; // 窗口高度

                            that.dragMousedown(e, doc, mousemoveHandler);
                        });
                        element.addEventListener("mouseup", e => {
                            /* 取消其他默认事件处理 */
                            e.preventDefault();
                            e.stopPropagation();

                            flag_popover_dragging = false;
                            that.dragMouseup(e, doc, mousemoveHandler);
                        });
                    }
                    // REF [JS拖动浮动DIV - boystar - 博客园](https://www.cnblogs.com/boystar/p/5231697.html)
                    // REF [JS鼠标事件完成元素拖拽（简单-高级） - 百度文库](https://wenku.baidu.com/view/0c56050c3269a45177232f60ddccda38376be161?bfetype=new)
                    /* 悬浮子窗口拖拽 */
                    function popoverDrag(e) {
                        // console.log(e);
                        /* 子窗口左上角将要移动到的位置坐标 */
                        let x = e.clientX - popover.drag.position.x;
                        let y = e.clientY - popover.drag.position.y;

                        /* 子窗口左上角可以可以移动到区域边缘 */
                        let window_width = doc.documentElement.clientWidth - block__popover.offsetWidth;
                        let window_height = doc.documentElement.clientHeight - block__popover.offsetHeight;

                        x = (x < 0) ? 0 : x;                          // 当子窗口移动到主窗口最左边时
                        x = (x > window_width) ? window_width : x;    // 当子窗口移动到主窗口最右边时
                        y = (y < 0) ? 0 : y;                          // 当子窗口移动到主窗口最上边时
                        y = (y > window_height) ? window_height : y;  // 当子窗口移动到主窗口最下边时

                        block__popover.style.left = `${x}px`;
                        block__popover.style.top = `${y}px`;
                    }

                    /* 子窗口尺寸拖动调整功能 */
                    function calcSize(e) { // 计算子窗口应该调整到的尺寸
                        /** 计算窗口将要调整到的宽度
                         *  鼠标当前位置(应优化为控件的中轴) - 窗口左边位置
                         *      = 鼠标相对于窗口左边的横向偏移量
                         *  鼠标相对于窗口左边的横向偏移量 - 拖动前鼠标相对于窗口左边的横向偏移量
                         *      = 鼠标宽度变化量
                         *  原宽度 - 鼠标宽度变化量
                         *      = 当前应当设置的宽度
                         */
                        let width = popover.drag.size.width + (e.clientX - block__popover.offsetLeft - popover.drag.position.x);
                        /* 窗口宽度上限: 页面可视宽度 - 窗口左边位置 */
                        let max_width = doc.documentElement.clientWidth - block__popover.offsetLeft;
                        width = (width < 0) ? 0 : width; // 当子窗口移动到主窗口最左边时
                        width = (width > max_width) ? max_width : width; // 当子窗口移动到主窗口最右边时

                        let height = popover.drag.size.height + (e.clientY - block__popover.offsetTop - popover.drag.position.y);
                        let max_height = doc.documentElement.clientHeight - block__popover.offsetTop;
                        height = (height < 0) ? 0 : height; // 当子窗口移动到主窗口最上边时
                        height = (height > max_height) ? max_height : height; // 当子窗口移动到主窗口最下边时
                        return {
                            width: width,
                            height: height,
                        };
                    }
                    function widthDrag(e) {
                        block__popover.style.width = `${calcSize(e).width}px`;
                    }
                    function heightDrag(e) {
                        block__popover.style.height = `${calcSize(e).height}px`;
                    }
                    function sizeDrag(e) {
                        const size = calcSize(e);
                        block__popover.style.width = `${size.width}px`;
                        block__popover.style.height = `${size.height}px`;
                    }
                    dragRegister(border, popoverDrag);
                    dragRegister(width_handle, widthDrag);
                    dragRegister(height_handle, heightDrag);
                    dragRegister(size_handle, sizeDrag);

                    /* 鼠标移出预览时关闭预览 */
                    const icon_pin = block__popover.querySelector('[data-type="pin"]');
                    function close(_) {
                        if (!flag_popover_dragging && !icon_pin.classList.contains("block__icon--active")) {
                            setTimeout(() => block__popover.remove(), TIMEOUT.CLOSE);
                        }
                    }
                    /* 预览钉住/取消钉住 */
                    function pin(e) {
                        if (icon_pin.classList.contains("block__icon--active")) {
                            // 如果钉住按钮被激活
                            icon_pin.classList.remove("block__icon--active");
                            block__popover.addEventListener("mouseleave", close);
                        } else {
                            icon_pin.classList.add("block__icon--active");
                            block__popover.removeEventListener('mouseleave', close);
                        }
                    }
                    icon_pin.addEventListener("click", pin); // 钉住按钮
                    border.addEventListener("dblclick", pin); // 标题栏双击

                    /* 关闭按钮单击 */
                    block__popover.querySelector('[data-type="close"]').addEventListener("click", _ => block__popover.remove());
                    /* 鼠标移出元素后自动关闭 */
                    setTimeout(() => block__popover.addEventListener("mouseleave", close), TIMEOUT.SHOW_MIN);

                    /* 在新窗口打开 */
                    // REF [Window.open() - Web APIs | MDN](https://developer.mozilla.org/en-US/docs/Web/API/Window/open)
                    // REF [Window open() 方法 | 菜鸟教程](https://www.runoob.com/jsref/met-win-open.html)
                    const icon_open_page = block__popover.querySelector('[data-type="open-page"]');
                    const icon_open_window = block__popover.querySelector('[data-type="open-window"]');
                    icon_open_page.addEventListener("click", _ => this.context.top.open(iframe.src));
                    icon_open_window.addEventListener("click", _ => this.context.top.open(
                        iframe.src,
                        iframe.src,
                        `
                        popup = true,
                        left = ${this.context.top.screenX + block__popover.offsetLeft},
                        top = ${this.context.top.screenY + block__popover.offsetTop},
                        width = ${block__popover.offsetWidth},
                        height = ${block__popover.offsetHeight},
                    `,
                    ));

                    doc.body.append(block__popover);
                }
            }
        }
        else {
            this.context.document.body.style.backgroundImage = "none";
            this.context.document.body.style.backgroundColor = "transparent";
        }

        /* 鼠标悬浮在某个元素内一段时间后触发 */
        // REF [javascript - Iterating over result of getElementsByClassName using Array.forEach - Stack Overflow](https://stackoverflow.com/questions/3871547/iterating-over-result-of-getelementsbyclassname-using-array-foreach)
        Array.from(document.getElementsByClassName(this.POPOVER_TRIGGER)).forEach(item => {
            item.addEventListener("mouseenter", () => {
                if (this.context.top.publish.popover.timeout) {
                    clearTimeout(this.context.top.publish.popover.timeout);
                }
                this.context.top.publish.popover.timeout = setTimeout(
                    () => this.context.top.publish.popover.handler(item),
                    this.context.top.publish.config.popover.timeout,
                );
            });
            item.addEventListener("mouseleave", () => {
                if (this.context.top.publish.popover.timeout) {
                    clearTimeout(this.context.top.publish.popover.timeout);
                }
            });
        });
    }

    /**
     * 拖动功能鼠标按下时的处理器
     * @params {Event} e: 鼠标事件
     * @params {HTMLElement} stage: 在哪个元素内拖拽
     * @params {function} mousemoveHandler: 鼠标移动事件的处理器
     */
    dragMousedown(e, stage, mousemoveHandler) {
        /* 取消其他默认事件处理 */
        e.preventDefault();
        e.stopPropagation();

        /* 避免 mousemove 事件在 iframe 中无法触发 */
        // REF [在 iframe 上无法捕获 mousemove](https://blog.csdn.net/DongFuPanda/article/details/109533365)
        stage.querySelectorAll('iframe').forEach(iframe => iframe.style.pointerEvents = 'none');

        stage.addEventListener("mousemove", mousemoveHandler, true);
    }


    /**
     * 拖动功能鼠标抬起时的处理器
     * @params {Event} e: 鼠标事件
     * @params {HTMLElement} stage: 在哪个元素内拖拽
     * @params {function} mousemoveHandler: 鼠标移动事件的处理器
     */
    dragMouseup(e, stage, mousemoveHandler) {
        /* 取消其他默认事件处理 */
        e.preventDefault();
        e.stopPropagation();

        stage.querySelectorAll('iframe').forEach(iframe => iframe.style.pointerEvents = 'auto');

        stage.removeEventListener("mousemove", mousemoveHandler, true);
    }
}
