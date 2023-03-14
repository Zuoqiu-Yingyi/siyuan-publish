/* 渲染
 * REF: https://github.com/siyuan-note/siyuan/blob/master/app/src/protyle/export/index.ts
 */

(() => {
    const previewElement = document.getElementById('preview');

    Protyle.highlightRender(previewElement, window.publish.render.protyle);
    Protyle.mathRender(previewElement, window.publish.render.protyle, false);
    Protyle.mermaidRender(previewElement, window.publish.render.protyle);
    Protyle.flowchartRender(previewElement, window.publish.render.protyle);
    Protyle.graphvizRender(previewElement, window.publish.render.protyle);
    Protyle.chartRender(previewElement, window.publish.render.protyle);
    Protyle.mindmapRender(previewElement, window.publish.render.protyle);
    Protyle.abcRender(previewElement, window.publish.render.protyle);
    Protyle.plantumlRender(previewElement, window.publish.render.protyle);

    document.querySelectorAll(".protyle-action__copy").forEach((item) => {
        item.addEventListener("click", (event) => {
            navigator.clipboard.writeText(item.parentElement.nextElementSibling.textContent.trimEnd());
            event.preventDefault();
            event.stopPropagation();
        })
    });
})();
