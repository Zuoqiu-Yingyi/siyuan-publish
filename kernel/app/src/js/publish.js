(() => {
    const url = new URL(window.location.href);
    const REG = {
        url: /^siyuan:\/\/blocks\/(\d{14}\-[0-9a-z]{7})/,
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

    /* 将块引用转化为超链接 */
    const publish_url = new URL(url);
    publish_url.pathname = "/block";
    document.querySelectorAll(`#preview span[data-type="block-ref"][data-id]`).forEach(item => {
        const id = item.dataset.id;
        const a = document.createElement("a");
        publish_url.searchParams.set("id", id);
        a.href = publish_url.href;
        // a.target = "_blank";
        item.parentElement.replaceChild(a, item);
        a.appendChild(item);
    });

    /* 将链接转化为超链接 */
    document.querySelectorAll(`#preview span[data-type="a"][data-href]`).forEach(item => {
        const a = document.createElement("a");
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

    /* 定位到指定的块并高亮 */
    const id = url.searchParams.get("id");
    if (id) {
        const block = document.querySelector(`[data-node-id="${id}"]`);
        if (block) {
            block.classList.add("protyle-wysiwyg--select"); // 高亮指定的块
            block.scrollIntoView(true); // 滚动到指定的块
        }
    }
})();
