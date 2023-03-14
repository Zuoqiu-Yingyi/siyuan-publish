/* 渲染
 * REF: https://github.com/siyuan-note/siyuan/blob/master/app/src/protyle/export/index.ts
 */

(() => {
    const protyleElement = document.getElementById('protyle');

    Protyle.highlightRender(protyleElement, window.publish.render.protyle);
    Protyle.mathRender(protyleElement, window.publish.render.protyle, false);
    Protyle.mermaidRender(protyleElement, window.publish.render.protyle);
    Protyle.flowchartRender(protyleElement, window.publish.render.protyle);
    Protyle.graphvizRender(protyleElement, window.publish.render.protyle);
    Protyle.chartRender(protyleElement, window.publish.render.protyle);
    Protyle.mindmapRender(protyleElement, window.publish.render.protyle);
    Protyle.abcRender(protyleElement, window.publish.render.protyle);
    Protyle.plantumlRender(protyleElement, window.publish.render.protyle);

    document.querySelectorAll(".protyle-action__copy").forEach((item) => {
        item.addEventListener("click", (event) => {
            navigator.clipboard.writeText(item.parentElement.nextElementSibling.textContent.trimEnd());
            event.preventDefault();
            event.stopPropagation();
        })
    });
})();
