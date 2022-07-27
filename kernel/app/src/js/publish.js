/**
 * 拖动功能鼠标按下时的处理器
 * @params {Event} e: 鼠标事件
 * @params {HTMLElement} target: 拖拽的目标元素
 * @params {HTMLElement} satge: 在哪个元素内拖拽
 * @params {function} mousemoveHandler: 鼠标移动事件的处理器
 */
function dragMousedown(e, target, satge, mousemoveHandler) {
    console.log("dragMousedown");
    e.stopPropagation(); // 阻止冒泡
    target.removeEventListener("mousedown", dragMousedown); // 避免重复触发
    
    /* 避免 mousemove 事件在 iframe 中无法触发 */
    // REF [在 iframe 上无法捕获 mousemove](https://blog.csdn.net/DongFuPanda/article/details/109533365)
    satge.querySelectorAll('iframe').forEach(iframe => iframe.style.pointerEvents = 'none');
    
    satge.addEventListener("mousemove", mousemoveHandler, true);
}

/**
 * 拖动功能鼠标抬起时的处理器
 * @params {Event} e: 鼠标事件
 * @params {HTMLElement} target: 拖拽的目标元素
 * @params {HTMLElement} satge: 在哪个元素内拖拽
 * @params {function} mousemoveHandler: 鼠标移动事件的处理器
 */
function dragMouseup(e, target, satge, mousemoveHandler) {
    e.stopPropagation(); // 阻止冒泡
    satge.querySelectorAll('iframe').forEach(iframe => iframe.style.pointerEvents = 'auto');
    satge.removeEventListener("mousemove", mousemoveHandler, true);
    target.addEventListener("mousedown", dragMousedown);
}

