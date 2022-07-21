(() => {
    const url = new URL(window.location.href);

    /* 将块引用转化为超链接 */
    const a_url = new URL(url);
    a_url.pathname = "/block";
    document.querySelectorAll(`#preview span[data-type="block-ref"][data-id]`).forEach(item => {
        const id = item.dataset.id;
        const a = document.createElement("a");
        a_url.searchParams.set("id", id);
        a.href = a_url.href;
        // a.target = "_blank";
        item.parentElement.replaceChild(a, item);
        a.appendChild(item);
    });

    /* 将链接转化为超链接 */
    document.querySelectorAll(`#preview span[data-type="a"][data-href]`).forEach(item => {
        const href = item.dataset.href;
        const a = document.createElement("a");
        a.href = href;
        // a.target = "_blank";
        item.parentElement.replaceChild(a, item);
        a.appendChild(item);
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
