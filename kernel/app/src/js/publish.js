(() => {
    const url = new URL(window.location.href);
    const REG = {
        url: /^siyuan:\/\/blocks\/(\d{14}\-[0-9a-z]{7})/,
    };

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

    /* 修改块超链接的 href */
    document.querySelectorAll(`#preview a[href ^="siyuan:"]`).forEach(item => {
        const href = item.href;
        item.href;
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