(() => {
    const url = new URL(window.location.href);
    const REG = {
        url: /^siyuan:\/\/blocks\/(\d{14}\-[0-9a-z]{7})/,
    };
    const POPOVER_TRIGGER = "popover-trigger"; // 可悬浮预览元素的类名
    const POPOVER_TIMEOUT = 1000; // 悬浮预览元素的触发延时
    const POPOVER_SIZE = { // 悬浮预览元素的尺寸
        width: window.top.document.documentElement.clientWidth / 3,
        height: window.top.document.documentElement.clientHeight / 2,
    };
    const TYPE_ICON_MAP = {
        NodeAudio: "#iconRecord",
        NodeBlockQueryEmbed: "#iconSQL",
        NodeBlockquote: "#iconQuote",
        NodeCodeBlock: "#iconCode",
        NodeDocument: "#iconFile",
        NodeHTMLBlock: "#iconHTML5",
        NodeHeading: {
            h1: "#iconH1",
            h2: "#iconH2",
            h3: "#iconH3",
            h4: "#iconH4",
            h5: "#iconH5",
            h6: "#iconH6",
        },
        NodeIFrame: "#iconLanguage",
        NodeList: {
            o: "#iconList",
            u: "#iconOrderedList",
            t: "#iconCheck",
        },
        NodeListItem: "#iconListItem",
        NodeMathBlock: "#iconMath",
        NodeParagraph: "#iconParagraph",
        NodeSuperBlock: "#iconSuper",
        NodeTable: "#iconTable",
        NodeThematicBreak: "#iconLine",
        NodeVideo: "#iconVideo",
        NodeWidget: "#iconBoth",
    }

    /* 定位到指定的块并高亮 */
    const id = url.searchParams.get("id");
    if (id) {
        const block = document.querySelector(`[data-node-id="${id}"]`);
        if (block) {
            block.classList.add("protyle-wysiwyg--select"); // 高亮指定的块
            block.scrollIntoView(true); // 滚动到指定的块
        }
    }

    /* 将块引用转化为超链接 */
    const publish_url = new URL(url);
    publish_url.pathname = "/block";
    document.querySelectorAll(`#preview span[data-type="block-ref"][data-id]`).forEach(item => {
        const id = item.dataset.id;
        const a = document.createElement("a");
        a.classList.add(POPOVER_TRIGGER);
        publish_url.searchParams.set("id", id);
        a.href = publish_url.href;
        // a.target = "_blank";
        item.parentElement.replaceChild(a, item);
        a.appendChild(item);
    });

    /* 将链接转化为超链接 */
    document.querySelectorAll(`#preview span[data-type="a"][data-href]`).forEach(item => {
        const a = document.createElement("a");
        a.classList.add(POPOVER_TRIGGER)
        let href = item.dataset.href;
        if (REG.url.test(href)) { // 思源块超链接转化为站点超链接
            const id = REG.url.exec(href)[1];
            publish_url.searchParams.set("id", id);
            href = publish_url.href;
        }
        a.href = href;
        // a.target = "_blank";
        item.parentElement.replaceChild(a, item);
        a.appendChild(item);
    });

    /* 为所有块添加悬浮复制超链接 */
    document.querySelectorAll(`#preview [data-node-id]`).forEach(item => {
        publish_url.searchParams.set("id", item.dataset.nodeId);
        const icon = typeof TYPE_ICON_MAP[item.dataset.type] === 'string'
            ? TYPE_ICON_MAP[item.dataset.type]
            : TYPE_ICON_MAP[item.dataset.type][item.dataset.subtype];
        const a = document.createElement("a");
        a.classList.add("copy-link");
        a.href = publish_url.href;
        a.title = publish_url.href;
        a.innerHTML = `<svg style="height: 1rem; width: 1rem"><use xlink:href="${icon}"></use></svg>`
        item.appendChild(a);
        // item.parentElement.insertBefore(a, item);
    });

    /* 超链接鼠标悬浮预览 */
    if (window.top === window) { // 只在顶层窗口执行
        /* 注册多窗口共用的属性与方法 */
        window.publish.mouse = {
            position: { // 当前鼠标位置
                x: 0,
                y: 0,
            },
        };
        /* 获得鼠标位置 */
        document.addEventListener("mousemove", e => {
            window.publish.mouse.position.x = e.pageX;
            window.publish.mouse.position.y = e.pageY;
        });
        window.publish.popover = {
            drag: { // 拖拽
                position: { // 相对位置
                    x: 0,
                    y: 0,
                },
            },
            timeout: null, // 定时器
            mouse_position: { x: 0, y: 0 }, // 鼠标位置
            z_index: 0, // 当前最高层级
            handler: (element) => { // 鼠标悬浮事件处理
                const doc = window.top.document; // 顶层窗口的 document
                const popover = window.top.publish.popover;
                const position = window.top.publish.mouse.position;
                const TIMEOUT = {
                    SHOW_MIN: 125, // 悬浮预览元素的显示最小时间
                    CLOSE: 0, // 悬浮预览元素的关闭延时
                }
                // console.log(element);
                const block__popover = doc.createElement("div"); // 悬浮预览显示元素
                block__popover.classList.add("block__popover", "block__popover--move", "block__popover--open");
                block__popover.style.zIndex = popover.z_index++;
                block__popover.style.width = `${POPOVER_SIZE.width}px`;
                block__popover.style.height = `${POPOVER_SIZE.height}px`;
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
                        block__popover.style.left = `${position.x - POPOVER_SIZE.width}px`;
                        block__popover.style.top = `${position.y}px`;
                        break;
                    case position.x < midline.x
                        && position.y > midline.y:
                        // 左下象限
                        block__popover.style.left = `${position.x}px`;
                        block__popover.style.top = `${position.y - POPOVER_SIZE.height}px`;
                        break;
                    case position.x >= midline.x
                        && position.y >= midline.y:
                        // 右下象限
                        block__popover.style.left = `${position.x - POPOVER_SIZE.width}px`;
                        block__popover.style.top = `${position.y - POPOVER_SIZE.height}px`;
                        break;
                }
                block__popover.innerHTML = `
                    <div class="block__icons block__icons--border">
                        <span class="fn__space fn__flex-1"></span>
                        <span data-type="pin" class="block__icon b3-tooltips b3-tooltips__sw" title="钉住">
                            <svg>
                                <use xlink:href="#iconPin"></use>
                            </svg>
                        </span>
                        <span class="fn__space"></span>
                        <span data-type="close" class="block__icon b3-tooltips b3-tooltips__sw" title="关闭">
                            <svg style="width: 10px">
                                <use xlink:href="#iconClose"></use>
                            </svg>
                        </span>
                    </div>
                    <div class="block__content">
                        <iframe src="${element.href}" border="0" frameborder="no" framespacing="0" allowfullscreen="true" class="fn__flex-1"></iframe>
                    </div>
                    <div class="block__nwse"></div>
                    <div class="block__ew"></div>
                    <div class="block__ns"></div>`;

                /* 标题栏可以拖动 */
                // REF [JS拖动浮动DIV - boystar - 博客园](https://www.cnblogs.com/boystar/p/5231697.html)
                // REF [JS鼠标事件完成元素拖拽（简单-高级） - 百度文库](https://wenku.baidu.com/view/0c56050c3269a45177232f60ddccda38376be161?bfetype=new)
                const border = block__popover.querySelector(".block__icons--border");
                const iframe = block__popover.querySelector("iframe");
                // 鼠标移动时
                var gragging = false;

                /* 标题栏拖动功能 */
                function borderDrag(e) {
                    // console.log(e);
                    let x = e.clientX - popover.drag.position.x;
                    let y = e.clientY - popover.drag.position.y;
                    let window_width = doc.documentElement.clientWidth - block__popover.offsetWidth;
                    let window_height = doc.documentElement.clientHeight - block__popover.offsetHeight;

                    x = (x < 0) ? 0 : x;                          // 当div1到窗口最左边时
                    x = (x > window_width) ? window_width : x;    // 当div1到窗口最右边时
                    y = (y < 0) ? 0 : y;                          // 当div1到窗口最上边时
                    y = (y > window_height) ? window_height : y;  // 当div1到窗口最下边时

                    block__popover.style.left = `${x}px`;
                    block__popover.style.top = `${y}px`;
                }
                border.addEventListener("mousedown", e => {
                    gragging = true; // 正在拖拽
                    block__popover.style.zIndex = popover.z_index++; // 将当前窗口置顶
                    /* 记录鼠标相对于小窗标题栏的位置 */
                    popover.drag.position.x = e.clientX - block__popover.offsetLeft; // 鼠标相对于预览左上角的横向偏移量(鼠标横坐标 - popover 的 左侧偏移量)
                    popover.drag.position.y = e.clientY - block__popover.offsetTop; // 鼠标相对于预览左上角的纵向偏移量(鼠标纵坐标 - popover 的 上侧偏移量)

                    dragMousedown(e, border, doc, borderDrag);
                });
                border.addEventListener("mouseup", e => {
                    grabging = false;
                    dragMouseup(e, border, doc, borderDrag);
                });

                /* 鼠标移出预览时关闭预览 */
                const icon_pin = block__popover.querySelector('[data-type="pin"]');
                function close(_) {
                    if (!gragging && !icon_pin.classList.contains("block__icon--active")) {
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
                // TODO 子窗口支持拖动边缘调整大小
                doc.body.append(block__popover);
            }
        }
    }
    /* 鼠标悬浮在某个元素内一段时间后触发 */
    // REF [javascript - Iterating over result of getElementsByClassName using Array.forEach - Stack Overflow](https://stackoverflow.com/questions/3871547/iterating-over-result-of-getelementsbyclassname-using-array-foreach)
    Array.from(document.getElementsByClassName(POPOVER_TRIGGER)).forEach(item => {
        item.addEventListener("mouseenter", function () {
            if (window.top.publish.popover.timeout) {
                clearTimeout(window.top.publish.popover.timeout);
            }
            window.top.publish.popover.timeout = setTimeout(() => window.top.publish.popover.handler(item), POPOVER_TIMEOUT);
        });
        item.addEventListener("mouseleave", function () {
            if (window.top.publish.popover.timeout) {
                clearTimeout(window.top.publish.popover.timeout);
            }
        });
    });
})();
