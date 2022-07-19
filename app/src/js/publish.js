(() => {
    const url = new URL(window.location.href);
    /* 将块引用转化为超链接 */
    document.querySelectorAll(`#preview span[data-type="block-ref"][data-id]`).forEach(item => {
        const id = item.dataset.id;
        const a = document.createElement("a");
        a.href = `${url.origin}/block?id=${id}`;
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